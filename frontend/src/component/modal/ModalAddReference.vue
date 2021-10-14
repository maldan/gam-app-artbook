<template>
  <div :class="$style.window">
    <img
      v-if="image"
      :src="image"
      alt=""
      style="width: 100%; max-width: 600px; margin-bottom: 10px"
    />

    <ui-input
      placeholder="Category..."
      style="margin-bottom: 10px"
      v-model="$store.state.modal.data.category"
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
        :disabled="!$store.state.modal.data.imageFile"
      />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  props: {},
  components: {},
  async mounted() {
    document.onpaste = (event) => {
      var items = event.clipboardData?.items || [];
      for (const index in items) {
        var item = items[index];
        if (item.kind === 'file') {
          var blob = item.getAsFile();
          var reader = new FileReader();
          reader.onload = (event) => {
            // @ts-ignore
            this.image = event.target?.result || '';
          };
          this.$store.state.modal.data.imageFile = blob;
          reader.readAsDataURL(blob as any);
        }
      }
    };
  },
  beforeUnmount() {
    document.onpaste = () => {};
  },
  methods: {},
  data() {
    return {
      image: '',
      ///imageFile: null as any,
    };
  },
});
</script>

<style lang="scss" module>
.window {
  min-width: 320px;
}
</style>
