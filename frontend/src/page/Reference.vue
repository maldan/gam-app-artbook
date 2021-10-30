<template>
  <div :class="$style.main">
    <ui-button
      text="Add image"
      icon="plus"
      @click="
        $store.dispatch('modal/show', {
          name: 'addReference',
          data: {
            imageFile: null,
            category: '',
            tags: '',
          },
          onSuccess() {
            $store.dispatch('reference/add');
          },
        })
      "
      style="width: 100%; margin-bottom: 10px"
    />

    <ui-input placeholder="Filter..." style="margin-bottom: 10px" v-model="filter" />

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
          @click.stop="
            $store.dispatch('modal/show', {
              name: 'approve',
              data: {
                id: item.id,
                title: 'Remove?',
              },
              onSuccess() {
                $store.dispatch('reference/delete');
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
              name: 'editReference',
              data: {
                id: item.id,
                category: item.category,
                tags: item.tags.join(', '),
              },
              onSuccess() {
                $store.dispatch('reference/update');
              },
            })
          "
          class="clickable"
          src="../asset/pencil.svg"
          alt=""
          style="position: absolute; right: 40px; top: 15px"
        />

        <img :class="$style.preview" :src="item.thumbnail" />

        <div :class="$style.title" v-html="item.category"></div>
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
  computed: {
    list(): any {
      let list = this.$store.state.reference.list.filter((x: any) => {
        const parts = this.filter.split(' ').filter(Boolean);
        if (parts.length === 0) return true;

        let amount = parts.length;

        for (let j = 0; j < parts.length; j++) {
          for (let i = 0; i < x.tags.length; i++) {
            if (x.tags[i].match(parts[j])) {
              amount -= 1;
              if (amount <= 0) return true;
            }
          }

          if (x.category.match(parts[j])) {
            amount -= 1;
            if (amount <= 0) return true;
          }
        }

        return false;
      });

      return list.sort((a: any, b: any) => a.category.localeCompare(b.category));
    },
  },
  async mounted() {
    this.$store.dispatch('reference/getList');
  },
  methods: {
    show(url: string) {
      window.open(url, '_blank');
    },
  },
  data: () => {
    return {
      filter: '',
    };
  },
});
</script>

<style lang="scss" module>
.main {
  padding: 10px;
  height: calc(100% - 70px);
  box-sizing: border-box;

  .item_list {
    display: grid;
    grid-template-columns: 1fr 1fr 1fr 1fr 1fr 1fr 1fr 1fr;
    gap: 10px;
    height: calc(100% - 70px);
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
        display: flex;

        .hilight {
          background: #fffb00;
          color: #2b2b2b;
        }
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

@media (max-width: 2400px) {
  .main {
    .item_list {
      display: grid;
      grid-template-columns: 1fr 1fr 1fr 1fr 1fr 1fr 1fr;
      gap: 10px;
    }
  }
}

@media (max-width: 2100px) {
  .main {
    .item_list {
      display: grid;
      grid-template-columns: 1fr 1fr 1fr 1fr 1fr 1fr;
      gap: 10px;
    }
  }
}

@media (max-width: 1600px) {
  .main {
    .item_list {
      display: grid;
      grid-template-columns: 1fr 1fr 1fr 1fr 1fr;
      gap: 10px;
    }
  }
}

@media (max-width: 1440px) {
  .main {
    .item_list {
      display: grid;
      grid-template-columns: 1fr 1fr 1fr 1fr;
      gap: 10px;
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
