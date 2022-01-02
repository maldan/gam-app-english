<template>
  <div class="main">
    <div v-if="$store.state.training.category === ''" class="category_list">
      <div
        @click="$store.dispatch('training/selectCategory', x.name)"
        class="category clickable"
        v-for="x in categoryList"
        :key="x.name"
      >
        <div class="name">{{ x.name }}</div>
        <div class="score">
          <div v-if="x.correct > 0" class="correct">{{ x.correct }}</div>
          <div class="total">{{ x.amount - x.correct }}</div>
        </div>
      </div>
    </div>
    <div class="card_container" v-else>
      <div class="stat">
        <div>Correct: {{ $store.state.statistics.correct?.length }}</div>
        <div>Wrong: {{ $store.state.statistics.wrong?.length }}</div>
        <div>
          Total:
          {{ $store.state.statistics.correct?.length + $store.state.statistics.wrong?.length }}
        </div>
      </div>
      <div class="button_group">
        <ui-button
          @click="$store.dispatch('training/selectCategory', '')"
          text="Leave"
          style="flex: none; margin-bottom: 10px"
        />
        <ui-button
          @click="
            $store.dispatch('modal/show', {
              name: 'approve',
              data: {
                title: 'Are you sure?',
              },
              func: () => {
                $store.dispatch('training/reset');
              },
            })
          "
          icon="delete"
          text="Reset"
          style="flex: none; margin-bottom: 10px"
        />
      </div>

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
  async mounted() {
    this.$store.dispatch('training/getCategoryList');
  },
  computed: {
    categoryList() {
      const list = this.$store.state.training.categoryList;
      list.sort((a: any, b: any) => {
        return a.name.localeCompare(b.name);
      });
      return list;
    },
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
    grid-template-columns: 1fr 1fr 1fr 1fr 1fr 1fr;
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
        display: flex;

        .correct {
          font-weight: bold;
          margin-right: 5px;
          color: #ff8843;
        }
      }
    }
  }

  .card_container {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100%;

    .stat {
      display: flex;

      > div {
        padding: 10px;
      }
    }
  }
}
</style>
