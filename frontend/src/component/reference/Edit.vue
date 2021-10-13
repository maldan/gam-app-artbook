<template>
  <div :class="$style.container">
    <div :class="$style.window">
      <Input placeholder="Category..." style="margin-bottom: 10px" v-model="category" />
      <Input placeholder="Tags..." style="margin-bottom: 10px" v-model="tags" />

      <div style="display: flex">
        <ui-button @click="$emit('close')" text="Cancel" style="margin-right: 5px" />
        <ui-button @click="submit()" text="Save" icon="add" style="margin-left: 5px" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '../../util/RestApi';
import TextArea from '../TextArea.vue';
import Input from '../Input.vue';

export default defineComponent({
  props: {
    id: String,
  },
  components: { TextArea, Input },
  async mounted() {
    const d = await RestApi.reference.get(this.id + '');
    this.category = d.category;
    this.tags = d.tags.join(', ');
  },
  methods: {
    async submit() {
      await RestApi.reference.update({
        id: this.id + '',
        category: this.category,
        tags: this.tags
          .split(',')
          .map((x: string) => x.trim())
          .filter(Boolean),
      });
      this.$emit('close');
    },
  },
  data() {
    return {
      category: '',
      tags: '',
    };
  },
});
</script>

<style lang="scss" module>
.container {
  position: fixed;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.6);
  z-index: 10;
  display: flex;
  align-items: center;
  justify-content: center;

  .window {
    min-width: 320px;
    background: #444444;
    border-radius: 4px;
    padding: 25px;
    color: #fefefe;
    box-shadow: 0px 1px 6px 2px #00000055;
  }
}
</style>
