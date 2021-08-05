<template>
  <div :class="$style.list">
    <Button
      text="Add image"
      icon="add"
      @click="isAdd = true"
      style="width: 100%; margin-bottom: 10px"
    />

    <!-- List -->
    <div :class="$style.item_list">
      <div v-for="item in item.imageList" :key="item.id" :class="$style.block">
        <img
          @click.stop="moveDown(item.id)"
          class="clickable"
          src="../../asset/move_down.svg"
          alt=""
          style="position: absolute; right: 75px; top: 12px"
          draggable="false"
        />
        <img
          @click.stop="(editId = item.id), (isEdit = true)"
          class="clickable"
          src="../../asset/pencil.svg"
          alt=""
          style="position: absolute; right: 42px; top: 12px"
          draggable="false"
        />
        <img
          @click.stop="remove(item.id)"
          class="clickable"
          src="../../asset/trash.svg"
          alt=""
          style="position: absolute; right: 12px; top: 12px"
          draggable="false"
        />

        <img :src="item.url" style="width: 80%; margin: 0 auto" draggable="false" />
        <div style="display: flex; flex: 1; margin-top: 10px">
          <div style="flex: 1">{{ $root.moment.utc(item.time * 1000).format('HH:mm:ss') }}</div>
          <div style="flex: 1">
            {{ $root.moment(item.created).format('DD MMM YYYY HH:mm:ss') }}
          </div>
        </div>
      </div>
    </div>

    <!-- Modal -->
    <Add :workId="item.id" v-if="isAdd" @close="(isAdd = false), refresh()" />
    <Edit :id="editId" :workId="item.id" v-if="isEdit" @close="(isEdit = false), refresh()" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '../../util/RestApi';
import Button from '../Button.vue';
import Add from './Add.vue';
import Edit from './Edit.vue';

export default defineComponent({
  props: {},
  components: { Button, Add, Edit },
  async mounted() {
    this.refresh();
  },
  methods: {
    async refresh() {
      this.item = await RestApi.work.get(this['$route'].params.id as string);
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
  .item_list {
    .block {
      width: 100%;
      font-size: 15px;
      background: #1b1b1b;
      padding: 10px;
      color: #c5c5c5;
      box-sizing: border-box;
      position: relative;
      margin-bottom: 10px;
      display: flex;
      flex-direction: column;
    }
  }
}
</style>
