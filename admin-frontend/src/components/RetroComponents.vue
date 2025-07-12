<template>
  <!-- 这个文件用于导出复古组件，实际组件在下面定义 -->
</template>

<script>
// 复古按钮组件
export const RetroButton = {
  name: 'RetroButton',
  props: {
    type: {
      type: String,
      default: 'primary',
      validator: value => ['primary', 'success', 'warning', 'danger'].includes(value)
    },
    size: {
      type: String,
      default: 'medium',
      validator: value => ['small', 'medium', 'large'].includes(value)
    },
    disabled: {
      type: Boolean,
      default: false
    }
  },
  template: `
    <button 
      :class="[
        'retro-button',
        'retro-button--' + type,
        'retro-button--' + size,
        { 'retro-button--disabled': disabled }
      ]"
      :disabled="disabled"
      @click="$emit('click', $event)"
    >
      <slot></slot>
    </button>
  `
}

// 复古输入框组件
export const RetroInput = {
  name: 'RetroInput',
  props: {
    value: {
      type: [String, Number],
      default: ''
    },
    placeholder: {
      type: String,
      default: ''
    },
    type: {
      type: String,
      default: 'text'
    },
    disabled: {
      type: Boolean,
      default: false
    }
  },
  template: `
    <input
      :class="[
        'retro-input',
        { 'retro-input--disabled': disabled }
      ]"
      :value="value"
      :placeholder="placeholder"
      :type="type"
      :disabled="disabled"
      @input="$emit('input', $event.target.value)"
      @focus="$emit('focus', $event)"
      @blur="$emit('blur', $event)"
    />
  `
}

// 复古卡片组件
export const RetroCard = {
  name: 'RetroCard',
  props: {
    title: {
      type: String,
      default: ''
    },
    hoverable: {
      type: Boolean,
      default: false
    }
  },
  template: `
    <div 
      :class="[
        'retro-card',
        { 'retro-card--hoverable': hoverable }
      ]"
    >
      <div v-if="title" class="retro-card__title">{{ title }}</div>
      <div class="retro-card__content">
        <slot></slot>
      </div>
    </div>
  `
}

// 复古表格组件
export const RetroTable = {
  name: 'RetroTable',
  props: {
    data: {
      type: Array,
      default: () => []
    },
    columns: {
      type: Array,
      default: () => []
    }
  },
  template: `
    <table class="retro-table">
      <thead>
        <tr>
          <th v-for="column in columns" :key="column.key">
            {{ column.label }}
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(row, index) in data" :key="index">
          <td v-for="column in columns" :key="column.key">
            {{ row[column.key] }}
          </td>
        </tr>
      </tbody>
    </table>
  `
}

// 复古加载组件
export const RetroLoading = {
  name: 'RetroLoading',
  props: {
    text: {
      type: String,
      default: '加载中...'
    }
  },
  template: `
    <div class="retro-loading-container">
      <div class="retro-loading"></div>
      <div class="retro-loading-text">{{ text }}</div>
    </div>
  `
}

// 复古通知组件
export const RetroNotification = {
  name: 'RetroNotification',
  props: {
    message: {
      type: String,
      required: true
    },
    type: {
      type: String,
      default: 'info',
      validator: value => ['info', 'success', 'warning', 'error'].includes(value)
    },
    duration: {
      type: Number,
      default: 3000
    },
    show: {
      type: Boolean,
      default: false
    }
  },
  template: `
    <transition name="retro-notification">
      <div 
        v-if="show"
        :class="[
          'retro-notification',
          'retro-notification--' + type
        ]"
      >
        <div class="retro-notification__icon">
          <span v-if="type === 'success'">✅</span>
          <span v-else-if="type === 'warning'">⚠️</span>
          <span v-else-if="type === 'error'">❌</span>
          <span v-else>ℹ️</span>
        </div>
        <div class="retro-notification__message">{{ message }}</div>
      </div>
    </transition>
  `,
  mounted() {
    if (this.duration > 0) {
      setTimeout(() => {
        this.$emit('close')
      }, this.duration)
    }
  }
}

