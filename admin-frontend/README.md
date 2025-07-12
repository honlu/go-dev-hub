# DevHub 管理后台（admin-frontend）

基于 Vue3 + Vite + Element Plus 的博客与题库管理后台。

## 功能模块
- 博客管理（增删改查、Markdown 编辑）
- 题库管理（增删改查、Markdown 编辑）
- 支持分类、标签、难度筛选
- 后端接口对接 `/api/admin/`

## 技术栈
- Vue3
- Vite
- Element Plus
- Axios

## 启动方式
```bash
npm install
npm run dev
```

默认开发端口：`http://localhost:5173`

## 生产部署
1. 构建静态文件：
   ```bash
   npm run build
   ```
2. 产物在 `dist/` 目录，可用 nginx、Caddy、静态服务器等部署
3. 典型 nginx 配置：
   ```nginx
   server {
     listen 80;
     server_name your-domain.com;
     root /path/to/admin-frontend/dist;
     location / {
       try_files $uri $uri/ /index.html;
     }
     location /api/ {
       proxy_pass http://localhost:8080;
     }
   }
   ```

## 目录结构
```
admin-frontend/
├── src/
│   ├── views/         # 页面
│   ├── components/    # 组件
│   ├── api/           # 接口封装
│   └── ...
├── public/
├── package.json
└── README.md
```

## 常见问题
- 前端接口 404：请确保 nginx 配置了 history 路由和 /api 代理
- 跨域问题：开发环境已同源，生产建议前后端同域或配置 CORS

## 后端接口文档
详见 `../backend/README.md`。

# Vue 3 + Vite

This template should help get you started developing with Vue 3 in Vite. The template uses Vue 3 `<script setup>` SFCs, check out the [script setup docs](https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup) to learn more.

Learn more about IDE Support for Vue in the [Vue Docs Scaling up Guide](https://vuejs.org/guide/scaling-up/tooling.html#ide-support).
