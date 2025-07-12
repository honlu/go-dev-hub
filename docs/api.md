# DevHub 后端 API 文档

## 1. 博客接口
### 管理接口（/api/admin/blogs）
- `GET /api/admin/blogs`：获取博客列表（分页、分类、状态筛选）
- `POST /api/admin/blogs`：新建博客
- `GET /api/admin/blogs/:id`：获取单个博客
- `PUT /api/admin/blogs/:id`：更新博客
- `DELETE /api/admin/blogs/:id`：删除博客

### 只读接口（/api/web/blogs）
- `GET /api/web/blogs`：获取已发布博客列表（分页、分类筛选）
- `GET /api/web/blogs/:id`：获取单个已发布博客

## 2. 题库接口
### 管理接口（/api/admin/questions）
- `GET /api/admin/questions`：获取题目列表（分页、分类、难度筛选）
- `POST /api/admin/questions`：新建题目
- `GET /api/admin/questions/:id`：获取单个题目
- `PUT /api/admin/questions/:id`：更新题目
- `DELETE /api/admin/questions/:id`：删除题目

### 只读接口（/api/web/questions）
- `GET /api/web/questions`：获取题目列表（分页、分类、难度筛选）
- `GET /api/web/questions/:id`：获取单个题目

## 3. 通用响应格式
```json
{
  "code": 200,
  "message": "操作成功",
  "data": { }
}
```

## 4. 错误码说明
- 200：成功
- 400：参数错误
- 404：未找到
- 500：服务器错误

## 5. 示例
### 获取博客列表
```
GET /api/web/blogs?page=1&page_size=10&category=Go
```
响应：
```json
{
  "code": 200,
  "message": "查询成功",
  "data": {
    "blogs": [
      { "id": 1, "title": "xxx", ... }
    ],
    "pagination": { "page": 1, "page_size": 10, "total": 100 }
  }
}
``` 