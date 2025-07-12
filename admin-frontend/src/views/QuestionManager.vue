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
        <el-button type="primary" @click="openEdit()">新建题目</el-button>
      </div>
      <el-table :data="list" style="width: 100%" border>
        <el-table-column prop="id" label="ID" width="60" />
        <el-table-column prop="title" label="题目">
          <template #default="scope">
            <el-link type="primary" @click="openDetail(scope.row)" :underline="false">
              {{ scope.row.title }}
            </el-link>
          </template>
        </el-table-column>
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
        <el-table-column label="操作" width="180">
          <template #default="scope">
            <el-button size="small" @click="openEdit(scope.row)">编辑</el-button>
            <el-button size="small" type="danger" @click="handleDelete(scope.row.id)">删除</el-button>
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

    <!-- 详情弹窗 -->
    <el-dialog title="题目详情" v-model="detailVisible" width="900px">
      <div v-if="detailData">
        <el-descriptions :column="2" border>
          <el-descriptions-item label="题目">{{ detailData.title }}</el-descriptions-item>
          <el-descriptions-item label="分类">{{ detailData.category }}</el-descriptions-item>
          <el-descriptions-item label="难度">
            <el-tag :type="difficultyType(detailData.difficulty)">
              {{ difficultyLabel(detailData.difficulty) }}
            </el-tag>
          </el-descriptions-item>
          <el-descriptions-item label="创建时间">{{ formatDate(detailData.created_at) }}</el-descriptions-item>
          <el-descriptions-item label="标签" :span="2">
            <el-tag v-for="tag in detailData.tags" :key="tag" size="small" style="margin-right: 4px;">{{ tag }}</el-tag>
          </el-descriptions-item>
        </el-descriptions>
        <el-divider content-position="left">答案</el-divider>
        <div 
          class="markdown-content" 
          style="max-height: 400px; overflow-y: auto; padding: 16px; background: #fff; border: 1px solid #dcdfe6; border-radius: 4px;"
          v-html="renderedAnswer"
        ></div>
      </div>
      <template #footer>
        <el-button @click="detailVisible = false">关闭</el-button>
        <el-button type="primary" @click="editFromDetail">编辑</el-button>
      </template>
    </el-dialog>

    <!-- 编辑弹窗 -->
    <el-dialog :title="editForm.id ? '编辑题目' : '新建题目'" v-model="editVisible" width="800px">
      <el-form :model="editForm" label-width="80px">
        <el-form-item label="题目" required>
          <el-input v-model="editForm.title" placeholder="请输入题目内容" />
        </el-form-item>
        <el-form-item label="分类" required>
          <el-select v-model="editForm.category" placeholder="请选择分类" style="width: 100%;" @change="onCategoryChange">
            <el-option v-for="category in categoryOptions" :key="category" :label="category" :value="category" />
          </el-select>
        </el-form-item>
        <el-form-item label="标签">
          <el-select 
            v-model="editForm.tags" 
            multiple 
            filterable 
            allow-create 
            default-first-option 
            placeholder="选择或输入标签"
            style="width: 100%;"
          >
            <el-option v-for="tag in currentCategoryTags" :key="tag" :label="tag" :value="tag" />
          </el-select>
          <div style="margin-top: 8px; font-size: 12px; color: #909399;">
            当前分类：{{ editForm.category || '未选择' }} | 可用标签：{{ currentCategoryTags.length }}个
          </div>
        </el-form-item>
        <el-form-item label="难度" required>
          <el-select v-model="editForm.difficulty" placeholder="请选择难度" style="width: 100%;">
            <el-option v-for="item in difficultyOptions" :key="item.value" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="答案" required>
          <el-input 
            type="textarea" 
            v-model="editForm.answer" 
            :rows="10" 
            placeholder="支持 Markdown 格式，请输入详细答案" 
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="editVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSave">保存</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, computed, watch } from 'vue';
import { fetchQuestions, createQuestion, updateQuestion, deleteQuestion } from '../api/question';
import { fetchCategories } from '../api/index';
import { ElMessage, ElMessageBox } from 'element-plus';
import { marked } from 'marked';

