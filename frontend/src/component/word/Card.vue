<template>
  <div :class="$style.card" class="shadow">
    <div :class="$style.name">
      <div>{{ word.name }}</div>
      <div :class="$style.info">(noun)</div>
      <div :class="$style.power">[{{ word.power }}]</div>
    </div>

    <!-- Word name -->
    <div
      :style="{
        opacity: $store.state.training.isShow || $store.state.training.isBlockButton ? 1 : 0,
      }"
      :class="$style.translate"
    >
      <div v-if="word.translate.noun?.length">
        <b>Noun:</b> {{ word.translate.noun?.join(', ') }}
      </div>
      <div v-if="word.translate.verb?.length">
        <b>Verb:</b> {{ word.translate.verb?.join(', ') }}
      </div>
      <div v-if="word.translate.adjective?.length">
        <b>Adjective:</b> {{ word.translate.adjective?.join(', ') }}
      </div>
    </div>

    <div
      v-if="!$store.state.training.isBlockButton"
      class="button_group"
      :class="$style.button_group"
    >
      <ui-button
        @click="$store.dispatch('training/dontKnowWord')"
        text="Don't know"
        icon="not_allowed"
        size="compact"
      />
      <ui-button
        @click="$store.dispatch('training/show', true)"
        text="Show"
        icon="show"
        size="compact"
      />
      <ui-button
        @click="$store.dispatch('training/knowWord')"
        text="Know"
        icon="check"
        size="compact"
      />
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
    return {};
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

    .info,
    .power {
      font-size: 14px;
      margin-left: 10px;
      color: #15acde;
      font-weight: normal;
    }

    .power {
      color: #deb915;
      font-weight: bold;
      font-size: 18px;
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

    b {
      color: #d6bd54;
    }
  }

  .button_group {
    width: 100%;
  }
}
</style>
