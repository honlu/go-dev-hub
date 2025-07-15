<template>
  <div>
    <el-card>
      <template #header>
        <div style="display: flex; align-items: center; gap: 10px;">
          <span class="title-icon">📊</span>
          <span>网站统计</span>
        </div>
      </template>
      
      <!-- 统计概览 -->
      <el-row :gutter="20" style="margin-bottom: 30px;">
        <el-col :span="8">
          <div class="stat-card">
            <div class="stat-item">
              <div class="stat-icon">👥</div>
              <div class="stat-content">
                <div class="stat-value" id="busuanzi_value_site_uv">-</div>
                <div class="stat-label">总访客数</div>
              </div>
            </div>
          </div>
        </el-col>
        
        <el-col :span="8">
          <div class="stat-card">
            <div class="stat-item">
              <div class="stat-icon">📈</div>
              <div class="stat-content">
                <div class="stat-value" id="busuanzi_value_site_pv">-</div>
                <div class="stat-label">总访问量</div>
              </div>
            </div>
          </div>
        </el-col>
        
        <el-col :span="8">
          <div class="stat-card">
            <div class="stat-item">
              <div class="stat-icon">📄</div>
              <div class="stat-content">
                <div class="stat-value" id="busuanzi_value_page_pv">-</div>
                <div class="stat-label">本页访问</div>
              </div>
            </div>
          </div>
        </el-col>
      </el-row>
      
      <!-- 内容统计 -->
      <el-row :gutter="20" style="margin-bottom: 30px;">
        <el-col :span="12">
          <div class="content-stat-card">
            <div class="content-stat-header">
              <span class="content-icon">📝</span>
              <span>博客内容</span>
            </div>
            <div class="content-stat-body">
              <div class="content-stat-item">
                <span class="content-stat-label">已发布博客：</span>
                <span class="content-stat-value">{{ publishedBlogCount }}</span>
              </div>
              <div class="content-stat-item">
                <span class="content-stat-label">博客分类：</span>
                <span class="content-stat-value">{{ categoryCount }}</span>
              </div>
            </div>
          </div>
        </el-col>
        
        <el-col :span="12">
          <div class="content-stat-card">
            <div class="content-stat-header">
              <span class="content-icon">❓</span>
              <span>题库内容</span>
            </div>
            <div class="content-stat-body">
              <div class="content-stat-item">
                <span class="content-stat-label">题目总数：</span>
                <span class="content-stat-value">{{ questionCount }}</span>
              </div>
              <div class="content-stat-item">
                <span class="content-stat-label">题目分类：</span>
                <span class="content-stat-value">{{ categoryCount }}</span>
              </div>
            </div>
          </div>
        </el-col>
      </el-row>
      
      <!-- 统计说明 -->
      <el-alert
        title="统计说明"
        type="info"
        :closable="false"
        style="margin-top: 20px;"
      >
        <template #default>
          <div style="color: var(--classic-green);">
            <p style="margin: 0 0 10px 0;">📊 <strong>访问统计</strong></p>
            <ul style="margin: 0; padding-left: 20px;">
              <li><strong>访客数（UV）</strong>：统计独立访客数量，同一IP在一天内多次访问只计算一次</li>
              <li><strong>访问量（PV）</strong>：统计页面访问次数，同一页面被多次访问会重复计算</li>
              <li><strong>本页访问</strong>：当前页面的访问次数</li>
            </ul>
            <p style="margin: 15px 0 0 0;">🔧 <strong>数据来源</strong>：访问统计数据由不蒜子统计提供，内容统计数据来自网站数据库。</p>
          </div>
        </template>
      </el-alert>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { fetchBlogs } from '../api/blog';
import { fetchQuestions } from '../api/question';
import { fetchCategories } from '../api/index';

const publishedBlogCount = ref(0);
const questionCount = ref(0);
const categoryCount = ref(0);

// 获取已发布博客总数
async function getPublishedBlogCount() {
  try {
    const res = await fetchBlogs({ page: 1, page_size: 1 });
    publishedBlogCount.value = res.data.pagination.total;
  } catch (error) {
    console.error('获取博客总数失败:', error);
    publishedBlogCount.value = 0;
  }
}

// 获取题目总数
async function getQuestionCount() {
  try {
    const res = await fetchQuestions({ page: 1, page_size: 1 });
    questionCount.value = res.data.pagination.total;
  } catch (error) {
    console.error('获取题目总数失败:', error);
    questionCount.value = 0;
  }
}

