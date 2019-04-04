<template>
  <b-navbar toggleable="md">
    <router-link class="navbar-brand" to="/">Go-vue-portal-starter</router-link>
    <b-navbar-nav>
      <AdminNav class="d-sm-none" v-if="app === 'admin'" />
      <MeNav class="d-sm-none" v-if="app === 'me'" />
    </b-navbar-nav>
    <b-navbar-toggle target="nav_collapse" />
    <b-collapse is-nav id="nav_collapse">
      <b-navbar-nav>
        <m-nav-item-dropdown
          class="services-dropdown"
          :text="$t('serviceNav.title')"
          right
        >
          <b-dropdown-form is="b-container">
            <b-row>
              <service-navbar-item
                name="serviceNav.storage.title"
                icon="database"
              >
                <m-nav-item to="/contact">{{
                  "serviceNav.storage.contact" | t
                }}</m-nav-item>
              </service-navbar-item>
              <service-navbar-item
                name="serviceNav.marketing.title"
                icon="bullhorn"
              >
                <m-nav-item to="/campaign">{{
                  "serviceNav.marketing.campaign" | t
                }}</m-nav-item>
              </service-navbar-item>
            </b-row>
          </b-dropdown-form>
        </m-nav-item-dropdown>
      </b-navbar-nav>
      <b-navbar-nav class="ml-auto">
        <!-- <m-nav-item-dropdown icon="bell" right no-caret>
          <b-dropdown-item href="#">None</b-dropdown-item>
        </m-nav-item-dropdown>-->
        <SignInMenu v-if="!me" />
        <UserMenu v-if="me" :me="me" />
      </b-navbar-nav>
    </b-collapse>
  </b-navbar>
</template>

<style>
.services-dropdown .nav-link {
  padding-top: 0;
  padding-bottom: 0;
  margin: 0.5rem 0;
}

@media (min-width: 768px) {
  .services-dropdown .dropdown-menu {
    left: 0;
    min-width: 400px;
  }
}
</style>

<script>
import UserMenu from "@/components/UserMenu";
import SignInMenu from "@/components/SignInMenu";
import AdminNav from "@/components/ServiceNavbar/navs/AdminNav";
import MeNav from "@/components/ServiceNavbar/navs/MeNav";
import ServiceNavbarItem from "./ServiceNavbarItem";

export default {
  components: {
    UserMenu,
    SignInMenu,
    AdminNav,
    MeNav,
    ServiceNavbarItem
  },
  props: {
    me: Object,
    app: String
  }
};
</script>
