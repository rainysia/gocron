class User {
  get () {
    return {
      'token': this.getToken(),
      'uid': this.getUid(),
      'username': this.getUsername(),
      'isAdmin': this.getIsAdmin(),
      'isStaff': this.getIsStaff()
    }
  }

  getToken () {
    return localStorage.getItem('token') || ''
  }

  setToken (token) {
    localStorage.setItem('token', token)
    return this
  }

  clear () {
    localStorage.clear()
  }

  getUid () {
    return localStorage.getItem('uid') || ''
  }

  setUid (uid) {
    localStorage.setItem('uid', uid)
    return this
  }

  getUsername () {
    return localStorage.getItem('username') || ''
  }

  setUsername (username) {
    localStorage.setItem('username', username)
    return this
  }

  getIsAdmin () {
    let isAdmin = localStorage.getItem('is_admin')
    return isAdmin === '1' || isAdmin === '2'
  }

  setIsAdmin (isAdmin) {
    localStorage.setItem('is_admin', isAdmin)
    return this
  }

  getIsStaff () {
    return localStorage.getItem('is_staff') === '1'
  }

  setIsStaff (isAdmin) {
    let isStaff = (isAdmin === 2) ? 1 : 0
    localStorage.setItem('is_staff', isStaff)
    return this
  }
}

export default new User()
