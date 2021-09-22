import Axios from 'axios';
import { ActionContext } from 'vuex';
import { MainTree } from '.';

export type WordStore = {
  list: any[];
};
export type WordActionContext = ActionContext<WordStore, MainTree>;

export default {
  namespaced: true,
  state: {
    list: [],
  },
  mutations: {
    SET_LIST(state: WordStore, payload: any) {
      state.list = payload;
    },
  },
  actions: {
    async getList(action: WordActionContext) {
      const list = (await Axios.get(`${action.rootState.main.API_URL}/word/list`)).data.response;
      action.commit('SET_LIST', list);
    },
    async add(action: WordActionContext) {
      (
        await Axios.post(`${action.rootState.main.API_URL}/word`, {
          name: action.rootState.modal.data.name.toLowerCase(),
          translate: {
            noun: action.rootState.modal.data.translate.noun
              .split(', ')
              .map((x: string) => x.trim()),
            verb: [],
            adjective: [],
          },
        })
      ).data.response;
      action.dispatch('getList');
    },
    async edit(action: WordActionContext) {
      (
        await Axios.patch(`${action.rootState.main.API_URL}/word`, {
          name: action.rootState.modal.data.name.toLowerCase(),
          translate: {
            noun: action.rootState.modal.data.translate.noun
              .split(', ')
              .map((x: string) => x.trim()),
            verb: [],
            adjective: [],
          },
        })
      ).data.response;
      action.dispatch('getList');
    },
    async delete(action: WordActionContext, payload: string) {
      (await Axios.delete(`${action.rootState.main.API_URL}/word?name=${payload}`)).data.response;
      action.dispatch('getList');
    },
  },
};
