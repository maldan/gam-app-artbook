<template>
  <div class="main">
    <Header />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Header from '../component/Header.vue';

export default defineComponent({
  components: { Header },
  async mounted() {},
  methods: {},
  data: () => {
    return {};
  },
});
</script>

<style lang="scss" scoped>
.main {
  padding: 10px;
}
</style>
