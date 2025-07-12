import service from './index';

export function fetchQuestions(params) {
    return service.get('/questions', { params });
}

export function fetchQuestion(id) {
    return service.get(`/questions/${id}`);
}

export function createQuestion(data) {
    return service.post('/questions', data);
}

export function updateQuestion(id, data) {
    return service.put(`/questions/${id}`, data);
}

export function deleteQuestion(id) {
    return service.delete(`/questions/${id}`);
} 