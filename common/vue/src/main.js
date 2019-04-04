import Vue from "vue";

import BootstrapVue from "bootstrap-vue";
import router from "./router";
import App from "./App.vue";
import VueUniversalRouter from "./VueUniversalRouter";

import "bootstrap/dist/css/bootstrap.css";
import "bootstrap-vue/dist/bootstrap-vue.css";

Vue.config.productionTip = false;

Vue.use(BootstrapVue);
Vue.use(VueUniversalRouter, router);

new Vue({
  render: h => h(App)
}).$mount("#app");
