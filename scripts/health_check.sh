#!/bin/bash

echo "DevHub 健康检查..."

# 检查后端服务
echo "检查后端服务..."
if curl -f http://localhost:8080/api/web/blogs > /dev/null 2>&1; then
    echo "✓ 后端服务正常"
else
    echo "✗ 后端服务异常"
    exit 1
fi

# 检查数据库连接
echo "检查数据库连接..."
PG_DSN=${POSTGRES_DSN:-"host=localhost user=postgres password=postgres dbname=devhub port=5432"}
if psql "$PG_DSN" -c "SELECT 1;" > /dev/null 2>&1; then
    echo "✓ 数据库连接正常"
else
    echo "✗ 数据库连接异常"
    exit 1
fi

# 检查前端服务（可选）
echo "检查前端服务..."
if curl -f http://localhost:3000 > /dev/null 2>&1; then
    echo "✓ 管理后台正常"
else
    echo "⚠ 管理后台未启动"
fi

if curl -f http://localhost:3001 > /dev/null 2>&1; then
    echo "✓ 只读前端正常"
else
    echo "⚠ 只读前端未启动"
fi

echo "健康检查完成" 