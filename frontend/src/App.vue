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
import ModalAddReference from '../src/component/modal/ModalAddReference.vue';
import ModalEditReference from '../src/component/modal/ModalEditReference.vue';
import ModalAddTraining from '../src/component/modal/ModalAddTraining.vue';
import ModalEditTraining from '../src/component/modal/ModalEditTraining.vue';
import ModalAddArt from '../src/component/modal/ModalAddArt.vue';
import ModalEditArt from '../src/component/modal/ModalEditArt.vue';

export default defineComponent({
  components: {
    ModalAddProject,
    ModalEditProject,
    ModalAddReference,
    ModalEditReference,
    ModalAddTraining,
    ModalEditTraining,
    ModalAddArt,
    ModalEditArt,
  },
  async mounted() {},
  methods: {},
  data: () => {
    return {
      // @ts-ignore
      API_URL: process.env.VUE_APP_API_URL || `${window.location.origin}/api`,
      moment: Moment,

      header: [
        { name: 'Statistics', url: '/' },
        { name: 'Project', url: '/project' },
        { name: 'Training', url: '/training' },
        { name: 'Reference', url: '/reference' },
      ],
    };
  },
});
</script>

<style lang="scss" scoped></style>
