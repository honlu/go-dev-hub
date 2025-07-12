#!/bin/bash
cd $(dirname $0)/..
cd backend && go run main.go models.go handlers.go
cd ../admin-frontend && npm run build
cd ../web-frontend && npm run build
echo "后端与前端已启动/打包" 