/* eslint-disable */
// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App.vue'
import router from './router'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
// import Axios from 'axios'
import axios from 'axios'
axios.defaults.withCredentials=true;//让ajax携带cookie

import VueAwesomeSwiper from 'vue-awesome-swiper'
// require styles
import 'swiper/dist/css/swiper.css'

import './assets/icon/iconfont.css'
// import store from '@/store'
// import store from './store'
// import Editable from '@/components/Editable.vue'
// import EditableColumn from '@/components/EditableColumn.vue'
import moment from 'moment'//导入文件

// import VueElementExtends from 'vue-element-extends'
// import 'vue-element-extends/lib/index.css'

import 'xe-utils'
import VXETable from 'vxe-table'
import 'vxe-table/lib/index.css'
import VXETablePluginElement from 'vxe-table-plugin-element'
import 'vxe-table-plugin-element/dist/style.css'

Vue.prototype.$moment = moment;//赋值使用
Vue.filter('dateformat', function(dataStr, pattern = 'YYYY-MM-DD HH:mm:ss') {
    return moment(dataStr).subtract(8,'hour').format(pattern)
})
// Vue.component(Editable.name, Editable)
// Vue.component(EditableColumn.name, EditableColumn)

Vue.prototype.$ajax = axios;
// Vue.prototype.$axios = axios;

// axios.defaults.baseURL = '/api';
axios.defaults.headers.post['Content-Type'] = 'application/json';

//https://blog.csdn.net/wuyan1001/article/details/84840703
var root = process.env.API_ROOT;
//请求拦截
axios.interceptors.request.use((config) => {
  //请求之前重新拼装url
  config.url = root + config.url;
  return config;
});

Vue.config.productionTip = false;

Vue.use(ElementUI);
// Vue.use(VueElementExtends);
Vue.use(VXETable);
VXETable.use(VXETablePluginElement);
Vue.use(VueAwesomeSwiper);/* , { default global options } */
// Vue.use(Element, {
//   size: Cookies.get('size') || 'medium', // set element-ui default size
//   i18n: (key, value) => i18n.t(key, value)
// })

new Vue({
  el: '#app',
  render: h => h(App),
  router,
  // store,
  components: { App },
  template: '<App/>'
})
