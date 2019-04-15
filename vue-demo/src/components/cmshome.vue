<template>
  <div id="app">
    <el-container ref="homePage"> 
      <!-- <el-header class="blueheader"> -->
        <!-- <h2 class="headlogo">EngineerCMS<img src="/static/logo.png"></h2> -->
        <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect" background-color="#409EFF" text-color="#fff" active-text-color="#ffd04b">
          <el-menu-item index="cmsindex"><i class="el-icon-ali-shouye" style="margin-right: 15px;color: #fff;font-size: 35px;"></i>Home</el-menu-item>
          <el-menu-item index="projects"><i class="el-icon-ali-fenlei" style="margin-right: 15px;color: #fff;font-size: 35px;"></i><a href="https://zsj.itdos.com/project/25001" target="_blank">Projects</a></el-menu-item>
          <el-menu-item index="onlyoffice"><i class="el-icon-ali-changyonggoupiaorenbianji" style="margin-right: 15px;color: #fff;font-size: 35px;"></i><a href="https://zsj.itdos.com/onlyoffice" target="_blank">Onlyoffice</a></el-menu-item>
          
          <el-submenu index="2">
            <template slot="title"><i class="el-icon-ali-sousuo" style="margin-right: 15px;color: #fff;font-size: 35px;"></i>Standard</template>
            <el-menu-item index="2-1"><a href="https://zsj.itdos.com/standard" target="_blank">查询</a></el-menu-item>
            <el-menu-item index="2-2"><a href="https://zsj.itdos.com/legislation" target="_blank">对标</a></el-menu-item>
          </el-submenu>
          <el-submenu index="3" style="position: absolute;right: 0;">
            <template slot="title"><i class="el-icon-setting" style="margin-right: 15px;color: #fff;font-size: 35px;"></i>工作台</template>
            <el-menu-item index="login" v-if="!logindata.islogin">登录</el-menu-item>
            <el-menu-item index="admin" v-if="logindata.isadmin"><a href="https://zsj.itdos.com/admin" target="_blank">后台</a></el-menu-item>
            <el-menu-item index="3-2" v-if="logindata.islogin">aboutme</el-menu-item>
            <el-menu-item index="3-3">帮助</el-menu-item>
            <el-menu-item index="3-4" v-if="logindata.islogin">重新登录</el-menu-item>
          </el-submenu>
        </el-menu>
        <el-main><!--  class="main-container" -->
          <!-- <el-col> --><!--  :span="24" -->
            <!-- <div style="margin-top:-20px">  -->
              <router-view style="margin-top:-20px"/>
            <!-- </div> -->
          <!-- </el-col> -->
        </el-main>
      <!-- </el-container>   -->
      <el-footer style="text-align: center;background-color: #E9EEF3;color: #B3C0D1;line-height: 60px;">
      Copyright © 2016-2019 EngineerCMS
      网站由
      <a target="_blank" href="https://github.com/3xxx">@3xxx</a>
      建设，并由
      <a target="_blank" href="http://golang.org">golang</a>
      和
      <a target="_blank" href="http://beego.me">beego</a>
      提供动力。
      <a target="_blank" href="https://github.com/3xxx"><strong><i class="icon-github-sign"></i> Github</strong></a>
      </el-footer>
    </el-container>

<el-dialog title="系统登录" :visible.sync="dialogFormVisible" center>
  <!-- 插入测试 -->
  <el-form :model="ruleForm2" status-icon :rules="rules2" ref="ruleForm2" label-width="100px" class="demo-ruleForm">
    <el-form-item label="账号" prop="uname">
      <el-input v-model.number="ruleForm2.uname" auto-complete="off" placeholder="账号"></el-input>
    </el-form-item>
    <el-form-item label="密码" prop="pass">
      <el-input type="password" v-model="ruleForm2.pass" auto-complete="off" placeholder="密码"></el-input>
    </el-form-item>
    <el-checkbox v-model="checked" checked class="remember">记住密码</el-checkbox>
    <el-form-item label="记住密码" prop="delivery">
      <el-switch v-model="ruleForm2.delivery"></el-switch>
    </el-form-item> 
    <span><a>忘记密码？</a></span>
  </el-form>
   <!-- 插入测试 -->
  <div slot="footer" class="dialog-footer">
    <el-button @click="dialogFormVisible = false; resetForm('ruleForm2')">取 消</el-button>
    <el-button type="primary" @click="submitForm('ruleForm2')">登 录</el-button>
  </div>
