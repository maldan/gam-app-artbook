import Axios from 'axios';
import Moment from 'moment';
import { ActionContext } from 'vuex';
import { MainTree } from '.';

export type ArtStore = {
  art: any;
};
export type ArtActionContext = ActionContext<ArtStore, MainTree>;

export default {
  namespaced: true,
  state: {
    art: {},
  },
  mutations: {
    SET_ART(state: ArtStore, art: any) {
      state.art = art;
    },
  },
  actions: {
    async get(action: ArtActionContext, id: string) {
      const art = (await Axios.get(`${action.rootState.main.API_URL}/art?id=${id}`)).data.response;
      action.commit('SET_ART', art);
    },
  },
};
