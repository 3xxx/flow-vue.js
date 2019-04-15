<template>
  <div id="app">
<!--  <el-container>
  <el-header>Header</el-header>
  <el-container>
    <el-aside width="200px">
      Aside
    </el-aside>
    <el-main>
      Main
      <div>*********************************</div>
      <div>****************</div></div>
      <div>**************</div>
      <div>*************</div>
      <div>********</div>
      <div>**********</div>
      <div>******</div>
      <div>*******************</div>
    </el-main>
  </el-container>
</el-container> -->

<!--   <el-container>
    <el-header class="blueheader">
      <h2 class="headlogo">EngineerCMS<img src="/static/logo.png"></h2>
      <el-col :span="3" class="userinfo">
        <el-dropdown>
          <i class="el-icon-setting" style="margin-right: 15px"></i>
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item>查看</el-dropdown-item>
            <el-dropdown-item>新增</el-dropdown-item>
            <el-dropdown-item>删除</el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
        <span>3xxx</span>
      </el-col>
    </el-header>
    <el-container>
      <el-aside style="width: auto;">
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
      <el-main class="home_main">
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
            <router-view></router-view>
          </div>
        </el-col>
      </el-main>
    </el-container>
  </el-container> -->
    <router-view />
  </div>
</template>
        <!--  style="flex: 0 0 230px;width: 230px;background:#eef1f6"  class="el-menu-vertical-demo"-->
          <!-- <el-submenu index="documents">
            <template slot="title"><i class="el-icon-ali-favorite"></i>Douments</template>
            <el-menu-item index="documents">documents</el-menu-item>
            <el-menu-item index="event">event</el-menu-item>
          </el-submenu> -->
<script>
/* eslint-disable */
  export default {
    name: 'App',
    data(){
      return {
        isCollapse: true,
        searchCriteria: '',
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
      // handleSelect(key, keyPath){
      //   switch(key){
      //     case 'readme':
      //       this.$router.push('/');
      //       break;
      //     case 'doctype':
      //       this.$router.push('/doctype');
      //       break;
      //     case 'docstate':
      //       this.$router.push('/docstate')
      //       break;
      //     case 'docaction':
      //       this.$router.push('/docaction')
      //       break;
      //     case 'transition':
      //       this.$router.push('/transition')
      //       break;
      //     case 'workflow':
      //       this.$router.push('/workflow')
      //       break;
      //     case 'node':
      //       this.$router.push('/node')
      //       break;
      //     case 'accesscontext':
      //       this.$router.push('/accesscontext')
      //       break;
      //     case 'user':
      //       this.$router.push('/user')
      //       break;
      //     case 'group':
      //       this.$router.push('/group')
      //       break;
      //     case 'userGroup':
      //       this.$router.push('/userGroup')
      //       break;            
      //     case 'role':
      //       this.$router.push('/role')
      //       break;
      //     case 'permission':
      //       this.$router.push('/permission')
      //       break;
      //     case 'groupRole':
      //       this.$router.push('/groupRole')
      //       break;
      //     case 'documents':
      //       this.$router.push('/documents')
      //       break;
      //     case 'event':
      //       this.$router.push('/event')
      //       break;
      //     case 'documentlist':
      //       this.$router.push('/documentlist')
      //       break;
      //     // case 'dumentdetail':
      //     //   this.$router.push('/documentlist/documentdetail')
      //     //   break;
      //     case 'documentdetail':
      //       this.$router.push('/documentdetail')
      //       break;
      //     case 'usermailbox':
      //       this.$router.push('/usermailbox')
      //       break;
      //     case 'groupmailbox':
      //       this.$router.push('/groupmailbox')
      //     // default:
      //     //   this.$router.push('/doctype');
      //     //   break;
      //   }
      // },
    },
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  /*#app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    margin-top: 60px;
  }*/

  .el-menu-vertical-demo:not(.el-menu--collapse) {
    width: 200px;
    min-height: 400px;
  }
  
  .con_section{
    position: absolute;
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
