<template>
  <b-nav-item-dropdown
    icon="user-circle"
    class="signin-menu-dropdown"
    ref="ddown"
    right
    no-caret
  >
    <template slot="button-content">
      <font-awesome-icon icon="user-circle" />
      <font-awesome-icon
        class="ml-1"
        :icon="opened ? 'chevron-up' : 'chevron-down'"
      />
    </template>
    <b-dropdown-form class="p-3" @submit.prevent="signin">
      <div class="text-center text-primary mb-4">
        <h5>{{ $t("signinMenu.title") }}</h5>
      </div>
      <b-form-group>
        <b-form-input
          type="text"
          :placeholder="$t('signinMenu.signin')"
          @focus="error = ''"
          v-model.trim="$v.formData.login.$model"
        />
        <b-form-invalid-feedback
          :state="!$v.formData.login.$dirty || $v.formData.login.required"
          >{{ "errors.field.required" | t }}</b-form-invalid-feedback
        >
      </b-form-group>
      <b-form-group>
        <b-form-input
          type="password"
          :placeholder="$t('signinMenu.password')"
          @focus="error = ''"
          v-model.trim="$v.formData.password.$model"
        />
        <b-form-invalid-feedback
          :state="!$v.formData.password.$dirty || $v.formData.password.required"
          >{{ "errors.field.required" | t }}</b-form-invalid-feedback
        >
      </b-form-group>
      <div class="text-danger" v-if="error">{{ error | t }}</div>
      <b-form-group class="text-right">
        <router-link class="nav-link" to="#">
          {{ $t("signinMenu.forgotPassword") }}
          <font-awesome-icon class="ml-1" icon="chevron-right" />
        </router-link>
      </b-form-group>
      <b-row>
        <b-col class="text-center">
          <b-img
            :title="$t('signinMenu.signinWith', { name: 'google' })"
            class="cursor-pointer"
            :src="images.Google"
            alt="google"
            height="28"
            @click="signinOauth2('google')"
          />
          <b-img
            :title="$t('signinMenu.signinWith', { name: 'facebook' })"
            class="cursor-pointer ml-2"
            :src="images.Facebook"
            alt="facebook"
            height="28"
            @click="signinOauth2('facebook')"
          />
        </b-col>
        <b-col cols="3" class="d-none d-sm-inline"
          ><span class="align-middle"
            >- {{ $t("signinMenu.divider") }} -</span
          ></b-col
        >
        <b-col>
          <b-button type="submit" size="sm" variant="primary">
            {{ $t("signinMenu.title") }}
          </b-button>
        </b-col>
      </b-row>
      <hr />
      <div>
        {{ $t("signinMenu.new") }}
        <router-link class="nav-link p-0 d-inline-block" to="/signup">
          {{ $t("signinMenu.signup") }}
        </router-link>
      </div>
    </b-dropdown-form>
  </b-nav-item-dropdown>
</template>

<style>
@media (min-width: 768px) {
  .signin-menu-dropdown .dropdown-menu {
    min-width: 370px;
    padding: 0px;
  }

  .signin-menu-dropdown .separator {
    top: 4px;
  }
}
</style>

<script>
import axios from "axios";
import { required } from "vuelidate/lib/validators";
import { ACL } from "@go-vue-portal-starter/vue";
import Google from "@/assets/google.png";
import Facebook from "@/assets/facebook.png";

const winProviderOptions = {
  google: "width=452,height=633",
  facebook: "width=580,height=400"
};

export default {
  props: {
    me: Object
  },
  validations: {
    formData: {
      login: {
        required
      },
      password: {
        required
      }
    }
  },

  data() {
    return {
      ACL,
      images: {
        Google,
        Facebook
      },
      opened: false,
      formData: {
        login: "",
        password: ""
      },
      error: ""
    };
  },
  mounted() {
    this.$root.$on("bv::dropdown::shown", e => {
      if (this.$el.id === e.$el.id) {
        this.opened = true;
      }
    });

    this.$root.$on("bv::dropdown::hidden", e => {
      if (this.$el.id === e.$el.id) {
        this.opened = false;
      }
    });
  },
  methods: {
    async signin() {
      this.$v.formData.$touch();

      if (this.$v.formData.$invalid) {
        return;
      }

      try {
        await axios.post("/auth/local", this.formData);

        this.$refs.ddown.hide(true);

        window.dispatchEvent(
          new CustomEvent("single-spa:routing-event", {
            detail: { force: true }
          })
        );
      } catch (error) {
        this.error = error.response.data;
      }
    },
    signinOauth2(provider) {
      this.$refs.ddown.hide(true);
      window.open(`/auth/${provider}`, "", winProviderOptions[provider]);
    }
  }
};
</script>
