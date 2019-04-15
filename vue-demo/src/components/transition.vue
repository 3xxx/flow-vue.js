<template>
  <div>
    <el-button-group style="float: left; margin:10px">
      <el-button type="primary" icon="el-icon-circle-plus-outline" size="small" @click="$refs.editable.insertAt({name: `New last ${Date.now()}`, flag: true, createDate: Date.now()}, -1)">新增</el-button>
      <el-button type="info" size="small" @click="$refs.editable.revert()">放弃更改</el-button>
      <el-button type="info" size="small" icon="el-icon-delete" @click="$refs.editable.clear()">清空数据</el-button>
    </el-button-group>
    <el-button-group  style="float: right; margin:10px">
      <el-button type="primary" icon="el-icon-circle-plus-outline" size="small">搜索</el-button>
      <el-button type="primary" icon="el-icon-refresh" size="small">刷新</el-button>
    </el-button-group>

    <el-editable ref="editable"
      :data.sync="transitiondata.transitions" border style="width: 100%" stripe>
      <el-editable-column label="序号" type="index" show-overflow-tooltip width="50"  align="center"></el-editable-column>
      <el-editable-column prop="DoctypeId" label="DOCTYPE" :editRender="{type: 'default'}" align="center">
        <template slot="edit" slot-scope="scope">
          <el-select v-model="scope.row.DoctypeId" clearable>
            <el-option
              v-for="item in doctypedata.doctypes"
              :key="item.ID"
              :label="item.Name"
              :value="item.ID">
            </el-option>
          </el-select>
        </template>
        <template slot-scope="scope">{{ getColumnLabel(scope.row.DoctypeId) }}</template>
      </el-editable-column>
      <el-editable-column prop="FromStateId" label="FromState" :editRender="{type: 'default'}" align="center">
        <template slot="edit" slot-scope="scope">
          <el-select v-model="scope.row.FromStateId" clearable>
            <el-option
              v-for="item in docstatedata"
              :key="item.ID"
              :label="item.Name"
              :value="item.ID">
            </el-option>
          </el-select>
        </template>
        <template slot-scope="scope">{{ getColumnLabel2(scope.row.FromStateId) }}</template>
      </el-editable-column>
      <el-editable-column  prop="DocactionId" label="DOCACTION" :editRender="{type: 'default'}" align="center">
        <template slot="edit" slot-scope="scope">
          <el-select v-model="scope.row.DocactionId" clearable>
            <el-option
              v-for="item in docactiondata"
              :key="item.ID"
              :label="item.Name"
              :value="item.ID">
            </el-option>
          </el-select>
        </template>
        <template slot-scope="scope">{{ getColumnLabel3(scope.row.DocactionId) }}</template>
      </el-editable-column>
      <el-editable-column prop="ToStateId" label="ToState" :editRender="{type: 'default'}" align="center">
        <template slot="edit" slot-scope="scope">
          <el-select v-model="scope.row.ToStateId" clearable>
            <el-option
              v-for="item in docstatedata"
              :key="item.ID"
              :label="item.Name"
              :value="item.ID">
            </el-option>
          </el-select>
        </template>
        <template slot-scope="scope">{{ getColumnLabel4(scope.row.ToStateId) }}</template>
      </el-editable-column>
      <el-editable-column label="操作" align="center">
        <template slot-scope="scope">
          <el-button-group>
          <el-button size="mini" @click="handleSubmit(scope.$index, scope.row)">Save</el-button>
          <el-button size="mini" type="danger" @click="deleteRow(scope.$index, transitiondata)">Delete</el-button>
          </el-button-group>
        </template>
      </el-editable-column>
    </el-editable>
    <el-pagination background
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
      :current-page="currentPage"
      :page-sizes="[10, 50, 100, 200]"
      :page-size="pageSize"
      layout="total, sizes, prev, pager, next, jumper"
      :total="transitiondata.total" style="float: right; margin:10px">
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
    name: 'transition',
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
          transitiondata: [
            {
              page: 1,
              total: 6,
              transitions:{
                DocactionId: 1,
                DoctypeId: 1,
                FromStateId: 2,
                Id: 1,
                ToStateId: 3
              }
            }
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
          docstatedata: [
            {
              ID:3,
              Name:'设计中……'
            },
            {
              ID:5,
              Name:'校核中……'
            }
          ],
          docactiondata: [
            {
              ID:4,
              Name:'提交设计',
              Reconfirm: false
            }
          ],
          search: '',
          sortList: [
            {'list': '观影兑换券', },
            {'list': '室内乐兑换券', },
            {'list': '沙龙兑换券', }
          ],
        };
      },
      mounted:function () {
        this.transition(this.currentPage);
        this.doctype(this.currentPage);
        this.docstate(this.currentPage);
        this.docaction(this.currentPage);
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
        transition(currentPage){
          axios({
            method: 'get',
            url: '/api/flowtransitionlist',//2.get通过params选项
            params:{
              page:currentPage,
              limit:this.pageSize
            }
          })
          .then(response => (this.transitiondata = response.data))
          .catch(function (error) {
            console.log(error);
          });
        },
        doctype(currentPage){
          axios({
            method: 'get',
            url: '/api/flowtypelist',//2.get通过params选项
            // params:{
            //   page:currentPage
            // }
          })
          .then(response => (this.doctypedata = response.data))
          .catch(function (error) {
            console.log(error);
          });
        },
        docstate(currentPage){
          axios({
            method: 'get',
            url: '/api/flowstatelist',//2.get通过params选项
            params:{
              page:currentPage
            }
          })
          .then(response => (this.docstatedata = response.data))
          .catch(function (error) {
            console.log(error);
          });
        },
        docaction(currentPage){
          axios({
            method: 'get',
            url: '/api/flowactionlist',//2.get通过params选项
            params:{
              page:currentPage
            }
          })
          .then(response => (this.docactiondata = response.data))
          .catch(function (error) {
            console.log(error);
          });
        },
        deleteRow(index, rows) {//删除改行
          //先删除数据库
          //再删除前端行
          rows.splice(index, 1);
        },
        addRow(transitiondata,event){
          transitiondata.push({ Id:'',})
        },
        handleSubmit(index, row) {
          console.log(row);
              axios({
                method: "POST",//请求方式
                url: "/api/flowtransition",//请求地址
                params:{
                  dtid:row.DoctypeId,
                  dsid1:row.FromStateId,
                  daid:row.DocactionId,
                  dsid2:row.ToStateId
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
                  this.transition(currentPage);
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

        toNext: function(index) {
          sessionStorage.ticketName =this.sortList[index].list;
          this.$router.push('/mine/tiketOrder');
        },
        handleSizeChange: function (size) {
          this.pageSize = size;
        },
        handleCurrentChange: function(currentPage){
          this.currentPage = currentPage;
          this.transition(currentPage);
        },
        getColumnLabel (value) {
          let selectItem = this.doctypedata.doctypes.find(item => item.ID === value)
          return selectItem ? selectItem.Name : null
        },
        getColumnLabel2 (value) {
          let selectItem = this.docstatedata.find(item => item.ID === value)
          return selectItem ? selectItem.Name : null
        },
        getColumnLabel3 (value) {
          let selectItem = this.docactiondata.find(item => item.ID === value)
          return selectItem ? selectItem.Name : null
        },
        
        getColumnLabel4 (value) {
          let selectItem = this.docstatedata.find(item => item.ID === value)
          return selectItem ? selectItem.Name : null
        },
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
