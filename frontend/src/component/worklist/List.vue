<template>
  <div :class="$style.list">
    <Button
      text="Create new project"
      icon="add"
      @click="isAdd = true"
      style="width: 100%; margin-bottom: 10px"
    />

    <!-- List -->
    <div :class="$style.item_list">
      <div
        @click="$router.push(`/work/${item.id}`)"
        class="clickable"
        v-for="item in list"
        :key="item.id"
        :class="$style.block"
      >
        <img
          @click.stop="(editId = item.id), (isEdit = true)"
          class="clickable"
          src="../../asset/pencil.svg"
          alt=""
          style="position: absolute; right: 12px; top: 12px"
        />

        <img :class="$style.preview" v-if="item.imageList[0]" :src="item.imageList[0].thumbnail" />

        <div :class="$style.title">{{ item.title }}</div>
        <div :class="$style.tags">
          <div v-for="x in item.tags" :class="$style.tag" :key="x">{{ x }}</div>
        </div>
      </div>
    </div>

    <!-- Modal -->
    <Add v-if="isAdd" @close="(isAdd = false), refresh()" />
    <Edit :id="editId" v-if="isEdit" @close="(isEdit = false), refresh()" />
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
      this.list = await RestApi.work.getList();
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

<style lang="scss" module>
.list {
  .item_list {
    display: grid;
    grid-template-columns: 1fr 1fr 1fr 1fr;
    gap: 10px;

    .block {
      width: 100%;
      font-size: 15px;
      background: #1b1b1b;
      padding: 10px;
      color: #c5c5c5;
      box-sizing: border-box;
      position: relative;

      .preview {
        display: block;
        width: 100%;
      }

      .title {
        background: #4a4a4a;
        padding: 5px;
        color: #dadada;
        font-weight: bold;
        margin-top: 10px;
      }

      .tags {
        display: flex;
        flex-wrap: wrap;

        .tag {
          background: #00bd0e;
          color: #fefefe;
          padding: 5px 10px;
          border-radius: 2px;
          font-size: 14px;
          font-weight: bold;
          margin-right: 5px;
          margin-top: 10px;
        }
      }
    }
  }
}
</style>
