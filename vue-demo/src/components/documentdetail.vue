<template>
  <div>
    <el-col :span="24" class="breadcrumb-container">
      <el-button-group style="float: left; margin:10px">
        <el-button type="primary" icon="el-icon-circle-plus-outline" size="small" @click.native="dialogFormVisible = true">添加</el-button>
        <el-button type="primary" icon="el-icon-share" size="small">分享</el-button>
        <el-button type="primary" icon="el-icon-delete" size="small">删除</el-button>
      </el-button-group>
      <el-button-group  style="float: right; margin:10px">
        <el-button type="primary" icon="el-icon-circle-plus-outline" size="small">搜索</el-button>
      <el-button type="primary" icon="el-icon-refresh" size="small">刷新</el-button>
      </el-button-group>
    </el-col>

    <el-editable ref="editable"
      :data="documentdetaildata" border style="width: 100%" stripe>
      <!-- <el-table-column label="DocID" prop="Document.ID" align="center"></el-table-column> -->
      <!-- <el-table-column label="DocType" prop="Document.DocType.ID" align="center"></el-table-column> -->
      <el-editable-column label="序号" type="index" show-overflow-tooltip width="50"  align="center"></el-editable-column>
      <el-editable-column label="DocState" prop="Document.DocState.Name" align="center"></el-editable-column>
      <el-editable-column label="recip GROUPs" prop="value1" align="center">
        <template slot-scope="scope">
          <el-select v-model="value1" multiple clearable>
            <el-option
              v-for="item in groupdata"
              :key="item.ID"
              :label="item.Name"
              :value="item.ID">
            </el-option>
          </el-select>
        </template>
      </el-editable-column>
      <el-table-column label="ACTIONS" prop="value2" align="center">
        <template slot-scope="scope">
          <el-select v-model="value2" clearable>
            <el-option
              v-for="item in documentdetaildata[0].Action"
              :key="item.ID"
              :label="item.Name"
              :value="item.ID">
            </el-option>
          </el-select>
        </template>
      </el-table-column>
      <el-editable-column label="TITLE" prop="Document.Title" align="center"></el-editable-column>
      <el-editable-column label="PATH" prop="Document.Path" size="mini" align="center"></el-editable-column>
      <el-editable-column label="CTIME" prop="Document.Ctime" size="mini" :formatter="formatter" align="center"></el-editable-column>
      <el-editable-column label="Text" prop="Text" :editRender="{Name: 'ElInput'}" align="center"></el-editable-column> 
      <el-editable-column  label="操作" align="center">
        <template slot-scope="scope">
          <el-button type="primary" size="mini" @click="handleSubmit(scope.$index, scope.row)">Apply</el-button>
          <!-- <el-button size="mini" type="danger" @click="deleteRow(scope.$index, documentdetaildata)">Delete</el-button> -->
        </template>
      </el-editable-column>
    </el-editable>
    <!-- <el-pagination background
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
      :current-page="currentPage"
      :page-sizes="[10, 50, 100, 200]"
      :page-size="pageSize"
      layout="total, sizes, prev, pager, next, jumper"
      :total="total" style="float: right; margin:10px">
    </el-pagination> -->
  <!-- </div> -->
  <el-col :span="24" class="breadcrumb-container">
    <p></p>
  </el-col>
    <!-- <template> -->
    <el-col :span="24" class="breadcrumb-container">
    <!-- <div class="block"> :timestamp="activity.timestamp"-->
      <el-timeline :reverse="reverse">
        <el-timeline-item
          v-for="(item, index) in documentdetaildata[0].History"
          :key='index'> 
          <el-card>
            <h4>Group：{{item.Group}}</h4>
            <p>Comment：{{item.Data}}</p>
            <p>于“ {{item.Ctime | dateformat}} ”将文档状态从“ {{item.FromState}} ”通过动作“ {{item.DocAction}} ”改为“ {{item.ToState}} ”</p>
          </el-card>
        </el-timeline-item>
      </el-timeline>
      <!-- <el-timeline>
      <el-timeline-item timestamp="2018/4/12" placement="top">
        <el-card>
          <h4>更新 Github 模板</h4>
          <p>王小虎 提交于 2018/4/12 20:46</p>
        </el-card>
      </el-timeline-item>
      <el-timeline-item timestamp="2018/4/3" placement="top">
        <el-card>
          <h4>更新 Github 模板</h4>
          <p>王小虎 提交于 2018/4/3 20:46</p>
        </el-card>
      </el-timeline-item>
      <el-timeline-item timestamp="2018/4/2" placement="top">
        <el-card>
          <h4>更新 Github 模板</h4>
          <p>王小虎 提交于 2018/4/2 20:46</p>
        </el-card>
      </el-timeline-item>
      </el-timeline> -->
    </el-col>  
  </div>
