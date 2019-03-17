<template>
  <div>
    <h1>1 设计说明</h1>
    <el-collapse v-model="activeNames">
      <el-collapse-item title="什么是文档流程？" name="1">
        <div>一个文档流程的本质就是改变文档状态state）；对应的，我们再设计一个改变状态的动作action；如果再给文档分类doctype，比如合同流程，采购流程的文档；再配合用户-角色-权限（accesscontex）；当一个文档处于一个state的时候，谁有权限permission操作action它？就要预设置好当前状态state下transition:从state1通过动作action1改为state2;那么一个文档如何按照指定的流程来往下走呢，就要先定义一系列节点node，再将这些节点组合到workflow下，见下面的数据结构和说明：</div>
        <pre>
          <code>
            Document Type : docType1 //1.定义文档类型-流程类型
            Document States : [
                docState1, docState2, docState3, docState4 // 2.定义4个状态
            ]
            Document Actions : [
              docAction12, docAction23, docAction34 // 3.定义修改状态的动作，实际应用中，除了向前传递，还要定义向后回退的动作。for the above document states
              docAction21, docAction31, docAction41  //每个都直接回退到1
            ]
            Document Type State Transitions : [
              docState1 --docAction12--&gt; docState2, //前进流程
              docState2 --docAction21--&gt; docState1, //回退流程
              docState2 --docAction23--&gt; docState3,
              docState3 --docAction31--&gt; docState1,
              docState3 --docAction34--&gt; docState4,
              docState4 --docAction41--&gt; docState1,
            ]

            Access Contexts : [
              accCtx1, accCtx2 // 主要是为了用户-角色-权限用的
            ]
            
            Workflow : {
              Name : wFlow1,  //一个doctype只能对应一个workflow
              Initial State : docState1  //一个流程的初始状态
            }
            Nodes : [
              node1: {
                Document Type : docType1,
                Workflow : wFlow1,
                Node Type : NodeTypeBegin, // 开始节点 note this
                From State : docState1,
                Access Context : accCtx1,
              },
              node2: {
                Document Type : docType1,
                Workflow : wFlow1,
                Node Type : NodeTypeLinear, // 中间节点可以是线性的，也可以是平行的 note this
                From State : docState2,
                Access Context : accCtx2, // a different context
              },
              node3: {
                Document Type : docType1,
                Workflow : wFlow1,
                Node Type : NodeTypeEnd, // note this
                From State : docState3,
                Access Context : accCtx1,
              },
            ]
          </code>
        </pre>
        <div>以上都是系统管理员设置的部分，当一个用户使用这个系统的时候，他在某个板块（doctype分类）新建（上传）一个文档document，系统自动根据这个分类，设置了文档的初始state，在这个state下，一般只有一个动作可选，就是传递给下一个状态的操作，还要选接受的人或组（也可以系统固定），还可以写上注释、说明，最后点击按钮提交apply；系统要做的，记录下来谁在什么时间操作了，把信息message发送给对应的用户的信箱mailbox，当对应的用户登录系统的时候，可以查看message。</div>
        <div>当用户点击信箱里某个信息，转到对应的文档详情detail上，能看到这个文档以往（历史）操作记录时间线，可以对这个文档进行操作：选择动作action，选择接受的组，提交apply。</div>
      </el-collapse-item>
    </el-collapse>
    <h1>2 设置部分</h1>
    <el-collapse v-model="activeNames">
      <el-collapse-item title="1.1 设置doctype docstate和docaction" name="2">
        <div>1.1.1 文档类型doctype，只是给定一个名字而已，每个doctype对应一个workflow流程，如果文档某些需要签署到总经理，有些则只要签署到部门经理，咋办？</div>
        <div>1.1.2 文档状态，即每一级给定一个状态，也只是定一个名字而已；</div>
        <div>1.1.3 用户动作，从一个状态到另一个状态对应的动作，只是一个名字而已；</div>
      </el-collapse-item>
      <el-collapse-item title="1.2 设置transition node 和workflow" name="3">
        <div>1.2.1 transition转变，就是定义一个状态到另一个状态以及动作；</div>
        <div>1.2.2 workflow流程，一种文档类型完整流程，将节点纳入到这个流程中，就完成了一个文档流程的设置；</div>
        <div>1.2.3 node节点，给流程定义完整流程中各个节点，包括节点类型：比如是单线过来的还是这个节点有多线过来，有开始节点，和结束节点……；</div>
      </el-collapse-item>
      <el-collapse-item title="1.3 设置accesscontex" name="4">
        <div>1.3.1 accesscontex访问环境，就是给某个流程定义一下用户-角色-权限；</div>
        <div>1.3.2 user用户，每个用户会自动生成一个单用户组，因为flow都是针对group的；</div>
        <div>1.3.3 group用户组；</div>
        <div>1.3.3 user_group将用户加入组；</div>
        <div>1.3.3 role定义角色；</div>
        <div>1.3.3 permission定义角色权限，即action；</div>
        <div>1.3.3 group_role将group加入角色；</div>
      </el-collapse-item>
    </el-collapse>
    <h1>3 用户部分</h1>
    <el-collapse v-model="activeNames">
      <el-collapse-item title="3.1 新建文档" name="5">
        <div>；</div>
        <div>。</div>
      </el-collapse-item>
      <el-collapse-item title="3.2 传递文档" name="6">
        <div>选择action决定文档是往前传递还是回退，加入说明</div>
        <div>查阅历史时间线。</div>
      </el-collapse-item>
      <el-collapse-item title="3.3 查阅消息" name="7">
        <div>；</div>
        <div>。</div>
      </el-collapse-item>

    </el-collapse>

  </div>
