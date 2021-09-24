<template>
  <div :class="$style.card" class="shadow">
    <div :class="$style.name">
      <div>{{ word.name }}</div>
      <div :class="$style.info">(noun)</div>
    </div>
    <div :style="{ opacity: isShow ? 1 : 0 }" :class="$style.translate">
      {{ word.translate.noun.join(', ') }}
    </div>
    <div class="button_group" :class="$style.button_group">
      <ui-button text="Don't know" icon="not_allowed" size="compact" />
      <ui-button @click="isShow = true" text="Show" icon="show" size="compact" />
      <ui-button text="Know" icon="check" size="compact" />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  props: {
    word: {
      type: Object,
      required: true,
    },
  },
  components: {},
  async mounted() {},
  methods: {},
  data: () => {
    return {
      isShow: false,
    };
  },
});
</script>

<style lang="scss" module>
.card {
  background: #444444;
  padding: 10px;
  width: 420px;
  border-radius: 4px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;

  .name {
    display: flex;
    align-items: center;

    .info {
      font-size: 14px;
      margin-left: 10px;
      color: #15acde;
      font-weight: normal;
    }
  }

  .name,
  .translate {
    margin-bottom: 10px;
    text-transform: uppercase;
    font-weight: bold;
    font-size: 28px;
    color: #afafaf;
  }

  .translate {
    font-size: 22px;
    color: #9cb458;
    margin-bottom: 15px;
  }

  .button_group {
    width: 100%;
  }
}
</style>
