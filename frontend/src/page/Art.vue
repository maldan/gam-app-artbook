<template>
  <div :class="$style.main">
    <ui-button
      text="Add art"
      icon="plus"
      @click="
        $store.dispatch('modal/show', {
          name: 'addArt',
          data: {
            projectId: $route.params.id,
            imageFile: null,
            time: '00:00',
            created: $root.moment().format('YYYY-MM-DD HH:mm:ss'),
          },
          onSuccess() {
            $store.dispatch('project/addImage');
          },
        })
      "
      style="width: 100%; margin-bottom: 10px"
    />

    <!-- List -->
    <div :class="$style.item_list">
      <div v-for="item in $store.state.project.art.imageList" :key="item.id" :class="$style.block">
        <img
          @click.stop="
            $store.dispatch('project/imageMoveDown', {
              imageId: item.id,
              projectId: $route.params.id,
            })
          "
          class="clickable"
          src="../asset/move_down.svg"
          alt=""
          style="position: absolute; right: 75px; top: 12px"
          draggable="false"
        />
        <img
          @click.stop="
            $store.dispatch('modal/show', {
              name: 'editArt',
              data: {
                projectId: $route.params.id,
                imageId: item.id,
                time: $root.moment.utc(item.time * 1000).format('HH:mm'),
                created: $root.moment(item.created).format('YYYY-MM-DD HH:mm:ss'),
              },
              onSuccess() {
                $store.dispatch('project/updateImage');
              },
            })
          "
          class="clickable"
          src="../asset/pencil.svg"
          alt=""
          style="position: absolute; right: 42px; top: 12px"
          draggable="false"
        />
        <img
          @click.stop="
            $store.dispatch('modal/show', {
              name: 'approve',
              data: {
                title: 'Remove image?',
                projectId: $route.params.id,
                imageId: item.id,
              },
              onSuccess() {
                $store.dispatch('project/deleteImage');
              },
            })
          "
          class="clickable"
          src="../asset/trash.svg"
          alt=""
          style="position: absolute; right: 12px; top: 12px"
          draggable="false"
        />

        <img :src="item.url" style="width: 80%; margin: 0 auto" draggable="false" />
        <div style="display: flex; flex: 1; margin-top: 10px">
          <div style="flex: 1">
            {{
              $root.moment
                .utc(item.time * 1000)
                .format('H %1 m %2')
                .replace('%1', 'hour')
                .replace('%2', 'mins')
            }}
          </div>
          <div style="flex: 1">
            {{ $root.moment(item.created).format('DD MMM YYYY HH:mm:ss') }}
          </div>
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
    this.$store.dispatch('project/get', this.$route.params.id as string);
  },
  methods: {},
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
    height: calc(100% - 35px);
    overflow-y: auto;

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
