<template>
  <div :class="$style.container">
    <div :class="$style.window">
      <Input placeholder="Priotiry..." style="margin-bottom: 10px" v-model="priority" />
      <TextArea placeholder="Description..." style="margin-bottom: 10px" v-model="description" />
      <Input placeholder="Created..." style="margin-bottom: 10px" v-model="created" />
      <Input placeholder="Deadline..." style="margin-bottom: 10px" v-model="deadline" />

      <div style="display: flex">
        <Button @click="$emit('close')" text="Cancel" style="margin-right: 5px" />
        <Button @click="submit()" text="Save" icon="add" style="margin-left: 5px" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '../../util/RestApi';
import Button from '../Button.vue';
import TextArea from '../TextArea.vue';
import Input from '../Input.vue';
import Moment from 'moment';

export default defineComponent({
  props: {
    id: String,
    date: Object,
  },
  components: { Button, TextArea, Input },
  async mounted() {
    /*const d = await RestApi.todo.get(this.id + '');
    this.priority = d.priority + '';
    this.description = d.description + '';
    this.created = Moment(d.created).format('YYYY-MM-DD HH:mm:ss');
    this.deadline = Moment(d.deadline).format('YYYY-MM-DD HH:mm:ss');*/
  },
  methods: {
    async submit() {
      /*await RestApi.todo.update({
        id: this.id + '',
        priority: Number(this.priority),
        description: this.description,
        created: this.created,
        deadline: this.deadline,
      });
      this.$emit('close');*/
    },
  },
  data() {
    return {
      priority: '',
      description: '',
      created: '',
      deadline: Moment().format('YYYY-MM-DD HH:mm:ss'),
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
