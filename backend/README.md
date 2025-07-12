# DevHub Backend

DevHub 后端服务，基于 Go + Gin + GORM + PostgreSQL 构建。

## 项目结构

```
backend/
├── main.go          # 主程序入口
├── models.go        # 数据模型定义
├── handlers.go      # API 处理函数
├── go.mod           # Go 模块文件
└── README.md        # 项目说明
```

## 数据模型

### Blog（博客）
- `id`: 主键
- `title`: 标题
- `content_md`: Markdown 内容
- `content_html`: HTML 内容（可选）
- `category`: 分类
- `tags`: 标签数组
- `status`: 状态（draft/published）
- `created_at`: 创建时间
- `updated_at`: 更新时间

### Question（题目）
- `id`: 主键
- `title`: 题目
- `answer`: 答案
- `category`: 分类
- `tags`: 标签数组
- `difficulty`: 难度
- `created_at`: 创建时间
- `updated_at`: 更新时间

## API 接口

### 管理接口（/api/admin/）

#### 博客管理
- `POST /api/admin/blogs` - 创建博客
- `GET /api/admin/blogs` - 获取博客列表（支持分页、分类、状态筛选）
- `GET /api/admin/blogs/:id` - 获取单个博客
- `PUT /api/admin/blogs/:id` - 更新博客
- `DELETE /api/admin/blogs/:id` - 删除博客

#### 题库管理
- `POST /api/admin/questions` - 创建题目
- `GET /api/admin/questions` - 获取题目列表（支持分页、分类、难度筛选）
- `GET /api/admin/questions/:id` - 获取单个题目
- `PUT /api/admin/questions/:id` - 更新题目
- `DELETE /api/admin/questions/:id` - 删除题目

### 只读接口（/api/web/）

#### 博客展示
- `GET /api/web/blogs` - 获取已发布的博客列表（支持分页、分类筛选）
- `GET /api/web/blogs/:id` - 获取单个已发布的博客

#### 题库展示
- `GET /api/web/questions` - 获取题目列表（支持分页、分类、难度筛选）
- `GET /api/web/questions/:id` - 获取单个题目

## 响应格式

```json
{
  "code": 200,
  "message": "操作成功",
  "data": {
    // 具体数据
  }
}
```

## 环境配置

### 数据库配置
通过环境变量 `POSTGRES_DSN` 配置数据库连接，默认值：
```
host=localhost user=admin password=password_123456 dbname=mydb port=5432 sslmode=disable TimeZone=Asia/Shanghai
```

### 启动服务
```bash
# 设置数据库连接（可选）
export POSTGRES_DSN="your_database_connection_string"

# 启动服务
go run main.go models.go handlers.go
```

服务将在 `http://localhost:8080` 启动。

## 开发说明

1. 项目使用 GORM 自动迁移数据表，即在数据库中创建表 
2. 支持 PostgreSQL 的 `text[]` 类型存储标签
3. 统一错误处理和响应格式
4. 支持分页查询和条件筛选
5. 只读接口仅返回已发布的博客内容 