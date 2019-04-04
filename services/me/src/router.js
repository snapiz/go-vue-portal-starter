import gql from "graphql-tag";
import { appName } from "./constants";
import pages from "./pages";

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
  query
};
