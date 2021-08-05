import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router';
import Main from '../page/Main.vue';
import WorkList from '../page/WorkList.vue';
import Work from '../page/Work.vue';
import Training from '../page/Training.vue';
import RefDB from '../page/RefDB.vue';

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Main',
    component: Main,
  },
  {
    path: '/workList',
    name: 'WorkList',
    component: WorkList,
  },
  {
    path: '/work/:id',
    name: 'Work',
    component: Work,
  },
  {
    path: '/training',
    name: 'Training',
    component: Training,
  },
  {
    path: '/refDB',
    name: 'RefDB',
    component: RefDB,
  },
];

const router = createRouter({
  history: createWebHashHistory(process.env.BASE_URL),
  routes,
});

export default router;
