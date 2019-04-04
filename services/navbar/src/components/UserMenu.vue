<template>
  <m-nav-item-dropdown class="user-menu-dropdown" right>
    <template slot="button-content">
      <b-img
        :src="me.avatar"
        :alt="me.displayName"
        rounded="circle"
        height="28"
        width="28"
      />
    </template>
    <b-dropdown-form is="b-container" class="bg-white p-3">
      <b-row>
        <b-col cols="4">
          <b-img
            :src="me.avatar"
            :alt="me.displayName"
            rounded="circle"
            height="68"
            width="68"
          />
        </b-col>
        <b-col cols="8">
          <span
            :title="me.displayName"
            class="d-inline-block text-truncate w-100 font-weight-bold"
            >{{ me.displayName }}</span
          >
          <router-link class="btn btn-primary btn-sm" to="/me/account">
            {{ $t("userMenu.account") }}
          </router-link>
        </b-col>
      </b-row>
    </b-dropdown-form>
    <ul class="list-unstyled p-2">
      <m-nav-item v-if="isAuthorized(ACL.STAFF)" to="/admin">
        <font-awesome-icon class="mr-1" icon="cogs" />
        {{ $t("userMenu.admin") }}
      </m-nav-item>
      <m-nav-item to="/me">
        <font-awesome-icon class="mr-1" icon="id-card" />
        {{ $t("meNav.profile") }}
      </m-nav-item>
      <m-nav-item to="/#">
        <font-awesome-icon class="mr-1" icon="question-circle" />
        {{ $t("userMenu.help") }}
      </m-nav-item>
      <b-dropdown-divider></b-dropdown-divider>
      <m-nav-item to="/logout">
        <font-awesome-icon class="mr-1" icon="sign-out-alt" />
        {{ $t("userMenu.logout") }}</m-nav-item
      >
    </ul>
  </m-nav-item-dropdown>
</template>

<style>
.user-menu-dropdown .dropdown-menu {
  min-width: 280px;
  padding: 0px;
  background-color: #f3f5f6;
}
</style>

<script>
import { ACL } from "@go-vue-portal-starter/vue";

export default {
  props: {
    me: Object
  },
  data() {
    return {
      ACL
    };
  },
  methods: {
    isAuthorized(acl) {
      return this.me && acl.indexOf(this.me.role) !== -1;
    }
  }
};
</script>
