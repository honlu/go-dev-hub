<script setup>
import HelloWorld from './components/HelloWorld.vue'
</script>

<template>
  <div id="app">
    <el-container style="height: 100vh; background: transparent;">
      <el-aside width="200px" class="retro-sidebar">
        <div class="retro-menu">
          <div class="retro-menu-item" 
               :class="{ active: $route.path === '/blogs' }"
               @click="$router.push('/blogs')">
            <span class="menu-icon">📝</span>
            博客列表
          </div>
          <div class="retro-menu-item" 
               :class="{ active: $route.path === '/questions' }"
               @click="$router.push('/questions')">
            <span class="menu-icon">❓</span>
            题库列表
          </div>
          <div class="retro-menu-item" 
               :class="{ active: $route.path === '/stats' }"
               @click="$router.push('/stats')">
            <span class="menu-icon">📊</span>
            网站统计
          </div>
        </div>
      </el-aside>
      <el-container>
        <el-header class="retro-header">
          <div class="retro-title">
            <span class="title-icon">🖥️</span>
            DevHub 用户前台
            <span class="title-subtitle">v1.0.0</span>
          </div>
        </el-header>
        <el-main class="retro-main">
          <router-view />
        </el-main>
        
        <!-- 底部统计信息 -->
        <el-footer height="auto" class="retro-footer">
          <div class="footer-stats">
            <div class="stats-container">
              <div class="stat-item">
                <span class="stat-icon">👥</span>
                <span class="stat-label">总访客数：</span>
                <span class="stat-value" id="busuanzi_value_site_uv">-</span>
              </div>
              <div class="stat-divider">|</div>
              <div class="stat-item">
                <span class="stat-icon">📈</span>
                <span class="stat-label">总访问量：</span>
                <span class="stat-value" id="busuanzi_value_site_pv">-</span>
              </div>
              <div class="stat-divider">|</div>
              <div class="stat-item">
                <span class="stat-icon">📄</span>
                <span class="stat-label">本页访问：</span>
                <span class="stat-value" id="busuanzi_value_page_pv">-</span>
              </div>
            </div>
          </div>
        </el-footer>
      </el-container>
    </el-container>
  </div>
</template>

<style scoped>
/* 侧边栏样式 */
.retro-sidebar {
  background: var(--classic-black) !important;
  border-right: var(--border-width) solid var(--classic-green);
  box-shadow: 2px 0 10px rgba(0, 255, 0, 0.2);
  position: relative;
  z-index: 2;
}

.retro-sidebar::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 2px;
  background: linear-gradient(90deg, var(--classic-green), var(--classic-cyan), var(--classic-green));
  animation: retro-pulse 3s ease-in-out infinite;
}

/* 菜单样式 */
.retro-menu {
  background: transparent;
  border: none;
  box-shadow: none;
  padding: 20px 0;
}

.retro-menu-item {
  color: var(--classic-green);
  padding: 15px 20px;
  font-family: 'Courier New', monospace;
  font-weight: bold;
  font-size: 14px;
  text-transform: uppercase;
  cursor: pointer;
  transition: all 0.3s ease;
  text-shadow: 0 0 3px var(--classic-green);
  border-bottom: 1px solid var(--classic-dark-gray);
  margin: 5px 10px;
  border-radius: var(--border-radius);
  position: relative;
  overflow: hidden;
}

.retro-menu-item::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(0, 255, 0, 0.1), transparent);
  transition: left 0.5s ease;
}

.retro-menu-item:hover::before {
  left: 100%;
}

.retro-menu-item:hover {
  background: var(--classic-dark-gray);
  color: var(--classic-yellow);
  text-shadow: 0 0 5px var(--classic-yellow);
  box-shadow: 
    inset 0 0 10px rgba(255, 255, 0, 0.2),
    0 0 10px rgba(255, 255, 0, 0.3);
  transform: translateX(5px);
}

.retro-menu-item.active {
  background: var(--classic-green);
  color: var(--classic-black);
  text-shadow: none;
  box-shadow: 
    inset 0 0 10px rgba(0, 0, 0, 0.3),
    0 0 15px rgba(0, 255, 0, 0.5);
  font-weight: bold;
}

