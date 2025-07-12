import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import router from './router.js'
import { RetroAudioPlugin } from './utils/RetroAudio.js'

const app = createApp(App)

// 注册复古音效插件
app.use(RetroAudioPlugin)

// 注册Element Plus
app.use(ElementPlus)

// 注册路由
app.use(router)

// 全局混入，为所有组件添加音效功能
app.mixin({
    mounted() {
        // 为所有按钮添加音效
        this.$nextTick(() => {
            const buttons = this.$el.querySelectorAll('button, .retro-button, .retro-menu-item, .retro-nav-item');
            buttons.forEach(button => {
                button.addEventListener('click', () => {
                    this.$playSound('buttonClick');
                });

                button.addEventListener('mouseenter', () => {
                    this.$playSound('hover');
                });
            });
        });
    },

    methods: {
        // 播放成功音效
        playSuccessSound() {
            this.$playSound('success');
        },

        // 播放错误音效
        playErrorSound() {
            this.$playSound('error');
        },

        // 播放警告音效
        playWarningSound() {
            this.$playSound('warning');
        },

        // 播放页面切换音效
        playPageTransitionSound() {
            this.$playSound('pageTransition');
        },

        // 播放确认音效
        playConfirmSound() {
            this.$playSound('confirm');
        },

        // 播放取消音效
        playCancelSound() {
            this.$playSound('cancel');
        }
    }
})

// 全局错误处理
app.config.errorHandler = (err, vm, info) => {
    console.error('Vue错误:', err);
    console.error('错误信息:', info);
    // 播放错误音效
    if (vm && vm.$playSound) {
        vm.$playSound('error');
    }
}

// 路由切换时播放音效
router.beforeEach((to, from, next) => {
    // 播放页面切换音效
    if (window.$playSound) {
        window.$playSound('pageTransition');
    }
    next();
});

app.mount('#app')
