<template>
  <service-layout>
    <template slot="layout-nav">
      <li class="p-2 pl-3 pt-3 font-weight-bold">{{ "name" | t }}</li>
      <router-link-nav-item to="/">
        <font-awesome-icon class="mr-2" icon="tachometer-alt" />
        {{ "pages.home.title" | t }}
      </router-link-nav-item>
      <router-link-nav-item to="/lists">
        <font-awesome-icon class="mr-2" icon="users" />
        {{ "pages.lists.title" | t }}
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
  }
};
</script>
