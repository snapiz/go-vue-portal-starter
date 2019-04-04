<template>
  <layout>
    <div id="signup">
      <b-card>
        <b-form class="w-signup" @submit.prevent="signup">
          <div class="text-center text-primary mb-4">
            <h5>{{ $t("pages.signup.title") }}</h5>
          </div>
          <b-form-group>
            <b-form-input
              v-model.trim="$v.formData.email.$model"
              :placeholder="$t('pages.signup.email')"
              icon="user"
              @focus="error = ''"
              icon-position="left"
            />
            <b-form-invalid-feedback
              :state="!$v.formData.email.$dirty || $v.formData.email.required"
              >{{ "errors.field.required" | t }}</b-form-invalid-feedback
            >
            <b-form-invalid-feedback
              :state="!$v.formData.email.$dirty || $v.formData.email.email"
              >{{ "errors.field.email" | t }}</b-form-invalid-feedback
            >
          </b-form-group>
          <b-form-group>
            <b-form-input
              :placeholder="$t('pages.signup.username')"
              v-model.trim="$v.formData.username.$model"
              icon="lock"
              icon-position="left"
              @focus="error = ''"
            />
            <b-form-invalid-feedback
              :state="
                !$v.formData.username.$dirty || $v.formData.username.required
              "
              >{{ "errors.field.required" | t }}</b-form-invalid-feedback
            >
            <b-form-invalid-feedback
              :state="
                !$v.formData.username.$dirty ||
                  ($v.formData.username.minLength &&
                    $v.formData.username.maxLength)
              "
              >{{
                "errors.field.between" | t({ min: 3, max: 50 })
              }}</b-form-invalid-feedback
            >
            <b-form-invalid-feedback
              :state="
                !$v.formData.username.$dirty || $v.formData.username.alphaNum
              "
              >{{ "errors.field.alphaNumeric" | t }}</b-form-invalid-feedback
            >
          </b-form-group>
          <b-form-group>
            <b-form-input
              type="password"
              :placeholder="$t('pages.signup.password')"
              v-model.trim="$v.formData.password.$model"
              icon="lock"
              icon-position="left"
              @focus="error = ''"
            />
            <b-form-invalid-feedback
              :state="
                !$v.formData.password.$dirty || $v.formData.password.required
              "
              >{{ "errors.field.required" | t }}</b-form-invalid-feedback
            >
            <b-form-invalid-feedback
              :state="
                !$v.formData.password.$dirty ||
                  ($v.formData.password.minLength &&
                    $v.formData.password.maxLength)
              "
              >{{
                "errors.field.between" | t({ min: 8, max: 20 })
              }}</b-form-invalid-feedback
            >
          </b-form-group>
          <b-form-group>
            <b-form-input
              type="password"
              :placeholder="$t('pages.signup.confirmPassword')"
              @focus="passwordFormError = ''"
              v-model.trim="$v.formData.confirmPassword.$model"
            />
            <b-form-invalid-feedback
              :state="
                !$v.formData.confirmPassword.$dirty ||
                  $v.formData.confirmPassword.custom
              "
              >{{ "errors.field.confirmPassword" | t }}</b-form-invalid-feedback
            >
          </b-form-group>

          <div class="text-danger" v-if="error">{{ error | t }}</div>
          <div class="text-right">
            <b-button type="submit" size="sm" variant="primary">
              {{ "pages.signup.title" | t }}
            </b-button>
          </div>
          <hr />
          <div>
            {{ $t("pages.signup.redirectLogin") }}
            <router-link class="nav-link p-0 d-inline-block" to="/signin">
              {{ $t("pages.signup.redirectAction") }}
            </router-link>
          </div>
        </b-form>
      </b-card>
    </div>
  </layout>
</template>

<style scoped>
#signup {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
}

@media (min-width: 768px) {
  .w-signup {
    width: 350px;
  }
}
</style>

<script>
import {
  required,
  minLength,
  maxLength,
  email,
  alphaNum
} from "vuelidate/lib/validators";

import axios from "axios";

export default {
  name: "signup",
  data: () => ({
    formData: {
      email: "",
      username: "",
      password: "",
      confirmPassword: ""
    },
    error: ""
  }),
  validations() {
    return {
      formData: {
        email: {
          required,
          email
        },
        username: {
          required,
          alphaNum,
          minLength: minLength(3),
          maxLength: maxLength(50)
        },
        password: {
          required,
          minLength: minLength(8),
          maxLength: maxLength(20)
        },
        confirmPassword: {
          custom: v => v === this.formData.password
        }
      }
    };
  },
  async beforeRouteEnter() {
    try {
      await axios.post("/auth/logout");
    } catch (error) {
      console.log("%cError logout", "color: orange;", error.message);
    }
  },
  methods: {
    async signup() {
      this.$v.formData.$touch();

      if (this.$v.formData.$invalid) {
        return;
      }

      try {
        await axios.post("/auth/local/new", this.formData);
        this.$router.context.history.push("/");
      } catch (error) {
        this.error = error.response.data;
      }
    }
  }
};
</script>
