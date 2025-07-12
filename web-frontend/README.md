# DevHub 前台只读展示（web-frontend）

基于 Vue3 + Vite + Element Plus 的博客与题库只读展示前端。

## 功能亮点
- 博客/题库列表与详情，支持多条件筛选、分页
- Markdown 渲染，内容美观易读
- 响应式布局，适配主流 PC 浏览器
- 仅展示和查询，无需登录、无编辑功能，适合公开访问
- 对接后端 `/api/web/` 只读接口，安全高效

## 技术栈
- Vue3
- Vite
- Element Plus
- Axios
- markdown-it

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
     root /path/to/web-frontend/dist;
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
web-frontend/
├── src/
│   ├── views/         # 页面
│   ├── api/           # 接口封装
│   └── ...
├── public/
├── package.json
└── README.md
```

## 常见问题
- 前端接口 404：请确保 nginx 配置了 history 路由和 /api 代理
- 跨域问题：开发环境已同源，生产建议前后端同域或配置 CORS

## SEO/移动端适配建议
- 如需 SEO，可用 SSR（如 Nuxt、Vite SSR）或预渲染方案
- 如需移动端适配，可引入响应式 CSS 框架或自定义媒体查询

## 后端接口文档
详见 `../backend/README.md`。
