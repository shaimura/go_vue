import Vue from 'vue';
import App from './App.vue';
import VuePrlx from 'vue-prlx'
import Loading from 'vue-loading-template';
import axios from 'axios';
import router from './router';



Vue.config.productionTip = false;

axios.defaults.baseURL = 'http://localhost:8888';
// axios.defaults.baseURL = 'http://h-shimu.velvet.jp:8888';
// axios.defaults.baseURL = ':8888';

Vue.use(VuePrlx);

Vue.use(Loading);


axios.interceptors.request.use(
  config => {
    return config;
  },
  error => {
    // "error"の場合は"catch"に行くようにする
    return Promise.reject(error);
  }
);
// サーバーに送った後
axios.interceptors.response.use(
  config => {
    return config;
  },
  error => {
    // "error"の場合は"catch"に行くようにする
    return Promise.reject(error);
  }
);


new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
