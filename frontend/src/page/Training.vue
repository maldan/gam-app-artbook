<template>
  <div :class="$style.main">
    <ui-button
      text="Add image"
      icon="plus"
      @click="
        $store.dispatch('modal/show', {
          name: 'addTraining',
          data: {
            imageFile: null,
            title: '',
            tags: '',
            time: '00:00',
            created: $root.moment().format('YYYY-MM-DD HH:mm:ss'),
          },
          onSuccess() {
            $store.dispatch('training/add');
          },
        })
      "
      style="width: 100%; margin-bottom: 10px"
    />

    <!-- List -->
    <div :class="$style.item_list">
      <div
        @click="show(item.url)"
        class="clickable"
        v-for="(item, i) in $store.state.training.list"
        :key="item.id"
        :class="$style.block"
      >
        <div :class="$style.number">{{ $store.state.training.list.length - i }}</div>
        <img
          @click.stop="
            $store.dispatch('modal/show', {
              name: 'approve',
              data: {
                id: item.id,
                title: 'Remove?',
              },
              onSuccess() {
                $store.dispatch('training/delete');
              },
            })
          "
          class="clickable"
          src="../asset/trash.svg"
          alt=""
          style="position: absolute; right: 15px; top: 15px"
        />
        <img
          @click.stop="
            $store.dispatch('modal/show', {
              name: 'editTraining',
              data: {
                id: item.id,
                title: item.title,
                tags: item.tags.join(', '),
                time: $root.moment.utc(item.time * 1000).format('HH:mm'),
                created: $root.moment(item.created).format('YYYY-MM-DD HH:mm:ss'),
              },
              onSuccess() {
                $store.dispatch('training/update');
              },
            })
          "
          class="clickable"
          src="../asset/pencil.svg"
          alt=""
          style="position: absolute; right: 40px; top: 15px"
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
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  components: {},
  async mounted() {
    this.$store.dispatch('training/getList');
  },
  methods: {
    show(url: string) {
      window.open(url, '_blank');
    },
  },
  data: () => {
    return {};
  },
});
</script>

<style lang="scss" module>
.main {
  padding: 10px;
  height: calc(100% - 60px);
  box-sizing: border-box;

  .item_list {
    display: grid;
    grid-template-columns: 1fr 1fr 1fr 1fr;
    gap: 10px;
    height: calc(100% - 35px);
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
  .main {
    .item_list {
      display: grid;
      grid-template-columns: 1fr 1fr 1fr;
      gap: 10px;
    }
  }
}

@media (max-width: 576px) {
  .main {
    .item_list {
      grid-template-columns: 1fr 1fr;
    }
  }
}
</style>
