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

import VueAwesomeSwiper from 'vue-awesome-swiper'
// require styles
import 'swiper/dist/css/swiper.css'

import './assets/icon/iconfont.css'
// import store from '@/store'
// import store from './store'
// import Editable from '@/components/Editable.vue'
// import EditableColumn from '@/components/EditableColumn.vue'
import moment from 'moment'//导入文件
import VueElementExtends from 'vue-element-extends'
import 'vue-element-extends/lib/index.css'
Vue.prototype.$moment = moment;//赋值使用
Vue.filter('dateformat', function(dataStr, pattern = 'YYYY-MM-DD HH:mm:ss') {
    return moment(dataStr).format(pattern)
})
// Vue.component(Editable.name, Editable)
// Vue.component(EditableColumn.name, EditableColumn)

Vue.prototype.$ajax = axios;

// Vue.prototype.$axios = Axios;
// axios.defaults.baseURL = '/api';
axios.defaults.headers.post['Content-Type'] = 'application/json';

Vue.config.productionTip = false;

Vue.use(ElementUI);
Vue.use(VueElementExtends);
Vue.use(VueAwesomeSwiper, /* { default global options } */);
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
