import UniversalRouter from "universal-router";
import { pick } from "lodash";
import idx from "idx";
import createHistory from "history/createBrowserHistory";
import { mergeQueries } from "./utils";
import ErrorPage from "@/components/ErrorPage";
import NotFoundPage from "@/components/NotFoundPage";
import ForbiddenPage from "@/components/ForbiddenPage";
import BlankPage from "@/components/BlankPage";

export const ACL = {
  ADMIN: "ADMIN",
  STAFF: ["STAFF", "ADMIN"],
  USER: ["ADMIN", "STAFF", "USER"]
};

export function createRouter(apollo, routes, options) {
  const {
    router,
    history: _history,
    ForbiddenPage: _ForbiddenPage,
    NotFoundPage: _NotFoundPage,
    ErrorPage: _ErrorPage,
    query,
    defaultAcl,
    baseUrl
  } = options;

  function createHref(to) {
    const basename = (baseUrl || "/").replace(/\/$/i, "");
    if (basename && to === "/") {
      to = "";
    }

    return `${basename}${to}`;
  }

  UniversalRouter.prototype.createHref = createHref;

  const _routes = routes.map(x => ({
    ...x,
    path: createHref(x.path),
    redirect: x.redirect && createHref(x.redirect)
  }));

  _routes.push({
    path: "(.*)",
    title: "errors.http.404",
    acl: false,
    query: false,
    navbar: false,
    component: _NotFoundPage || NotFoundPage
  });

  function resolveRoute(ctx) {
    const { route, next } = ctx;
    const acl =
      typeof route.acl === "undefined" ? defaultAcl || ACL.USER : route.acl;

    if (typeof route.children === "function") {
      return route.children().then(x => {
        route.children = x.default;

        return next();
      });
    }

    if (route.redirect) {
      return {
        redirect: route.redirect,
        ctx
      };
    }

    if (typeof route.title === "undefined") {
      return next();
    }

    if (window.opener) {
      window.opener.postMessage({ key: "signin_success" });
      window.close();

      return {
        component: BlankPage,
        data: {},
        title: "",
        ctx
      };
    }

    let componentPromise = route.component ? route.component : null;

    if (typeof componentPromise === "object") {
      componentPromise = Promise.resolve(componentPromise);
    } else if (typeof componentPromise === "function") {
      componentPromise = route.component().then(x => x.default);
    }

    const dataPromise =
      route.query !== false
        ? apollo.query({
            query: mergeQueries(query, route.query),
            fetchPolicy: "no-cache",
            ...pick(
              route.queryOptions || {},
              "children",
              "variables",
              "pollInterval",
              "notifyOnNetworkStatusChange",
              "fetchPolicy",
              "errorPolicy",
              "ssr",
              "displayName",
              "skip",
              "onCompleted",
              "onError",
              "context",
              "partialRefetch"
            )
          })
        : Promise.resolve({ me: {} });

    return Promise.all([componentPromise, dataPromise]).then(
      ([component, resp]) => {
        const { title, beforeRender } = route;
        const { data } = resp;

        if (acl) {
          const role = idx(resp, x => x.data.me.role);
          const redirectLogin = encodeURIComponent(route.path || "/");

          if (!role) {
            return {
              redirect: `/signin?redirect=${redirectLogin}`,
              ctx
            };
          }

          if (acl.indexOf(role) === -1) {
            ctx.route.navbar = false;

            return {
              component: _ForbiddenPage || ForbiddenPage,
              data: {},
              title: "errors.http.403",
              ctx
            };
          }
        }

        if (!component) {
          return next();
        }

        if (beforeRender) {
          beforeRender({ data, ctx });
        }

        return { component, data, title, ctx };
      }
    );
  }

  function errorHandler(error, ctx) {
    ctx.route.navbar = false;

    return {
      component: _ErrorPage || ErrorPage,
      data: {},
      title: "errors.http.500",
      ctx
    };
  }

  const history = createHistory(_history);

  return new UniversalRouter(_routes, {
    resolveRoute,
    errorHandler,
    context: { history, apollo },
    ...router
  });
}
