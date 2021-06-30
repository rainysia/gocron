package authldap

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/go-ldap/ldap"
	"github.com/ouqiang/gocron/internal/modules/app"
)

type UserAuthLdap struct {
	Username                 string
	Password                 string
	Email                    string
	IsAdmin, IsUser, IsGuest int8 // IsAdmin 1:管理员, 0:不是 IsUser 1:普通用户, IsGuest 1:访客 , 不在组的不能访问
}

func SetAuth(admin int8, user int8, guest int8) UserAuthLdap {
	return UserAuthLdap{
		IsAdmin: admin,
		IsUser:  user,
		IsGuest: guest,
	}
}

// Guest can visit
func (ual *UserAuthLdap) CanVist() bool {
	return false
}

// User can execute
func (ual *UserAuthLdap) CanExecute() bool {
	return false
}

// Admin can execute/modify/delete
func (ual *UserAuthLdap) CanDoEveryThing() bool {
	return false
}

// LDAP auth
func (ual *UserAuthLdap) VerifyLdapAuth() (err error) {
	l, err := ldap.DialURL(app.Setting.Ldap.Addr)

	if err != nil {
		log.Println(err)
		return err
	}
	defer l.Close()

	// Reconnect with TLS
	err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
	if err != nil {
		log.Println(err)
		return err
	}

	// First bind with a read only user
	err = l.Bind(app.Setting.Ldap.User, app.Setting.Ldap.Password)
	if err != nil {
		log.Println(err)
		return err
	}

	// Search for the given username
	searchRequest := ldap.NewSearchRequest(
		app.Setting.Ldap.Dn,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf("(&(objectClass=%s)(uid=%s))", app.Setting.Ldap.ObjClassUser, ldap.EscapeFilter(ual.Username)),
		[]string{},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Println(err)
		return err
	}

	if len(sr.Entries) != 1 {
		log.Println("User does not exist or too many entries returned")
		return errors.New("User does not exist or too many entries returned")
	}

	user := sr.Entries[0]
	userdn := user.DN

	// Bind as the user to verify their password
	err = l.Bind(userdn, ual.Password)
	if err != nil {
		log.Println(err)
		return err
	}

	// Rebind as the read only user for any further queries
	err = l.Bind(app.Setting.Ldap.User, app.Setting.Ldap.Password)
	if err != nil {
		log.Println(err)
		return err
	}
	//log.Printf(" passed %s", ldap.EscapeFilter(userdn))
	email := user.GetAttributeValue("mail")
	(*ual).Email = email

	// Query Group to check IsAdmin
	searchAdminGroupRequest := ldap.NewSearchRequest(
		app.Setting.Ldap.DnAdmin, // Query specified DN
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf("(&(objectClass=%s))", app.Setting.Ldap.ObjClassMember),
		[]string{},
		nil,
	)
	agsr, err := l.Search(searchAdminGroupRequest)
	if err != nil {
		log.Println(err)
		return err
	}

	if len(agsr.Entries) != 1 {
		log.Println("Admin group does not exist or too many entries returned")
		return errors.New("Admin group does not exist or too many entries returned")
	}
	adminGroupData := agsr.Entries[0]
	adminMembers := adminGroupData.GetAttributeValues("member")

	if len(adminMembers) > 0 {
		for _, v := range adminMembers {
			if strings.Contains(v, fmt.Sprintf("cn=%s", ual.Username)) {
				(*ual).IsAdmin = 1
			}
		}
	}

	// Query Group to check IsUser
	searchUserGroupRequest := ldap.NewSearchRequest(
		app.Setting.Ldap.DnUser,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf("(&(objectClass=%s))", app.Setting.Ldap.ObjClassMember),
		[]string{},
		nil,
	)
	ugsr, err := l.Search(searchUserGroupRequest)
	if err != nil {
		log.Println(err)
		return err
	}

	if len(ugsr.Entries) != 1 {
		log.Println("User group does not exist or too many entries returned")
		return errors.New("User group does not exist or too many entries returned")
	}
	userGroupData := ugsr.Entries[0]
	userMembers := userGroupData.GetAttributeValues("member")

	if len(userMembers) > 0 {
		for _, v := range userMembers {
			if strings.Contains(v, fmt.Sprintf("cn=%s", ual.Username)) {
				(*ual).IsUser = 1
			}
		}
	}

	// Query Group to check IsGuest
	searchGuestGroupRequest := ldap.NewSearchRequest(
		app.Setting.Ldap.DnGuest,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf("(&(objectClass=%s))", app.Setting.Ldap.ObjClassMember),
		[]string{},
		nil,
	)
	ggsr, err := l.Search(searchGuestGroupRequest)
	if err != nil {
		log.Println(err)
		return err
	}

	if len(ggsr.Entries) != 1 {
		log.Println("Guest group does not exist or too many entries returned")
		return errors.New("Guest group does not exist or too many entries returned")
	}
	guestGroupData := ggsr.Entries[0]
	guestMembers := guestGroupData.GetAttributeValues("member")

	if len(guestMembers) > 0 {
		for _, v := range guestMembers {
			if strings.Contains(v, fmt.Sprintf("cn=%s", ual.Username)) {
				(*ual).IsGuest = 1
			}
		}
	}

	return nil
}