import service from './index';

export function fetchBlogs(params) {
    return service.get('/blogs', { params });
}

export function fetchBlog(id) {
    return service.get(`/blogs/${id}`);
}

export function createBlog(data) {
    return service.post('/blogs', data);
}

export function updateBlog(id, data) {
    return service.put(`/blogs/${id}`, data);
}

export function deleteBlog(id) {
    return service.delete(`/blogs/${id}`);
} 