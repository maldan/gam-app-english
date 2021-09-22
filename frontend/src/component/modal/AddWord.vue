<template>
  <div>
    <ui-input
      @keydown.enter="$store.dispatch('modal/ok')"
      placeholder="Word..."
      v-model="$store.state.modal.data.name"
    />
    <ui-input
      @keydown.enter="$store.dispatch('modal/ok')"
      placeholder="Translate as noun..."
      v-model="$store.state.modal.data.translate.noun"
    />
    <ui-input
      @keydown.enter="$store.dispatch('modal/ok')"
      placeholder="Translate as verb..."
      v-model="$store.state.modal.data.translate.verb"
    />
    <ui-input
      @keydown.enter="$store.dispatch('modal/ok')"
      placeholder="Translate as adjective..."
      v-model="$store.state.modal.data.translate.adjective"
    />
    <div class="button_group" style="margin-top: 5px">
      <ui-button text="Cancel" @click="$store.dispatch('modal/close')" />
      <ui-button
        :disabled="!$store.state.modal.data.name"
        text="Add"
        icon="plus"
        @click="$store.dispatch('modal/ok')"
      />
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
