<!-- 00 用户信箱侧栏 -->
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
          <el-menu default-active="usertobeprocessed" @select="handleSelect" :collapse="isCollapse">
            <el-submenu index="usertobeprocessed">
              <template slot="title">
                <i class="el-icon-ali-viewlist"></i>
                <span slot="title">我的事务</span>
              </template>
              <el-menu-item index="usertobeprocessed"><i class="el-icon-news"></i>待处理的事务</el-menu-item>
              <el-menu-item index="userhaveprocessed"><i class="el-icon-news"></i>已处理的事务</el-menu-item>
              <el-menu-item index="usertobesubmit"><i class="el-icon-news"></i>我发起的事务</el-menu-item>
            </el-submenu>

            <el-submenu index="grouptobeprocessed">
              <template slot="title">
                <i class="el-icon-connection"></i>
                <span slot="title">用户组</span>
              </template>
              <!-- <el-menu-item index="usermailbox2">usermailbox</el-menu-item> -->
              <!-- <el-menu-item index="groupmailbox2">groupmailbox</el-menu-item> -->
              <el-menu-item index="grouptobeprocessed"><i class="el-icon-news"></i>待处理</el-menu-item>
              <el-menu-item index="grouphaveprocessed"><i class="el-icon-news"></i>已处理</el-menu-item>
              <el-menu-item index="groupprocessed"><i class="el-icon-news"></i>已结束</el-menu-item>
            </el-submenu>
            <el-submenu index="userunreadnews">
              <template slot="title">
                <i class="el-icon-message"></i>
                <span slot="title">消  息</span>
              </template>
              <el-menu-item index="userunreadnews"><i class="el-icon-news"></i>未读</el-menu-item>
              <el-menu-item index="userreadnews"><i class="el-icon-news"></i>已读</el-menu-item>
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
    name: 'home_user',
    data(){
      return {
        isCollapse: false,
        searchCriteria: '',
        clientHeight:'',
      }
    },
    mounted(){
      // 获取浏览器可视区域高度
      this.clientHeight = `${document.documentElement.clientHeight}`          //document.body.clientWidth;
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
            this.$router.push('/flow/readme');
            break;
          // case 'usermailbox2':
          //   this.$router.push('/flow/usermailbox2')
          //   break;
          case 'usertobeprocessed':
            this.$router.push('/flow/usertobeprocessed')
            break;
          case 'userhaveprocessed':
            this.$router.push('/flow/userhaveprocessed')
            break;
          case 'usertobesubmit':
            this.$router.push('/flow/usertobesubmit')
            break;
          // case 'groupmailbox2':
          //   this.$router.push('/flow/groupmailbox2')
          case 'grouptobeprocessed':
            this.$router.push('/flow/grouptobeprocessed')
            break;
          case 'grouphaveprocessed':
            this.$router.push('/flow/grouphaveprocessed')
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