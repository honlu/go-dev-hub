<template>
  <div>
    <el-card>
      <div style="display: flex; justify-content: space-between; align-items: center; margin-bottom: 16px;">
        <div>
          <el-input v-model="search.title" placeholder="搜索题目" style="width: 200px; margin-right: 8px;" @keyup.enter="handleTitleSearch" />
          <el-select v-model="search.category" placeholder="分类" clearable style="width: 120px; margin-right: 8px;">
            <el-option v-for="item in categoryOptions" :key="item" :label="item" :value="item" />
          </el-select>
          <el-select v-model="search.difficulty" placeholder="难度" clearable style="width: 120px;">
            <el-option v-for="item in difficultyOptions" :key="item.value" :label="item.label" :value="item.value" />
          </el-select>
        </div>
      </div>
      <el-table :data="list" style="width: 100%" border @row-click="goDetail">
        <el-table-column prop="title" label="题目" />
        <el-table-column prop="category" label="分类" width="100" />
        <el-table-column prop="tags" label="标签" width="150">
          <template #default="scope">
            <el-tag v-for="tag in scope.row.tags" :key="tag" size="small" style="margin-right: 4px;">{{ tag }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="difficulty" label="难度" width="80">
          <template #default="scope">
            <el-tag :type="difficultyType(scope.row.difficulty)">
              {{ difficultyLabel(scope.row.difficulty) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间" width="160">
          <template #default="scope">
            {{ formatDate(scope.row.created_at) }}
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 16px; text-align: right;">
        <el-pagination
          background
          layout="prev, pager, next, jumper"
          :total="pagination.total"
          :page-size="pagination.pageSize"
          :current-page="pagination.page"
          @current-change="page => { pagination.page = page; fetchList(); }"
        />
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, watch } from 'vue';
import { useRouter } from 'vue-router';
import { fetchQuestions } from '../api/question';
import { fetchCategories } from '../api/index';

const router = useRouter();
const list = ref([]);
const pagination = reactive({ page: 1, pageSize: 10, total: 0 });
const search = reactive({ title: '', category: '', difficulty: '' });
const categoryOptions = ref([]);
const difficultyOptions = ref([]);
const categoryTags = ref({});

// 监听筛选条件变化（不包括标题搜索）
watch(
  () => [search.category, search.difficulty],
  () => {
    pagination.page = 1; // 重置到第一页
    fetchList();
  },
  { deep: true }
);

async function loadCategories() {
  try {
    const res = await fetchCategories();
    if (res.code === 200) {
      categoryOptions.value = res.data.categories;
      difficultyOptions.value = res.data.difficulties;
      categoryTags.value = res.data.category_tags;
    }
  } catch (error) {
    console.error('获取分类失败:', error);
    // 如果获取失败，使用默认分类
    categoryOptions.value = ['Go开发', '系统架构', '微服务', '云原生', '数据库', '前端技术', 'DevOps', '性能优化', '算法', '其他'];
    difficultyOptions.value = [
      { label: '简单', value: 'easy' },
      { label: '中等', value: 'medium' },
      { label: '困难', value: 'hard' }
    ];
  }
}

function fetchList() {
  fetchQuestions({
    page: pagination.page,
    page_size: pagination.pageSize,
    category: search.category,
    difficulty: search.difficulty,
    title: search.title,
  }).then(res => {
    list.value = res.data.questions;
    pagination.total = res.data.pagination.total;
  });
}

// 标题搜索函数（回车触发）
function handleTitleSearch() {
  pagination.page = 1; // 重置到第一页
  fetchList();
}

function goDetail(row) {
  router.push(`/questions/${row.id}`);
}

function formatDate(ts) {
  if (!ts) return '';
  const d = new Date(ts * 1000);
  return d.toLocaleString();
}

function difficultyLabel(val) {
  if (val === 'easy') return '简单';
  if (val === 'medium') return '中等';
  if (val === 'hard') return '困难';
  return val;
}
function difficultyType(val) {
  if (val === 'easy') return 'success';
  if (val === 'medium') return 'warning';
  if (val === 'hard') return 'danger';
  return '';
}

onMounted(async () => {
  await loadCategories();
  fetchList();
});
</script> 