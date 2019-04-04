<template>
  <div id="layout">
    <div class="service-layout-sidebar text-light d-none d-md-block">
      <perfect-scrollbar tag="ul" class="list-unstyled">
        <slot name="layout-nav" />
      </perfect-scrollbar>
    </div>
    <perfect-scrollbar tag="main">
      <b-container class="p-4">
        <slot />
      </b-container>
    </perfect-scrollbar>
  </div>
</template>

<style>
.service-layout-sidebar {
  height: calc(100vh - 50px);
  width: 217px;
  position: fixed;
  top: 50px;
  left: 0;
  z-index: 100;
  background-color: #455161;
}

.service-layout-sidebar ul {
  height: calc(100vh - 50px);
}

.service-layout-sidebar li.nav-item:hover {
  background: #374451;
  transition: all 0.3s;
}

.service-layout-sidebar li.nav-item:hover a {
  color: white;
}

.service-layout-sidebar li > a {
  color: #939daa;
  text-decoration: none;
}

main {
  position: fixed;
  width: 100%;
  height: calc(100vh - 50px);
}

@media (min-width: 768px) {
  main {
    left: 217px;
    width: calc(100% - 217px);
  }
}
</style>

<script>
import PerfectScrollbar from "@/components/PerfectScrollbar";
export default {
  props: {
    me: Object
  },
  components: {
    PerfectScrollbar
  }
};
</script>
