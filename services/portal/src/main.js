import "whatwg-fetch";
import "bluebird";
import "systemjs/dist/system";
import * as singleSpa from "single-spa";

function pathPrefix(prefix) {
  return function(location) {
    return location.pathname.indexOf(`${prefix}`) === 0;
  };
}

async function loadApp(name, appURL, path) {
  singleSpa.registerApplication(
    name,
    () => System.import(appURL).then(m => m.default),
    typeof path === "string" ? pathPrefix(path) : path
  );
}

async function run() {
  await Promise.all([
    loadApp("navbar", "/navbar/app.js", () => true),
    ...["admin", "me", "contact", "campaign"].map(x => loadApp(x, `/${x}/app.js`, `/${x}`)),
    loadApp("main", "/main/app.js", location =>
      singleSpa
        .getAppNames()
        .reduce(
          (obj, x) =>
            x !== "main" && obj ? !location.pathname.includes(x) : obj,
          true
        )
    )
  ]);

  singleSpa.start();
}

window.addEventListener('single-spa:before-first-mount', () => {
  const el = document.querySelector(".gooey");
  el.parentNode.removeChild(el);
});

run();
