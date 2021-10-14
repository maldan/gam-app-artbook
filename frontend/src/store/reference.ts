import Axios from 'axios';
import Moment from 'moment';
import { ActionContext } from 'vuex';
import { MainTree } from '.';

export type ReferenceStore = {
  list: any[];
};
export type ReferenceActionContext = ActionContext<ReferenceStore, MainTree>;

export default {
  namespaced: true,
  state: {
    list: [],
  },
  mutations: {
    SET_LIST(state: ReferenceStore, list: any[]) {
      state.list = list;
    },
  },
  actions: {
    async getList(action: ReferenceActionContext) {
      const list = (await Axios.get(`${action.rootState.main.API_URL}/reference/list`)).data
        .response;
      action.commit('SET_LIST', list);
    },
    async add(action: ReferenceActionContext) {
      const formData = new FormData();
      formData.append('category', action.rootState.modal.data.category);
      formData.append('tags', action.rootState.modal.data.tags);
      formData.append('image', action.rootState.modal.data.imageFile);

      await Axios.post(`${action.rootState.main.API_URL}/reference`, formData);

      action.dispatch('getList');
    },
    async update(action: ReferenceActionContext) {
      await Axios.patch(`${action.rootState.main.API_URL}/reference`, {
        id: action.rootState.modal.data.id,
        category: action.rootState.modal.data.category,
        tags: action.rootState.modal.data.tags
          .split(',')
          .map((x: string) => x.trim())
          .filter(Boolean),
      });

      action.dispatch('getList');
    },
    async delete(action: ReferenceActionContext) {
      await Axios.delete(
        `${action.rootState.main.API_URL}/reference?id=${action.rootState.modal.data.id}`,
      );
      action.dispatch('getList');
    },
  },
};
