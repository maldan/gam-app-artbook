import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router';
import Statistics from '../page/Statistics.vue';
import Project from '../page/Project.vue';
import Art from '../page/Art.vue';
import Training from '../page/Training.vue';
import Reference from '../page/Reference.vue';

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Statistics',
    component: Statistics,
  },
  {
    path: '/project',
    name: 'Project',
    component: Project,
  },
  {
    path: '/project/:id',
    name: 'Art',
    component: Art,
  },
  {
    path: '/training',
    name: 'Training',
    component: Training,
  },
  {
    path: '/reference',
    name: 'Reference',
    component: Reference,
  },
];

const router = createRouter({
  history: createWebHashHistory(process.env.BASE_URL),
  routes,
});

export default router;
