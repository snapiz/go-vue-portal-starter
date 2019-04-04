<template>
  <div id="layout">
    <slot></slot>
  </div>
</template>

<script>
export default {
  name: "Layout",
  props: {
    me: Object
  }
};
</script>
