#!/bin/bash
set -e

echo "开始部署 DevHub..."

# 备份数据库
echo "备份数据库..."
./scripts/backup_db.sh

# 构建后端
echo "构建后端..."
cd backend
go build -o devhub-server main.go models.go handlers.go
cd ..

# 构建前端
echo "构建管理后台..."
cd admin-frontend
npm run build
cd ..

echo "构建只读前端..."
cd web-frontend
npm run build
cd ..

# 停止旧服务
echo "停止旧服务..."
pkill -f devhub-server || true

# 启动新服务
echo "启动新服务..."
cd backend
nohup ./devhub-server > ../backend.log 2>&1 &
cd ..

# 健康检查
echo "健康检查..."
sleep 5
if curl -f http://localhost:8080/api/web/blogs > /dev/null 2>&1; then
    echo "部署成功！"
else
    echo "部署失败，请检查日志"
    exit 1
fi 