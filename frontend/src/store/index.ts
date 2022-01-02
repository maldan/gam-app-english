import { createStore, Store } from 'vuex';
import { InjectionKey } from 'vue';
import main, { MainStore } from './main';
import modal, { ModalStore } from '../gam_sdk_ui/vue/store/modal';
import word, { WordStore } from './word';
import training, { TrainingStore } from './training';
import statistics, { StatisticsStore } from './statistics';

export type MainTree = {
  main: MainStore;
  word: WordStore;
  modal: ModalStore;
  training: TrainingStore;
  statistics: StatisticsStore;
};

// define injection key
export const key: InjectionKey<Store<MainTree>> = Symbol();

export default createStore({
  modules: { main, word, modal, statistics, training },
});
