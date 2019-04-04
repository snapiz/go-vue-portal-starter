const { appName } = require("./src/constants");
const { configureWebpack } = require("@go-vue-portal-starter/vue");

module.exports = configureWebpack(appName);
