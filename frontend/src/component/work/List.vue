<template>
  <div :class="$style.list">
    <!-- Modal -->
    <Add :workId="item.id" v-if="isAdd" @close="(isAdd = false), refresh()" />
    <Edit :id="editId" :workId="item.id" v-if="isEdit" @close="(isEdit = false), refresh()" />
  </div>
</template>

<script lang="ts">
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
      this.item = await RestApi.art.get(this['$route'].params.id as string);
    },
    async remove(id: string) {
      if (confirm('Are you sure?')) {
        await RestApi.image.delete(this['$route'].params.id as string, id);
      }
      this.refresh();
    },
    async moveDown(id: string) {
      await RestApi.image.moveDown(this['$route'].params.id as string, id);
      this.refresh();
    },
  },
  data: () => {
    return {
      isAdd: false,
      isEdit: false,
      editId: '',
      item: {
        id: '',
        title: '',
        tags: [],
        imageList: [],
      },
    };
  },
});
</script>

<style lang="scss" module>
.list {
  height: calc(100% - 50px);
}
</style>
