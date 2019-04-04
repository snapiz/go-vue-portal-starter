<template>
  <div :id="getAppName" v-waitForT v-if="display">
    <main-navbar id="main-navbar" v-if="isMainApp()" :me="me" />
    <service-navbar
      id="service-navbar"
      v-if="!isMainApp()"
      :me="me"
      :app="currentAppName"
    />
  </div>
</template>

<script>
import idx from "idx";
import MainNavbar from "@/components/MainNavbar/MainNavbar";
import ServiceNavbar from "@/components/ServiceNavbar/ServiceNavbar";
import { appName } from "./constants";

export default {
  name: "app",
  components: {
    MainNavbar,
    ServiceNavbar
  },
  data() {
    return {
      me: null,
      display: false,
      currentAppName: null
    };
  },
  computed: {
    getAppName() {
      return appName;
    }
  },
  created() {
    window.addEventListener("single-spa:component:render", this.render);
    window.addEventListener("go-vue-portal-starter:navbar:me", this.renderMe);
  },
  destroyed() {
    window.removeEventListener("single-spa:component:render", this.render);
    window.removeEventListener("go-vue-portal-starter:navbar:me", this.renderMe);
  },
  methods: {
    render(e) {
      const { singleSpa } = this.props;
      const { route } = e.detail;
      this.me = idx(route, x => x.data.me);
      this.display =
        route.component.name === "error"
          ? false
          : idx(route, x => x.ctx.route.navbar) !== false;

      const { pathname } = window.location;
      this.currentAppName = singleSpa
        .getAppNames()
        .filter(x => pathname.indexOf(`/${x}`) === 0)[0];
    },
    renderMe(e) {
      this.me = e.detail;
    },
    isMainApp() {
      return !this.currentAppName;
    }
  }
};
</script>
