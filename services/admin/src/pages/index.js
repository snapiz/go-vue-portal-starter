export default [
  {
    path: "/",
    title: "pages.home.title",
    component: () => import(/* webpackChunkName: 'home' */ "./Home")
  }
];
