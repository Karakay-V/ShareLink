import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'
import UploadView from '../components/UploadView.vue'
import ResultsView from '../components/ResultsView.vue'
import SelectView from '../components/SelectView.vue'
import AuthView from '../components/Admin/AuthView.vue'
import FilesView from '../components/Admin/FilesView.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'select',
    component: SelectView,
    props: route => ({
        classroomProp: route.query.classroom,
    }),
  },
  {
    path: '/search',
    name: 'result',
    component: ResultsView,
    props: route => ({
        lesson: route.query.lesson,
        classroom: route.query.classroom,
    }),
  },
  {
    path: '/lesson/:lessonId',
    name: 'lesson',
    component: UploadView,
    props: route => ({
        lesson: route.params.lessonId,
        classroom: route.query.classroom,
    }),
  },
  {
    path: '/auth',
    name: 'auth',
    component: AuthView,
  },
  {
    path: '/files',
    name: 'files',
    component: FilesView,
    props: route => ({
        classroomProp: route.query.classroom,
    }),
  },
  {
    path: '/:pathMatch(.*)*',
    redirect: '/',
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
