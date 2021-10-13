import { createStore } from 'vuex';
import main, { MainStore } from './main';
import statistics, { StatisticsStore } from './statistics';
import art, { ArtStore } from './art';
import modal, { ModalStore } from '../gam_sdk_ui/vue/store/modal';

export type MainTree = {
  main: MainStore;
  modal: ModalStore;
  statistics: StatisticsStore;
  art: ArtStore;
};

export default createStore({
  modules: { main, modal, art, statistics },
});
