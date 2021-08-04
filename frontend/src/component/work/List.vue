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
      <div
        @click="$router.push(`/work/${item.id}`)"
        class="clickable"
        v-for="item in item.imageList"
        :key="item.id"
        :class="$style.block"
      ></div>
    </div>

    <!-- Modal -->
    <Add v-if="isAdd" :workId="item.id" @close="(isAdd = false), refresh()" />
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
      this.item = await RestApi.work.get(this['$route'].params.id as string);
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

      .header {
        display: flex;

        .left,
        .right,
        .left2 {
          padding: 10px 15px;
          background: #0000004d;
          border-radius: 6px 6px 0 0;
          color: #979797;
          font-weight: bold;

          span {
            color: #979797;
            font-style: italic;
            font-weight: 300;
          }
        }

        .left {
          color: #cfda1e;
        }

        .left2 {
          margin-left: 10px;
          color: #1edaab;
          font-weight: bold;
        }

        img {
          margin-left: 15px;
        }

        .right {
          margin-left: 15px;
        }
      }

      .body {
        padding: 10px 15px;
        background: #0000004d;
        border-radius: 0 0 6px 6px;
        color: #979797;
      }

      &:last-child {
        margin-bottom: 0;
      }
    }
  }
}
</style>
