import gql from "graphql-tag";

export default [
  {
    path: "/",
    title: "components.layout.menu.profile",
    query: gql`
      {
        me {
          displayName
          picture
        }
      }
    `,
    component: () => import(/* webpackChunkName: 'profile' */ "./Profile")
  },
  {
    path: "/account",
    title: "components.layout.menu.account",
    query: gql`
      {
        me {
          displayName
          picture
          username
          hasPassword
        }
      }
    `,
    component: () => import(/* webpackChunkName: 'account' */ "./Account")
  }
];
