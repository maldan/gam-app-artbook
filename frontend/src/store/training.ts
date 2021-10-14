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

    async add(action: TrainingActionContext) {
      const formData = new FormData();
      formData.append('title', action.rootState.modal.data.title);
      formData.append('tags', action.rootState.modal.data.tags);
      formData.append('image', action.rootState.modal.data.imageFile);
      formData.append(
        'time',
        Moment.duration(action.rootState.modal.data.time + ':00').asSeconds() + '',
      );
      formData.append('created', action.rootState.modal.data.created);

      await Axios.post(`${action.rootState.main.API_URL}/training`, formData);

      action.dispatch('getList');
    },

    async update(action: TrainingActionContext) {
      await Axios.patch(`${action.rootState.main.API_URL}/training`, {
        id: action.rootState.modal.data.id,
        title: action.rootState.modal.data.title,
        tags: action.rootState.modal.data.tags
          .split(',')
          .map((x: string) => x.trim())
          .filter(Boolean),
        time: Moment.duration(action.rootState.modal.data.time + ':00').asSeconds(),
        created: action.rootState.modal.data.created,
      });

      action.dispatch('getList');
    },

    async delete(action: TrainingActionContext) {
      await Axios.delete(
        `${action.rootState.main.API_URL}/training?id=${action.rootState.modal.data.id}`,
      );
      action.dispatch('getList');
    },
  },
};
