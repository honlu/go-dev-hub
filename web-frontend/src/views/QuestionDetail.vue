<template>
  <div>
    <el-card v-if="question">
      <h2>{{ question.title }}</h2>
      <div style="color: #888; margin-bottom: 8px;">
        分类：{{ question.category }}
        <el-tag v-for="tag in question.tags" :key="tag" size="small" style="margin-left: 8px;">{{ tag }}</el-tag>
        <el-tag :type="difficultyType(question.difficulty)" style="margin-left: 8px;">{{ difficultyLabel(question.difficulty) }}</el-tag>
        <span style="float: right;">{{ formatDate(question.created_at) }}</span>
      </div>
      <el-divider />
      <div v-html="renderedContent" class="markdown-body" />
    </el-card>
    <el-empty v-else description="未找到题目" />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { fetchQuestion } from '../api/question';
import MarkdownIt from 'markdown-it';

const route = useRoute();
const question = ref(null);
const renderedContent = ref('');
const md = new MarkdownIt();

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

onMounted(() => {
  fetchQuestion(route.params.id).then(res => {
    question.value = res.data;
    renderedContent.value = md.render(res.data.answer || '');
  });
});
</script>

<style>
.markdown-body {
  font-size: 16px;
  line-height: 1.8;
  word-break: break-all;
}
</style> 