import { ApolloClient } from "apollo-client";
import { InMemoryCache } from "apollo-cache-inmemory";
import { HttpLink } from "apollo-link-http";
import { ApolloLink } from "apollo-link";

import RouterLink from "@/components/RouterLink";
import RouterLinkNavItem from "@/components/RouterLinkNavItem";
import { createRouter } from "./createRouter";

export function install(Vue, options) {
  if (install.installed) return;
  install.installed = true;

  Vue.component("router-link", RouterLink);
  Vue.component("router-link-nav-item", RouterLinkNavItem);

  const {
    routes,
    query,
    errorHandler,
    defaultAcl,
    baseUrl,
    router,
    history
  } = options;

  const apollo = new ApolloClient({
    link: ApolloLink.from([
      new HttpLink({
        uri: `${baseUrl || ""}/graphql`,
        credentials: "same-origin"
      })
    ]),
    cache: new InMemoryCache()
  });

  Vue.prototype.$apollo = apollo;
  Vue.prototype.$router = createRouter(apollo, routes, {
    query,
    errorHandler,
    defaultAcl,
    baseUrl,
    router,
    history
  });
}

export default {
  install
};
