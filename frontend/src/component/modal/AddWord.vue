<template>
  <div>
    <ui-input
      placeholder="Word..."
      v-model="$store.state.modal.data.name"
      @keyup="$store.dispatch('word/checkExists', $store.state.modal.data)"
      @keyup.enter="$store.dispatch('word/translate', $store.state.modal.data)"
    />
    <ui-input placeholder="Category..." v-model="$store.state.modal.data.category" />
    <ui-input placeholder="Translate as noun..." v-model="$store.state.modal.data.translate.noun" />
    <ui-input placeholder="Translate as verb..." v-model="$store.state.modal.data.translate.verb" />
    <ui-input
      placeholder="Translate as adjective..."
      v-model="$store.state.modal.data.translate.adjective"
    />
    <div
      v-if="$store.state.modal.data.isExists"
      style="margin-top: 6px; color: #33ad53; padding: 3px"
    >
      Word already exists
    </div>
    <div class="button_group" style="margin-top: 5px">
      <ui-button text="Cancel" @click="$store.dispatch('modal/close')" />
      <ui-button
        text="Translate"
        @click="$store.dispatch('word/translate', $store.state.modal.data)"
      />
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

<style lang="scss" scoped>
.window {
}
</style>
