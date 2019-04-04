<template>
  <service-layout>
    <template slot="layout-nav">
      <router-link-nav-item to="#" v-for="(item, i) in items" :key="i">{{
        item
      }}</router-link-nav-item>
    </template>
    <slot />
  </service-layout>
</template>

<script>
import ServiceLayout from "@/components/ServiceLayout";

export default {
  components: {
    ServiceLayout
  },
  props: {
    me: Object
  },
  data() {
    return {
      items: Array.from({ length: 100 }, () => "Lorem ipsum dolor")
    };
  }
};
</script>
