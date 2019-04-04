<template>
  <b-col cols="12" sm="6">
    <ul class="p-0 list-unstyled">
      <li class="service-navbar-item-header">
        <font-awesome-icon class="mr-2" :icon="icon" />
        <span>{{ name | t }}</span>
      </li>
      <slot />
    </ul>
  </b-col>
</template>

<style scoped>
.service-navbar-item-header {
  font-weight: bold;
  color: #455161;
}
</style>

<script>
export default {
  props: {
    name: String,
    icon: String
  }
};
</script>
