<template>
  <div class="main">
    <!-- <Card v-if="$store.state.word.list[0]" :word="$store.state.word.list[0]" /> -->
    <div class="category_list">
      <div class="category clickable" v-for="x in $store.state.training.categoryList" :key="x.name">
        <div class="name">{{ x.name }}</div>
        <div class="score">{{ x.score || 0 }}</div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import List from '../component/word/List.vue';
import Card from '../component/word/Card.vue';

export default defineComponent({
  components: { List, Card },
  async mounted() {
    this.$store.dispatch(`word/getList`);
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
    align-items: center;
    justify-content: center;
  }
}
</style>
