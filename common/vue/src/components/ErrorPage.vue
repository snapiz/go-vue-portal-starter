<template>
  <div class="error text-center d-flex justify-content-center flex-column">
    <div class="error-code text-primary">500</div>
    <h2 class="error-header">{{ $t("errors.pages.errorHeader") }}</h2>
    <div class="error-desc">
      {{ $t("errors.pages.errorHeader") }}
    </div>
  </div>
</template>

<style scoped>
.error {
  background-color: #ecf2f6 !important;
  height: 100vh;
}
.error .error-code {
  letter-spacing: 4px;
  color: #269af1;
  font-size: 280px;
  line-height: 280px;
  font-weight: 300;
  margin-bottom: 48px;
}
.error .error-header {
  font-size: 40px;
  color: #000;
  padding: 0 15px;
}
.error .error-desc {
  font-size: 20px;
  color: #000;
  padding: 0 15px;
}
</style>

<script>
export default {
  name: "Error"
};
</script>
