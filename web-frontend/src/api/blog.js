import service from './index';

export function fetchBlogs(params) {
    return service.get('/blogs', { params });
}

export function fetchBlog(id) {
    return service.get(`/blogs/${id}`);
} 