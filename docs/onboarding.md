# DevHub 新成员快速上手指南

## 1. 环境准备
- Go 1.21+
- Node.js 18+
- PostgreSQL 15+
- Git

## 2. 项目克隆与启动
```bash
git clone <repository-url>
cd dev_hub

# 启动后端
cd backend
go mod download
go run main.go models.go handlers.go

# 启动管理后台
cd ../admin-frontend
npm install
npm run dev

# 启动只读前端
cd ../web-frontend
npm install
npm run dev
```

## 3. 开发流程
1. 创建功能分支：`git checkout -b feature/xxx`
2. 开发功能，遵循代码规范
3. 提交代码：`git commit -m "feat: xxx"`
4. 推送分支：`git push origin feature/xxx`
5. 创建 Pull Request

## 4. 代码规范
- 后端：遵循 Go 官方代码规范
- 前端：使用 ESLint + Prettier
- 提交信息：使用 Conventional Commits

## 5. 常用命令
```bash
# 后端
go run main.go models.go handlers.go
go test ./...

# 前端
npm run dev
npm run build
npm run lint

# 数据库
./scripts/backup_db.sh
./scripts/restore_db.sh db_backup_2024-01-01.sql
```

## 6. 文档资源
- 项目总览：docs/overview.md
- API 文档：docs/api.md
- 数据库设计：docs/db.md
- 部署手册：docs/deploy.md
- 常见问题：docs/faq.md

## 7. 联系方式
- 技术问题：提 Issue 或联系项目负责人
- 文档更新：同步更新 docs/ 目录 