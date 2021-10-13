<template>
  <div :class="$style.list"></div>
</template>

<script lang="ts">
import Moment from 'moment';
import { defineComponent } from 'vue';
import { RestApi } from '../../util/RestApi';
import Add from './Add.vue';
import Edit from './Edit.vue';

export default defineComponent({
  props: {},
  components: { Add, Edit },
  async mounted() {
    this.refresh();
  },
  methods: {
    async refresh() {
      this.list = await RestApi.art.getList();
    },
    async remove(id: string) {
      if (confirm('Are you sure?')) {
        // await RestApi.todo.delete(id);
      }
      this.refresh();
    },
  },
  data: () => {
    return {
      isAdd: false,
      isEdit: false,
      editId: '',
      list: [] as any[],
    };
  },
});
</script>

<style lang="scss" module></style>
