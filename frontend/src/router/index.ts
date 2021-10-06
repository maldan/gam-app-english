import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router';
import Main from '../page/Main.vue';
import Training from '../page/Training.vue';
import Statistics from '../page/Statistics.vue';

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Main',
    component: Main,
  },
  {
    path: '/training',
    name: 'Training',
    component: Training,
  },
  {
    path: '/statistics',
    name: 'Statistics',
    component: Statistics,
  },
];

const router = createRouter({
  history: createWebHashHistory(process.env.BASE_URL),
  routes,
});

export default router;
