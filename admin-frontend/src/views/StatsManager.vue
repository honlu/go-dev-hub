<template>
  <div>
    <el-card>
      <template #header>
        <div style="display: flex; align-items: center; gap: 10px;">
          <span class="title-icon">📊</span>
          <span>访问统计</span>
        </div>
      </template>
      
      <!-- 统计卡片 -->
      <el-row :gutter="20" style="margin-bottom: 20px;">
        <el-col :span="6">
          <el-card shadow="hover" class="stat-card">
            <div class="stat-item">
              <div class="stat-icon">👥</div>
              <div class="stat-content">
                <div class="stat-value" id="busuanzi_value_site_uv">-</div>
                <div class="stat-label">总访客数</div>
              </div>
            </div>
          </el-card>
        </el-col>
        
        <el-col :span="6">
          <el-card shadow="hover" class="stat-card">
            <div class="stat-item">
              <div class="stat-icon">📈</div>
              <div class="stat-content">
                <div class="stat-value" id="busuanzi_value_site_pv">-</div>
                <div class="stat-label">总访问量</div>
              </div>
            </div>
          </el-card>
        </el-col>
        
        <el-col :span="6">
          <el-card shadow="hover" class="stat-card">
            <div class="stat-item">
              <div class="stat-icon">📝</div>
              <div class="stat-content">
                <div class="stat-value">{{ blogCount }}</div>
                <div class="stat-label">博客总数</div>
              </div>
            </div>
          </el-card>
        </el-col>
        
        <el-col :span="6">
          <el-card shadow="hover" class="stat-card">
            <div class="stat-item">
              <div class="stat-icon">❓</div>
              <div class="stat-content">
                <div class="stat-value">{{ questionCount }}</div>
                <div class="stat-label">题目总数</div>
              </div>
            </div>
          </el-card>
        </el-col>
      </el-row>
      
      <!-- 统计说明 -->
      <el-alert
        title="统计说明"
        type="info"
        :closable="false"
        style="margin-bottom: 20px;"
      >
        <template #default>
          <ul style="margin: 0; padding-left: 20px;">
            <li>访客数（UV）：统计独立访客数量，同一IP在一天内多次访问只计算一次</li>
            <li>访问量（PV）：统计页面访问次数，同一页面被多次访问会重复计算</li>
            <li>数据由不蒜子统计提供，实时更新</li>
            <li>博客和题目数据来自本地数据库</li>
          </ul>
        </template>
      </el-alert>
      
      <!-- 访问统计趋势图占位 -->
      <el-card shadow="hover" style="margin-top: 20px;">
        <template #header>
          <span>访问趋势（功能预留）</span>
        </template>
        <div style="height: 300px; display: flex; align-items: center; justify-content: center; color: #909399;">
          <div style="text-align: center;">
            <el-icon size="48"><TrendCharts /></el-icon>
            <div style="margin-top: 10px;">访问趋势图表功能预留</div>
            <div style="font-size: 12px; margin-top: 5px;">可集成ECharts等图表库实现</div>
          </div>
        </div>
      </el-card>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { fetchBlogs } from '../api/blog';
import { fetchQuestions } from '../api/question';
import { TrendCharts } from '@element-plus/icons-vue';

const blogCount = ref(0);
const questionCount = ref(0);

// 获取博客总数
async function getBlogCount() {
  try {
    const res = await fetchBlogs({ page: 1, page_size: 1 });
    blogCount.value = res.data.pagination.total;
  } catch (error) {
    console.error('获取博客总数失败:', error);
    blogCount.value = 0;
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

onMounted(async () => {
  await getBlogCount();
  await getQuestionCount();
});
</script>

<style scoped>
.stat-card {
  border: 2px solid var(--classic-green);
  background: linear-gradient(145deg, var(--classic-dark-gray), var(--classic-black));
  transition: all 0.3s ease;
}

.stat-card:hover {
  box-shadow: 0 0 20px rgba(0, 255, 0, 0.3);
  transform: translateY(-2px);
}

.stat-item {
  display: flex;
  align-items: center;
  gap: 15px;
  padding: 10px;
}

.stat-icon {
  font-size: 36px;
  width: 60px;
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--classic-green);
  border-radius: 50%;
  color: var(--classic-black);
}

.stat-content {
  flex: 1;
}

.stat-value {
  font-size: 28px;
  font-weight: bold;
  color: var(--classic-yellow);
  font-family: 'Courier New', monospace;
  text-shadow: 0 0 5px var(--classic-yellow);
  margin-bottom: 5px;
}

.stat-label {
  font-size: 14px;
  color: var(--classic-green);
  font-weight: bold;
  text-transform: uppercase;
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
</style> 