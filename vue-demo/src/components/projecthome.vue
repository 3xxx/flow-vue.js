<!-- 树状目录  class="block"-->
<template>
	<div>
		<el-container>
    	<!-- <p>使用 scoped slot</p> :default-checked-keys="resourceCheckedKey" :props="resourceProps"-->
    	<el-aside style="width: auto;" class="sidebar-container">
    		<el-tree ref="tree" show-checkbox node-key="id" :data="projtreedata" default-expand-all @node-click="handleNodeClick">
    	    <span class="custom-tree-node" slot-scope="{ node, data }">
    	      <span>{{ node.label }}</span>
    	      <span>
    	        <a class="my-badge">{{data.tags[0]}}</a>
    	      </span>
    	    </span>
    		</el-tree>
			</el-aside>
	
			<el-main class="main-container">
				<!-- <router-view /> -->
        <router-view :tree-id="treeId"></router-view>
			</el-main>
		</el-container> 
	</div>
</template>

<script type="text/javascript">
/* eslint-disable */
  const axios = require('axios');
  export default { // 这里需要将模块引出，可在其他地方使用
    name: 'projecthome',
      data() {
        return {
        	treeId:"",
        	projtreedata: [
          	{
            	id: 26159,
            	label: '一级 1',
            	num: 10,
            	children: [
            		{
            		  id: 4,
            		  label: '二级 1-1',
            		  num: 1,
            		  children: [
            		  	{
            		  	  id: 9,
            		  	  num: 0,
            		  	  label: '三级 1-1-1'
            		  	},
            		  	{
            		  	  id: 10,
            		  	  num: 5,
            		  	  label: '三级 1-1-2'
            		  	}
            			]
            		}
          		]
          	}, 
          	{
            	id: 2,
            	label: '一级 2',
            	num: 0,
            	children: [
            		{
                    id: 5,
                    num: 30,
                    label: '二级 2-1'
            		}, 
            		{
                  id: 6,
                  num: 0,
                  label: '二级 2-2'
           			}
           		]
           	}, 
           	{
            	id: 3,
            	num: 0,
            	label: '一级 3',
            	children:[
            		{
              		id: 7,
              		num: 0,
              		label: '二级 3-1'
            		}, 
            		{
                 id: 8,
                 num: 0,
                 label: '二级 3-2'
            		}
          		]
        		}
					],
					// resourceDialogisShow: false,
    			// resourceCheckedKey: [1,3],//通过接口获取的需要默认展示的数组 [1,3,15,18,...]
    			// resourceProps: {
    			//   children: "subArr",
    			//   label: "chineseName"
    			// },
          total: 50,
          currentPage: 1,
          pageSize: 10,
          projid: 26159,

          /*插入form方法*/
          dialogTableVisible: false,
          dialogFormVisible: false,
          formLabelWidth: '120px',
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
        };
      },

      mounted:function () {
        this.getprojecttree();
        this.treeId=this.projtreedata[0].id;
        console.log(this.projtreedata[0].id);
      },
      methods:{
      	getKey(id) {
        	let that = this;
        	let params = {
        	  roleId: id
        	};
        	that.$get(this.$api.getConfigure(), params, function (res) {
        	  if (res.code == "0000") {
        	    if (res.result) {
        	      let data = res.result;
        	      let arr = [];
        	      for (let item of data) {
        	        arr.push(item.id)
        	      }
        	      that.$refs.tree.setCheckedKeys(arr);//获取已经设置的资源后渲染
        	    }
        	  }
        	});
   			},

   			handleNodeClick(data){
   				// console.log(data.label);
   				this.treeId=data.id;
   				// console.log(this.treeId);
   			},

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
                url: "/flowtype",//请求地址
                params:{
                  name:this.ruleForm2.typename,
                },
                data: {
                  name:this.ruleForm2.typename,
                  // "thirdapp_id":1//请求参数
                }
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

        deleteRow(index, rows) {//删除改行
          rows.splice(index, 1);
        },
        addRow(docstatedata,event){
          docstatedata.push({ ID:'', Name: ''})
        },

        getprojecttree(){
          axios({
            method: 'get',
            url: '/projapi/getprojecttree/'+this.projid,//2.get通过params选项
          })
          .then(response => (this.projtreedata = response.data))
          .catch(function (error) {
            console.log(error);
          });
        },
      }
  };
</script>

<style scoped>
  .block{
    width: 230px;
  }
  .my-badge {
    color: #fff;
    background: #99a9bf;
    padding: 3px 8px;
    font-size: 12px;
    line-height: 12px;
    border-radius: 20px;
    text-align: center;
  }
  .custom-tree-node {
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: space-between;
    font-size: 14px;
    padding-right: 8px;
  }
</style>