</template>

<script type="text/javascript">
/* eslint-disable */
  const axios = require('axios');
  export default { // 这里需要将模块引出，可在其他地方使用
    name: 'readme',
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
          activeNames: ['1'],
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
          doctypedata:[
            {ID:1,Name:"图纸"},
            {ID:2,Name:"合同"}
          ],
          search: '',
        };
      },
      mounted:function () {
        this.doctype(this.currentPage);
        // const that = this;
        // 获取浏览器可视区域高度
        // this.clientHeight = `${document.documentElement.clientHeight}`; //document.body.clientWidth;
        // console.log(this.clientHeight);
        // window.onresize = function temp() {
          // this.clientHeight = `${document.documentElement.clientHeight}`;
        // };
      },
      methods:{
        submitForm(formName) {
          this.$refs[formName].validate((valid) => {
            if (valid) {
              axios({
                // headers: {
                //   'X-Requested-With': 'XMLHttpRequest',
                //   'Content-Type': 'application/json; charset=UTF-8',
                //   'Access-Control-Allow-Origin': '*'
                // },//设置跨域请求头
                method: "POST",//请求方式
                url: "/api/flowtype",//请求地址
                params:{
                  name:this.ruleForm2.typename,
                },
                // data: {
                  // name:this.ruleForm2.typename,
                  // "thirdapp_id":1//请求参数
                // }
              })
              // .then(response => (this.posts = response.data.articles))
              .then(function (response) {
                console.log(response);
                if (response=="err") {
                  //提交成功做的动作
                  this.$message({
                    type: 'success',
                    message: '提交成功' 
                  });
                  //刷新表格
                  this.flowtypelist();
                  this.dialogFormVisible = false;                 
                } else {
                  //写入失败！
                  this.$message.error('写入失败！');
                }
              })
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
          console.log(index, row.Name);
        },
        handleDelete(index, row) {
          console.log(index, row);
        },
        deleteRow(index, rows) {//删除改行
          rows.splice(index, 1);
        },
        addRow(doctypedata,event){
          doctypedata.push({ ID:'', Name: ''})
        },
        handleSubmit(index, row) {
          console.log(row);
              axios({
                method: "POST",//请求方式
                url: "/api/flowtype",//请求地址
                params:{
                  name:row.Name,
                },
                // data: {
                //   dtid:row.DoctypeId,
                //   dsid1:row.FromStateId,
                //   daid:row.DocactionId,
                //   dsid2:row.ToStateId
                // }
              })
              // .then(response => (this.posts = response.data.articles))
              .then(function (response) {
                console.log(response);
                if (response=="err") {
                  //提交成功做的动作
                  this.$message({
                    type: 'success',
                    message: '提交成功' 
                  });
                  //刷新表格
                  this.doctype(currentPage);
                  this.dialogFormVisible = false;                 
                } else {
                  //写入失败！
                  this.$message.error('写入失败！');
                }
              })
              .catch(function (error) {
                console.log(error);
              });
        },
        sendGetByStr(){
          axios.get(`api/v1/wx/getlistarticles?page=1`)//1.get通过直接发字符串拼接
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
        doctype(currentPage){
          axios({
            method: 'get',
            url: '/api/flowtypelist',//2.get通过params选项
            params:{
              page:currentPage,
              limit:this.pageSize
            }
          })
          .then(response => (this.doctypedata = response.data))
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
        },
        handleSizeChange: function (size) {
          this.pageSize = size;
        },
        handleCurrentChange: function(currentPage){
          this.currentPage = currentPage;
          this.doctype(currentPage);
        },

        // addRow(tableData,event){//新增一行
          //之前一直想不到怎么新增一行空数据，最后幸亏一位朋友提示：表格新增一行，其实就是源数据的新增，
          //从源数据入手就可以实现了，于是 恍然大悟啊！
          // tableData.push({ name: '', xpath: '',desc:'',value:'',defVal:'' })
        // },
      }
  };
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  #vue{
    color: green;
    font-size: 28px;
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
    background: #67c23a;
    color: #fff;
  }
  .el-menu-item.is-active {
      color: #67c23a;
  }
  .headlogo{
    float: left;
    height: 60px;
    margin: 0 20px;
    width: 300px;
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
