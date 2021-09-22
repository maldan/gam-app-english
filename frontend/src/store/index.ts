import { createStore } from 'vuex';
import main, { MainStore } from './main';
import modal, { ModalStore } from './modal';
import word, { WordStore } from './word';

export type MainTree = { main: MainStore; word: WordStore; modal: ModalStore };

export default createStore({
  modules: { main, word, modal },
});
