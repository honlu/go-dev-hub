#!/bin/bash
PG_DSN=${POSTGRES_DSN:-"host=localhost user=postgres password=postgres dbname=devhub port=5432"}
pg_dump "$PG_DSN" > db_backup_$(date +%F).sql
echo "数据库已备份" 