</el-dialog>

  </div>
 
</template>

<script>
  const axios = require('axios');
  export default {
    name: 'cmshome',
    data() {
        var checkName = (rule, value, callback) => {
          if (value === '') {
            callback(new Error('请输入账号'));
          } else {
            callback();
          }
        };
        var checkNum = (rule, value, callback) => {
          if (!value) {
            return callback(new Error('账号不能为空'));
          }
          setTimeout(() => {
            if (!Number.isInteger(value)) {
              callback(new Error('请输入数字值'));
            } else {
              var myreg=/^[1][3,4,5,7,8][0-9]{9}$/;
              if (!myreg.test(value) ) {
                callback(new Error('请输入正确的手机号码'));
              } else {
                callback();
              }
            }
          }, 1000);
        };
        var validatePass = (rule, value, callback) => {
          if (value === '') {
            callback(new Error('请输入密码'));
          } else {
            callback();
          }
        };
        return {
          isCollapse: true,
        searchCriteria: '',
        clientHeight:'',
        activeIndex: 'cmsindex',

          total: 50,
          currentPage: 1,
          pageSize: 10,
          loginPower:false,
          checked: true,
          /*插入form方法*/
          /*设定规则指向*/
          ruleForm2: {
            name:'',
            pass: '',
            num: '',
            delivery: false,
          },
          rules2: {
            typename: [
              { validator: checkName, trigger: 'blur' }
            ],
            pass: [
              { validator: validatePass, trigger: 'blur' }
            ],
            num: [
              { validator: checkNum, trigger: 'blur' }
            ]
          },
          /*插入form方法*/
          dialogTableVisible: false,
          dialogFormVisible: false,
          form: {
            name: '',
            type: [],
            resource: '',
            desc: ''
          },
          formLabelWidth: '120px',
          posts:[],
          numbers:0,

          // dialogFormVisible: false,
          form: {
            name: '',
            region: '',
            date1: '',
            date2: '',
            delivery: false,
            type: [],
            resource: '',
            desc: ''
          },
          documentsdata:[
            {
              Id:1,
              AcId:1,
              DocstateId:1,
              GroupId:1,
              Title:'documents title',
              Data:'数据',
              Path:'path',
              Ctime:'2019-02-09',
            }
          ],
          accesscontextdata: [
            {
              Id:1,
              Name:'我是accesscontext',
            }
          ],
          docstatedata:[
            {ID:1,Name:"图纸"},
            {ID:2,Name:"合同"}
          ],
          doctypedata:[
            {
              page: 1,
              total: 3,
              doctypes:
              {
                ID: 1,
                Name: "draw"
              }
            }
          ],
          groupdata:[
            {ID:1,Name:"秦晓川",GroupType:"S"},
            {ID:2,Name:"校核组",GroupType:"G"}
          ],
          value1:'',
          value2:'',
          value3:'',
          // islogin:false,
          logindata:{
            islogin:false,
            isadmin:false
          }
        };
      },
    mounted(){
      //检查是否登录
      this.isLogin();
      // 获取浏览器可视区域高度
      this.clientHeight = `${document.documentElement.clientHeight}`;//document.body.clientWidth;
      //console.log(self.clientHeight);
      window.onresize = function temp() {
        this.clientHeight = `${document.documentElement.clientHeight}`;
      };
    },
    watch: {
      // 如果 `clientHeight` 发生改变，这个函数就会运行
      clientHeight: function () {
        this.changeFixed(this.clientHeight)
      }
    },
    computed: {
      activeName: {
        get () {
          return this.$route.name
        },
        set (value) {
          this.$router.replace({name: value})
        }
      }
    },
    methods:{
      submitForm(formName) {
          this.$refs[formName].validate((valid) => {
            if (valid) {
              axios({
                method: "POST",//请求方式
                url: "https://zsj.itdos.com/v1/wx/loginpost",//请求地址
                params:{
                  uname:this.ruleForm2.uname,
                  pwd:this.ruleForm2.pass,
                },
              }).then((response) => {
                if (response.data.islogin === 1) {
                  this.$message.error('用户名或密码错误，请重新登录！')
                } else {
                  // console.log(response.data);
                  //关闭对话框
                  this.dialogFormVisible = false;
                  this.logindata.islogin=true;
                }
              })
              // .then(response => (this.posts = response.data.articles))
              // .then(function (response) {
              //   console.log(response);
              //   if (response.data.islogin==0) {
              //     //提交成功做的动作
              //     this.$message({
              //       type: 'success',
              //       message: '登录成功' 
              //     });
              //     //关闭对话框
              //     this.dialogFormVisible = false;
              //   } else {
              //     //写入失败！
              //     this.$message.error('登录失败！');
              //   }
              // })
              .catch(function (error) {
                console.log(error);
              });

            } else {
              console.log('error submit!!');
              return false;
            }
          }); 
      },
      resetForm(formName) {
        this.$refs[formName].resetFields();
      },
      handleIconClick(ev) {
        console.log(ev);
      },
      handleSelect(key, keyPath){
        // console.log(key);
        switch(key){
          case 'cmsindex':
            this.$router.push('/');
            break;
          case 'projects':
            this.$router.push('/projects');
            break;
          case 'project':
            this.$router.push('/project')
            break;
          case 'login':
            this.dialogFormVisible = true;
            break;
        }
      },
      handleCommand(command) {
        // this.$message('click on item ' + command);
        switch(command){
          case 'a':
            this.$router.push('/login');
            break;
          case 'b':
            this.$router.push('/readme');
            break;
          case 'c':
            this.$router.push('/chat')
            break;
          // case 'docaction':
          //   this.$router.push('/docaction')
          //   break;
        }
      },
      changeFixed(clientHeight){ //动态修改样式
        // console.log(clientHeight);
        // console.log(this.$refs.homePage);
        // console.log(this.$refs.homePage.$el.style.height);这里虽然显示不出来
        this.$refs.homePage.$el.style.height = clientHeight-20+'px';
      },
        isLogin(){
          axios({
            method: 'get',
            url: 'https://zsj.itdos.com/v1/wx/islogin',
          })
          .then(response => (this.logindata = response.data))
          .catch(function (error) {
            console.log(error);
          });
        },
    },
  }
