<!-- <template>
  <div id="vue">Hello Vue.js! {{ message }}</div>
</template> -->
<template>
  <div>
    <el-button-group style="float: left; margin:10px">
      <!-- <el-button type="primary" icon="el-icon-circle-plus-outline" size="small" @click="$refs.editable.insertAt({name: `New last ${Date.now()}`, flag: true, createDate: Date.now()}, -1)">新增m</el-button> -->
      <el-button type="primary" icon="el-icon-circle-plus-outline" size="small" @click.native="dialogFormVisible = true">添加m</el-button>
      <el-button type="info" size="small" @click="$refs.editable.revert()">放弃更改</el-button>
      <el-button type="info" size="small" icon="el-icon-delete" @click="$refs.editable.clear()">清空数据</el-button>
    </el-button-group>
    <el-button-group  style="float: right; margin:10px">
      <el-button type="primary" icon="el-icon-circle-plus-outline" size="small">搜索</el-button>
      <el-button type="primary" icon="el-icon-refresh" size="small">刷新</el-button>
    </el-button-group>
    <vxe-toolbar></vxe-toolbar>

    <vxe-table ref="xTable"
      :data.sync="usergroupdata" border style="width: 100%" stripe :edit-config="{trigger: 'click', mode: 'cell'}"
      @edit-actived="editActivedEvent"
      @edit-closed="editClosedEvent">
      <vxe-table-column title="序号" type="index" show-overflow-tooltip width="50"  align="center"></vxe-table-column>
      <!-- <vxe-table-column label="GroupType" prop="Group.GroupType" :editRender="{Name: 'ElInput'}" align="cente"></vxe-table-column> -->
      <vxe-table-column field="Group.GroupType" title="GroupType" :editRender="{type: 'default'}" align="center">
        <template slot="edit" slot-scope="scope">
          <el-select v-model="scope.row.Group.GroupType" clearable>
            <el-option
              v-for="item in grouptypedata"
              :key="item.ID"
              :label="item.Name"
              :value="item.Name">
            </el-option>
          </el-select>
        </template>
        <template slot-scope="scope">{{ getColumnLabel(scope.row.Group.GroupType) }}</template>
      </vxe-table-column>

      <vxe-table-column title="GrouName" field="Group.Name" :edit-render="{name: 'input'}" align="center"></vxe-table-column>

      <vxe-table-column title="Users" field="Users[0].LastName" :editRender="{type: 'default'}" align="center">
        <template slot="edit" slot-scope="scope">
          <el-select v-model="scope.row.Users[0].ID" clearable>
            <el-option
              v-for="item in scope.row.Users"
              :key="item.ID"
              :label="item.FirstName+item.LastName"
              :value="item.ID">
            </el-option>
          </el-select>
        </template>
        <!-- <template slot-scope="scope">{{ getColumnLabel2(scope.row.Users[0].LastName) }}</template> -->
      </vxe-table-column>
      <vxe-table-column  title="操作" align="center">
        <template slot-scope="scope">
          <el-button-group>
            <el-button size="mini" @click="handleSubmit(scope.$index, scope.row)">Save</el-button>
            <el-button size="mini" type="danger" @click="deleteRow(scope.$index, usergroupdata)">Delete</el-button>
          </el-button-group>
        </template>
      </vxe-table-column>
    </vxe-table>
    <el-pagination background
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
      :current-page="currentPage"
      :page-sizes="[10, 50, 100, 200]"
      :page-size="pageSize"
      layout="total, sizes, prev, pager, next, jumper"
      :total="total" style="float: right; margin:10px">
    </el-pagination>

    <el-dialog title="添加usergroup" :visible.sync="dialogFormVisible" center>
      <el-form :model="ruleForm2" status-icon :rules="rules2" ref="ruleForm2" label-width="100px" class="demo-ruleForm">
        <el-form-item label="group" prop="groupid2">
          <!-- <template > -->
            <el-select v-model="ruleForm2.groupid2" clearable>
              <el-option
                v-for="item in groupdata.groups"
                :key="item.ID"
                :label="item.Name"
                :value="item.ID">
              </el-option>
            </el-select>
          <!-- </template> -->
        </el-form-item>

        <el-form-item label="users" prop="userid2">
          <!-- <template> -->
            <el-select v-model="ruleForm2.userid2" multiple clearable placeholder="请选择">
              <el-option
                v-for="item in userdata.users"
                :key="item.ID"
                :label="item.FirstName+item.LastName"
                :value="item.ID">
              </el-option>
            </el-select>
          <!-- </template> -->
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
    name: 'userGroup',
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
            groupid2:'',
            userid2:'',
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
          usergroupdata:[
            {
              ID:1,
              Group:
              {
                ID:1,
                Name:'设计组',
                GroupType:'G'
              },
              Users:[
                {
                  ID:1,
                  FirstName:"秦",
                  LastName:'晓川1',
                  Email:'xc-qin1@163.com',
                  Active:true
                },
                {
                  ID:2,
                  FirstName:"秦",
                  LastName:'晓川2',
                  Email:'xc-qin2@163.com',
                  Active:true
                }
              ]
            }
          ],
          grouptypedata: [
            {
              ID:1,
              Name:'S'
            },
            {
              ID:2,
              Name:'G'
            },
          ],
          userdata:[
            {ID:1,FirstName:"秦",LastName:'晓川1',Email:'xc-qin1@163.com',Active:true},
            {ID:2,FirstName:"秦",LastName:'晓川2',Email:'xc-qin2@163.com',Active:true}
          ],
          groupdata:[
            {ID:1,Name:"秦晓川",GroupType:"S"},
            {ID:2,Name:"校核组",GroupType:"G"}
          ],
          userid2:[],
          groupid2:''
        };
      },
      mounted:function () {
        this.usergroup(this.currentPage);
        this.group(this.currentPage);
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
                url: "/flowusergroup",//请求地址
                params:{
                  gid:this.ruleForm2.groupid2,
                  uid:this.ruleForm2.userid2
                },
                // data: {
                  // name:this.ruleForm2.typename,
                  // "thirdapp_id":1//请求参数
                // }
              })
              .then((response) => {
                if (response != "err") {
                  // this.$Message.info('用户名或密码错误，请送心')
                  //提交成功做的动作
                  this.$message({
                    type: 'success',
                    message: '提交成功' 
                  });
                  //刷新表格
                  // this.docaction(currentPage);
                  this.dialogFormVisible = false;
                } else {
                  // console.log(response.data)
                  //写入失败！
                  this.$message.error('写入失败！');
                }
              })
              // .then(response => (this.posts = response.data.articles))
              // .then(function (response) {
              //   console.log(response);
              //   if (response=="err") {
              //     //提交成功做的动作
              //     this.$message({
              //       type: 'success',
              //       message: '提交成功' 
              //     });
              //     //刷新表格
              //     this.flowtypelist();
              //     this.dialogFormVisible = false;                 
              //   } else {
              //     //写入失败！
              //     this.$message.error('写入失败！');
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
        addRow(usergroupdata,event){
          usergroupdata.push({ Id:'',Group:{ID:'',Name:'',GroupType:'G'},Users:[{ID:'',FirstName: '', LastName: '', Email: '', Active: true}]})
        },
        handleSubmit(index, row) {
          console.log(row);
              axios({
                method: "POST",//请求方式
                url: "/flowusergroup",//请求地址
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
              .then((response) => {
                if (response != "err") {
                  // this.$Message.info('用户名或密码错误，请送心')
                  //提交成功做的动作
                  this.$message({
                    type: 'success',
                    message: '提交成功' 
                  });
                  //刷新表格
                  // this.docaction(currentPage);
                  this.dialogFormVisible = false;
                } else {
                  // console.log(response.data)
                  //写入失败！
                  this.$message.error('写入失败！');
                }
              })
              // .then(response => (this.posts = response.data.articles))
              // .then(function (response) {
              //   console.log(response);
              //   if (response=="err") {
              //     //提交成功做的动作
              //     this.$message({
              //       type: 'success',
              //       message: '提交成功' 
              //     });
              //     //刷新表格
              //     this.user(currentPage);
              //     this.dialogFormVisible = false;                 
              //   } else {
              //     //写入失败！
              //     this.$message.error('写入失败！');
              //   }
              // })
              .catch(function (error) {
                console.log(error);
              });
        },
        usergroup(currentPage){
          axios({
            method: 'get',
            url: '/flowgroupuserslist',//2.get通过params选项
            params:{
              page:currentPage
            }
          })
          .then(response => (this.usergroupdata = response.data))
          .catch(function (error) {
            console.log(error);
          });
        },
        group(currentPage){
          axios({
            method: 'get',
            url: '/flowgrouplist',//2.get通过params选项
            params:{
              page:currentPage
            }
          })
          .then(response => (this.groupdata = response.data))
          .catch(function (error) {
            console.log(error);
          });
        },
        user(currentPage){
          axios({
            method: 'get',
            url: '/flowuserlist',//2.get通过params选项
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
        toNext: function(index) {
          sessionStorage.ticketName =this.sortList[index].list;
          this.$router.push('/mine/tiketOrder');
        },
        handleSizeChange: function (size) {
          this.pageSize = size;
        },
        handleCurrentChange: function(currentPage){
          this.currentPage = currentPage;
          this.usergroup(currentPage);
        },
        getColumnLabel (value) {
          let selectItem = this.grouptypedata.find(item => item.Name === value)
          return selectItem ? selectItem.Name : null
        },
        getColumnLabel2 (value) {
          let selectItem = this.grouptypedata.find(item => item.Name === value)
          return selectItem ? selectItem.Name : null
        },
        editActivedEvent ({ row, column }, event) {
          console.log(`打开 ${column.title} 列编辑`)
        },
        editClosedEvent ({ row, column }, event) {
          console.log(`关闭 ${column.title} 列编辑`)
        },
        insertEvent (row) {
          let record = {
            sex: '1'
          }
          this.$refs.xTable.insertAt(record, row)
            .then(({ row }) => this.$refs.xTable.setActiveCell(row, 'sex'))
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
