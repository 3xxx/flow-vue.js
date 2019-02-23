<!-- <template>
  <div id="vue">Hello Vue.js! {{ message }}</div>
</template> -->
<template>
  <div>
    <el-col :span="24" class="breadcrumb-container">
      <el-button-group style="float: left; margin:10px">
        <el-button type="primary" size="small" @click="addRow(userdata)">新增</el-button>
        <el-button type="primary" icon="el-icon-circle-plus-outline" size="small" @click.native="dialogFormVisible = true">添加</el-button>
        <el-button type="primary" icon="el-icon-share" size="small">分享</el-button>
        <el-button type="primary" icon="el-icon-delete" size="small">删除</el-button>
      </el-button-group>
      <el-button-group  style="float: right; margin:10px">
        <el-button type="primary" icon="el-icon-circle-plus-outline" size="small">搜索</el-button>
        <!-- <el-input placeholder="请输入内容" v-model="input5" class="input-with-select" size="small" type="primary"> -->
        <!-- <el-button type="primary" slot="append" icon="el-icon-search" size="small"></el-button> -->
        <!-- </el-input> -->
        <!-- <el-input v-model="search" size="mini" placeholder="输入关键字搜索"/> -->
      <el-button type="primary" icon="el-icon-refresh" size="small">刷新</el-button>
        <!-- <el-button type="primary" icon="el-icon-delete" size="small">导出</el-button> .slice((currentPage-1)*pageSize,currentPage*pageSize)-->
      </el-button-group>
    </el-col>

    <el-table
      :data="userdata" border style="width: 100%" stripe>
      <el-table-column label="ID" prop="ID" align="center"></el-table-column>
      <el-table-column label="Name" prop="FirstName" align="center">
        <template slot-scope="scope">
          <el-input size="mini" v-model="scope.row.FirstName"></el-input>
        </template>
      </el-table-column>
      <el-table-column label="Name" prop="LastName" align="center">
        <template slot-scope="scope">
          <el-input size="mini" v-model="scope.row.LastName"></el-input>
        </template>
      </el-table-column>
      <el-table-column label="Name" prop="Email" align="center">
        <template slot-scope="scope">
          <el-input size="mini" v-model="scope.row.Email"></el-input>
        </template>
      </el-table-column>
      <el-table-column prop="Active" label="ACTIVE" :formatter="formatter">
        <template slot-scope="scope">
          <el-select v-model="scope.row.Active" clearable>
            <el-option
              v-for="item in useractivedata"
              :key="item.ID"
              :label="item.Name"
              :value="item.Value">
            </el-option>
          </el-select>
        </template>
      </el-table-column>

      <el-table-column  label="操作" align="center">
        <template slot-scope="scope">
          <el-button size="mini" @click="handleSubmit(scope.$index, scope.row)">Save</el-button>
          <el-button size="mini" type="danger" @click="deleteRow(scope.$index, userdata)">Delete</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination background
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
      :current-page="currentPage"
      :page-sizes="[10, 50, 100, 200]"
      :page-size="pageSize"
      layout="total, sizes, prev, pager, next, jumper"
      :total="total" style="float: right; margin:10px">
    </el-pagination>

    <el-dialog title="定义doctype" :visible.sync="dialogFormVisible" center>
    <!-- 插入测试 -->
      <el-form :model="ruleForm2" status-icon :rules="rules2" ref="ruleForm2" label-width="100px" class="demo-ruleForm">
        <el-form-item label="名称" prop="typename">
          <el-input v-model.number="ruleForm2.typename" auto-complete="off" placeholder="输入名称"></el-input>
        </el-form-item>
      </el-form>
      <!-- 插入测试 -->
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false; resetForm('ruleForm2')">取 消</el-button>
        <el-button type="primary" @click="submitForm('ruleForm2')">确 定</el-button>
      </div>
    </el-dialog>

  </div>
</template>

<script type="text/javascript">
/* eslint-disable */
  const axios = require('axios');
  export default { // 这里需要将模块引出，可在其他地方使用
    name: 'user',
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
          formLabelWidth: '120px',
          userdata:[
            {ID:1,FirstName:"秦",LastName:'晓川1',Email:'xc-qin1@163.com',Active:true},
            {ID:2,FirstName:"秦",LastName:'晓川2',Email:'xc-qin2@163.com',Active:true}
          ],
          search: '',
          useractivedata: [
            {
              ID:1,
              Value:true,
              Name:'true'
            },
            {
              ID:0,
              Value:false,
              Name:'false'
            },
          ],
        };
      },
      mounted:function () {
        this.user(this.currentPage);
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
                url: "http://127.0.0.1:8081/v1/admin/flowtype",//请求地址
                params:{
                  name:this.ruleForm2.typename,
                },
                data: {
                  name:this.ruleForm2.typename,
                  // "thirdapp_id":1//请求参数
                }
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
                url: "http://127.0.0.1:8081/v1/admin/flowuser",//请求地址
                params:{
                  firstname:row.FirstName,
                  lastname:row.LastName,
                  email:row.Email,
                  active:row.Active
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
                  this.user(currentPage);
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
        user(currentPage){
          axios({
            method: 'get',
            url: 'http://127.0.0.1:8081/v1/admin/flowuserlist',//2.get通过params选项
            params:{
              page:currentPage
            }
          })
          .then(response => (this.userdata = response.data))
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
        handleSizeChange: function (size) {
          this.pageSize = size;
        },
        handleCurrentChange: function(currentPage){
          this.currentPage = currentPage;
          this.user(currentPage);
        },
        formatter(row, column) {
          // return row.address;
          return String(row.Active);
        }
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
