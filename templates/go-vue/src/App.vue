<template>
  <div :id="getAppName">
    <AppRender />
  </div>
</template>

<script>
import { AppRender } from "@go-vue-portal-starter/vue";
import { appName } from "./constants";

export default {
  name: "app",
  components: {
    AppRender
  },
  computed: {
    getAppName() {
      return appName;
    }
  }
};
</script>
