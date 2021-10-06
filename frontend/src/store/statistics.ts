import Axios from 'axios';
import Moment from 'moment';
import { ActionContext } from 'vuex';
import { MainTree } from '.';

export type StatisticsStore = {
  correct: number;
  wrong: number;
  yearMap: any;
};
export type StatisticsActionContext = ActionContext<StatisticsStore, MainTree>;

export default {
  namespaced: true,
  state: {
    correct: 0,
    wrong: 0,
    yearMap: {},
  },
  mutations: {
    SET_TODAY(state: StatisticsStore, payload: any) {
      state.correct = payload.correct;
      state.wrong = payload.wrong;
    },
    SET_YEAR_MAP(state: StatisticsStore, payload: any) {
      state.yearMap = payload;
    },
  },
  actions: {
    async getToday(action: StatisticsActionContext) {
      const today = (await Axios.get(`${action.rootState.main.API_URL}/statistics/today`)).data
        .response;
      action.commit('SET_TODAY', today);
    },
    async getYearMap(action: StatisticsActionContext) {
      const yearMap = (
        await Axios.get(
          `${action.rootState.main.API_URL}/statistics/yearMap?date=${Moment().format(
            'YYYY-MM-DD',
          )}`,
        )
      ).data.response;
      action.commit('SET_YEAR_MAP', yearMap);
    },
  },
};
