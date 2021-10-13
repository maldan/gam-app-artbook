<template>
  <div :class="$style.window">
    <ui-input
      placeholder="Title..."
      style="margin-bottom: 10px"
      v-model="$store.state.modal.data.title"
    />
    <ui-input
      placeholder="Tags..."
      style="margin-bottom: 10px"
      v-model="$store.state.modal.data.tags"
    />

    <div style="display: flex">
      <ui-button @click="$store.dispatch('modal/close')" text="Cancel" style="margin-right: 5px" />
      <ui-button
        @click="$store.dispatch('modal/ok')"
        text="Add"
        icon="plus"
        style="margin-left: 5px"
      />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  props: {},
  components: {},
  async mounted() {},
  methods: {},
  data() {
    return {};
  },
});
</script>

<style lang="scss" module>
.window {
  min-width: 320px;
}
</style>
