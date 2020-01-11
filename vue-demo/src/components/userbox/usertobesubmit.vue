<!-- 用户待提交的流程 -->
<template>
  <div>
    <el-button-group style="float: left; margin:10px">
      <el-button type="primary" icon="el-icon-circle-plus-outline" size="small" @click.native="dialogFormVisible = true">添加</el-button>
      <el-button type="info" size="small" @click="$refs.editable.revert()">放弃更改</el-button>
      <el-button type="info" size="small" icon="el-icon-delete" @click="$refs.editable.clear()">清空数据</el-button>
    </el-button-group>
    <el-button-group  style="float: right; margin:10px">
      <el-button type="primary" icon="el-icon-circle-plus-outline" size="small">搜索</el-button>
      <el-button type="primary" icon="el-icon-refresh" size="small">刷新</el-button>
    </el-button-group>
    <vxe-toolbar></vxe-toolbar>
    <el-col>
      <el-button-group style="float: left; margin:10px">
          <el-select v-model="dtvalue" placeholder="请选择doctype" @change="changedtValue" clearable>
            <el-option
              v-for="item in doctypedata.doctypes"
              :key="item.ID"
              :label="item.Name"
              :value="item.ID">
            </el-option>
          </el-select>
          <el-select v-model="acvalue" placeholder="请选择accCTX" @change="changeacValue" clearable>
            <el-option
              v-for="item in accesscontextdata.accesscontexts"
              :key="item.ID"
              :label="item.Name"
              :value="item.ID">
            </el-option>
          </el-select>
          <el-select v-model="dsvalue" placeholder="请选择docstate" @change="changedsValue" clearable>
            <el-option
              v-for="item in docstatedata.docstates"
              :key="item.ID"
              :label="item.Name"
              :value="item.ID">
            </el-option>
          </el-select>
      </el-button-group>
    </el-col>

    <vxe-table stripe border ref="xTable" :data.sync="documentsdata.docs" style="width: 100%">
      <vxe-table-column title="序号" type="index" show-overflow-tooltip width="50"  align="center"></vxe-table-column>
      <vxe-table-column title="TITLE文档名" field="Title" align="center"></vxe-table-column>
      <vxe-table-column title="DOCSTATE" field="DocState.ID" align="center">
        <template slot="edit" slot-scope="scope">
          <el-select v-model="scope.row.DocState.ID" clearable>
            <el-option
              v-for="item in docstatedata.docstates"
              :key="item.ID"
              :label="item.Name"
              :value="item.ID">
            </el-option>
          </el-select>
        </template>
        <template slot-scope="scope">{{ getColumnLabel3(scope.row.DocState.ID) }}</template>
      </vxe-table-column>
      <vxe-table-column title="GROUP" field="Group.Name" align="center">
        <template slot="edit" slot-scope="scope">
          <el-select v-model="scope.row.Group.ID" clearable>
            <el-option
              v-for="item in groupdata.groups"
              :key="item.ID"
              :label="item.Name"
              :value="item.ID">
            </el-option>
          </el-select>
        </template>
        <!-- <template slot-scope="scope">{{ getColumnLabel4(scope.row.Group.ID) }}</template> -->
      </vxe-table-column>
      <vxe-table-column title="PATH" field="Path" size="mini" align="center"></vxe-table-column>
      <vxe-table-column title="CTIME" field="Ctime" size="mini" :formatter="formatter" align="center">
      </vxe-table-column>
      <vxe-table-column  title="操作" align="center">
        <template slot-scope="scope">
          <el-button-group>
            <el-button size="mini" @click="detail(scope.$index, scope.row)">detail</el-button>
            <el-button size="mini" type="danger" @click="deleteRow(scope.$index, documentsdata)">Del</el-button>
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
      :total="documentsdata.total" style="float: right; margin:10px">
    </el-pagination>

    <el-dialog title="添加document" :visible.sync="dialogFormVisible" center>
      <el-form :model="ruleForm2" status-icon :rules="rules2" ref="ruleForm2" label-width="100px" class="demo-ruleForm">
        <el-form-item label="DOCTYPE" prop="dtid">
          <template >
            <el-select v-model="ruleForm2.dtid2" clearable>
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
            <el-select v-model="ruleForm2.acid2" clearable>
              <el-option
                v-for="item in accesscontextdata.accesscontexts"
                :key="item.ID"
                :label="item.Name"
                :value="item.ID">
              </el-option>
            </el-select>
          </template>
        </el-form-item>
        <el-form-item label="GROUP" prop="gid">
          <template>
            <el-select v-model="ruleForm2.gid2" value-key="ID" clearable>
              <el-option
                v-for="item in groupdata.groups"
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
    name: 'usertobesubmit',
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
            // name:'',
            // pass: '',
            // num: '',
            // delivery: false,
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
              ID:1,
              DocType:{
                ID: 4,
                Name: "变更立项"
              },
              Path: "",
              AccessContext: {
                "ID": 1
              },
              DocState: {
                "ID": 8,
                "Name": "校核中..."
              },
              Group: {
                "ID": 12,
                "Name": "校核人员组",
                "GroupType": ""
              },
              Ctime: "2019-02-09T21:52:40Z",
              Title: "这个是测试数据",
              Path: "/path/document"
            }
          ],
          accesscontextdata: [
            {
              ID:1,
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
          dtid2:'',
          acid2:'',
          gid2:'',
          acvalue:'',
          dtvalue:'',
          dsvalue:'',
          dtid:'',
          dsid:'',
          acid:'',
          message:''
        };
      },
      mounted:function () {
        // this.documents(3,this.currentPage);
        this.accesscontext(this.currentPage);
        // this.group(this.currentPage);
        this.docstate(this.currentPage);
        this.doctype(this.currentPage);
        // this.documents(this.currentPage)
      },
      methods:{
        detail(index, row){
          // console.log(row.ID);
          // console.log(row.DocType.ID);
          this.$router.push({
            path: '/flow/documentdetail2',
            query: {docid: row.ID,dtid:row.DocType.ID}
          })
        },
        submitForm(formName) {
          this.$refs[formName].validate((valid) => {
            if (valid) {
              axios({
                method: "POST",//请求方式
                url: "/flowdoc",//请求地址
                params:{
                  // form:this.ruleForm2,
                  // dsid:this.ruleForm2.DocstateId,
                  dtid:this.ruleForm2.dtid2,
                  // dtid:this.ruleForm2.DoctypeId,
                  acid:this.ruleForm2.acid2,
                  // // gid:this.ruleForm2.GroupId,
                  gid:this.ruleForm2.gid2,
                  docname:this.ruleForm2.docname,
                  docdata:this.ruleForm2.docdata,
                },
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
              //   console.log(response.data.err);
              //   if (response.data.err=="ok") {
              //     //提交成功做的动作
              //     this.$message({
              //       type: 'success',
              //       message: '提交成功' 
              //     });
              //     //刷新表格
              //     this.documents(3,this.currentPage);
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
        addRow(documentsdata,event){
          documentsdata.push({ Id:'',Group:''})
        },
        handleSubmit(index, row) {
          // console.log(row);
              axios({
                method: "POST",//请求方式
                url: "/flowdocument",//请求地址
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
                // console.log(response);
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
        documents(currentPage){
          axios({
            method: 'get',
            url: '/flowdocumentlist2',//2.get通过params选项
            params:{
              page:currentPage,
              limit:this.pageSize,
              dtid:this.dtid,
              acid:this.acid,
              dsid:this.dsid
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
            url: '/flowtypelist',//2.get通过params选项
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
            url: '/flowstatelist',//2.get通过params选项
            // params:{
            //   page:currentPage
            // }
          })
          .then(response => (this.docstatedata = response.data))
          .catch(function (error) {
            console.log(error);
          });
        },
        accesscontext(currentPage){
          axios({
            method: 'get',
            url: '/flowaccesscontextlist',//2.get通过params选项
            // params:{
            //   page:currentPage
            // }
          })
          .then(response => (this.accesscontextdata = response.data))
          .catch(function (error) {
            console.log(error);
          });
        },
        group(currentPage){
          axios({
            method: 'get',
            url: '/flowgrouplist',//2.get通过params选项
            // params:{
            //   page:currentPage
            // }
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
          // this.documents(3,this.currentPage)
        },
        handleCurrentChange: function(currentPage){
          this.currentPage = currentPage;
          this.documents(currentPage);
        },
        formatter:function(row, column){
          var date = row.cellValue;
          // console.log(date)
          if (date === undefined) {
            return "";
          }
          // return util.formatDate.format(new Date(date), 'yyyy-MM-dd');
          // this.$moment().format('YYYY-MM-DD HH:mm:ss')
          return this.$moment(date).subtract(8,'hour').format("YYYY-MM-DD HH:mm");
          // https://blog.csdn.net/ysq0317/article/details/81089962
          // vue的话，在moment.js的官网里，是给了安装方法的
          // cnpm install moment --save   
          // 然后再入口文件 main.js中导入并使用
          // import moment from 'moment'//导入文件
          // Vue.prototype.$moment = moment;//赋值使用
        },
        // getColumnLabel (value) {
        //   let selectItem = this.accesscontextdata.accesscontexts.find(item => item.ID === value)
        //   return selectItem ? selectItem.Name : null
        // },
        // getColumnLabel2 (value) {
        //   let selectItem = this.doctypedata.doctypes.find(item => item.ID === value)
        //   return selectItem ? selectItem.Name : null
        // },
        getColumnLabel3 (value) {
          console.log(value)
          let selectItem = this.docstatedata.docstates.find(item => item.ID === value)
          return selectItem ? selectItem.Name : null
        },
        // getColumnLabel4 (value) {
        //   let selectItem = this.groupdata.groups.find(item => item.ID === value)
        //   return selectItem ? selectItem.Name : null
        // },
        changedtValue(value) {
          // console.log(value);
          this.dtid = value;
          if (this.acid!='' & this.dsid!=''){
            this.documents(this.currentPage);
          }
          // let obj = {};
          // obj = this.options.find((item)=>{
          //     return item.value === value;
          // });
          // console.log(obj.label);
        },
        changeacValue(value) {
          // console.log(value);
          this.acid = value;
          console.log(this.dtid);
          console.log(this.dsid);
          if (this.dtid!='' & this.dsid!=''){
            this.documents(this.currentPage);
          }
        },
        changedsValue(value) {
          // console.log(value);
          this.dsid = value;
          if (this.dtid!='' & this.dtid!=''){
            this.documents(this.currentPage);
          }
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
