<template>
  <div class="main">
    <ui-header :items="header" />
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

    <ui-modal v-if="$store.state.modal.name">
      <component :is="'modal-' + $store.state.modal.name" />
    </ui-modal>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import List from '../component/word/List.vue';
import ModalAddWord from '../component/modal/AddWord.vue';
import ModalEditWord from '../component/modal/EditWord.vue';
import ModalApprove from '../component/modal/Approve.vue';

export default defineComponent({
  components: { List, ModalAddWord, ModalApprove, ModalEditWord },
  async mounted() {},
  methods: {},
  data: () => {
    return {
      header: [
        { name: 'Word DB', url: '/' },
        { name: 'Training', url: '/training' },
        { name: 'Statistics', url: '/statistics' },
      ],
    };
  },
});
</script>

<style lang="scss" scoped>
.main {
}
</style>
