<template>
  <div>
    <div>{{ $store.state.modal.data.title }}</div>
    <div class="button_group" style="margin-top: 5px">
      <ui-button text="Cancel" @click="$store.dispatch('modal/close')" />
      <ui-button text="Ok" icon="plus" @click="$store.dispatch('modal/ok')" />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  props: {},
  components: {},
  async mounted() {},
  methods: {},
  data: () => {
    return {};
  },
});
</script>

<style lang="scss" scoped></style>
