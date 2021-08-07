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
        @click="show(item.url)"
        class="clickable"
        v-for="(item, i) in list"
        :key="item.id"
        :class="$style.block"
      >
        <div :class="$style.number">{{ list.length - i }}</div>
        <img
          @click.stop="(editId = item.id), (isEdit = true)"
          class="clickable"
          src="../../asset/pencil.svg"
          alt=""
          style="position: absolute; right: 12px; top: 12px"
        />

        <img :class="$style.preview" :src="item.thumbnail" />

        <div :class="$style.title">{{ item.title }}</div>
        <div :class="$style.info">
          <div>
            {{
              $root.moment
                .utc(item.time * 1000)
                .format('H %1 m %2')
                .replace('%1', 'h')
                .replace('%2', 'm')
            }}
          </div>
          <div>{{ $root.moment(item.created).format('DD MMM YYYY') }}</div>
        </div>
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
import Moment from 'moment';
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
      this.list = await RestApi.training.getList();
    },
    async remove(id: string) {
      if (confirm('Are you sure?')) {
        // await RestApi.todo.delete(id);
      }
      this.refresh();
    },
    show(url: string) {
      window.open(url, '_blank');
    },
  },
  data: () => {
    return {
      isAdd: false,
      isEdit: false,
      editId: '',
      list: [] as any[],
      selected: null as any,
    };
  },
});
</script>

<style lang="scss" module>
.list {
  height: calc(100% - 50px);

  .item_list {
    display: grid;
    grid-template-columns: 1fr 1fr 1fr 1fr;
    gap: 10px;
    height: calc(100% - 60px);
    overflow-y: auto;
    grid-auto-rows: max-content;

    .block {
      width: 100%;
      font-size: 15px;
      background: #1b1b1b;
      padding: 10px;
      color: #c5c5c5;
      box-sizing: border-box;
      position: relative;
      height: max-content;

      .preview {
        display: block;
        width: 100%;
      }

      .title,
      .info {
        background: #4a4a4a;
        padding: 5px;
        color: #dadada;
        font-weight: bold;
        margin-top: 10px;
      }

      .number {
        position: absolute;
        left: 15px;
        top: 15px;
        background: #2b2b2bcc;
        padding: 2px 10px;
        border-radius: 4px;
        font-weight: bold;
      }

      .info {
        display: flex;

        > div {
          flex: 1;

          &:first-child {
            color: #00ee0c;
          }

          &:last-child {
            text-align: right;
            color: #fecf00;
          }
        }
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

@media (max-width: 1024px) {
  .list {
    .item_list {
      display: grid;
      grid-template-columns: 1fr 1fr 1fr;
      gap: 10px;
    }
  }
}

@media (max-width: 576px) {
  .list {
    .item_list {
      grid-template-columns: 1fr 1fr;
    }
  }
}
</style>
