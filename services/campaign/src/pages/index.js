export default [
  {
    path: "/",
    title: "Campaign",
    component: () => import(/* webpackChunkName: 'home' */ "./Home")
  }
];