</template>

<script type="text/javascript">
/* eslint-disable */
  const axios = require('axios');
  export default { // 这里需要将模块引出，可在其他地方使用
    name: 'documentdetail',
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
          reverse: true,//时间线倒序
          activities2: [{
            content: '支持使用图标',
            timestamp: '2018-04-12 20:46',
            size: 'large',
            type: 'primary',
            icon: 'el-icon-more'
          }, {
            content: '支持自定义颜色',
            timestamp: '2018-04-03 20:46',
            color: '#0bbd87'
          }, {
            content: '支持自定义尺寸',
            timestamp: '2018-04-02 20:46',
            size: 'large'
          }, {
            content: '默认样式的节点',
            timestamp: '2018-04-01 20:46'
          }],
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
          documentdetaildata:[
            {
              Document: {
                ID: 5,
                DocType: {
                  ID: 14,
                  Name: "秘密"
                },
                Path: "",
                AccessContext: {
                  ID: 1
                },
                DocState: {
                  ID: 10,
                  Name: "核定中..."
                },
                Group: {
                  ID: 7,
                  Name: "email1@example.com",
                  GroupType: ""
                },
                Ctime: "2019-03-09T20:09:56Z",
                Title: "布置图啊1",
                Data: "data"
              },
              Action: null,
              History: [
                {
                  FromState: "设计中...",
                  DocAction: "提交设计",
                  ToState: "校核中...",
                  Group: "email4@example.com",
                  Data: "no comments",
                  Ctime: "2019-03-09T18:31:02Z"
                },
                {
                  FromState: "校核中...",
                  DocAction: "提交审查",
                  ToState: "审查中...",
                  Group: "email4@example.com",
                  Data: "校核完毕，提交审查",
                  Ctime: "2019-03-09T19:59:20Z"
                }
              ],
              Text: ""
            }
          ],
          // docid:(()=>{
          //   /* $route.query获取shopid */
          //   return this.$route.query.docid
          //   })(),
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
          docactiondata:[
            {ID:1,Name:"提交设计"},
            {ID:2,Name:"提交校核"}
          ],
          value1:[],
          value2:'',
        };
      },
      mounted:function () {
        this.docaction(this.currentPage);
        this.group(this.currentPage);
        this.documentdetail();
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
                  title:this.ruleForm2.docname,
                  data:this.ruleForm2.docdata,
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
                url: "/api/flownext",//请求地址
                params:{
                  // docaction:row,
                  docid:row.Document.ID,
                  dtid:row.Document.DocType.ID,
                  daid:this.value2,
                  gid:this.value1,
                  text:row.Text
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
                  this.documentdetail();
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
        documentdetail(){
          this.docid = this.$route.query.docid;
          this.dtid = this.$route.query.dtid;
          axios({
            method: 'get',
            url: '/api/flowdocumentdetail',//2.get通过params选项
            params:{
              docid:this.docid,
              dtid:this.dtid
            }
          })
          .then(response => (this.documentdetaildata = response.data))
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
          var date = row.Document.Ctime;
          // console.log(date)
          if (date === undefined) {
            return "";
          }
          // return util.formatDate.format(new Date(date), 'yyyy-MM-dd');
          // this.$moment().format('YYYY-MM-DD HH:mm:ss')
          return this.$moment(date).format("YYYY-MM-DD HH:mm:ss");
          // https://blog.csdn.net/ysq0317/article/details/81089962
          // vue的话，在moment.js的官网里，是给了安装方法的
          // cnpm install moment --save   
          // 然后再入口文件 main.js中导入并使用
          // import moment from 'moment'//导入文件
          // Vue.prototype.$moment = moment;//赋值使用
        },
         getColumnLabel (value) {
          let selectItem = this.groupdata.find(item => item.ID === value)
          return selectItem ? selectItem.Name : null
        },
        getColumnLabel2 (value) {
          let selectItem = this.documentdetaildata[0].Action.find(item => item.ID === value)
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
