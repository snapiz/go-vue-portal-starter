<style>
html,
body {
  height: 100%;
}
body {
  background-color: #fafafa;
}
</style>

<template>
  <AppRender />
</template>

<script>
import AppRender from "./components/AppRender";

export default {
  name: "app",
  components: {
    AppRender
  }
};
</script>
