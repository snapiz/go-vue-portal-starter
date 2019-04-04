export default [
  {
    path: "",
    title: "Go vue starter - Home",
    acl: false,
    query: false,
    component: () => import(/* webpackChunkName: 'home' */ "./Home")
  },
  {
    path: "/about",
    title: "Go vue starter - About",
    acl: false,
    query: false,
    component: () => import(/* webpackChunkName: 'home' */ "./About")
  }
];
