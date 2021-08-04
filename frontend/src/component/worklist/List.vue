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

        <div>{{ item.title }}</div>
        <div>{{ item.tags.join(', ') }}</div>
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
      list: [],
    };
  },
});
</script>

<style lang="scss" module>
.list {
  .item_list {
    display: grid;
    grid-template-columns: 1fr 1fr 1fr 1fr 1fr;
    gap: 10px;

    .block {
      width: 100%;
      font-size: 15px;
      background: #1b1b1b;
      padding: 10px;
      color: #c5c5c5;
      box-sizing: border-box;
      position: relative;

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
