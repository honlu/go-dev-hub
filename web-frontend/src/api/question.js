import service from './index';

export function fetchQuestions(params) {
    return service.get('/questions', { params });
}

export function fetchQuestion(id) {
    return service.get(`/questions/${id}`);
} 