import { createStore } from 'vuex';
import main, { MainStore } from './main';
import statistics, { StatisticsStore } from './statistics';
import project, { ProjectActionContext } from './project';
import training, { TrainingStore } from './training';
import art, { ArtStore } from './art';
import reference, { ReferenceStore } from './reference';
import modal, { ModalStore } from '../gam_sdk_ui/vue/store/modal';

export type MainTree = {
  main: MainStore;
  modal: ModalStore;
  statistics: StatisticsStore;
  project: ProjectActionContext;
  training: TrainingStore;
  art: ArtStore;
  reference: ReferenceStore;
};

export default createStore({
  modules: { main, modal, project, art, reference, training, statistics },
});
