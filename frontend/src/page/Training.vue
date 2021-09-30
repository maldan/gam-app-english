<template>
  <div class="main">
    <div v-if="$store.state.training.category === ''" class="category_list">
      <div
        @click="$store.dispatch('training/selectCategory', x.name)"
        class="category clickable"
        v-for="x in $store.state.training.categoryList"
        :key="x.name"
      >
        <div class="name">{{ x.name }}</div>
        <div class="score">{{ x.score || 0 }}</div>
      </div>
    </div>
    <div class="card_container" v-else>
      <ui-button
        @click="$store.dispatch('training/selectCategory', '')"
        text="Leave"
        style="flex: none; margin-bottom: 10px"
      />

      <Card v-if="$store.state.training.word.name" :word="$store.state.training.word" />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import List from '../component/word/List.vue';
import Card from '../component/word/Card.vue';

export default defineComponent({
  components: { List, Card },
  async mounted() {},
  methods: {},
  data: () => {
    return {};
  },
});
</script>

<style lang="scss" scoped>
.main {
  padding: 10px;
  height: calc(100% - 70px);

  .category_list {
    display: grid;
    grid-template-columns: 1fr 1fr 1fr 1fr;
    gap: 10px;

    .category {
      background: #2b2b2b;
      color: #b1b1b1;
      padding: 10px;
      border-radius: 4px;
      display: flex;
      align-items: center;
      justify-content: center;
      user-select: none;
      text-transform: uppercase;
      flex-direction: column;

      .name {
        font-weight: bold;
      }

      .score {
        margin-top: 10px;
      }
    }
  }

  .card_container {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100%;
  }
}
</style>
