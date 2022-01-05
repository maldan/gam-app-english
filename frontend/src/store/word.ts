import Axios from 'axios';
import { ActionContext } from 'vuex';
import { MainTree } from '.';

export interface Translate {
  noun: string[] | string;
  verb: string[] | string;
  adjective: string[] | string;
}

export interface Word {
  name: string;
  translate: Translate;
  isExists?: boolean;
}

export type WordStore = {
  list: Word[];
};
export type WordActionContext = ActionContext<WordStore, MainTree>;

export default {
  namespaced: true,
  state: {
    list: [],
  },
  mutations: {
    SET_LIST(state: WordStore, payload: Word[]): void {
      state.list = payload;
    },
  },
  actions: {
    async checkExists(action: WordActionContext, word: Word): Promise<void> {
      try {
        await Axios.get(`${action.rootState.main.API_URL}/word?name=${word.name}`);
        word.isExists = true;
      } catch {
        word.isExists = false;
      }
    },
    async getList(action: WordActionContext): Promise<void> {
      const list = (await Axios.get(`${action.rootState.main.API_URL}/word/list`)).data.response;
      action.commit('SET_LIST', list);
    },
    async add(action: WordActionContext): Promise<void> {
      await Axios.post(`${action.rootState.main.API_URL}/word`, {
        name: action.rootState.modal.data.name.toLowerCase(),
        category:
          action.rootState.modal.data.category
            .toLowerCase()
            .split(',')
            .map((x: string) => x.trim())
            .filter(Boolean) || [],
        translate: {
          noun:
            action.rootState.modal.data.translate?.noun
              ?.split(',')
              ?.map((x: string) => x.trim())
              ?.filter(Boolean) || [],
          verb:
            action.rootState.modal.data.translate?.verb
              ?.split(',')
              ?.map((x: string) => x.trim())
              ?.filter(Boolean) || [],
          adjective:
            action.rootState.modal.data.translate?.adjective
              ?.split(',')
              ?.map((x: string) => x.trim())
              ?.filter(Boolean) || [],
        },
      });
      await action.dispatch('getList');
    },
    async edit(action: WordActionContext): Promise<void> {
      await Axios.patch(`${action.rootState.main.API_URL}/word`, {
        name: action.rootState.modal.data.name.toLowerCase(),
        category:
          action.rootState.modal.data.category
            .toLowerCase()
            .split(',')
            .map((x: string) => x.trim())
            .filter(Boolean) || [],
        translate: {
          noun:
            action.rootState.modal.data.translate?.noun
              ?.split(',')
              ?.map((x: string) => x.trim())
              ?.filter(Boolean) || [],
          verb:
            action.rootState.modal.data.translate?.verb
              ?.split(',')
              ?.map((x: string) => x.trim())
              ?.filter(Boolean) || [],
          adjective:
            action.rootState.modal.data.translate?.adjective
              ?.split(',')
              ?.map((x: string) => x.trim())
              ?.filter(Boolean) || [],
        },
      });
      await action.dispatch('getList');
    },
    async delete(action: WordActionContext, payload: string): Promise<void> {
      await Axios.delete(`${action.rootState.main.API_URL}/word?name=${payload}`);
      await action.dispatch('getList');
    },
    async play(action: WordActionContext, payload: string): Promise<void> {
      await new Audio(`${action.rootState.main.API_URL}/word/play?name=${payload}`).play();
    },
    async translate(action: WordActionContext, word: Word): Promise<void> {
      const translatedWord = (
        await Axios.get(`${action.rootState.main.API_URL}/word/translate?name=${word.name}`)
      ).data.response as Word;
      word.translate.noun = (translatedWord.translate.noun as string[]).join(', ');
      word.translate.verb = (translatedWord.translate.verb as string[]).join(', ');
      word.translate.adjective = (translatedWord.translate.adjective as string[]).join(', ');
    },
  },
};
