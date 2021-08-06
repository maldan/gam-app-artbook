<template>
  <div class="main">
    <Header />

    <div class="block" style="margin-top: 10px">
      <Schedule />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Header from '../component/Header.vue';
import Schedule from '../component/worklist/Schedule.vue';

export default defineComponent({
  components: { Header, Schedule },
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
