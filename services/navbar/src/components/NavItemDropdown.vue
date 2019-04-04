<template>
  <b-nav-item-dropdown ref="dropdown" :right="right" no-caret>
    <template slot="button-content">
      <span v-if="text">{{ text }}</span>
      <font-awesome-icon v-if="icon" :icon="icon" />
      <slot name="button-content" />
      <font-awesome-icon
        v-if="!noCaret"
        class="ml-1"
        :icon="opened ? 'chevron-up' : 'chevron-down'"
      />
    </template>
    <slot />
  </b-nav-item-dropdown>
</template>

<script>
export default {
  props: {
    me: Object,
    right: Boolean,
    noCaret: Boolean,
    icon: String,
    text: String
  },
  data() {
    return {
      opened: false
    };
  },
  mounted() {
    if (this.noCaret) {
      return;
    }

    this.$root.$on("m::dropdown::hide", this.mHide);
    this.$root.$on("bv::dropdown::shown", this.shown);
    this.$root.$on("bv::dropdown::hidden", this.hidden);
  },
  beforeDestroy() {
    this.$root.$off("m::dropdown::hide", this.mHide);
    this.$root.$off("bv::dropdown::shown", this.shown);
    this.$root.$off("bv::dropdown::hidden", this.hidden);
  },
  methods: {
    mHide() {
      this.$refs.dropdown.hide();
    },
    shown(e) {
      if (this.$el.id === e.$el.id) {
        this.opened = true;
      }
    },
    hidden(e) {
      if (this.$el.id === e.$el.id) {
        this.opened = false;
      }
    }
  }
};
</script>
