// 1980年代复古音效系统

class RetroAudio {
    constructor() {
        this.audioContext = null;
        this.sounds = {};
        this.enabled = true;
        this.volume = 0.3;

        this.init();
    }

    init() {
        try {
            // 创建音频上下文
            this.audioContext = new (window.AudioContext || window.webkitAudioContext)();
            this.createSounds();
        } catch (error) {
            console.warn('音频上下文创建失败，音效将被禁用:', error);
            this.enabled = false;
        }
    }

    createSounds() {
        // 按钮点击音效 - 经典的"滴"声
        this.sounds.buttonClick = () => this.createBeep(800, 0.1, 'square');

        // 悬停音效 - 轻微的"嗡"声
        this.sounds.hover = () => this.createBeep(1200, 0.05, 'sine');

        // 成功音效 - 上升的音调
        this.sounds.success = () => this.createSuccessSound();

        // 错误音效 - 下降的音调
        this.sounds.error = () => this.createErrorSound();

        // 警告音效 - 重复的短音
        this.sounds.warning = () => this.createWarningSound();

        // 页面切换音效 - 扫描线效果
        this.sounds.pageTransition = () => this.createScanSound();

        // 加载音效 - 循环的滴答声
        this.sounds.loading = () => this.createLoadingSound();

        // 确认音效 - 双音调
        this.sounds.confirm = () => this.createConfirmSound();

        // 取消音效 - 单音调下降
        this.sounds.cancel = () => this.createCancelSound();
    }

    // 创建基础蜂鸣声
    createBeep(frequency, duration, type = 'square') {
        if (!this.enabled || !this.audioContext) return;

        const oscillator = this.audioContext.createOscillator();
        const gainNode = this.audioContext.createGain();

        oscillator.connect(gainNode);
        gainNode.connect(this.audioContext.destination);

        oscillator.frequency.setValueAtTime(frequency, this.audioContext.currentTime);
        oscillator.type = type;

        gainNode.gain.setValueAtTime(0, this.audioContext.currentTime);
        gainNode.gain.linearRampToValueAtTime(this.volume, this.audioContext.currentTime + 0.01);
        gainNode.gain.exponentialRampToValueAtTime(0.001, this.audioContext.currentTime + duration);

        oscillator.start(this.audioContext.currentTime);
        oscillator.stop(this.audioContext.currentTime + duration);
    }

    // 成功音效 - 上升的三音调
    createSuccessSound() {
        if (!this.enabled || !this.audioContext) return;

        const frequencies = [523, 659, 784]; // C5, E5, G5
        const duration = 0.15;

        frequencies.forEach((freq, index) => {
            setTimeout(() => {
                this.createBeep(freq, duration, 'sine');
            }, index * 100);
        });
    }

    // 错误音效 - 下降的三音调
    createErrorSound() {
        if (!this.enabled || !this.audioContext) return;

        const frequencies = [784, 659, 523]; // G5, E5, C5
        const duration = 0.2;

        frequencies.forEach((freq, index) => {
            setTimeout(() => {
                this.createBeep(freq, duration, 'sawtooth');
            }, index * 150);
        });
    }

    // 警告音效 - 重复的短音
    createWarningSound() {
        if (!this.enabled || !this.audioContext) return;

        for (let i = 0; i < 3; i++) {
            setTimeout(() => {
                this.createBeep(440, 0.1, 'triangle');
            }, i * 200);
        }
    }

    // 扫描线音效 - 模拟CRT扫描
    createScanSound() {
        if (!this.enabled || !this.audioContext) return;

        const oscillator = this.audioContext.createOscillator();
        const gainNode = this.audioContext.createGain();

        oscillator.connect(gainNode);
        gainNode.connect(this.audioContext.destination);

        oscillator.type = 'sawtooth';
        oscillator.frequency.setValueAtTime(50, this.audioContext.currentTime);
        oscillator.frequency.exponentialRampToValueAtTime(200, this.audioContext.currentTime + 0.3);

        gainNode.gain.setValueAtTime(0, this.audioContext.currentTime);
        gainNode.gain.linearRampToValueAtTime(this.volume * 0.5, this.audioContext.currentTime + 0.01);
        gainNode.gain.exponentialRampToValueAtTime(0.001, this.audioContext.currentTime + 0.3);

        oscillator.start(this.audioContext.currentTime);
        oscillator.stop(this.audioContext.currentTime + 0.3);
    }

    // 加载音效 - 循环的滴答声
    createLoadingSound() {
        if (!this.enabled || !this.audioContext) return;

        const interval = setInterval(() => {
            this.createBeep(600, 0.05, 'square');
        }, 500);

        // 5秒后停止
        setTimeout(() => {
            clearInterval(interval);
        }, 5000);
    }

    // 确认音效 - 双音调
    createConfirmSound() {
        if (!this.enabled || !this.audioContext) return;

        setTimeout(() => this.createBeep(523, 0.1, 'sine'), 0);
        setTimeout(() => this.createBeep(659, 0.1, 'sine'), 100);
    }

    // 取消音效 - 单音调下降
    createCancelSound() {
        if (!this.enabled || !this.audioContext) return;

        const oscillator = this.audioContext.createOscillator();
        const gainNode = this.audioContext.createGain();

        oscillator.connect(gainNode);
        gainNode.connect(this.audioContext.destination);

        oscillator.type = 'sawtooth';
        oscillator.frequency.setValueAtTime(400, this.audioContext.currentTime);
        oscillator.frequency.exponentialRampToValueAtTime(200, this.audioContext.currentTime + 0.2);

        gainNode.gain.setValueAtTime(0, this.audioContext.currentTime);
        gainNode.gain.linearRampToValueAtTime(this.volume, this.audioContext.currentTime + 0.01);
        gainNode.gain.exponentialRampToValueAtTime(0.001, this.audioContext.currentTime + 0.2);

        oscillator.start(this.audioContext.currentTime);
        oscillator.stop(this.audioContext.currentTime + 0.2);
    }

    // 播放音效的公共方法
    play(soundName) {
        if (!this.enabled) return;

        const sound = this.sounds[soundName];
        if (sound) {
            sound();
        }
    }

    // 设置音量
    setVolume(volume) {
        this.volume = Math.max(0, Math.min(1, volume));
    }

    // 启用/禁用音效
    setEnabled(enabled) {
        this.enabled = enabled;
    }

    // 获取音效状态
    isEnabled() {
        return this.enabled;
    }

    // 获取音量
    getVolume() {
        return this.volume;
    }
}

// 创建全局音效实例
const retroAudio = new RetroAudio();

// 导出音效实例和便捷方法
export default retroAudio;

// 便捷的播放方法
export const playSound = (soundName) => retroAudio.play(soundName);

// 音效类型常量
export const SOUNDS = {
    BUTTON_CLICK: 'buttonClick',
    HOVER: 'hover',
    SUCCESS: 'success',
    ERROR: 'error',
    WARNING: 'warning',
    PAGE_TRANSITION: 'pageTransition',
    LOADING: 'loading',
    CONFIRM: 'confirm',
    CANCEL: 'cancel'
};

// Vue插件形式
export const RetroAudioPlugin = {
    install(app) {
        app.config.globalProperties.$retroAudio = retroAudio;
        app.config.globalProperties.$playSound = playSound;
    }
}; 