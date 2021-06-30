<template>
  <el-container>
    <el-main>
      <el-form ref="form" :model="form" :rules="formRules" label-width="100px" style="width: 700px;">
        <h3>数据库配置</h3>
        <el-form-item label="数据库选择" prop="db_type">
          <el-select v-model.trim="form.db_type" @change="update_port">
            <el-option
              v-for="item in dbList"
              :key="item.value"
              :label="item.label"
              :value="item.value">
            </el-option>
          </el-select>
        </el-form-item>
        <el-row>
          <el-col :span="12">
            <el-form-item label="主机名" prop="db_host">
              <el-input v-model="form.db_host"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="端口" prop="db_port">
              <el-input v-model.number="form.db_port"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item label="用户名" prop="db_username">
              <el-input v-model="form.db_username"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="密码" prop="db_password">
              <el-input v-model="form.db_password" type="password"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item label="数据库名称" prop="db_name">
              <el-input v-model="form.db_name" placeholder="如果数据库不存在, 需提前创建"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="表前缀" prop="db_table_prefix">
              <el-input v-model="form.db_table_prefix"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <h3>管理员账号配置</h3>
        <el-row>
          <el-col :span="12">
            <el-form-item label="账号" prop="admin_username">
              <el-input v-model="form.admin_username"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="邮箱" prop="admin_email">
              <el-input v-model="form.admin_email"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item label="密码" prop="admin_password">
              <el-input v-model="form.admin_password" type="password"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="确认密码" prop="confirm_admin_password">
              <el-input v-model="form.confirm_admin_password" type="password"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <h3>LDAP配置</h3>
        <el-row>
          <el-col :span="12">
            <el-form-item label="地址" prop="ldap_addr">
              <el-input v-model="form.ldap_addr"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="DN" prop="ldap_dn">
              <el-input v-model="form.ldap_dn"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item label="管理员帐号" prop="ldap_admin">
              <el-input v-model="form.ldap_admin"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="管理员密码" prop="ldap_password">
              <el-input v-model="form.ldap_password" type="password"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item label="DN-管理员" prop="ldap_dn_admin">
              <el-input v-model="form.ldap_dn_admin"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="DN-用户" prop="ldap_dn_user">
              <el-input v-model="form.ldap_dn_user"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="DN-访客" prop="ldap_dn_guest">
              <el-input v-model="form.ldap_dn_guest"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item label="objClsUser" prop="ldap_obj_class_user">
              <el-input v-model="form.ldap_obj_class_user"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="objClsMember" prop="ldap_obj_class_member">
              <el-input v-model="form.ldap_obj_class_member"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item>
          <el-button type="primary" @click="submit()">安装</el-button>
        </el-form-item>
      </el-form>
    </el-main>
  </el-container>
</template>

<script>
import installService from '../../api/install'
export default {
  name: 'index',
  data () {
    return {
      form: {
        db_type: 'mysql',
        db_host: '127.0.0.1',
        db_port: 3306,
        db_username: '',
        db_password: '',
        db_name: '',
        db_table_prefix: '',
        admin_username: '',
        admin_password: '',
        confirm_admin_password: '',
        admin_email: '',
        ldap_addr: '',
        ldap_dn: '',
        ldap_admin: '',
        ldap_password: '',
        ldap_dn_admin: '',
        ldap_dn_user: '',
        ldap_dn_guest: '',
        ldap_obj_class_user: '',
        ldap_obj_class_member: ''
      },
      formRules: {
        db_type: [
          {required: true, message: '请选择数据库', trigger: 'blur'}
        ],
        db_host: [
          {required: true, message: '请输入数据库主机名', trigger: 'blur'}
        ],
        db_port: [
          {type: 'number', required: true, message: '请输入数据库端口', trigger: 'blur'}
        ],
        db_username: [
          {required: true, message: '请输入数据库用户名', trigger: 'blur'}
        ],
        db_password: [
          {required: true, message: '请输入数据库密码', trigger: 'blur'}
        ],
        db_name: [
          {required: true, message: '请输入数据库名称', trigger: 'blur'}
        ],
        admin_username: [
          {required: true, message: '请输入管理员账号', trigger: 'blur'}
        ],
        admin_email: [
          {type: 'email', required: true, message: '请输入管理员邮箱', trigger: 'blur'}
        ],
        ldap_addr: [
          {required: true, message: '请输入LDAP地址, 例如: ldap://10.0.0.1:389', trigger: 'blur'}
        ],
        ldap_dn: [
          {required: true, message: '请输入LDAP DN, 例如: ou=Users,dc=domain,dc=com', trigger: 'blur'}
        ],
        ldap_admin: [
          {required: true, message: '请输入LDAP管理员帐号, 例如: cn=gocron,ou=Service,dc=domain,dc=com', trigger: 'blur'}
        ],
        ldap_password: [
          {required: true, message: '请输入LDAP管理员密码', trigger: 'blur'}
        ],
        ldap_dn_admin: [
          {required: true, message: '请输入LDAP管理员DN, 例如: cn=gocron-admins,ou=GoCron-Prod,ou=Groups,dc=domain,dc=com', trigger: 'blur'}
        ],
        ldap_dn_user: [
          {required: true, message: '请输入LDAP用户DN, 例如: cn=gocron-users,ou=GoCron-Prod,ou=Groups,dc=domain,dc=com', trigger: 'blur'}
        ],
        ldap_dn_guest: [
          {required: true, message: '请输入LDAP访客DN, 例如: cn=gocron-guests,ou=GoCron-Prod,ou=Groups,dc=domain,dc=com', trigger: 'blur'}
        ],
        ldap_obj_class_user: [
          {required: true, message: '请输入LDAP用户ObjClass, 例如: inetOrgPerson', trigger: 'blur'}
        ],
        ldap_obj_class_member: [
          {required: true, message: '请输入LDAP组ObjClass, 例如: groupOfMembers', trigger: 'blur'}
        ],
        admin_password: [
          {required: true, message: '请输入管理员密码', trigger: 'blur'},
          {min: 6, message: '长度至少6个字符', trigger: 'blur'}
        ],
        confirm_admin_password: [
          {required: true, message: '请再次输入管理员密码', trigger: 'blur'},
          {min: 6, message: '长度至少6个字符', trigger: 'blur'}
        ]
      },
      dbList: [
        {
          value: 'mysql',
          label: 'MySQL'
        },
        {
          value: 'postgres',
          label: 'PostgreSql'
        }
      ],
      default_ports: {
        'mysql': 3306,
        'postgres': 5432
      }
    }
  },
  methods: {
    update_port (dbType) {
      console.log(dbType)
      console.log(this.default_ports[dbType])
      this.form['db_port'] = this.default_ports[dbType]
      console.log(this.form['db_port'])
    },
    submit () {
      this.$refs['form'].validate((valid) => {
        if (!valid) {
          return false
        }
        this.save()
      })
    },
    save () {
      installService.store(this.form, () => {
        this.$router.push('/')
      })
    }
  }
}
</script>
