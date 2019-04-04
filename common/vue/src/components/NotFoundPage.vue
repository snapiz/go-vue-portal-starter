<template>
  <div class="not-found text-center d-flex justify-content-center flex-column">
    <div class="error-code text-primary">404</div>
    <h2 class="error-header">{{ $t("errors.pages.notFoundHeader") }}</h2>
    <div class="error-desc">
      {{ $t("errors.pages.notFoundDesc") }}
    </div>
  </div>
</template>

<style scoped>
.not-found {
  background-color: #ecf2f6;
  height: 100vh;
}
.not-found .error-code {
  letter-spacing: 4px;
  color: #269af1;
  font-size: 280px;
  line-height: 280px;
  font-weight: 300;
  margin-bottom: 48px;
}
.not-found .error-header {
  font-size: 40px;
  color: #000;
  padding: 0 15px;
}
.not-found .error-desc {
  font-size: 20px;
  color: #000;
  padding: 0 15px;
}
</style>

<script>
export default {
  name: "NotFound"
};
</script>