</script>

<style scoped>
  .el-menu-vertical-demo:not(.el-menu--collapse) {
    width: 200px;
    min-height: 400px;
  }

  .con_section{
    top: 0px;
    bottom: 0px;
    left:0px;
    width:100%;
  }
  .blueheader {
    height: 60px;
    line-height: 60px;
    background: #409EFF;
    color: #fff;
  }
  .el-menu-item.is-active {
    color: #409EFF;
  }
  .headlogo{
    float: left;
    height: 60px;
    margin: 0 20px;
    width: 300px;

    text-align: center;

    line-height:60px;
  }
  .headlogo img{
    height: 40px;
    display: inline-block;
    vertical-align: middle;
  }
  ul.el-menu {
    background: #e4e8f1;
  }
  .userinfo{
    
    right: 0;
  }
  .el-submenu__title{
    background:#eef1f6;
  }
  .el_main{
    padding:0px;
  }
  .home_main{
    padding:10px;
  }
  .breadcrumb-container .title {
    width: 200px;
    float: left;
    color: #475669;
    font-size: 13px;
  }
  .breadcrumb-inner {
    float: right;
    font-size: 13px;
  }
  .el-breadcrumb__inner, .el-breadcrumb__inner a {
    font-weight: 400;
  }

  .el-footer a { 
    text-decoration:none
  } 

</style>