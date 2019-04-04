<template>
  <m-nav-item-dropdown
    class="admin-nav-dropdown"
    right
    :text="$t('serviceNav.menu')"
  >
    <ul class="list-unstyled p-2">
      <m-nav-item to="/me">{{ "meNav.profile" | t }}</m-nav-item>
      <m-nav-item to="/me/account">{{ "meNav.account" | t }}</m-nav-item>
    </ul>
  </m-nav-item-dropdown>
</template>
<style>
.admin-nav-dropdown .dropdown-menu {
  width: 100%;
  position: fixed;
  top: 60px;
}
</style>

<script>
export default {
  props: {
    me: Object
  }
};
</script>