// 复古模态框组件
export const RetroModal = {
  name: 'RetroModal',
  props: {
    visible: {
      type: Boolean,
      default: false
    },
    title: {
      type: String,
      default: ''
    },
    width: {
      type: String,
      default: '500px'
    }
  },
  template: `
    <transition name="retro-modal">
      <div v-if="visible" class="retro-modal" @click.self="$emit('close')">
        <div class="retro-modal-content" :style="{ maxWidth: width }">
          <div v-if="title" class="retro-modal__title">{{ title }}</div>
          <div class="retro-modal__body">
            <slot></slot>
          </div>
          <div class="retro-modal__footer">
            <slot name="footer">
              <retro-button @click="$emit('close')">关闭</retro-button>
            </slot>
          </div>
        </div>
      </div>
    </transition>
  `
}

// 复古进度条组件
export const RetroProgress = {
  name: 'RetroProgress',
  props: {
    percentage: {
      type: Number,
      default: 0,
      validator: value => value >= 0 && value <= 100
    },
    textInside: {
      type: Boolean,
      default: false
    }
  },
  template: `
    <div class="retro-progress">
      <div 
        class="retro-progress-bar"
        :style="{ width: percentage + '%' }"
      >
        <span v-if="textInside" class="retro-progress__text">{{ percentage }}%</span>
      </div>
      <span v-if="!textInside" class="retro-progress__text">{{ percentage }}%</span>
    </div>
  `
}

export default {
  name: 'RetroComponents',
  components: {
    RetroButton,
    RetroInput,
    RetroCard,
    RetroTable,
    RetroLoading,
    RetroNotification,
    RetroModal,
    RetroProgress
  }
}
</script>

<style scoped>
/* 复古按钮样式变体 */
.retro-button--primary {
  border-color: var(--crt-green);
  color: var(--crt-green);
}

.retro-button--success {
  border-color: var(--crt-green);
  color: var(--crt-green);
}

.retro-button--warning {
  border-color: var(--crt-yellow);
  color: var(--crt-yellow);
}

.retro-button--danger {
  border-color: var(--crt-red);
  color: var(--crt-red);
}

.retro-button--small {
  padding: 8px 16px;
  font-size: 12px;
}

.retro-button--large {
  padding: 16px 32px;
  font-size: 16px;
}

.retro-button--disabled {
  opacity: 0.5;
  cursor: not-allowed;
  pointer-events: none;
}

/* 复古输入框样式变体 */
.retro-input--disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

/* 复古卡片样式变体 */
.retro-card--hoverable:hover {
  transform: translateY(-5px);
  box-shadow: 
    0 8px 16px rgba(0, 0, 0, 0.5),
    0 0 30px rgba(0, 255, 0, 0.4);
}

.retro-card__title {
  color: var(--crt-yellow);
  font-weight: bold;
  font-size: 18px;
  margin-bottom: 15px;
  text-shadow: 0 0 5px var(--crt-yellow);
}

.retro-card__content {
  color: var(--crt-green);
}

/* 复古加载容器 */
.retro-loading-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 15px;
  padding: 30px;
}

.retro-loading-text {
  color: var(--crt-green);
  font-family: 'Courier New', monospace;
  font-weight: bold;
  text-shadow: 0 0 5px var(--crt-green);
  animation: text-shadow 2s ease-in-out infinite;
}

/* 复古通知样式变体 */
.retro-notification--success {
  border-color: var(--crt-green);
  color: var(--crt-green);
}

.retro-notification--warning {
  border-color: var(--crt-yellow);
  color: var(--crt-yellow);
}

.retro-notification--error {
  border-color: var(--crt-red);
  color: var(--crt-red);
}

.retro-notification__icon {
  margin-right: 10px;
  font-size: 16px;
}

.retro-notification__message {
  font-weight: bold;
}

/* 复古模态框样式 */
.retro-modal__title {
  color: var(--crt-yellow);
  font-weight: bold;
  font-size: 20px;
  margin-bottom: 20px;
  text-align: center;
  text-shadow: 0 0 10px var(--crt-yellow);
}

.retro-modal__body {
  color: var(--crt-green);
  margin-bottom: 20px;
}

.retro-modal__footer {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}

/* 复古进度条文本 */
.retro-progress__text {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  color: var(--crt-black);
  font-weight: bold;
  font-size: 12px;
}

/* 过渡动画 */
.retro-notification-enter-active,
.retro-notification-leave-active {
  transition: all 0.3s ease;
}

.retro-notification-enter-from,
.retro-notification-leave-to {
  opacity: 0;
  transform: translateX(100%);
}

.retro-modal-enter-active,
.retro-modal-leave-active {
  transition: all 0.3s ease;
}

.retro-modal-enter-from,
.retro-modal-leave-to {
  opacity: 0;
  transform: scale(0.8);
}
</style> 