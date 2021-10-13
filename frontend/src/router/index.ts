import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router';
import Main from '../page/Main.vue';
import ArtList from '../page/ArtList.vue';
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
    path: '/art',
    name: 'Art List',
    component: ArtList,
  },
  {
    path: '/art/:id',
    name: 'Art',
    component: Work,
  },
  {
    path: '/training',
    name: 'Training',
    component: Training,
  },
  {
    path: '/db',
    name: 'DB',
    component: RefDB,
  },
];

const router = createRouter({
  history: createWebHashHistory(process.env.BASE_URL),
  routes,
});

export default router;
