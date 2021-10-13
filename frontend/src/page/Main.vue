<template>
  <div class="main">
    <div class="block">
      <ui-schedule
        :map="$store.state.statistics.yearMap"
        :max="6"
        :date="$store.state.statistics.date"
        :yearRange="[-6, -5, -4, -3, -2, -1, 0]"
        color="#f57202"
        @select="$store.dispatch('statistics/setDate', $event)"
      />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  components: {},
  async mounted() {
    this.$store.dispatch('statistics/getYearMap');
  },
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
