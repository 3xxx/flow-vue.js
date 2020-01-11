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

    <vxe-toolbar></vxe-toolbar>

    <vxe-table stripe border ref="xTable"
      :data.sync="documentdetaildata" style="width: 100%" :edit-config="{trigger: 'click', mode: 'cell'}">
      <!-- <el-table-column label="DocID" prop="Document.ID" align="center"></el-table-column> -->
      <!-- <el-table-column label="DocType" prop="Document.DocType.ID" align="center"></el-table-column> -->
      <vxe-table-column title="序号" type="index" show-overflow-tooltip width="50"  align="center"></vxe-table-column>
      <vxe-table-column title="TITLE文档名" field="Document.Title" align="center"></vxe-table-column>
      <vxe-table-column title="DocState" field="Document.DocState.Name" align="center"></vxe-table-column>
      <vxe-table-column title="成果" field="Document.Data" size="mini" align="center">
        <!-- <template slot-scope="scope">
          <a :href="'http://localhost/pdf?id='+scope.row.Document.Data" target="_blank" class="buttonText">{{scope.row.Document.Data}}</a>
        </template> -->
        <!-- <el-button type="primary" icon="el-icon-circle-plus-outline" size="small" @click.native="dialogFormVisible = true">查看</el-button> -->
        <template slot-scope="scope">
          <el-button @click.native.prevent="showattachment(scope.$index, scope.row)" type="text" size="small">查看</el-button>
        </template>
        <!-- 这里是弹框，成果里附件列表 -->
        <!-- <template v-slot="{row}">
          <a href="'/pdf?id='+{{ row.Document.Data }}" target="_black">我是超链接：{{ row.Document.Data }}</a>
        </template> -->
      </vxe-table-column>
      
      <vxe-table-column title="CTIME" field="Document.Ctime" size="mini" :formatter="formatter" align="center"></vxe-table-column>

      <vxe-table-column title="ACTIONS" field="value2" align="center">
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
      </vxe-table-column>
      <vxe-table-column title="recip GROUPs" field="value1" align="center">
        <template slot-scope="scope">
          <el-select v-model="value1" multiple clearable>
            <el-option
              v-for="item in groupdata.groups"
              :key="item.ID"
              :label="item.Name"
              :value="item.ID">
            </el-option>
          </el-select>
        </template>
      </vxe-table-column>
      <!-- <vxe-table-column title="PATH" field="Document.Path" size="mini" align="center"></vxe-table-column> -->
      <vxe-table-column title="Text意见" field="Text" :edit-render="{name: 'input'}" align="center"></vxe-table-column> 
      <vxe-table-column  title="操作" align="center">
        <template slot-scope="scope">
          <el-button type="primary" size="mini" @click="handleSubmit(scope.$index, scope.row)">Apply</el-button>
          <!-- <el-button size="mini" type="danger" @click="deleteRow(scope.$index, documentdetaildata)">Delete</el-button> -->
        </template>
      </vxe-table-column>
    </vxe-table>
  <el-col :span="24" class="breadcrumb-container">
    <p></p>
  </el-col>

    <el-col :span="24" class="breadcrumb-container">
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
    </el-col>
    <el-col>
      <!-- <div>
        <textarea id="code" style="width: 50%;display:none;" rows="11">
        st=>start: 开始|past
        e=>end: 结束
        
        input=>inputoutput: 初设、招标文件；审图意见|past
        design=>operation: 设计|past
        check=>operation: 校核|past
        countersign=>condition: 是否跨专业?|past
        countersignYes=>operation: 专业会签|past
        
        review=>operation: 审查|current
        auditcond=>condition: 是否核定
        audit=>operation: 核定
        approvalcond=>condition: 是否批准
        approval=>operation: 批准
        print=>inputoutput: 出版|future
        
        st->input->design->check->countersign
        countersign(yes,right)->countersignYes->review
        countersign(no,left)->review->auditcond
        auditcond(yes)->audit->approvalcond
        auditcond(no,right)->print
        approvalcond(no,right)->print
        approvalcond(yes)->approval->print->e
        </textarea>
      </div> -->
    <!-- <div><button id="run" type="button" style="display:none;">Run</button></div> -->
    <el-row type="flex" class="row-bg" justify="center">
      <el-col :span="6">
        <div class="grid-content">
          <div class="grid-content" id="canvas" style="width: 50%;"></div>
        </div>
      </el-col> 
    </el-row>   
  </el-col>

    <el-dialog title="pdf附件列表" :visible.sync="dialogFormVisible" center>
      <template>
        <vxe-table stripe border ref="xTable"
      :data.sync="pdfsdata" style="width: 100%" :edit-config="{trigger: 'click', mode: 'cell'}">
        <!-- <el-table :data="pdfsdata" style="width: 100%"> -->
          <vxe-table-column type="index" width="50" align="center"></vxe-table-column align="center">
          <vxe-table-column field="Title" title="名称" width="180" align="center"></vxe-table-column>
          <vxe-table-column field="Link" title="链接" width="180" align="center">
            <template slot-scope="scope">
              <a :href="'http://localhost/'+scope.row.Link" target="_blank" class="buttonText"><i class="el-icon-paperclip"></i></a>
            </template>
          </vxe-table-column>
          <vxe-table-column field="FileSize" title="大小(kb)" align="center"></vxe-table-column>
        </vxe-table>
      </template>
    </el-dialog>

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
          docactiondata:[
            {ID:1,Name:"提交设计"},
            {ID:2,Name:"提交校核"}
          ],
          pdfsdata:[
            {
              "Id": 349,
              "Title": "SL1205FT-586-02 SL02号工作井.PDF",
              "Link": "attachment/20190728测试上传文件名修改/可行性研究/SL1205FT-586-02 SL02号工作井.PDF",
              "FileSize": 2974,
              "Downloads": 0,
              "Created": "2019-07-28T02:00:06.8660349Z",
              "Updated": "2019-07-28T02:00:06.8660349Z"
            }
          ],
          value1:[],
          value2:'',
          messageID:'',
          messageid:'',
        };
      },
      props:['messageId'],//这个来不及传值吧
      mounted:function () {
        const s2 = document.createElement('script');
        s2.type = 'text/javascript';
        s2.src="/static/raphael-min.js"
        document.body.appendChild(s2);
        const s3 = document.createElement('script');
        s3.type = 'text/javascript';
        s3.src="/static/jquery.min.js"
        document.body.appendChild(s3);
        const s4 = document.createElement('script');
        s4.type = 'text/javascript';
        s4.src="/static/flowchart-latest.js"
        document.body.appendChild(s4);
        // insertBefore(newItem,existingItem);
        // 取到路由带过来的参数 ,要分清上个页面里用params还是query传值
        // console.log(this.$route.params.messageid)
        // console.log(this.$route.query.messageid)
        this.messageID=this.$route.query.messageid;
        this.docaction(this.currentPage);
        this.group(this.currentPage);
        this.documentdetail();
        // this.flowchart();
      },
      watch: {
        // 监测路由变化,只要变化了就调用获取路由参数方法将数据存储本组件即可
        // '$route': 'getParams',
        messageId: function (newQuestion) {
          this.projproducts(this.currentPage);
        },
      },
      methods:{
        submitForm(formName) {
          this.$refs[formName].validate((valid) => {
            if (valid) {
              axios({
                method: "POST",//请求方式
                url: "/flowdoc",//请求地址
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
                // console.log(response);
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
              axios({
                method: "POST",//请求方式
                url: "/flownext",//请求地址
                params:{
                  // docaction:row,
                  docid:row.Document.ID,
                  dtid:row.Document.DocType.ID,
                  daid:this.value2,
                  gid:this.value1,
                  text:row.Text,
                  messageid:this.messageID
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
                  this.documentdetail();
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
              //     this.documentdetail();
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
        documentdetail(){
          var that=this
          this.docid = this.$route.query.docid;
          this.dtid = this.$route.query.dtid;
          axios({
            method: 'get',
            url: '/flowdocumentdetail',
            params:{
              docid:this.docid,
              dtid:this.dtid
            }
          })
          .then(function(response){
            that.documentdetaildata = response.data;
            axios.get('/liucheng', {
              params: {
                docstate: that.documentdetaildata[0].Document.DocState.Name
              }
            })
            .then(function (response) {
              console.log(response.data);
              var code=response.data
              var chart;
              if (chart) {
                chart.clean();
              }
              chart = flowchart.parse(code);
              // console.log(chart)
              chart.drawSVG('canvas', {
                // 'x': 30,
                // 'y': 50,
                'line-width': 1,//3,
                'maxWidth': 3,//ensures the flowcharts fits within a certian width
                'line-length': 50,
                'text-margin': 10,
                'font-size': 14,
                'font': 'normal',
                'font-family': 'Helvetica',
                'font-weight': 'normal',
                'font-color': 'black',
                'line-color': 'black',
                'element-color': 'black',
                'fill': 'white',
                'yes-text': 'yes',
                'no-text': 'no',
                'arrow-end': 'block',
                'scale': 1,
                'symbols': {
                  'start': {
                    'font-color': 'red',
                    'element-color': 'green',
                    'fill': 'yellow'
                  },
                  'end':{
                    'class': 'end-element'
                  }
                },
                'flowstate' : {
                  'past' : { 'fill' : '#CCCCCC', 'font-size' : 12},
                  'current' : {'fill' : 'blue', 'font-color' : 'red', 'font-weight' : 'bold'},
                  'future' : { 'fill' : '#FFFF99'},
                  'request' : { 'fill' : 'blue'},
                  'invalid': {'fill' : '#444444'},
                  'approved' : { 'fill' : '#58C4A3', 'font-size' : 12, 'yes-text' : 'APPROVED', 'no-text' : 'n/a' },
                  'rejected' : { 'fill' : '#C45879', 'font-size' : 12, 'yes-text' : 'n/a', 'no-text' : 'REJECTED' }
                }
              });
            $('[id^=sub1]').click(function(){
              alert('info here');
            });
            })
          })
          .catch(function (error) {
            console.log(error);
          });
        },
        doctype(currentPage){
          axios({
            method: 'get',
            url: '/flowtypelist',
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
        docaction(currentPage){
          axios({
            method: 'get',
            url: '/flowactionlist',//2.get通过params选项
            // params:{
            //   page:currentPage
            // }
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
          // var date = row.Document.Ctime;
          // console.log(date)
          if (row === undefined) {
            return "";
          }
          // return util.formatDate.format(new Date(date), 'yyyy-MM-dd');
          // this.$moment().format('YYYY-MM-DD HH:mm:ss')
          return this.$moment(row.cellValue).subtract(8,'hour').format("YYYY-MM-DD HH:mm:ss");
          // https://blog.csdn.net/ysq0317/article/details/81089962
          // vue的话，在moment.js的官网里，是给了安装方法的
          // cnpm install moment --save   
          // 然后再入口文件 main.js中导入并使用
          // import moment from 'moment'//导入文件
          // Vue.prototype.$moment = moment;//赋值使用
        },
        attachmentformatter:function(row, column){
          if (row === undefined) {
            return "";
          }
          return '<a href="/pdf?id='+row.cellValue+'" title="打开pdf" target="_blank"><i class="fa fa-file-pdf-o"></i></a>';
        },
        showattachment:function(index, row){
          // console.log(row.Document.Data)
          axios({
            method: 'get',
            url: '/project/product/pdf/'+row.Document.Data,
          })
          .then(response => (this.pdfsdata = response.data))
          .catch(function (error) {
            console.log(error);
          });
          this.dialogFormVisible = true; 
        },
        getColumnLabel (value) {
          let selectItem = this.groupdata.groups.find(item => item.ID === value)
          return selectItem ? selectItem.Name : null
        },
        getColumnLabel2 (value) {
          let selectItem = this.documentdetaildata[0].Action.find(item => item.ID === value)
          return selectItem ? selectItem.Name : null
        },
        // flowchart (){
        //   axios.get('/liucheng', {
        //     params: {
        //       docstate: this.documentdetaildata[0].Document.DocState.Name
        //     }
        //   })
        //   .then(function (response) {
        //     console.log(response.data);
        //     var code=response.data
        //   // 'st=>start: 开始|past\n'+
        //   // 'e=>end: 结束\n'+
          
        //   // 'input=>inputoutput: 初设、招标文件；审图意见|past\n'+
        //   // 'design=>operation: 设计|past\n'+
        //   // 'check=>operation: 校核|past\n'+
        //   // 'countersign=>condition: 是否跨专业?|past\n'+
        //   // 'countersignYes=>operation: 专业会签|past\n'+
          
        //   // 'review=>operation: 审查|current\n'+
        //   // 'auditcond=>condition: 是否核定\n'+
        //   // 'audit=>operation: 核定\n'+
        //   // 'approvalcond=>condition: 是否批准\n'+
        //   // 'approval=>operation: 批准\n'+
        //   // 'print=>inputoutput: 出版|future\n'+
          
        //   // 'st->input->design->check->countersign\n'+
        //   // 'countersign(yes,right)->countersignYes->review\n'+
        //   // 'countersign(no,left)->review->auditcond\n'+
        //   // 'auditcond(yes)->audit->approvalcond\n'+
        //   // 'auditcond(no,right)->print\n'+
        //   // 'approvalcond(no,right)->print\n'+
        //   // 'approvalcond(yes)->approval->print->e\n'
        //   //**********************************
        //       var chart;
        //       if (chart) {
        //         chart.clean();
        //       }
        //       chart = flowchart.parse(code);
        //       // console.log(chart)
        //       chart.drawSVG('canvas', {
        //         // 'x': 30,
        //         // 'y': 50,
        //         'line-width': 1,//3,
        //         'maxWidth': 3,//ensures the flowcharts fits within a certian width
        //         'line-length': 50,
        //         'text-margin': 10,
        //         'font-size': 14,
        //         'font': 'normal',
        //         'font-family': 'Helvetica',
        //         'font-weight': 'normal',
        //         'font-color': 'black',
        //         'line-color': 'black',
        //         'element-color': 'black',
        //         'fill': 'white',
        //         'yes-text': 'yes',
        //         'no-text': 'no',
        //         'arrow-end': 'block',
        //         'scale': 1,
        //         'symbols': {
        //           'start': {
        //             'font-color': 'red',
        //             'element-color': 'green',
        //             'fill': 'yellow'
        //           },
        //           'end':{
        //             'class': 'end-element'
        //           }
        //         },
        //         'flowstate' : {
        //           'past' : { 'fill' : '#CCCCCC', 'font-size' : 12},
        //           'current' : {'fill' : 'blue', 'font-color' : 'red', 'font-weight' : 'bold'},
        //           'future' : { 'fill' : '#FFFF99'},
        //           'request' : { 'fill' : 'blue'},
        //           'invalid': {'fill' : '#444444'},
        //           'approved' : { 'fill' : '#58C4A3', 'font-size' : 12, 'yes-text' : 'APPROVED', 'no-text' : 'n/a' },
        //           'rejected' : { 'fill' : '#C45879', 'font-size' : 12, 'yes-text' : 'n/a', 'no-text' : 'REJECTED' }
        //         }
        //       });
        //     $('[id^=sub1]').click(function(){
        //       alert('info here');
        //     });
        //   })
        //   .catch(function (error) {
        //     console.log(error);
        //   })
        //   .finally(function () {
        //     // always executed
        //   });  
        // },
      }
  };
  
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .end-element { fill : #FFCCFF; }

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
