import { createStore } from 'vuex';
import main, { MainStore } from './main';
import modal, { ModalStore } from './modal';
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

export default createStore({
  modules: { main, word, modal, statistics, training },
});
