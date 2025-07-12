# DevHub 脚本工具说明

## 根目录 scripts/
- start_all.sh：一键启动后端和前端
- backup_db.sh：数据库备份
- restore_db.sh：数据库恢复

## backend/scripts/
- migrate.sh：数据库迁移（可选）
- import_data.sh：数据导入（可选）

## admin-frontend/scripts/
- build.sh：一键打包前端
- lint.sh：代码检查

## web-frontend/scripts/
- build.sh：一键打包前端
- lint.sh：代码检查

---

### 示例：start_all.sh
```bash
#!/bin/bash
cd backend && nohup go run main.go models.go handlers.go > backend.log 2>&1 &
cd ../admin-frontend && npm run build
cd ../web-frontend && npm run build
echo "后端与前端已启动/打包"
```

### 示例：backup_db.sh
```bash
#!/bin/bash
PG_DSN=${POSTGRES_DSN:-"host=localhost user=postgres password=postgres dbname=devhub port=5432"}
pg_dump "$PG_DSN" > db_backup_$(date +%F).sql
echo "数据库已备份"
``` 