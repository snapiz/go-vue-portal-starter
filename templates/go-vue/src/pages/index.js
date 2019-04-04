export default [
  {
    path: "/",
    title: "",
    component: () => import(/* webpackChunkName: 'home' */ "./Home")
  }
];
