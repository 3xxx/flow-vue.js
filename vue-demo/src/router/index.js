import Vue from 'vue'
import Router from 'vue-router'
// import HelloWorld from '@/components/HelloWorld'
import readme from '@/components/readme'
import doctype from '@/components/doctype'
import docstate from '@/components/docstate'
import docaction from '@/components/docaction'
import App from '@/App'
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

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',component: App,//默认子路由不能有name属性，name: 'App',
      children: [
        { path: '/', component: readme, name: 'README' },
        { path: '/doctype', component: doctype, name: 'DOCTYPE' },
        { path: '/docstate', component: docstate, name: 'DOCSTATE' },
        { path: '/docaction', component: docaction, name: 'DOCACTION' },
        { path: '/transition', component: transition, name: 'TRANSITION' },
        { path: '/workflow', component: workflow, name: 'workflow' },
        { path: '/node', component: node, name: 'node' },
        { path: '/accesscontext', component: accesscontext, name: 'accesscontext' },
        { path: '/user', component: user, name: 'user' },
        { path: '/group', component: group, name: 'group' },
        { path: '/userGroup', component: userGroup, name: 'userGroup' },
        { path: '/role', component: role, name: 'role' },
        { path: '/permission', component: permission, name: 'permission' },
        { path: '/groupRole', component: groupRole, name: 'groupRole' },
        { path: '/documents', component: documents, name: 'documents' },
        { path: '/event', component: event, name: 'event' },
        { path: '/documentlist', component: documentlist, name: 'documentlist' },
        { path: '/documentdetail', component: documentdetail, name: 'documentdetail' },
        { path: '/usermailbox', component: usermailbox, name: 'usermailbox' },
        { path: '/groupmailbox', component: groupmailbox, name: 'groupmailbox' },
        // { path: '/documentlist', component: documentlist, name: 'documentlist',
        //   children: [
        //     { path: '/documentdetail', component: documentdetail, name: 'documentdetail' }
        //   ]
        // },
      ]
    }
  ]
})
