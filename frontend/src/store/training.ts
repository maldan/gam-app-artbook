import Axios from 'axios';
import Moment from 'moment';
import { ActionContext } from 'vuex';
import { MainTree } from '.';

export type TrainingStore = {
  art: any;
  list: any[];
};
export type TrainingActionContext = ActionContext<TrainingStore, MainTree>;

export default {
  namespaced: true,
  state: {
    art: {},
    list: [],
  },
  mutations: {
    SET_LIST(state: TrainingStore, list: any[]) {
      state.list = list;
    },
  },
  actions: {
    async getList(action: TrainingActionContext) {
      const list = (await Axios.get(`${action.rootState.main.API_URL}/training/list`)).data
        .response;
      action.commit('SET_LIST', list);
    },
  },
};
