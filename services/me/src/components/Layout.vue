<template>
  <service-layout>
    <template slot="layout-nav">
      <li class="p-2 pl-3 pt-3 font-weight-bold">
        {{ "components.layout.menu.title" | t }}
      </li>
      <router-link-nav-item to="/">
        <font-awesome-icon class="mr-2" icon="id-card" />
        {{ "components.layout.menu.profile" | t }}
      </router-link-nav-item>
      <router-link-nav-item to="/account">
        <font-awesome-icon class="mr-2" icon="user-cog" />
        {{ "components.layout.menu.account" | t }}
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
