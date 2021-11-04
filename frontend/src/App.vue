<template>
  <ui-header :items="header" />
  <router-view />
  <ui-modal v-if="$store.state.modal.name">
    <component :is="'modal-' + $store.state.modal.name" />
  </ui-modal>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import ModalAddWord from './component/modal/AddWord.vue';
import ModalEditWord from './component/modal/EditWord.vue';
import ModalApprove from './component/modal/Approve.vue';

export default defineComponent({
  components: { ModalAddWord, ModalApprove, ModalEditWord },
  async mounted() {},
  methods: {},
  data: () => {
    return {
      // @ts-ignore
      API_URL: process.env.VUE_APP_API_URL || `${window.location.origin}/api`,

      header: [
        { name: 'Word DB', url: '/' },
        { name: 'Training', url: '/training' },
        { name: 'Statistics', url: '/statistics' },
      ],
    };
  },
});
</script>

<style lang="scss" scoped></style>
