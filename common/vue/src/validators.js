import { helpers } from "vuelidate/lib/validators";

export const alphaNumValidator = helpers.regex("alphaNum", /^[a-zA-Z0-9\s]*$/);
