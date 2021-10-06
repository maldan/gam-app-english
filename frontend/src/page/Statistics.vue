<template>
  <div class="main">
    <ui-schedule :map="$store.state.statistics.yearMap" :max="500" />
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
