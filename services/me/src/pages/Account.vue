<template>
  <layout>
    <b-row>
      <b-col offset-md="3" md="6">
        <b-card
          v-if="!username"
          class="mb-4"
          :title="$t('pages.account.setUsername')"
        >
          <b-form class="mt-4" @submit.prevent="setUsername">
            <b-form-group
              :label="$t('pages.account.username')"
              label-for="username"
            >
              <b-form-input
                id="username"
                @focus="usernameFormError = ''"
                v-model.trim="$v.usernameFormData.username.$model"
              />
              <b-form-invalid-feedback
                :state="
                  !$v.usernameFormData.username.$dirty ||
                    $v.usernameFormData.username.required
                "
                >{{ "errors.field.required" | t }}</b-form-invalid-feedback
              >
              <b-form-invalid-feedback
                :state="
                  !$v.usernameFormData.username.$dirty ||
                    $v.usernameFormData.username.alphaNum
                "
                >{{ "errors.field.alphaNumeric" | t }}</b-form-invalid-feedback
              >
              <b-form-invalid-feedback
                :state="
                  !$v.usernameFormData.username.$dirty ||
                    ($v.usernameFormData.username.minLength &&
                      $v.usernameFormData.username.maxLength)
                "
              >
                {{ "errors.field.between" | t({ min: 3, max: 50 }) }}
              </b-form-invalid-feedback>
            </b-form-group>
            <div class="text-danger" v-if="usernameFormError">
              {{ usernameFormError | t }}
            </div>
            <div class="text-right">
              <b-button type="submit" variant="primary">{{
                "pages.account.set" | t
              }}</b-button>
            </div>
          </b-form>
        </b-card>
        <b-card
          class="mb-4"
          :title="
            $t(
              hasPassword
                ? 'pages.account.changePassword'
                : 'pages.account.setPassword'
            )
          "
        >
          <b-form class="mt-4" @submit.prevent="updatePassword">
            <b-form-group
              v-if="hasPassword"
              :label="$t('pages.account.oldPassword')"
              label-for="old-password"
            >
              <b-form-input
                id="old-password"
                type="password"
                @focus="passwordFormError = ''"
                v-model.trim="$v.passwordFormData.oldPassword.$model"
              />
              <b-form-invalid-feedback
                :state="
                  !$v.passwordFormData.oldPassword.$dirty ||
                    $v.passwordFormData.oldPassword.required
                "
                >{{ "errors.field.required" | t }}</b-form-invalid-feedback
              >
              <b-form-invalid-feedback
                :state="
                  !$v.passwordFormData.oldPassword.$dirty ||
                    ($v.passwordFormData.oldPassword.minLength &&
                      $v.passwordFormData.oldPassword.maxLength)
                "
              >
                {{ "errors.field.between" | t({ min: 8, max: 20 }) }}
              </b-form-invalid-feedback>
            </b-form-group>
            <b-form-group
              :label="$t('pages.account.newPassword')"
              label-for="new-password"
            >
              <b-form-input
                id="new-password"
                type="password"
                @focus="passwordFormError = ''"
                v-model.trim="$v.passwordFormData.newPassword.$model"
              />
              <b-form-invalid-feedback
                :state="
                  !$v.passwordFormData.newPassword.$dirty ||
                    $v.passwordFormData.newPassword.required
                "
                >{{ "errors.field.required" | t }}</b-form-invalid-feedback
              >
              <b-form-invalid-feedback
                :state="
                  !$v.passwordFormData.newPassword.$dirty ||
                    ($v.passwordFormData.newPassword.minLength &&
                      $v.passwordFormData.newPassword.maxLength)
                "
              >
                {{ "errors.field.between" | t({ min: 8, max: 20 }) }}
              </b-form-invalid-feedback>
            </b-form-group>
            <b-form-group
              :label="$t('pages.account.confirmPassword')"
              label-for="confirm-password"
            >
              <b-form-input
                id="confirm-password"
                type="password"
                @focus="passwordFormError = ''"
                v-model.trim="$v.passwordFormData.confirmPassword.$model"
              />
              <b-form-invalid-feedback
                :state="
                  !$v.passwordFormData.confirmPassword.$dirty ||
                    $v.passwordFormData.confirmPassword.custom
                "
                >{{
                  "errors.field.confirmPassword" | t
                }}</b-form-invalid-feedback
              >
            </b-form-group>
            <div class="text-danger" v-if="passwordFormError">
              {{ passwordFormError | t }}
            </div>
            <div class="text-right">
              <b-button type="submit" variant="primary">{{
                "pages.account.updatePassword" | t
              }}</b-button>
            </div>
          </b-form>
        </b-card>
      </b-col>
    </b-row>
  </layout>
</template>

<script>
import {
  required,
  minLength,
  maxLength,
  alphaNum
} from "vuelidate/lib/validators";

import gql from "graphql-tag";
import { getGraphQLError } from "@go-vue-portal-starter/vue";

export default {
  name: "account",
  data() {
    const { username, hasPassword } = this.route.data.me;

    return {
      username,
      hasPassword,
      usernameFormData: {
        username: ""
      },
      usernameFormError: "",
      passwordFormData: {
        oldPassword: "",
        newPassword: "",
        confirmPassword: ""
      },
      passwordFormError: ""
    };
  },
  validations() {
    let passwordFormData = {
      newPassword: {
        required,
        minLength: minLength(8),
        maxLength: maxLength(20)
      },
      confirmPassword: {
        custom: v => v === this.passwordFormData.newPassword
      }
    };

    if (this.hasPassword) {
      passwordFormData.oldPassword = {
        required,
        minLength: minLength(8),
        maxLength: maxLength(20)
      };
    }

    return {
      usernameFormData: {
        username: {
          required,
          alphaNum,
          minLength: minLength(3),
          maxLength: maxLength(50)
        }
      },
      passwordFormData
    };
  },
  methods: {
    setUsername() {
      this.$v.usernameFormData.$touch();

      if (this.$v.usernameFormData.$invalid) {
        return;
      }

      const { displayName, picture } = this.route.data.me;

      this.$apollo
        .mutate({
          mutation: gql`
            mutation($input: UpdateUserInput!) {
              updateUser(input: $input) {
                user {
                  username
                }
              }
            }
          `,
          variables: {
            input: {
              ...this.usernameFormData,
              displayName,
              picture: picture || "",
              clientMutationId: `pages_Account_${new Date().getTime()}`
            }
          }
        })
        .then(resp => {
          this.username = resp.data.updateUser.user.username;
          this.error = null;
        })
        .catch(error => {
          this.usernameFormError = getGraphQLError(error);
        });
    },
    updatePassword() {
      this.$v.passwordFormData.$touch();

      if (this.$v.passwordFormData.$invalid) {
        return;
      }

      const {
        oldPassword: currentPassword,
        newPassword: password
      } = this.passwordFormData;

      this.$apollo
        .mutate({
          mutation: gql`
            mutation($input: ChangePasswordInput!) {
              changePassword(input: $input) {
                user {
                  hasPassword
                }
              }
            }
          `,
          variables: {
            input: {
              password,
              currentPassword,
              clientMutationId: `pages_Profile_${new Date().getTime()}`
            }
          }
        })
        .then(resp => {
          this.hasPassword = resp.data.changePassword.user.hasPassword;
          this.passwordFormError = null;
          this.passwordFormData = {
            oldPassword: "",
            newPassword: "",
            confirmPassword: ""
          };

          this.$v.$reset();
        })
        .catch(error => {
          this.passwordFormError = getGraphQLError(error);
        });
    }
  }
};
</script>
