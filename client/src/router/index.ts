import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'
import UploadView from '../components/UploadView.vue'
import SearchView from '../components/SearchView.vue'
import ResultsView from '../components/ResultsView.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'search',
    component: SearchView,
  },
  {
    path: '/search',
    name: 'result',
    component: ResultsView,
    props: route => ({
        lessonId: route.query.lesson,
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
    path: '/:pathMatch(.*)*',
    redirect: '/',
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
