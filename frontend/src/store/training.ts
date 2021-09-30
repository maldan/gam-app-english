import Axios from 'axios';
import { ActionContext } from 'vuex';
import { MainTree } from '.';

export type TrainingStore = {
  word: {
    name: string;
  };
  category: string;
  isBlockButton: boolean;
};
export type TrainingActionContext = ActionContext<TrainingStore, MainTree>;

export default {
  namespaced: true,
  state: {
    word: {
      name: '',
    },
    category: '',
    isBlockButton: false,
    categoryList: [
      { name: 'a', score: 1 },
      { name: 'b' },
      { name: 'c' },
      { name: 'd' },
      { name: 'e' },
      { name: 'f' },
      { name: 'g' },
      { name: 'h' },
      { name: 'i' },
      { name: 'j' },
      { name: 'k' },
      { name: 'l' },
      { name: 'm' },
      { name: 'n' },
      { name: 'o' },
      { name: 'p' },
      { name: 'q' },
      { name: 'r' },
      { name: 's' },
      { name: 't' },
      { name: 'u' },
      { name: 'v' },
      { name: 'w' },
      { name: 'x' },
      { name: 'y' },
      { name: 'z' },
    ],
  },
  mutations: {
    SET_CATEGORY(state: TrainingStore, payload: string) {
      state.category = payload;
    },
    SET_WORD(state: TrainingStore, payload: any) {
      state.word = payload;
    },
    SET_BLOCK(state: TrainingStore, payload: boolean) {
      state.isBlockButton = payload;
    },
  },
  actions: {
    selectCategory(action: TrainingActionContext, payload: string) {
      action.commit('SET_CATEGORY', payload);
      action.dispatch('getWord');
    },
    async getWord(action: TrainingActionContext) {
      const word = (
        await Axios.get(
          `${action.rootState.main.API_URL}/training/word?category=${action.state.category}`,
        )
      ).data.response;
      action.commit('SET_WORD', word);
    },
    async knowWord(action: TrainingActionContext) {
      await Axios.post(`${action.rootState.main.API_URL}/training/knowWord`, {
        name: action.state.word.name,
      });

      action.dispatch('next');
    },
    async dontKnowWord(action: TrainingActionContext) {
      await Axios.post(`${action.rootState.main.API_URL}/training/dontKnowWord`, {
        name: action.state.word.name,
      });

      action.dispatch('next');
    },
    async next(action: TrainingActionContext) {
      action.commit('SET_BLOCK', true);

      setTimeout(async () => {
        await action.dispatch('getWord');
        action.commit('SET_BLOCK', false);
      }, 500);
    },
  },
};
