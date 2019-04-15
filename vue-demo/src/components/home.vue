<template>
  <div id="app">
    <el-container ref="homePage"> 
      <el-header class="blueheader">
        <h2 class="headlogo">EngineerCMS<img src="/static/logo.png"></h2>
        <el-col :span="3" class="userinfo">
          <el-dropdown @command="handleCommand">
            <i class="el-icon-setting" style="margin-right: 15px"></i>
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item command="a">登录</el-dropdown-item>
              <!-- <div @click="state = !state">{{state ? '启用' : '停用'}}</div>
              <a class="button1" @click ="changeStatus"> {{btnStatus?'启用':'停用'}}</a>
              data(){
                return {
                  btnStatus: true,
                }
              },
              methods:{
                changeStatus(){
                  this.btnStatus = this.btnStatus ? false : true ;
                }
              } -->
              <el-dropdown-item command="b">关于</el-dropdown-item>
              <el-dropdown-item command="c">留言</el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>
          <span>3xxx</span>
        </el-col>
      </el-header>

      <el-container>
        <el-aside style="width: auto;" class="sidebar-container">
          <el-radio-group v-model="isCollapse" style="margin-bottom: 10px;margin-top: 10px;" size="mini">
            <el-radio-button :label="false">展开</el-radio-button>
            <el-radio-button :label="true">收起</el-radio-button>
          </el-radio-group>
            <el-menu default-active="doctype" @select="handleSelect" :collapse="isCollapse">
            <el-submenu index="doctype">
              <template slot="title">
                <i class="el-icon-ali-set"></i>
                <span slot="title">Doctype-state-action</span>
              </template>
              <el-menu-item index="doctype">Doctype</el-menu-item>
              <el-menu-item index="docstate">Docstate</el-menu-item>
              <el-menu-item index="docaction">Docaction</el-menu-item>
            </el-submenu>
            <el-submenu index="transition">
              <template slot="title">
                <i class="el-icon-ali-all"></i>
                <span slot="title">Transison</span>
              </template>
              <el-menu-item index="transition">Doctype_transition</el-menu-item>
              <el-menu-item index="workflow">Workflow</el-menu-item>
              <el-menu-item index="node">workflow_node</el-menu-item>
            </el-submenu>
            <el-submenu index="user">
              <template slot="title">
                <i class="el-icon-ali-account"></i>
                <span slot="title">Accesscontext</span>
              </template>
              <el-menu-item index="accesscontext">accesscontext</el-menu-item>
              <el-menu-item index="user">user</el-menu-item>
              <el-menu-item index="group">group</el-menu-item>
              <el-menu-item index="userGroup">user_group</el-menu-item>
              <el-menu-item index="role">role</el-menu-item>
              <el-menu-item index="permission">permission</el-menu-item>
              <el-menu-item index="groupRole">group_role</el-menu-item>
            </el-submenu>
            <el-submenu index="documentlist">
              <template slot="title">
                <i class="el-icon-ali-viewlist"></i>
                <span slot="title">Douments</span>
              </template>
              <el-menu-item index="documentlist">documentlist</el-menu-item>
            </el-submenu>
            <el-submenu index="mailbox">
              <template slot="title">
                <i class="el-icon-message"></i>
                <span slot="title">mailbox</span>
              </template>
              <el-menu-item index="usermailbox">usermailbox</el-menu-item>
              <el-menu-item index="groupmailbox">groupmailbox</el-menu-item>
            </el-submenu>
          </el-menu>
        </el-aside>

        <el-main class="main-container">
          <el-col :span="24" class="breadcrumb-container">
            <strong class="title">{{$route.name}}</strong>
            <el-breadcrumb separator="/" class="breadcrumb-inner">
              <el-breadcrumb-item v-for="item in $route.matched" :key="item.path">
                {{ item.name }}
              </el-breadcrumb-item>
            </el-breadcrumb>
          </el-col>
          <el-col :span="24">
            <div style="margin-top:10px">
              <router-view />
            </div>
          </el-col>
        </el-main>
      </el-container>  
    </el-container>
  </div>
</template>

<style>
/*  .el-header {
    background-color: #B3C0D1;
    color: #333;
    line-height: 60px;
  }
  
  .el-aside {
    color: #333;
  }*/
</style>

<script>
  export default {
    name: 'home',
    data(){
      return {
        isCollapse: true,
        searchCriteria: '',
        clientHeight:'',
      }
    },
    mounted(){
      // 获取浏览器可视区域高度
      this.clientHeight =   `${document.documentElement.clientHeight}`          //document.body.clientWidth;
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
      handleIconClick(ev) {
        console.log(ev);
      },
      handleSelect(key, keyPath){
        switch(key){
          case 'readme':
            this.$router.push('/flow');
            break;
          case 'doctype':
            this.$router.push('/flow/doctype');
            break;
          case 'docstate':
            this.$router.push('/flow/docstate')
            break;
          case 'docaction':
            this.$router.push('/flow/docaction')
            break;
          case 'transition':
            this.$router.push('/flow/transition')
            break;
          case 'workflow':
            this.$router.push('/flow/workflow')
            break;
          case 'node':
            this.$router.push('/flow/node')
            break;
          case 'accesscontext':
            this.$router.push('/flow/accesscontext')
            break;
          case 'user':
            this.$router.push('/flow/user')
            break;
          case 'group':
            this.$router.push('/flow/group')
            break;
          case 'userGroup':
            this.$router.push('/flow/userGroup')
            break;            
          case 'role':
            this.$router.push('/flow/role')
            break;
          case 'permission':
            this.$router.push('/flow/permission')
            break;
          case 'groupRole':
            this.$router.push('/flow/groupRole')
            break;
          case 'documents':
            this.$router.push('/flow/documents')
            break;
          case 'event':
            this.$router.push('/flow/event')
            break;
          case 'documentlist':
            this.$router.push('/flow/documentlist')
            break;
          case 'documentdetail':
            this.$router.push('/flow/documentdetail')
            break;
          case 'usermailbox':
            this.$router.push('/flow/usermailbox')
            break;
          case 'groupmailbox':
            this.$router.push('/flow/groupmailbox')
        }
      },
      handleCommand(command) {
        // this.$message('click on item ' + command);
        switch(command){
          case 'a':
            this.$router.push('/flow/login');
            break;
          case 'b':
            this.$router.push('/flow/readme');
            break;
          case 'c':
            this.$router.push('/flow/chatme')
            break;
          // case 'docaction':
          //   this.$router.push('/docaction')
          //   break;
        }
      },
      changeFixed(clientHeight){ //动态修改样式
        // console.log(clientHeight);
        // console.log(this.$refs.homePage.$el.style.height);
        this.$refs.homePage.$el.style.height = clientHeight-20+'px';
      },
    },
  }

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
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
    position: absolute;
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
</style>