import Axios from 'axios';
import { ActionContext } from 'vuex';
import { MainTree } from '.';

export type TrainingStore = {
  word: {
    name: string;
  };
  category: string;
  isBlockButton: boolean;
  isShow: boolean;
  categoryList: any[];
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
    isShow: false,
    categoryList: [],
  },
  mutations: {
    SET_CATEGORY_LIST(state: TrainingStore, payload: any) {
      state.categoryList = payload;
    },
    SET_CATEGORY(state: TrainingStore, payload: string) {
      state.category = payload;
    },
    SET_WORD(state: TrainingStore, payload: any) {
      state.word = payload;
    },
    SET_BLOCK(state: TrainingStore, payload: boolean) {
      state.isBlockButton = payload;
    },
    SET_SHOW(state: TrainingStore, payload: boolean) {
      state.isShow = payload;
    },
  },
  actions: {
    selectCategory(action: TrainingActionContext, payload: string) {
      action.commit('SET_CATEGORY', payload);
      action.dispatch('getWord');
    },
    show(action: TrainingActionContext, payload: string) {
      action.commit('SET_SHOW', payload);
    },
    async getCategoryList(action: TrainingActionContext) {
      const list = (await Axios.get(`${action.rootState.main.API_URL}/training/categoryList`)).data
        .response;
      action.commit('SET_CATEGORY_LIST', list);
    },
    async getWord(action: TrainingActionContext) {
      const word = (
        await Axios.get(
          `${action.rootState.main.API_URL}/training/word?category=${action.state.category}`,
        )
      ).data.response;
      action.commit('SET_WORD', word);

      action.dispatch('word/play', word.name, { root: true });
      action.dispatch('statistics/getToday', null, { root: true });
      action.dispatch('show', false);
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
        action.dispatch('show', false);
      }, 500);
    },
    async reset(action: TrainingActionContext) {
      await Axios.post(`${action.rootState.main.API_URL}/training/reset`, {
        category: action.state.category,
      });

      action.dispatch('next');
    },
  },
};
