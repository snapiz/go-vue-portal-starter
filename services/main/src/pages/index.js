export default [
  // {
  //   path: "/",
  //   title: "pages.home.title",
  //   acl: false,
  //   component: () => import(/* webpackChunkName: 'home' */ "./Home")
  // },
  {
    path: "/",
    redirect: "/contact"
  },
  {
    path: "/logout",
    redirect: "/signin"
  },
  {
    path: "/signin",
    title: "pages.signin.title",
    acl: false,
    navbar: false,
    component: () => import(/* webpackChunkName: 'sign-in' */ "./SignIn")
  },
  {
    path: "/signup",
    title: "pages.signup.title",
    acl: false,
    navbar: false,
    component: () => import(/* webpackChunkName: 'sign-up' */ "./SignUp")
  }
];
