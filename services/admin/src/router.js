import gql from "graphql-tag";
import { appName } from "./constants";
import pages from "./pages";
import { ACL } from "@go-vue-portal-starter/vue";

const query = gql`
  {
    me {
      id
      displayName
      avatar
      role
    }
  }
`;

export default {
  routes: [...pages],
  baseUrl: "/" + appName,
  defaultAcl: ACL.STAFF,
  query
};
