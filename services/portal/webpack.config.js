const path = require("path");
const HtmlWebpackPlugin = require("html-webpack-plugin");
const CleanWebpackPlugin = require("clean-webpack-plugin");

const proxy = [
  { name: "main", api: 3002, web: 9002 },
  { name: "navbar", web: 9001 },
  { name: "admin", api: 3003, web: 9003 },
  { name: "me", api: 3004, web: 9004 },
  { name: "contact", api: 3005, web: 9005 },
  { name: "campaign", api: 3006, web: 9006 }
].reduce(
  (obj, x) => {
    const apiUrl = `http://localhost:${x.api}`;
    const webUrl = `http://localhost:${x.web}`;
    if (x.name === "main") {
      obj["/graphql"] = {
        target: apiUrl
      };
    } else if (x.api) {
      obj[`/${x.name}/graphql`] = {
        target: apiUrl,
        pathRewrite: { [`^/${x.name}`]: "" }
      };
    }

    obj[`/${x.name}/app.js`] = {
      target: webUrl,
      pathRewrite: { [`^/${x.name}`]: "" }
    };

    obj[`/${x.name}/**/*.js`] = {
      target: webUrl
    };

    obj[`/${x.name}/assets`] = {
      target: webUrl
    };

    obj[`/${x.name}/locales`] = {
      target: webUrl,
      pathRewrite: { [`^/${x.name}`]: "" }
    };

    return obj;
  },
  {
    "/auth/**": {
      target: "http://localhost:3000"
    }
  }
);

module.exports = {
  mode: "production",

  entry: {
    main: "./src/main.js"
  },

  output: {
    filename: "[name].[hash].js",
    path: path.resolve("dist"),
    publicPath: "/"
  },

  module: {
    rules: [{ parser: { system: false } }]
  },

  plugins: [
    new CleanWebpackPlugin(["dist"]),
    new HtmlWebpackPlugin({
      template: "public/index.html",
      favicon: "public/favicon.ico"
    })
  ],

  performance: { hints: false },
  devServer: {
    historyApiFallback: true,
    watchOptions: { aggregateTimeout: 300, poll: 1000 },
    headers: {
      "Access-Control-Allow-Origin": "*",
      "Access-Control-Allow-Methods": "GET, POST, PUT, DELETE, PATCH, OPTIONS",
      "Access-Control-Allow-Headers":
        "X-Requested-With, content-type, Authorization"
    },
    proxy
  }
};
