<template>
  <layout>
    <b-row>
      <b-col offset-md="3" md="6">
        <b-card class="mb-4" :title="$t('pages.profile.public')">
          <b-form class="mt-4" @submit.prevent="updateProfile">
            <b-form-group
              :label="$t('pages.profile.displayName')"
              label-for="display-name"
            >
              <b-form-input
                id="display-name"
                @focus="error = ''"
                v-model.trim="$v.formData.displayName.$model"
              />
              <b-form-invalid-feedback
                :state="
                  !$v.formData.displayName.$dirty ||
                    $v.formData.displayName.required
                "
                >{{ "errors.field.required" | t }}</b-form-invalid-feedback
              >
              <b-form-invalid-feedback
                :state="
                  !$v.formData.displayName.$dirty ||
                    $v.formData.displayName.alphaNum
                "
                >{{ "errors.field.alphaNumeric" | t }}</b-form-invalid-feedback
              >
              <b-form-invalid-feedback
                :state="
                  !$v.formData.displayName.$dirty ||
                    ($v.formData.displayName.minLength &&
                      $v.formData.displayName.maxLength)
                "
              >
                {{ "errors.field.between" | t({ min: 3, max: 50 }) }}
              </b-form-invalid-feedback>
            </b-form-group>
            <b-form-group
              :label="$t('pages.profile.avatar')"
              label-for="avatar"
            >
              <b-form-input
                id="avatar"
                @focus="error = ''"
                v-model.trim="$v.formData.picture.$model"
              />
              <b-form-invalid-feedback
                :state="!$v.formData.picture.$dirty || $v.formData.picture.url"
                >{{ "errors.field.url" | t }}</b-form-invalid-feedback
              >
            </b-form-group>
            <div class="text-danger" v-if="error">{{ error | t }}</div>
            <div class="text-right">
              <b-button type="submit" variant="primary">{{
                "pages.profile.update" | t
              }}</b-button>
            </div>
          </b-form>
        </b-card>
      </b-col>
    </b-row>
  </layout>
</template>

<script>
import gql from "graphql-tag";
import { required, minLength, maxLength, url } from "vuelidate/lib/validators";

import { alphaNumValidator as alphaNum, getGraphQLError } from "@go-vue-portal-starter/vue";

export default {
  name: "profile",
  data() {
    const { displayName, picture } = this.route.data.me;

    return {
      formData: {
        displayName,
        picture: picture || ""
      },
      error: ""
    };
  },
  validations: {
    formData: {
      displayName: {
        required,
        alphaNum,
        minLength: minLength(3),
        maxLength: maxLength(50)
      },
      picture: { url }
    }
  },
  methods: {
    updateProfile() {
      this.$v.formData.$touch();

      if (this.$v.formData.$invalid) {
        return;
      }

      this.$apollo
        .mutate({
          mutation: gql`
            mutation($input: UpdateUserInput!) {
              updateUser(input: $input) {
                user {
                  id
                  username
                  displayName
                  picture
                }
              }
            }
          `,
          variables: {
            input: {
              ...this.formData,
              clientMutationId: `pages_Profile_${new Date().getTime()}`
            }
          }
        })
        .then(resp => {
          window.dispatchEvent(
            new CustomEvent("go-vue-portal-starter:navbar:me", {
              detail: {
                ...this.route.data.me,
                ...resp.data.updateUser.user
              }
            })
          );
          this.error = null;
        })
        .catch(error => {
          this.error = getGraphQLError(error);
        });
    }
  }
};
</script>
