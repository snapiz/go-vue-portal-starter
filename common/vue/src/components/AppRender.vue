<template>
  <component :is="currentComponent" :route="currentRoute"></component>
</template>

<script>
import { pick } from "lodash";
export default {
  name: "AppRender",
  data: () => ({
    currentComponent: null,
    currentRoute: null
  }),
  created() {
    this.render();

    const {
      context: { history }
    } = this.$router;

    this.unlisten = history.listen(() => {
      this.render();
    });

    this.routingEventFn = e => {
      this.render(e);
    };

    window.addEventListener("single-spa:routing-event", this.routingEventFn);
  },
  destroyed() {
    this.unlisten && this.unlisten();
    this.routingEventFn &&
      window.removeEventListener(
        "single-spa:routing-event",
        this.routingEventFn
      );
  },
  methods: {
    render(e) {
      const { pathname } = window.location;

      if (
        this.currentPathname === pathname &&
        (!e || !e.detail || !e.detail.force)
      ) {
        return;
      }

      if (this.$i18n && !this.$i18n.i18next.isInitialized) {
        window.requestAnimationFrame(() => {
          this.render(e);
        });
        return;
      }

      this.currentPathname = pathname;

      this.$router.resolve(pathname).then(route => {
        if (!route) {
          throw `Unable to resolve ${pathname}`;
        }

        const {
          title,
          component,
          redirect,
          ctx: { history, route: r },
          data
        } = route;

        if (redirect) {
          history.push(redirect);

          return;
        }

        window.dispatchEvent(
          new CustomEvent("single-spa:component:render", { detail: { route } })
        );

        if (title) {
          window.document.title =
            (this.$i18n && this.$i18n.i18next.t(title, data)) || title;
        } else {
          window.document.title = "";
        }

        const promise = component.beforeRouteEnter
          ? component.beforeRouteEnter.bind(this)(route)
          : Promise.resolve("");

        return promise.then(() => {
          this.currentComponent = {
            ...component,
            props: {
              route: Object
            }
          };

          this.currentRoute = {
            data,
            ...pick(r, "query", "queryOptions")
          };
        });
      });
    }
  }
};
</script>
