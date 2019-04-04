<template>
  <layout>
    <div id="signin">
      <b-card>
        <b-form class="w-signin" @submit.prevent="auth">
          <div class="text-center text-primary mb-4">
            <h5>{{ $t("pages.signin.title") }}</h5>
          </div>
          <b-form-group>
            <b-form-input
              type="text"
              :placeholder="$t('pages.signin.signin')"
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
              :placeholder="$t('pages.signin.password')"
              @focus="error = ''"
              v-model.trim="$v.formData.password.$model"
            />
            <b-form-invalid-feedback
              :state="
                !$v.formData.password.$dirty || $v.formData.password.required
              "
              >{{ "errors.field.required" | t }}</b-form-invalid-feedback
            >
          </b-form-group>
          <div class="text-danger" v-if="error">{{ error | t }}</div>
          <b-form-group class="text-right">
            <router-link class="nav-link" to="#">
              {{ $t("pages.signin.forgotPassword") }}
              <font-awesome-icon class="ml-1" icon="chevron-right" />
            </router-link>
          </b-form-group>
          <b-row>
            <b-col class="text-center">
              <b-img
                :title="$t('pages.signin.signinWith', { name: 'google' })"
                class="cursor-pointer"
                :src="images.Google"
                alt="google"
                height="28"
                @click="signinOauth2('google')"
              />
              <b-img
                :title="$t('pages.signin.signinWith', { name: 'facebook' })"
                class="cursor-pointer ml-2"
                :src="images.Facebook"
                alt="facebook"
                height="28"
                @click="signinOauth2('facebook')"
              />
            </b-col>
            <b-col cols="3"
              ><span class="align-middle"
                >- {{ $t("pages.signin.divider") }} -</span
              ></b-col
            >
            <b-col>
              <b-button type="submit" size="sm" variant="primary">
                {{ $t("pages.signin.title") }}
              </b-button>
            </b-col>
          </b-row>
          <hr />
          <div>
            {{ $t("pages.signin.new") }}
            <router-link class="nav-link p-0 d-inline-block" to="/signup">
              {{ $t("pages.signup.title") }}
            </router-link>
          </div>
        </b-form></b-card
      >
    </div>
  </layout>
</template>

<style>
#signin {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
}

@media (min-width: 768px) {
  .w-signin {
    width: 350px;
  }
}
</style>

<script>
import { required } from "vuelidate/lib/validators";
import { getUrlParameter } from "@go-vue-portal-starter/vue";
import axios from "axios";
import Google from "@/assets/google.png";
import Facebook from "@/assets/facebook.png";

const winProviderOptions = {
  google: "width=452,height=633",
  facebook: "width=580,height=400"
};

export default {
  name: "signin",
  data: () => ({
    images: {
      Google,
      Facebook
    },
    formData: {
      login: "",
      password: ""
    },
    error: ""
  }),
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
  async beforeRouteEnter() {
    try {
      await axios.post("/auth/logout");
    } catch (error) {
      console.log("%cError logout", "color: orange;", error.message);
    }
  },
  methods: {
    async auth() {
      this.$v.formData.$touch();

      if (this.$v.formData.$invalid) {
        return;
      }

      try {
        await axios.post("/auth/local", this.formData);

        this.$router.context.history.push(getUrlParameter("redirect") || "/");
      } catch (error) {
        this.error = error.response.data;
      }
    },
    async signinOauth2(provider) {
      try {
        window.open(`/auth/${provider}`, "", winProviderOptions[provider]);
      } catch (error) {
        console.error(error);
      }
    }
  }
};
</script>
