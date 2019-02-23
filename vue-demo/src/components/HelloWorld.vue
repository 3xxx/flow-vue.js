<template>
  <div>
  <el-menu default-active="activeIndex" class="el-menu-vertical-demo" @open="handleOpen" @close="handleClose" :collapse="isCollapse" router>
  <el-menu-item index="index">
    <i class="el-icon-menu"></i>
    <span slot="title">首页</span>
  </el-menu-item>
  <el-submenu index="1">
    <template slot="title">
      <i class="el-icon-location"></i>
      <span slot="title" @click="jump('project')">项目</span>
    </template>
    <!-- <el-menu-item-group> -->
      <el-menu-item index="bar" @click="jump('design')">设计</el-menu-item>
      <el-menu-item index="1-2">建设</el-menu-item>
      <el-menu-item index="1-3">监理</el-menu-item>
      <!-- <el-menu-item index="1-4">施工</el-menu-item> -->
    <!-- </el-menu-item-group> -->
  </el-submenu>
  <el-submenu index="2">
    <template slot="title">
      <i class="el-icon-printer"></i>
      <span slot="title">标段</span>
    </template>
    <!-- <el-menu-item-group title="A标"> -->
      <el-submenu index="2-1">
        <span slot="title">A标</span>
        <el-menu-item index="2-1-1">A1标</el-menu-item>
        <el-menu-item index="2-1-2">A2标</el-menu-item>
        <el-menu-item index="2-1-3">A3标</el-menu-item>
        <el-menu-item index="2-1-4">A4标</el-menu-item>
        <el-menu-item index="2-1-5">A5标</el-menu-item>
        <el-menu-item index="2-1-6">A6标</el-menu-item>
      </el-submenu>
    <!-- </el-menu-item-group> -->
    <!-- <el-menu-item-group title="B标"> -->
      <el-submenu index="2-2">
        <span slot="title">B标</span>
        <el-menu-item index="2-2-1">B1标</el-menu-item>
        <el-menu-item index="2-2-2">B2标</el-menu-item>
        <el-menu-item index="2-2-3">B3标</el-menu-item>
        <el-menu-item index="2-2-4">B4标</el-menu-item>
      </el-submenu>
    <!-- </el-menu-item-group> -->
    <!-- <el-menu-item-group title="B标"> -->
      <el-submenu index="2-3">
        <span slot="title">C标</span>
        <el-menu-item index="2-3-1">C1标</el-menu-item>
        <el-menu-item index="2-3-2">C2标</el-menu-item>
        <el-menu-item index="2-3-3">C3标</el-menu-item>
      </el-submenu>
    <!-- </el-menu-item-group> -->
    <!-- <el-menu-item-group title="B标"> -->
      <el-submenu index="2-4">
        <span slot="title">D标</span>
        <el-menu-item index="2-4-1">D1标</el-menu-item>
        <el-menu-item index="2-4-2">D2标</el-menu-item>
      </el-submenu>
    <!-- </el-menu-item-group>     -->
  </el-submenu>

  <el-menu-item index="3" @click="jump('onlyoffice')">
    <i class="el-icon-upload"></i>
    <span slot="title">onlyoffice</span>
  </el-menu-item>
  <el-menu-item index="4" disabled>
    <i class="el-icon-document"></i>
    <span slot="title">视频</span>
  </el-menu-item>
  <el-menu-item index="5">
    <i class="el-icon-setting"></i>
    <span slot="title">日记</span>
  </el-menu-item>
</el-menu>

  <el-header class="grid-content bg-purple-light">
    <el-row>
      <el-col :span="12" style="text-align: left;" class="cusheader">
        <!-- <el-switch v-model="isCollapse" active-color="#13ce66" inactive-color="#ff4949">
        </el-switch> -->
        <i @click="changeCollapse(isCollapse)" :class="{'el-icon-d-arrow-right':isCollapse,'el-icon-d-arrow-left':!isCollapse}"></i>
        <!-- <comment v-model="isCollapse"></comment> -->
      </el-col>
      <el-col :span="12" style="text-align: right;" class="cusheader">
        <el-dropdown>
          <i class="el-icon-setting" style="margin-right: 5px"></i>
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item @click.native="dialogFormVisible = true">登陆</el-dropdown-item>
            <el-dropdown-item>大事记</el-dropdown-item>
            <el-dropdown-item>查阅</el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
        <span>设代处</span>
      </el-col>
    </el-row>
  </el-header>

</div>

<!--   <div class="hello">
    <h1>{{ msg }}</h1>
    <h1>
      <router-link to="day01">跳转至HelloVue</router-link>
    </h1>
    <h2>Essential Links</h2>
    <ul>
      <li>
        <a
          href="https://vuejs.org"
          target="_blank"
        >
          Core Docs
        </a>
      </li>
      <li>
        <a
          href="https://forum.vuejs.org"
          target="_blank"
        >
          Forum
        </a>
      </li>
      <li>
        <a
          href="https://chat.vuejs.org"
          target="_blank"
        >
          Community Chat
        </a>
      </li>
      <li>
        <a
          href="https://twitter.com/vuejs"
          target="_blank"
        >
          Twitter
        </a>
      </li>
      <br>
      <li>
        <a
          href="http://vuejs-templates.github.io/webpack/"
          target="_blank"
        >
          Docs for This Template
        </a>
      </li>
    </ul>
    <h2>Ecosystem</h2>
    <ul>
      <li>
        <a
          href="http://router.vuejs.org/"
          target="_blank"
        >
          vue-router
        </a>
      </li>
      <li>
        <a
          href="http://vuex.vuejs.org/"
          target="_blank"
        >
          vuex
        </a>
      </li>
      <li>
        <a
          href="http://vue-loader.vuejs.org/"
          target="_blank"
        >
          vue-loader
        </a>
      </li>
      <li>
        <a
          href="https://github.com/vuejs/awesome-vue"
          target="_blank"
        >
          awesome-vue
        </a>
      </li>
    </ul>
  </div> -->
