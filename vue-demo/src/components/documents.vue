<!-- <template>
  <div id="vue">Hello Vue.js! {{ message }}</div>
</template> -->
<template>
  <div>
    <el-button-group style="float: left; margin:10px">
      <!-- <el-button type="primary" icon="el-icon-circle-plus-outline" size="small" @click="$refs.editable.insert({name: `New ${Date.now()}`, flag: true, createDate: Date.now()})">新增一行</el-button> -->
      <!-- <el-button type="primary" icon="el-icon-circle-plus-outline" size="small" @click="$refs.editable.insertAt({name: `New last ${Date.now()}`, flag: true, createDate: Date.now()}, -1)">新增</el-button> -->
      <el-button type="primary" icon="el-icon-circle-plus-outline" size="small" @click.native="dialogFormVisible = true">添加m</el-button>
      <el-button type="info" size="small" @click="$refs.editable.revert()">放弃更改</el-button>
      <el-button type="info" size="small" icon="el-icon-delete" @click="$refs.editable.clear()">清空数据</el-button>
    </el-button-group>
    <el-button-group  style="float: right; margin:10px">
      <el-button type="primary" icon="el-icon-circle-plus-outline" size="small">搜索</el-button>
      <el-button type="primary" icon="el-icon-refresh" size="small">刷新</el-button>
    </el-button-group>

    <el-editable ref="editable"
      :data="documentsdata" border style="width: 100%" stripe>
      <el-editable-column label="AccCtx" prop="AcId" size="mini" :editRender="{type: 'default'}" align="center">
        <template slot="edit" slot-scope="scope">
          <el-select v-model="scope.row.AcId" clearable>
            <el-option
              v-for="item in accesscontextdata"
              :key="item.ID"
              :label="item.Name"
              :value="item.ID">
            </el-option>
          </el-select>
        </template>
        <template slot-scope="scope">{{ getColumnLabel(scope.row.AcId) }}</template>
      </el-editable-column>
      <el-editable-column label="DOCSTATE" prop="DocstateId" size="mini" :editRender="{type: 'default'}" align="center">
        <template slot="edit" slot-scope="scope">
          <el-select v-model="scope.row.DocstateId" clearable>
            <el-option
              v-for="item in docstatedata"
              :key="item.ID"
              :label="item.Name"
              :value="item.ID">
            </el-option>
          </el-select>
        </template>
        <template slot-scope="scope">{{ getColumnLabel2(scope.row.DocstateId) }}</template>
      </el-editable-column>
      <el-editable-column label="GROUP" prop="GroupId" size="mini" :editRender="{type: 'default'}" align="center">
        <template slot="edit" slot-scope="scope">
          <el-select v-model="scope.row.GroupId" clearable>
            <el-option
              v-for="item in groupdata"
              :key="item.ID"
              :label="item.Name"
              :value="item.ID">
            </el-option>
          </el-select>
        </template>
        <template slot-scope="scope">{{ getColumnLabel3(scope.row.GroupId) }}</template>
      </el-editable-column>
      <el-editable-column label="TITLE" prop="Title" size="mini" :editRender="{Name: 'ElInput'}" align="center">
      </el-editable-column>
      <el-editable-column label="DATA" prop="Data" size="mini" :editRender="{Name: 'ElInput'}" align="center">
      </el-editable-column>
      <el-editable-column label="PATH" prop="Path" size="mini" :editRender="{Name: 'ElInput'}" align="center">
      </el-editable-column>
      <el-editable-column label="CTIME" prop="Ctime" size="mini" :formatter="formatter">
      </el-editable-column>
      <el-editable-column  label="操作" align="center">
        <template slot-scope="scope">
          <el-button-group>
            <el-button size="mini" @click="handleSubmit(scope.$index, scope.row)">Save</el-button>
            <el-button size="mini" type="danger" @click="deleteRow(scope.$index, documentsdata)">Delete</el-button>
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
      :total="total" style="float: right; margin:10px">
    </el-pagination>

    <el-dialog title="添加document" :visible.sync="dialogFormVisible" center>
      <el-form :model="ruleForm2" status-icon :rules="rules2" ref="ruleForm2" label-width="100px" class="demo-ruleForm">
        <el-form-item label="DOCTYPE" prop="dtid">
          <template >
            <el-select v-model="ruleForm2.value1" clearable>
              <el-option
                v-for="item in doctypedata.doctypes"
                :key="item.ID"
                :label="item.Name"
                :value="item.ID">
              </el-option>
            </el-select>
          </template>
        </el-form-item>
        <el-form-item label="ACCESSCONTEXT" prop="acid">
          <template>
            <el-select v-model="ruleForm2.value2" clearable>
              <el-option
                v-for="item in accesscontextdata"
                :key="item.ID"
                :label="item.Name"
                :value="item.ID">
              </el-option>
            </el-select>
          </template>
        </el-form-item>
        <el-form-item label="GROUP" prop="gid">
          <template>
            <el-select v-model="ruleForm2.value3" value-key="ID" clearable>
              <el-option
                v-for="item in groupdata"
                :key="item.ID"
                :label="item.Name"
                :value="item.ID">
              </el-option>
            </el-select>
          </template>
        </el-form-item>
        <el-form-item label="TITLE" prop="docname" class="transparentIcon" style='width: 320px;'>
          <el-input v-model.number="ruleForm2.docname" auto-complete="off" placeholder="输入名称"></el-input>
        </el-form-item>
        <el-form-item label="DATA" prop="docdata" class="transparentIcon" style='width: 320px;'>
          <el-input v-model.number="ruleForm2.docdata" auto-complete="off" placeholder="输入名称"></el-input>
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
    name: 'documents',
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
            {ID:1,Name:"图纸"},
            {ID:2,Name:"合同"}
          ],
          groupdata:[
            {ID:1,Name:"秦晓川",GroupType:"S"},
            {ID:2,Name:"校核组",GroupType:"G"}
          ],
          value1:'',
          value2:'',
          value3:'',
        };
      },
      mounted:function () {
        this.documents(3,this.currentPage);
        this.accesscontext(this.currentPage);
        this.group(this.currentPage);
        this.docstate(this.currentPage);
        this.doctype(this.currentPage);
      },
      methods:{
        submitForm(formName) {
          this.$refs[formName].validate((valid) => {
            if (valid) {
              axios({
                method: "POST",//请求方式
                url: "/api/flowdoc",//请求地址
                params:{
                  // acid:this.ruleForm2.AcId,
                  // dsid:this.ruleForm2.DocstateId,
                  dtid:this.ruleForm2.value1,
                  // dtid:this.ruleForm2.DoctypeId,
                  acid:this.ruleForm2.value2,
                  // gid:this.ruleForm2.GroupId,
                  gid:this.ruleForm2.value3,
                  docname:this.ruleForm2.docname,
                  docdata:this.ruleForm2.docdata,
                  // path:this.ruleForm2.Path
                  // ctime:row.Ctime
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
                  this.documents(3,this.currentPage);
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
        addRow(documentsdata,event){
          documentsdata.push({ Id:'',Group:''})
        },
        handleSubmit(index, row) {
          console.log(row);
              axios({
                method: "POST",//请求方式
                url: "/api/flowdocument",//请求地址
                params:{
                  acid:row.AcId,
                  // dsid:row.DocstateId,
                  dtid:row.DoctypeId,
                  gid:row.GroupId,
                  title:row.Title,
                  data:row.Data,
                  path:row.Path
                  // ctime:row.Ctime
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
        documents(dtid,currentPage){
          axios({
            method: 'get',
            url: '/api/flowdoclist',//2.get通过params选项
            params:{
              page:currentPage,
              limit:this.pageSize,
              dtid:dtid
            }
          })
          .then(response => (this.documentsdata = response.data))
          .catch(function (error) {
            console.log(error);
          });
        },
        doctype(currentPage){
          axios({
            method: 'get',
            url: '/api/flowtypelist',//2.get通过params选项
            params:{
              page:currentPage
            }
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
        accesscontext(currentPage){
          axios({
            method: 'get',
            url: '/api/flowaccesscontextlist',//2.get通过params选项
            params:{
              page:currentPage
            }
          })
          .then(response => (this.accesscontextdata = response.data))
          .catch(function (error) {
            console.log(error);
          });
        },
        group(currentPage){
          axios({
            method: 'get',
            url: '/api/flowgrouplist',//2.get通过params选项
            params:{
              page:currentPage
            }
          })
          .then(response => (this.groupdata = response.data))
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
          this.documents(3,this.currentPage)
        },
        handleCurrentChange: function(currentPage){
          this.currentPage = currentPage;
          this.documents(currentPage);
        },
        formatter:function(row, column){
          var date = row.Ctime;
          // console.log(date)
          if (date === undefined) {
            return "";
          }
          // return util.formatDate.format(new Date(date), 'yyyy-MM-dd');
          // this.$moment().format('YYYY-MM-DD HH:mm:ss')
          return this.$moment(date).format("YYYY-MM-DD HH:mm");
          // https://blog.csdn.net/ysq0317/article/details/81089962
          // vue的话，在moment.js的官网里，是给了安装方法的
          // cnpm install moment --save   
          // 然后再入口文件 main.js中导入并使用
          // import moment from 'moment'//导入文件
          // Vue.prototype.$moment = moment;//赋值使用
        },
        getColumnLabel (value) {
          let selectItem = this.accesscontextdata.find(item => item.ID === value)
          return selectItem ? selectItem.Name : null
        },
        getColumnLabel2 (value) {
          let selectItem = this.docstatedata.find(item => item.ID === value)
          return selectItem ? selectItem.Name : null
        },
        getColumnLabel3 (value) {
          let selectItem = this.groupdata.find(item => item.ID === value)
          return selectItem ? selectItem.Name : null
        }
        // attr5ChangeEvent (scope) {
        //   let list = this.$refs.editable.getRecords()
        //   this.attr5Options = this.typeOptions.filter(item => !list.some(row => row.attr5 === item.label))
        //   this.$refs.editable.updateStatus(scope)
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
