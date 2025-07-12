# DevHub —— 博客与题库统一平台

## 项目简介
DevHub 是一个面向技术团队和个人的全栈知识平台，集成了技术博客写作与展示、技术题库管理与练习，支持后台管理和前台只读展示，助力知识沉淀与技能成长。

## 技术架构
- **前端管理后台**：Vue3 + Vite + Element Plus（目录：admin-frontend）
- **前端只读展示**：Vue3 + Vite + Element Plus（目录：web-frontend）
- **后端服务**：Go + Gin + GORM + PostgreSQL（目录：backend）
- **数据库**：PostgreSQL
- **自动化/运维**：Docker、docker-compose、GitHub Actions、脚本工具

## 功能亮点
- 博客系统：Markdown 编辑、分类、标签、发布、搜索、详情页
- 题库系统：题目管理、难度、分类、标签、Markdown 支持、练习/展示
- 后台管理：可视化增删改查、表格、筛选、批量操作
- 前台展示：只读、筛选、分页、内容美观
- API 设计：RESTful，分离 /api/admin/ 与 /api/web/
- 自动化部署与备份、健康检查、CI/CD

## 目录结构
```
dev_hub/
├── backend/              # Go + Gin 后端服务
│   ├── scripts/          # 后端脚本
├── admin-frontend/       # Vue3 管理后台
│   ├── scripts/          # 前端脚本
├── web-frontend/         # Vue3 只读展示端
│   ├── scripts/          # 前端脚本
├── docs/                 # 项目文档（API、DB、部署、FAQ等）
├── scripts/              # 根目录通用脚本（启动、备份、健康检查等）
├── .github/workflows/    # CI/CD 配置
├── docker-compose.yml    # 一键部署
├── env.example           # 环境变量模板
└── readme.md             # 项目说明（本文件）
```

## 快速上手
```bash
# 克隆项目
git clone <repository-url>
cd go-dev-hub

# 先安装好db、创建库和表、配置登录密码，具体看backend/.env or backend/.main.go中连接串配置
# 一键启动开发环境
./scripts/start_all.sh

# 访问
# 管理后台：http://localhost:5173 或 http://localhost:3000（Docker）
# 只读前端：http://localhost:3001（Docker）
# 后端 API：http://localhost:8080
```

## 生产部署
- 推荐使用 Docker 或 docker-compose 一键部署
```
# 运行构建
docker-compose up --build
# 如果构建过程出现网络问题，可以暂时关闭BuildKit，执行
DOCKER_BUILDKIT=0 docker-compose build
docker-compose up
# 可能会产生「悬空镜像」
docker builder prune # 清除
```
- 详见 `docs/deploy.md`、`docker-compose.yml`、`scripts/deploy.sh`

## 自动化与运维
- CI/CD：GitHub Actions 自动测试、构建、部署
- 健康检查：`./scripts/health_check.sh`
- 数据库备份/恢复：`./scripts/backup_db.sh`、`./scripts/restore_db.sh`
- Swagger API 文档（可选）：见 backend/swagger.go

## 文档入口
- 项目总览：`docs/overview.md`
- API 文档：`docs/api.md`
- 数据库设计：`docs/db.md`
- 部署手册：`docs/deploy.md`
- 脚本说明：`docs/scripts.md`
- 常见问题：`docs/faq.md`
- 新手指南：`docs/onboarding.md`

## 团队协作
- 代码规范、分支管理、提交规范见 `docs/onboarding.md`
- 欢迎 issue、PR 及功能建议

---
如需扩展权限系统、全文检索、移动端适配、SSR、CI/CD、监控等，详见 docs/ 目录或联系项目维护者。