.menu-icon {
  margin-right: 10px;
  font-size: 16px;
  animation: retro-pulse 2s ease-in-out infinite;
}

/* 头部样式 */
.retro-header {
  background: var(--classic-black) !important;
  border-bottom: var(--border-width) solid var(--classic-green);
  box-shadow: 0 2px 10px rgba(0, 255, 0, 0.2);
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
  z-index: 2;
}

.retro-header::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 2px;
  background: linear-gradient(90deg, var(--classic-green), var(--classic-cyan), var(--classic-green));
  animation: retro-pulse 3s ease-in-out infinite;
}

.retro-title {
  color: var(--classic-yellow);
  font-family: 'Courier New', monospace;
  font-weight: bold;
  font-size: 24px;
  text-transform: uppercase;
  text-align: center;
  text-shadow: 0 0 5px var(--classic-yellow);
  animation: text-glow 3s ease-in-out infinite;
  display: flex;
  align-items: center;
  gap: 15px;
}

.title-icon {
  font-size: 28px;
  animation: retro-pulse 1.5s ease-in-out infinite;
}

.title-subtitle {
  font-size: 12px;
  color: var(--classic-green);
  text-shadow: 0 0 3px var(--classic-green);
  margin-left: 10px;
  opacity: 0.8;
}

/* 主内容区域 */
.retro-main {
  background: transparent !important;
  position: relative;
  z-index: 2;
  padding: 20px;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .retro-title {
    font-size: 18px;
    flex-direction: column;
    gap: 5px;
  }
  
  .title-icon {
    font-size: 20px;
  }
  
  .retro-menu-item {
    font-size: 12px;
    padding: 10px 15px;
  }
}

/* 深色模式优化 */
@media (prefers-color-scheme: dark) {
  .retro-sidebar {
    background: var(--classic-black) !important;
  }
}

/* 减少动画模式 */
@media (prefers-reduced-motion: reduce) {
  .retro-menu-item::before {
    display: none;
  }
  
  .retro-menu-item:hover {
    transform: none;
  }
}

/* 底部统计样式 */
.retro-footer {
  background: var(--classic-black) !important;
  border-top: var(--border-width) solid var(--classic-green);
  padding: 15px 20px;
  position: relative;
  z-index: 2;
}

.retro-footer::before {
  content: '';
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  height: 2px;
  background: linear-gradient(90deg, var(--classic-green), var(--classic-cyan), var(--classic-green));
  animation: retro-pulse 3s ease-in-out infinite;
}

.footer-stats {
  display: flex;
  justify-content: center;
  align-items: center;
}

.stats-container {
  display: flex;
  align-items: center;
  gap: 20px;
  font-family: 'Courier New', monospace;
  font-size: 14px;
}

.stat-item {
  display: flex;
  align-items: center;
  gap: 5px;
  color: var(--classic-green);
}

.stat-icon {
  font-size: 16px;
  animation: retro-pulse 2s ease-in-out infinite;
}

.stat-label {
  color: var(--classic-green);
  font-weight: bold;
}

.stat-value {
  color: var(--classic-yellow);
  font-weight: bold;
  text-shadow: 0 0 3px var(--classic-yellow);
  min-width: 30px;
  text-align: center;
}

.stat-divider {
  color: var(--classic-green);
  font-weight: bold;
  opacity: 0.6;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .stats-container {
    flex-direction: column;
    gap: 10px;
    font-size: 12px;
  }
  
  .stat-divider {
    display: none;
  }
}
</style>

<style>
/* 全局样式覆盖 */
.el-container {
  background: transparent !important;
}

.el-aside {
  background: transparent !important;
}

.el-header {
  background: transparent !important;
  border-bottom: none !important;
}

.el-main {
  background: transparent !important;
}

/* 隐藏Element Plus默认样式 */
.el-menu {
  background: transparent !important;
  border: none !important;
}

.el-menu-item {
  background: transparent !important;
  border: none !important;
  color: inherit !important;
}

.el-menu-item:hover {
  background: transparent !important;
  color: inherit !important;
}

.el-menu-item.is-active {
  background: transparent !important;
  color: inherit !important;
}
</style>