// 获取分类总数
async function getCategoryCount() {
  try {
    const res = await fetchCategories();
    if (res.code === 200) {
      categoryCount.value = res.data.categories.length;
    }
  } catch (error) {
    console.error('获取分类总数失败:', error);
    categoryCount.value = 0;
  }
}

onMounted(async () => {
  await getPublishedBlogCount();
  await getQuestionCount();
  await getCategoryCount();
});
</script>

<style scoped>
.stat-card {
  background: linear-gradient(145deg, var(--classic-dark-gray), var(--classic-black));
  border: 2px solid var(--classic-green);
  border-radius: var(--border-radius);
  padding: 20px;
  transition: all 0.3s ease;
  position: relative;
  overflow: hidden;
}

.stat-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 2px;
  background: linear-gradient(90deg, var(--classic-green), var(--classic-cyan), var(--classic-green));
  animation: retro-pulse 3s ease-in-out infinite;
}

.stat-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 10px 20px rgba(0, 255, 0, 0.2);
}

.stat-item {
  display: flex;
  align-items: center;
  gap: 15px;
}

.stat-icon {
  font-size: 40px;
  width: 70px;
  height: 70px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--classic-green);
  border-radius: 50%;
  color: var(--classic-black);
  animation: retro-pulse 2s ease-in-out infinite;
}

.stat-content {
  flex: 1;
  text-align: center;
}

.stat-value {
  font-size: 36px;
  font-weight: bold;
  color: var(--classic-yellow);
  font-family: 'Courier New', monospace;
  text-shadow: 0 0 8px var(--classic-yellow);
  margin-bottom: 8px;
  display: block;
}

.stat-label {
  font-size: 14px;
  color: var(--classic-green);
  font-weight: bold;
  text-transform: uppercase;
  letter-spacing: 1px;
}

.content-stat-card {
  background: linear-gradient(145deg, var(--classic-dark-gray), var(--classic-black));
  border: 2px solid var(--classic-green);
  border-radius: var(--border-radius);
  overflow: hidden;
  transition: all 0.3s ease;
}

.content-stat-card:hover {
  transform: translateY(-3px);
  box-shadow: 0 8px 16px rgba(0, 255, 0, 0.2);
}

.content-stat-header {
  background: var(--classic-green);
  color: var(--classic-black);
  padding: 12px 16px;
  font-weight: bold;
  font-size: 16px;
  display: flex;
  align-items: center;
  gap: 10px;
}

.content-icon {
  font-size: 20px;
}

.content-stat-body {
  padding: 20px;
}

.content-stat-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
  padding: 8px 0;
  border-bottom: 1px solid var(--classic-dark-gray);
}

.content-stat-item:last-child {
  margin-bottom: 0;
  border-bottom: none;
}

.content-stat-label {
  color: var(--classic-green);
  font-weight: bold;
  font-size: 14px;
}

.content-stat-value {
  color: var(--classic-yellow);
  font-weight: bold;
  font-size: 18px;
  font-family: 'Courier New', monospace;
  text-shadow: 0 0 5px var(--classic-yellow);
}

.title-icon {
  font-size: 20px;
  animation: retro-pulse 2s ease-in-out infinite;
}

/* 深色主题适配 */
:deep(.el-card) {
  background: var(--classic-black);
  border-color: var(--classic-green);
}

:deep(.el-card__header) {
  background: var(--classic-dark-gray);
  border-bottom-color: var(--classic-green);
  color: var(--classic-green);
  font-weight: bold;
}

:deep(.el-card__body) {
  background: var(--classic-black);
  color: var(--classic-green);
}

:deep(.el-alert) {
  background: var(--classic-dark-gray);
  border-color: var(--classic-green);
}

:deep(.el-alert__title) {
  color: var(--classic-green);
}

:deep(.el-alert__content) {
  color: var(--classic-green);
}

/* 响应式设计 */
@media (max-width: 768px) {
  .stat-item {
    flex-direction: column;
    text-align: center;
    gap: 10px;
  }
  
  .stat-icon {
    font-size: 30px;
    width: 50px;
    height: 50px;
  }
  
  .stat-value {
    font-size: 24px;
  }
  
  .content-stat-item {
    flex-direction: column;
    align-items: flex-start;
    gap: 5px;
  }
}
</style> 