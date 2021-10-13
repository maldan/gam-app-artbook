import Axios from 'axios';
import Moment from 'moment';
import { ActionContext } from 'vuex';
import { MainTree } from '.';

export type ArtStore = {
  list: any[];
};
export type ArtActionContext = ActionContext<ArtStore, MainTree>;

export default {
  namespaced: true,
  state: {
    list: [],
  },
  mutations: {
    SET_LIST(state: ArtStore, list: any[]) {
      state.list = list;
    },
  },
  actions: {
    async getList(action: ArtActionContext) {
      const list = (await Axios.get(`${action.rootState.main.API_URL}/art/list`)).data.response;
      action.commit('SET_LIST', list);
    },
    async addProject(action: ArtActionContext) {
      (
        await Axios.post(`${action.rootState.main.API_URL}/art`, {
          title: action.rootState.modal.data.title,
          tags: action.rootState.modal.data.tags
            .split(',')
            .map((x: string) => x.trim())
            .filter(Boolean),
        })
      ).data.response;
      action.dispatch('getList');
    },
    async updateProject(action: ArtActionContext) {
      (
        await Axios.patch(`${action.rootState.main.API_URL}/art`, {
          id: action.rootState.modal.data.id,
          title: action.rootState.modal.data.title,
          tags: action.rootState.modal.data.tags
            .split(',')
            .map((x: string) => x.trim())
            .filter(Boolean),
        })
      ).data.response;
      action.dispatch('getList');
    },
  },
};
