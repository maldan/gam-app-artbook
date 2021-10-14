import Axios from 'axios';
import Moment from 'moment';
import { ActionContext } from 'vuex';
import { MainTree } from '.';

export type ProjectStore = {
  art: any;
  list: any[];
};
export type ProjectActionContext = ActionContext<ProjectStore, MainTree>;

export default {
  namespaced: true,
  state: {
    art: {},
    list: [],
  },
  mutations: {
    SET_ART(state: ProjectStore, art: any) {
      state.art = art;
    },
    SET_LIST(state: ProjectStore, list: any[]) {
      state.list = list;
    },
  },
  actions: {
    async get(action: ProjectActionContext, id: string) {
      const art = (await Axios.get(`${action.rootState.main.API_URL}/project?id=${id}`)).data
        .response;
      action.commit('SET_ART', art);
    },
    async getList(action: ProjectActionContext) {
      const list = (await Axios.get(`${action.rootState.main.API_URL}/project/list`)).data.response;
      action.commit('SET_LIST', list);
    },
    async add(action: ProjectActionContext) {
      (
        await Axios.post(`${action.rootState.main.API_URL}/project`, {
          title: action.rootState.modal.data.title,
          tags: action.rootState.modal.data.tags
            .split(',')
            .map((x: string) => x.trim())
            .filter(Boolean),
        })
      ).data.response;
      action.dispatch('getList');
    },
    async update(action: ProjectActionContext) {
      (
        await Axios.patch(`${action.rootState.main.API_URL}/project`, {
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
