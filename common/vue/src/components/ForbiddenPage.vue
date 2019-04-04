<template>
  <div class="forbidden text-center d-flex justify-content-center flex-column">
    <div class="error-code text-primary">403</div>
    <h2 class="error-header">{{ $t("errors.pages.forbiddenHeader") }}</h2>
    <div class="error-desc">
      {{ $t("errors.pages.forbiddenDesc") }}
    </div>
  </div>
</template>

<style scoped>
.forbidden {
  background-color: #ecf2f6;
  height: 100vh;
}
.forbidden .error-code {
  letter-spacing: 4px;
  color: #269af1;
  font-size: 280px;
  line-height: 280px;
  font-weight: 300;
  margin-bottom: 48px;
}
.forbidden .error-header {
  font-size: 40px;
  color: #000;
  padding: 0 15px;
}
.forbidden .error-desc {
  font-size: 20px;
  color: #000;
  padding: 0 15px;
}
</style>

<script>
export default {
  name: "Forbidden"
};
</script>
