<template>
  <layout name="main">
    <h1>Home Page</h1>
    <router-link to="/about">Go to about page</router-link>
  </layout>
</template>

<script>
import RouterLink from "@/components/RouterLink";
import Layout from "@/components/Layout";

export default {
  name: "home",
  components: {
    RouterLink,
    Layout
  },
  computed: {
    navigations() {
      return [
        { name: "MAIN" },
        { name: "Main1", link: "#", icon: "mail" },
        { name: "Main2", link: "#", icon: "linkify" },
        { name: "APPS" },
        { name: "App1", link: "#", icon: "mail" }
      ];
    }
  }
};
</script>
