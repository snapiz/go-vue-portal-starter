<template>
  <div class="about">
    <h1>About Page</h1>
    <router-link to="/">Go to home page</router-link>
  </div>
</template>

<script>
import RouterLink from "@/components/RouterLink";

export default {
  name: "about",
  components: {
    RouterLink
  }
};
</script>
