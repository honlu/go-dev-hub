# DevHub 部署手册

## 1. 本地开发环境
- 后端：
  ```bash
  cd backend
  go run main.go models.go handlers.go
  ```
- 管理后台：
  ```bash
  cd admin-frontend
  npm install
  npm run dev
  ```
- 只读前端：
  ```bash
  cd web-frontend
  npm install
  npm run dev
  ```

## 2. 生产环境部署
- 后端：
  ```bash
  cd backend
  go build -o devhub-server main.go models.go handlers.go
  ./devhub-server
  ```
- 前端（管理后台/只读端）：
  ```bash
  cd admin-frontend && npm run build
  cd web-frontend && npm run build
  # 用 nginx/Caddy/静态服务器托管 dist 目录
  ```
- nginx 配置：
  ```nginx
  server {
    listen 80;
    server_name your-domain.com;
    root /path/to/dist;
    location / {
      try_files $uri $uri/ /index.html;
    }
    location /api/ {
      proxy_pass http://localhost:8080;
    }
  }
  ```

## 3. docker-compose 部署（可选）
- 提供 docker-compose.yml，包含后端、前端、数据库服务

## 4. CI/CD
- 推荐 GitHub Actions/Jenkins/GitLab CI
- 自动化测试、构建、部署

## 5. 备份与恢复
- 数据库备份：见 scripts/backup_db.sh
- 数据库恢复：见 scripts/restore_db.sh 