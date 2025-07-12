<template>
  <div>
    <el-card v-if="blog">
      <h2>{{ blog.title }}</h2>
      <div style="color: #888; margin-bottom: 8px;">
        分类：{{ blog.category }}
        <el-tag v-for="tag in blog.tags" :key="tag" size="small" style="margin-left: 8px;">{{ tag }}</el-tag>
        <span style="float: right;">{{ formatDate(blog.created_at) }}</span>
      </div>
      <el-divider />
      <div v-html="renderedContent" class="markdown-body" />
    </el-card>
    <el-empty v-else description="未找到博客" />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { fetchBlog } from '../api/blog';
import MarkdownIt from 'markdown-it';

const route = useRoute();
const blog = ref(null);
const renderedContent = ref('');
const md = new MarkdownIt();

function formatDate(ts) {
  if (!ts) return '';
  const d = new Date(ts * 1000);
  return d.toLocaleString();
}

onMounted(() => {
  fetchBlog(route.params.id).then(res => {
    blog.value = res.data;
    renderedContent.value = md.render(res.data.content_md || '');
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