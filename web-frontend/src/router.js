import { createRouter, createWebHistory } from 'vue-router';
import BlogList from './views/BlogList.vue';
import BlogDetail from './views/BlogDetail.vue';
import QuestionList from './views/QuestionList.vue';
import QuestionDetail from './views/QuestionDetail.vue';
import StatsView from './views/StatsView.vue';

const routes = [
    { path: '/', redirect: '/blogs' },
    { path: '/blogs', component: BlogList, name: '博客列表' },
    { path: '/blogs/:id', component: BlogDetail, name: '博客详情' },
    { path: '/questions', component: QuestionList, name: '题库列表' },
    { path: '/questions/:id', component: QuestionDetail, name: '题目详情' },
    { path: '/stats', component: StatsView, name: '网站统计' },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router; 