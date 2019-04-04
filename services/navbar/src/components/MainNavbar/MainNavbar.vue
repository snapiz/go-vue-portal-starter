<template>
  <b-navbar toggleable="md">
    <router-link class="navbar-brand" to="/">Go-vue-portal-starter</router-link>
    <b-navbar-toggle target="nav_collapse" />
    <b-collapse is-nav id="nav_collapse">
      <b-navbar-nav class="ml-auto">
        <m-nav-item v-if="me" to="/contact">
          {{ $t("mainNav.goToConsole") }}
        </m-nav-item>
        <!-- <m-nav-item-dropdown icon="bell" right no-caret>
          <b-dropdown-item href="#">None</b-dropdown-item>
        </m-nav-item-dropdown>-->
        <SignInMenu v-if="!me" />
        <UserMenu v-if="me" :me="me" />
      </b-navbar-nav>
    </b-collapse>
  </b-navbar>
</template>

<script>
import UserMenu from "@/components/UserMenu";
import SignInMenu from "@/components/SignInMenu";

export default {
  components: {
    UserMenu,
    SignInMenu
  },
  props: {
    me: Object
  }
};
</script>
