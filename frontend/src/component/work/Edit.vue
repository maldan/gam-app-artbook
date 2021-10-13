<template>
  <div :class="$style.container">
    <div :class="$style.window">
      <Input placeholder="Time..." style="margin-bottom: 10px" v-model="time" />
      <Input placeholder="Created..." style="margin-bottom: 10px" v-model="created" />

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
import Moment from 'moment';

export default defineComponent({
  props: {
    id: String,
    workId: String,
  },
  components: { TextArea, Input },
  async mounted() {
    const d = await RestApi.image.get(this.workId + '', this.id + '');
    this.time = Moment.utc(d.time * 1000).format('HH:mm');
    this.created = Moment(d.created).format('YYYY-MM-DD HH:mm:ss');
  },
  methods: {
    async submit() {
      await RestApi.image.update({
        id: this.id + '',
        workId: this.workId + '',
        time: Moment.duration(this.time + ':00').asSeconds(),
        created: this.created,
      });
      this.$emit('close');
    },
  },
  data() {
    return {
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
