<template>
  <div class="main">
    <ui-button
      @click="
        $store.dispatch('modal/show', {
          name: 'addWord',
          data: { name: '', translate: {} },
          func: () => {
            $store.dispatch('word/add');
          },
        })
      "
      text="Add"
      icon="plus"
    />
    <List />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import List from '../component/word/List.vue';

export default defineComponent({
  components: { List },
  async mounted() {},
  methods: {},
  data: () => {
    return {};
  },
});
</script>

<style lang="scss" scoped>
.main {
}
</style>
