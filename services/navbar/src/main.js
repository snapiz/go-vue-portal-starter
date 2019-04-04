import Vue from "vue";
import Vuelidate from "vuelidate";
import BootstrapVue from "bootstrap-vue";
import singleSpaVue from "single-spa-vue";
import axios from "axios";
import i18next from "i18next";
import LngDetector from "i18next-browser-languagedetector";
import i18nextFetchBackend from "i18next-fetch-backend";
import VueI18Next from "@panter/vue-i18next";
import { library } from "@fortawesome/fontawesome-svg-core";
import {
  faUserCircle,
  faBell,
  faChevronUp,
  faChevronDown,
  faBullhorn,
  faChevronRight,
  faQuestionCircle,
  faSignOutAlt,
  faIdCard,
  faCogs,
  faDatabase
} from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { getUrlParameter } from "@go-vue-portal-starter/vue";
import RouterLink from "@/components/RouterLink";
import NavItem from "@/components/NavItem";
import NavItemDropdown from "@/components/NavItemDropdown";

library.add(
  faUserCircle,
  faBell,
  faChevronUp,
  faChevronDown,
  faBullhorn,
  faDatabase,
  faChevronRight,
  faQuestionCircle,
  faSignOutAlt,
  faIdCard,
  faCogs
);

import "bootstrap/dist/css/bootstrap.css";
import "bootstrap-vue/dist/bootstrap-vue.css";
import "./main.css";
import "@go-vue-portal-starter/vue/dist/index.css";

import { appName } from "./constants";
import App from "./App.vue";

axios.defaults.headers["Content-Type"] = "application/json";

Vue.config.productionTip = false;

Vue.use(Vuelidate);
Vue.use(VueI18Next);
Vue.use(BootstrapVue);

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
Vue.component("font-awesome-icon", FontAwesomeIcon);
Vue.component("router-link", RouterLink);
Vue.component("m-nav-item", NavItem);
Vue.component("m-nav-item-dropdown", NavItemDropdown);

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

  window.addEventListener("message", e => {
    if (e.data.key !== "signin_success") {
      return;
    }

    const { pathname } = window.location;
    const path = getUrlParameter("redirect") || "/";

    if (path === pathname) {
      window.dispatchEvent(
        new CustomEvent("single-spa:routing-event", { detail: { force: true } })
      );
    } else {
      props.singleSpa.navigateToUrl(path);
    }
  });

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
