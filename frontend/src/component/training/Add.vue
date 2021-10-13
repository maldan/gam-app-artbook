<template>
  <div :class="$style.container">
    <div :class="$style.window">
      <img
        v-if="image"
        :src="image"
        alt=""
        style="width: 100%; max-width: 600px; margin-bottom: 10px"
      />

      <Input placeholder="Title..." style="margin-bottom: 10px" v-model="title" />
      <Input placeholder="Tags..." style="margin-bottom: 10px" v-model="tags" />
      <Input placeholder="Time..." style="margin-bottom: 10px" v-model="time" />
      <Input placeholder="Created..." style="margin-bottom: 10px" v-model="created" />

      <div style="display: flex">
        <ui-button @click="$emit('close')" text="Cancel" style="margin-right: 5px" />
        <ui-button @click="submit()" text="Add" icon="add" style="margin-left: 5px" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '../../util/RestApi';
import TextArea from '../TextArea.vue';
import Input from '../Input.vue';
import Moment from 'moment';

export default defineComponent({
  props: {
    date: Object,
  },
  components: { TextArea, Input },
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
          this.imageFile = blob;
          reader.readAsDataURL(blob as any);
        }
      }
    };
  },
  beforeUnmount() {
    document.onpaste = () => {};
  },
  methods: {
    async submit() {
      await RestApi.training.add(
        this.title,
        this.tags
          .split(',')
          .map((x: string) => x.trim())
          .filter(Boolean),
        Moment.duration(this.time + ':00').asSeconds(),
        this.created,
        this.imageFile,
      );
      this.$emit('close');
    },
  },
  data() {
    return {
      image: '',
      imageFile: null as any,

      title: '',
      tags: '',
      time: '00:00',
      created: Moment().format('YYYY-MM-DD HH:mm:ss'),
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
