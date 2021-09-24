import Axios from 'axios';
import { ActionContext } from 'vuex';
import { MainTree } from '.';

export type TrainingStore = {
  list: any[];
};
export type WordActionContext = ActionContext<TrainingStore, MainTree>;

export default {
  namespaced: true,
  state: {
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
  mutations: {},
  actions: {},
};
