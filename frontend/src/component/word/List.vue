<template>
  <div>
    <div class="word" v-for="x in $store.state.word.list" :key="x.word">
      <ui-icon
        @click="$store.dispatch('word/play', x.name)"
        class="clickable"
        name="play"
        width="17px"
        style="margin-right: 10px"
      />
      <div>{{ x.name }}</div>

      <div class="translate" v-if="x.translate.noun?.length">
        Noun: {{ x.translate.noun?.join(', ') }}
      </div>
      <div class="translate" v-if="x.translate.verb?.length">
        Verb: {{ x.translate.verb?.join(', ') }}
      </div>
      <div class="translate" v-if="x.translate.adjective?.length">
        Adjective: {{ x.translate.adjective?.join(', ') }}
      </div>

      <ui-icon
        class="clickable"
        @click="
          $store.dispatch('modal/show', {
            name: 'editWord',
            data: {
              name: x.name,
              category: x.category,
              translate: {
                noun: x.translate.noun.join(', '),
                verb: x.translate.verb.join(', '),
                adjective: x.translate.adjective.join(', '),
              },
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
  props: {},
  components: {},
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
.word {
  display: flex;
  background: #1b1b1b;
  padding: 10px;
  margin-bottom: 5px;
  align-items: center;

  .translate {
    color: #a1a1a1;
    margin-left: 10px;
  }
}
</style>
