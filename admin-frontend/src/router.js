import { createRouter, createWebHistory } from 'vue-router';
import BlogManager from './views/BlogManager.vue';
import QuestionManager from './views/QuestionManager.vue';

const routes = [
    { path: '/', redirect: '/blogs' },
    { path: '/blogs', component: BlogManager, name: '博客管理' },
    { path: '/questions', component: QuestionManager, name: '题库管理' },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router; 