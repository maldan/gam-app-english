<template>
  <div class="main">
    <!-- Header -->
    <div class="header">
      <ui-input placeholder="Filter..." v-model="filter" />
      <div class="counter">{{ wordList.length }}</div>
      <ui-button
        @click="
          $store.dispatch('modal/show', {
            name: 'addWord',
            data: { name: '', category: '', translate: {} },
            func: () => {
              $store.dispatch('word/add');
            },
          })
        "
        text="Add"
        icon="plus"
        style="flex: 0"
      />
    </div>

    <!-- Word list -->
    <div v-if="filter.length >= 2" class="word" v-for="x in wordList" :key="x.word">
      <ui-icon
        @click="$store.dispatch('word/play', x.name)"
        class="clickable"
        name="play"
        width="17px"
        style="margin-right: 10px"
      />
      <div class="name">{{ x.name }}</div>
      <div class="category">{{ x.category.join(', ') }}</div>
      <div class="category">{{ x.power }}</div>

      <div class="translate">
        <div v-if="x.translate.noun?.length">Noun: {{ x.translate.noun?.join(', ') }}</div>
        <div v-if="x.translate.verb?.length">Verb: {{ x.translate.verb?.join(', ') }}</div>
        <div v-if="x.translate.adjective?.length">
          Adjective: {{ x.translate.adjective?.join(', ') }}
        </div>
      </div>

      <ui-icon
        class="clickable"
        @click="
          $store.dispatch('modal/show', {
            name: 'editWord',
            data: {
              name: x.name,
              category: x.category.join(', '),
              translate: {
                noun: x.translate.noun.join(', '),
                verb: x.translate.verb.join(', '),
                adjective: x.translate.adjective.join(', '),
              },
              isExists: false,
            },
            func: () => {
              $store.dispatch('word/edit');
            },
          })
        "
        name="pencil"
        width="17px"
        style="margin-left: auto"
      />

      <ui-icon
        class="clickable"
        @click="
          $store.dispatch('modal/show', {
            name: 'approve',
            data: {
              title: `Delete '${x.name}'?`,
            },
            func: () => {
              $store.dispatch('word/delete', x.name);
            },
          })
        "
        name="delete"
        width="17px"
        style="margin-left: 15px"
      />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  components: {},
  computed: {
    wordList() {
      return this.$store.state.word.list.filter((x: any) =>
        x.name.match(new RegExp(`^${this.filter}`)),
      );
    },
  },
  async mounted() {
    this.$store.dispatch(`word/getList`);
  },
  methods: {},
  data: () => {
    return {
      filter: '',
    };
  },
});
</script>

<style lang="scss" scoped>
.main {
  .header {
    display: flex;
    padding: 10px;
    align-items: center;

    .counter {
      padding: 10px;
    }
  }

  .word {
    display: flex;
    background: #1b1b1b;
    padding: 10px;
    margin-bottom: 5px;
    align-items: center;

    .name {
      flex: 1;
    }

    .category {
      margin-left: 10px;
      flex: 1;
    }

    .translate {
      color: #a1a1a1;
      margin-left: 10px;
      flex: 1;
    }
  }
}
</style>