</template>

<script>
/* eslint-disable */
export default {
  name: 'HelloWorld',
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
          total: 5,
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

          isCollapse: true,
          bannerHeight:200,

          clientHeight:'',
          // calleft:0
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
          formLabelWidth: '120px',
          tableData: [],
          search: '',
          sortList: [
            {'list': '观影兑换券', },
            {'list': '室内乐兑换券', },
            {'list': '沙龙兑换券', }
          ],
        };
      },

      methods:{
        submitForm(formName) {
          this.$refs[formName].validate((valid) => {
            if (valid) {
              this.axios({
                headers: {
                  'X-Requested-With': 'XMLHttpRequest',
                  'Content-Type': 'application/json; charset=UTF-8',
                  'Access-Control-Allow-Origin': '*'
                },//设置跨域请求头
                method: "POST",//请求方式
                url: "/v1/admin/flowtype",//请求地址
                params:{
                  name:this.ruleForm2.typename,
                },
                data: {
                  name:this.ruleForm2.typename,
                  // "thirdapp_id":1//请求参数
                }
              })
              .then(response => (this.posts = response.data.articles))
              .catch(function (error) {
                console.log(error);
              });
              //提交成功做的动作
              this.$message({
                type: 'success',
                message: '提交成功' 
              });
              //刷新表格
              this.flowtypelist();
              this.dialogFormVisible = false;
            } else {
              console.log('error submit!!');
              return false;
            }
          }); 
        },

        resetForm(formName) {
          this.$refs[formName].resetFields();
        },

        changeWidthL(key, keyPath) {
          console.log(key, keyPath);
        },

        handleOpen(key, keyPath) {
          console.log(key, keyPath);
        },
        handleClose(key, keyPath) {
          console.log(key, keyPath);
        },
        handleEdit(index, row) {
          console.log(index, row);
        },
        handleDelete(index, row) {
          console.log(index, row);
        },
        sendGetByStr(){
          this.axios.get(`api/v1/wx/getlistarticles?page=1`)//1.get通过直接发字符串拼接
            // .then(function (response) {
            //   console.log(response);
            //   console.log(response.data.info);
            // }
            .then(response => (this.info = response.data.info)
            )
            .catch(function (error) {
              console.log(error);
            });
        },
        flowtypelist(){
          this.axios.get(`/v1/admin/flowtypelist`,{//2.get通过params选项
            params:{
                page:1
            }
          })
          .then(response => (this.tableData = response.data))
          .catch(function (error) {
            console.log(error);
          });
        },
        //html剔除富文本标签，留下纯文本
        getSimpleText(html){
          var re1 = new RegExp("<.+?>","g");//匹配html标签的正则表达式，"g"是搜索匹配多个符合的内容
          var msg = html.replace(re1,'');//执行替换成空字符
          return '<br>'+msg.substring(0,20);
        },
        even: function (numbers) {
          // console.log(numbers);
          // console.log(numbers % 2 === 0);
          // return numbers.filter(function (number) {
            return numbers % 2 === 0
          // })
        },
        even1: function (numbers) {
          // console.log(numbers);
          // console.log(numbers % 2 === 0);
          // return numbers.filter(function (number) {
            return numbers === 0
          // })
        },
        changeCollapse:function(isCollapse){
          if (isCollapse==true) {
            this.isCollapse=false
          }else{
            this.isCollapse=true
          }
        },
        skip(a){
          this.$router.push(a)
        },
        jump(select){
          console.log(select);
          let routeUrl = this.$router.resolve({
            path: "/onlyoffice",
            // query: {id:96}
          });
          switch (select) {
            case ("onlyoffice"):
              routeUrl = this.$router.resolve({
                path: "/onlyoffice",
                // query: {id:96}
              });
              window.open(routeUrl .href);
              break;
            case ("project"):
              routeUrl = this.$router.resolve({
                path: "/project",
                // query: {id:96}
              });
              window.open(routeUrl .href);
              break;
            case ("design"):
              routeUrl = this.$router.resolve({
                path: "/design",
                // query: {id:96}
              });
              window.open(routeUrl .href);
              break;
            default:
              routeUrl = this.$router.resolve({
                path: "/index",
                // query: {id:96}
              });
              window.open(routeUrl .href);
              break;
          }
          //this.$router.push({path: '/cart?goodsId=12'})
          //this.$router.go(-2)
          //后退两步
          // let routeUrl = this.$router.resolve({
          //   path: "/onlyoffice",
          //   query: {id:96}
          // });
          // window.open(routeUrl .href, '_blank');
          // window.open(routeUrl .href);
        },
        toNext: function(index) {
          sessionStorage.ticketName =this.sortList[index].list;
          this.$router.push('/mine/tiketOrder');
        }
      }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
