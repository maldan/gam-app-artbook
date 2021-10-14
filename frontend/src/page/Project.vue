<template>
  <div :class="$style.main">
    <ui-button
      text="Create new project"
      icon="plus"
      @click="
        $store.dispatch('modal/show', {
          name: 'addProject',
          data: {
            title: '',
            tags: '',
          },
          onSuccess() {
            $store.dispatch('project/add');
          },
        })
      "
      style="width: 100%; margin-bottom: 10px"
    />

    <!-- List -->
    <div :class="$style.item_list">
      <div
        @click="$router.push(`/project/${item.id}`)"
        class="clickable"
        v-for="(item, i) in $store.state.project.list"
        :key="item.id"
        :class="$style.block"
      >
        <div :class="$style.number">{{ $store.state.project.list.length - i }}</div>
        <img
          @click.stop="
            $store.dispatch('modal/show', {
              name: 'approve',
              data: {
                id: item.id,
                title: 'Remove?',
              },
              onSuccess() {
                $store.dispatch('project/delete');
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
              name: 'editProject',
              data: {
                id: item.id,
                title: item.title,
                tags: item.tags.join(', '),
              },
              onSuccess() {
                $store.dispatch('project/update');
              },
            })
          "
          class="clickable"
          src="../asset/pencil.svg"
          alt=""
          style="position: absolute; right: 40px; top: 15px"
        />

        <img
          :class="$style.preview"
          v-if="item.imageList[0]"
          :src="item.imageList[item.imageList.length - 1].thumbnail"
        />

        <div :class="$style.title">{{ item.title }}</div>
        <div v-if="item.imageList[0]" :class="$style.info">
          <div>{{ totalTime(item.imageList) }}</div>
          <div>{{ lastTime(item.imageList) }}</div>
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
import Moment from 'moment';

export default defineComponent({
  components: {},
  async mounted() {
    this.$store.dispatch('project/getList');
  },
  methods: {
    totalTime(imageList: any[]) {
      let out = 0;
      for (let i = 0; i < imageList.length; i++) {
        out += ~~imageList[i].time;
      }
      const h = ~~(out / 3600);
      const m = (out / 60) % 60;
      return `${h} h ${m} m`;
    },
    lastTime(imageList: any[]) {
      let times = (JSON.parse(JSON.stringify(imageList)) as any).sort(
        (a: any, b: any) => new Date(b.created).getTime() - new Date(a.created).getTime(),
      );
      return Moment(times[0].created).format('DD MMM YYYY');
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
  height: calc(100% - 45px);
  box-sizing: border-box;

  //.list {
  //  height: calc(100% - 50px);

  .item_list {
    display: grid;
    grid-template-columns: 1fr 1fr 1fr 1fr;
    gap: 10px;
    height: calc(100% - 50px);
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
//}
</style>
