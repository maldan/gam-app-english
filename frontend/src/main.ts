import { createApp } from 'vue';
import Store from './store';
import App from './App.vue';
import router from './router';
import UI from '../src/gam_sdk_ui/vue/ui';
import Event from '../src/gam_sdk_ui/vue/event';
import '../src/gam_sdk_ui/vue/main.scss';
import './main.scss';

createApp(App).use(router).use(UI).use(Event).use(Store).mount('#app');
