import Vue from "vue";
import Vuelidate from "vuelidate";
import BootstrapVue from "bootstrap-vue";
import singleSpaVue from "single-spa-vue";
import i18next from "i18next";
import LngDetector from "i18next-browser-languagedetector";
import i18nextFetchBackend from "i18next-fetch-backend";
import VueI18Next from "@panter/vue-i18next";
import { library } from "@fortawesome/fontawesome-svg-core";
import { faUserSecret } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

import { appName } from "./constants";
import router from "./router";
import App from "./App.vue";
import Layout from "./components/Layout.vue";
import { VueUniversalRouter } from "@go-vue-portal-starter/vue";

import "@go-vue-portal-starter/vue/dist/index.css";

library.add(faUserSecret);

Vue.config.productionTip = false;

Vue.use(VueI18Next);
Vue.use(Vuelidate);
Vue.use(BootstrapVue);
Vue.use(VueUniversalRouter, router);

i18next
  .use(LngDetector)
  .use(i18nextFetchBackend)
  .init({
    whitelist: ["en", "fr"],
    load: "languageOnly",
    fallbackLng: "en",
    backend: {
      loadPath: `/${appName}/locales/{{lng}}.json`
    }
  });

Vue.filter("t", (value, args) => (value && i18next.t(value, args)) || "");
Vue.component("layout", Layout);
Vue.component("font-awesome-icon", FontAwesomeIcon);

const i18n = new VueI18Next(i18next);

const vueLifecycles = singleSpaVue({
  Vue,
  appOptions: {
    el: "#" + appName,
    i18n,
    render: h => h(App)
  }
});

function createDomElement() {
  let el = document.getElementById(appName);

  if (!el) {
    el = document.createElement("div");
    el.id = appName;
    document.body.appendChild(el);
  }
  return el;
}

export const bootstrap = [vueLifecycles.bootstrap];

export async function mount(props) {
  createDomElement();

  Vue.mixin({
    data: function() {
      return {
        props
      };
    }
  });

  return vueLifecycles.mount(props);
}

export const unmount = [vueLifecycles.unmount];
