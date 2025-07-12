import axios from 'axios';

const service = axios.create({
    baseURL: '/api/web', // 只读接口前缀
    timeout: 10000,
});

service.interceptors.response.use(
    response => response.data,
    error => {
        return Promise.reject(error.response?.data || error);
    }
);

// 获取分类枚举
export function fetchCategories() {
    return service.get('/categories');
}

export default service; 