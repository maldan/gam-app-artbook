<template>
  <ui-header :items="header" />
  <router-view />
  <ui-modal v-if="$store.state.modal.name">
    <component :is="'modal-' + $store.state.modal.name" />
  </ui-modal>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Moment from 'moment';
import ModalAddProject from '../src/component/modal/ModalAddProject.vue';
import ModalEditProject from '../src/component/modal/ModalEditProject.vue';

export default defineComponent({
  components: { ModalAddProject, ModalEditProject },
  async mounted() {},
  methods: {},
  data: () => {
    return {
      // @ts-ignore
      API_URL: process.env.VUE_APP_API_URL || `${window.location.origin}/api`,
      moment: Moment,

      header: [
        { name: 'Statistics', url: '/' },
        { name: 'Art', url: '/art' },
        { name: 'Training', url: '/training' },
        { name: 'DB', url: '/db' },
      ],
    };
  },
});
</script>

<style lang="scss" scoped></style>
