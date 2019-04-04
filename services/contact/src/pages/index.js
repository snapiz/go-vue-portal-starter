export default [
  {
    path: "/",
    title: "description",
    component: () => import(/* webpackChunkName: 'home' */ "./Home")
  },
  {
    path: "/lists",
    title: "pages.lists.title",
    component: () => import(/* webpackChunkName: 'lists' */ "./Lists")
  }
];
