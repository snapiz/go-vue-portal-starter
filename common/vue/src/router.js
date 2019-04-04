import gql from "graphql-tag";
import pages from "./pages";
import ErrorPage from "./pages/Error.vue";

const query = gql`
  {
    me {
      id
      displayName
      picture
      role
    }
  }
`;

function errorHandler() {
  return {
    title: "Error",
    component: ErrorPage
  };
}
export default {
  routes: [...pages],
  query,
  errorHandler
};
