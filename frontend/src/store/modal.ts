import { ActionContext } from 'vuex';
import { MainTree } from '.';

export type ModalStore = {
  name: string;
  data: any;
  func: () => void;
};
export type ModalActionContext = ActionContext<ModalStore, MainTree>;

export default {
  namespaced: true,
  state: {
    name: '',
    data: {},
    func: () => {},
  },
  mutations: {
    SET_MODAL(state: ModalStore, payload: any) {
      state.name = payload.name;
      state.data = payload.data;
      state.func = payload.func || function () {};
    },
  },
  actions: {
    show({ commit }: any, payload: any) {
      commit('SET_MODAL', payload);
    },
    close({ commit }: any) {
      commit('SET_MODAL', {});
    },
    ok({ state, commit }: any) {
      state?.func();
      commit('SET_MODAL', {});
    },
  },
};
