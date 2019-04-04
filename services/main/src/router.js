import gql from "graphql-tag";
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
  query
};
