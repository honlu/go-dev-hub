#!/bin/bash
if [ -z "$1" ]; then
  echo "用法: $0 <备份文件.sql>"
  exit 1
fi
PG_DSN=${POSTGRES_DSN:-"host=localhost user=postgres password=postgres dbname=devhub port=5432"}
psql "$PG_DSN" < "$1"
echo "数据库已恢复" 