const list = ref([]);
const pagination = reactive({ page: 1, pageSize: 10, total: 0 });
const search = reactive({ title: '', category: '', difficulty: '' });
const categoryOptions = ref([]);
const difficultyOptions = ref([]);
const categoryTags = ref({});

const editVisible = ref(false);
const editForm = reactive({ id: null, title: '', category: '', tags: [], difficulty: 'easy', answer: '' });

const detailVisible = ref(false);
const detailData = ref(null);

// 当前分类的标签列表
const currentCategoryTags = computed(() => {
  if (!editForm.category || !categoryTags.value[editForm.category]) {
    return [];
  }
  return categoryTags.value[editForm.category];
});

// 渲染后的 Markdown 答案
const renderedAnswer = computed(() => {
  if (!detailData.value || !detailData.value.answer) {
    return '';
  }
  return marked(detailData.value.answer);
});

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

function openDetail(row) {
  detailData.value = row;
  detailVisible.value = true;
}

function editFromDetail() {
  detailVisible.value = false;
  openEdit(detailData.value);
}

function openEdit(row) {
  if (row) {
    Object.assign(editForm, row);
  } else {
    Object.assign(editForm, { id: null, title: '', category: '', tags: [], difficulty: 'easy', answer: '' });
  }
  editVisible.value = true;
}

function onCategoryChange() {
  // 当分类改变时，清空已选择的标签，避免标签与分类不匹配
  editForm.tags = [];
}

function handleSave() {
  const data = { ...editForm };
  if (!data.title) {
    ElMessage.error('题目不能为空');
    return;
  }
  if (!data.category) {
    ElMessage.error('请选择分类');
    return;
  }
  if (!data.difficulty) {
    ElMessage.error('请选择难度');
    return;
  }
  if (!data.answer) {
    ElMessage.error('答案不能为空');
    return;
  }
  if (data.id) {
    updateQuestion(data.id, data).then(() => {
      ElMessage.success('更新成功');
      editVisible.value = false;
      fetchList();
    });
  } else {
    createQuestion(data).then(() => {
      ElMessage.success('创建成功');
      editVisible.value = false;
      fetchList();
    });
  }
}

function handleDelete(id) {
  ElMessageBox.confirm('确定要删除该题目吗？', '提示', { type: 'warning' })
    .then(() => {
      deleteQuestion(id).then(() => {
        ElMessage.success('删除成功');
        fetchList();
      });
    })
    .catch(() => {});
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

<style scoped>
.markdown-content {
  line-height: 1.6;
}

.markdown-content h1,
.markdown-content h2,
.markdown-content h3,
.markdown-content h4,
.markdown-content h5,
.markdown-content h6 {
  margin-top: 1.5em;
  margin-bottom: 0.5em;
  font-weight: bold;
}

.markdown-content h1 {
  font-size: 1.8em;
  border-bottom: 2px solid #eee;
  padding-bottom: 0.3em;
}

.markdown-content h2 {
  font-size: 1.5em;
  border-bottom: 1px solid #eee;
  padding-bottom: 0.3em;
}

.markdown-content h3 {
  font-size: 1.3em;
}

.markdown-content p {
  margin-bottom: 1em;
}

.markdown-content ul,
.markdown-content ol {
  margin-bottom: 1em;
  padding-left: 2em;
}

.markdown-content li {
  margin-bottom: 0.5em;
}

.markdown-content blockquote {
  border-left: 4px solid #ddd;
  padding-left: 1em;
  margin: 1em 0;
  color: #666;
  background-color: #f9f9f9;
}

.markdown-content code {
  background-color: #f4f4f4;
  padding: 2px 4px;
  border-radius: 3px;
  font-family: Monaco, 'Courier New', monospace;
}

.markdown-content pre {
  background-color: #f4f4f4;
  padding: 1em;
  border-radius: 4px;
  overflow-x: auto;
  margin: 1em 0;
}

.markdown-content pre code {
  background-color: transparent;
  padding: 0;
}

.markdown-content table {
  border-collapse: collapse;
  width: 100%;
  margin: 1em 0;
}

.markdown-content th,
.markdown-content td {
  border: 1px solid #ddd;
  padding: 8px;
  text-align: left;
}

.markdown-content th {
  background-color: #f2f2f2;
  font-weight: bold;
}
</style> 