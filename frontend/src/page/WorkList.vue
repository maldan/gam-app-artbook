<template>
  <div class="main">
    <Header />

    <List style="margin-top: 10px" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Header from '../component/Header.vue';
import List from '../component/worklist/List.vue';

export default defineComponent({
  components: { Header, List },
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
