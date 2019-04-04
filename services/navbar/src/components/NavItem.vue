<template>
  <li class="nav-item">
    <router-link class="nav-link" :to="to">
      <slot />
    </router-link>
  </li>
</template>

<script>
export default {
  name: "m-nav-item",
  props: {
    to: String
  }
};
</script>
