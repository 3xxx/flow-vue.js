import Vue from 'vue'
import Router from 'vue-router'
// import HelloWorld from '@/components/HelloWorld'
// 路由页面导入
import login from '@/components/Login.vue'
import chatme from '@/components/chatme.vue'
import home from '@/components/Home.vue'
import readme from '@/components/readme'
import doctype from '@/components/doctype'
import docstate from '@/components/docstate'
import docaction from '@/components/docaction'
// import App from '@/App'_导致重影
import transition from '@/components/transition'
import workflow from '@/components/workflow'
import node from '@/components/node'
import accesscontext from '@/components/accesscontext'
import user from '@/components/user'
import group from '@/components/group'
import userGroup from '@/components/userGroup'
import role from '@/components/role'
import permission from '@/components/permission'
import groupRole from '@/components/groupRole'
import documents from '@/components/documents'
import event from '@/components/event'
import documentlist from '@/components/documentlist'
import documentdetail from '@/components/documentdetail'
import usermailbox from '@/components/usermailbox'
import groupmailbox from '@/components/groupmailbox'

import cmshome from '@/components/cmshome'
import cmsindex from '@/components/cmsindex'
import projects from '@/components/projects'
import project from '@/components/project'

// 懒加载方式，当路由被访问的时候才加载对应组件
Vue.use(Router)

export default new Router({
  routes: [
    // {
    //   path: '/cms',
    //   // name: '首页',
    //   // redirect: {
    //   //   name: 'cmshome'
    //   // },
    //   children: [      
    //     { path: '/cms', component: cmsindex, name: 'cmsindex' },
    //   //   { path: '/docstate', component: docstate, name: 'DOCSTATE' },
    //   //   { path: '/docaction', component: docaction, name: 'DOCACTION' },
    //   ]
    // },
    {
      path: '/',
      // name: '项目',
      redirect: {
        name: '首页'
      },
      component: cmshome,
      children: [  
        { path: '/', component: cmsindex, name: 'cmsindex' },    
        { path: '/projects', component: projects, name: 'projects' },
        { path: '/project', component: project, name: 'project' },
        // { path: '/docaction', component: docaction, name: 'DOCACTION' },
      ]
    },

    {
      path: '/login',
      name: '登录',
      component: login
    },
    {
      path: '/readme',
      name: '关于',
      component: readme
    },
    {
      path: '/flow/chatme',
      name: '留言',
      component: chatme
    },
    {
      path: '/flow',
      redirect: {
        name: 'home'
      },
      // name: 'readme',//默认子路由不能有name属性，name: 'App',
      component: home,
      children: [
        { path: '/flow', component: readme, name: 'README' },
      ]
    },
    {
      path: '/flow',
      name: 'doctype-state-action',
      component: home,//默认子路由不能有name属性，name: 'App',
      children: [      
        { path: '/flow/doctype', component: doctype, name: 'DOCTYPE' },
        { path: '/flow/docstate', component: docstate, name: 'DOCSTATE' },
        { path: '/flow/docaction', component: docaction, name: 'DOCACTION' },
      ]
    },
    {
      path: '/flow',
      name: 'AccessContext',
      component: home,
      children: [
        { path: '/flow/transition', component: transition, name: 'TRANSITION' },
        { path: '/flow/workflow', component: workflow, name: 'workflow' },
        { path: '/flow/node', component: node, name: 'node' },
        { path: '/flow/accesscontext', component: accesscontext, name: 'accesscontext' },
        { path: '/flow/user', component: user, name: 'user' },
        { path: '/flow/group', component: group, name: 'group' },
        { path: '/flow/userGroup', component: userGroup, name: 'userGroup' },
        { path: '/flow/role', component: role, name: 'role' },
        { path: '/flow/permission', component: permission, name: 'permission' },
        { path: '/flow/groupRole', component: groupRole, name: 'groupRole' },
      ]
    },
    {
      path: '/flow',
      name: 'Documents',
      component: home,
      children: [
        { path: '/flow/documents', component: documents, name: 'documents' },
        { path: '/flow/event', component: event, name: 'event' },
      ]
    },
    {
      path: '/flow',
      name: 'Documentlist',
      component: home,
      children: [
        { path: '/flow/documentlist', component: documentlist, name: 'documentlist' },
        { path: '/flow/documentdetail', component: documentdetail, name: 'documentdetail' },
      ]
    },
    {
      path: '/flow',
      name: 'mailbox',
      component: home,
      children: [
        { path: '/flow/usermailbox', component: usermailbox, name: 'usermailbox' },
        { path: '/flow/groupmailbox', component: groupmailbox, name: 'groupmailbox' },
      ]
    }
  ]
})
