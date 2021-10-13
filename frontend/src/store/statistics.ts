import Axios from 'axios';
import Moment from 'moment';
import { ActionContext } from 'vuex';
import { MainTree } from '.';

export type StatisticsStore = {
  date: Date;
  yearMap: any;
};
export type StatisticsActionContext = ActionContext<StatisticsStore, MainTree>;

export default {
  namespaced: true,
  state: {
    date: new Date(),
    yearMap: {},
  },
  mutations: {
    SET_YEAR_MAP(state: StatisticsStore, payload: any) {
      state.yearMap = payload;
    },
    SET_DATE(state: StatisticsStore, date: Date) {
      state.date = date;
    },
  },
  actions: {
    async setDate(action: StatisticsActionContext, date: Date) {
      const isRefresh = date.getFullYear() != action.state.date.getFullYear();
      action.commit('SET_DATE', date);
      if (isRefresh) {
        action.dispatch('getYearMap', date);
      }
    },
    async getYearMap(action: StatisticsActionContext) {
      const yearMap = (
        await Axios.get(
          `${action.rootState.main.API_URL}/statistics/yearMap?date=${Moment(
            action.state.date,
          ).format('YYYY-MM-DD')}`,
        )
      ).data.response;
      action.commit('SET_YEAR_MAP', yearMap);
    },
  },
};
