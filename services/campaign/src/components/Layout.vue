<template>
  <service-layout>
    <template slot="layout-nav">
      <li class="p-2 pl-3 pt-3 font-weight-bold">Campaign</li>
      <router-link-nav-item to="#">
        <font-awesome-icon class="mr-2" icon="user-secret" />
        Item1
      </router-link-nav-item>
    </template>
    <slot />
  </service-layout>
</template>

<script>
import { ServiceLayout } from "@go-vue-portal-starter/vue";

export default {
  components: {
    ServiceLayout
  },
  props: {
    me: Object
  }
};
</script>
