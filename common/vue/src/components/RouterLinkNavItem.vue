<template>
  <li class="p-0 nav-item">
    <slot name="pre" />
    <router-link class="d-inline-block w-100 p-2 pl-3" :to="to">
      <slot></slot>
    </router-link>
  </li>
</template>

<script>
import RouterLink from "@/components/RouterLink";

export default {
  name: "router-link-nav-item",
  props: {
    to: String
  },
  components: {
    RouterLink
  }
};
</script>
