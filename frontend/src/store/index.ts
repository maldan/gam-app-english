import { createStore } from 'vuex';
import main, { MainStore } from './main';
import modal, { ModalStore } from './modal';
import word, { WordStore } from './word';
import training, { TrainingStore } from './training';

export type MainTree = {
  main: MainStore;
  word: WordStore;
  modal: ModalStore;
  training: TrainingStore;
};

export default createStore({
  modules: { main, word, modal, training },
});
