<script setup>
import { ref, onMounted } from 'vue'
import { apiService } from '../services/api'

// 响应式数据
const inputMessage = ref('')
const messages = ref([])
const sessionId = ref('')
const isLoading = ref(false)
const error = ref(null)

// 初始化会话
const initSession = () => {
  sessionId.value = localStorage.getItem('session_id') || generateSessionId()
  localStorage.setItem('session_id', sessionId.value)
}

// 生成会话ID
const generateSessionId = () => {
  return 'session-' + Date.now().toString(36) + Math.random().toString(36).substr(2, 5)
}

// 处理消息提交
const handleSubmit = async () => {
  if (!inputMessage.value.trim() || isLoading.value) return
  
  try {
    isLoading.value = true
    error.value = null

    // 添加用户消息
    messages.value.push({
      role: 'user',
      content: inputMessage.value,
      timestamp: new Date().toISOString()
    })

    // 调用Go服务
    const response = await apiService.sendMessage(
      inputMessage.value,
      sessionId.value
    )

    // 更新消息列表
    messages.value.push({
      role: 'assistant',
      content: response.content,
      timestamp: new Date().toISOString()
    })

    inputMessage.value = ''
  } catch (err) {
    error.value = err.message
    // 错误时重置会话
    sessionId.value = ''
    localStorage.removeItem('session_id')
  } finally {
    isLoading.value = false
  }
}

// 组件挂载时初始化
onMounted(initSession)
</script>

<template>
<div class="bigfather">
  <div class="chat-container">
    <!-- 简洁头部 -->
    <header class="app-header">
      <h1 class="logo">慧编助手</h1>
      <div class="session-info">
        <span class="session-id">会话 #{{ sessionId.slice(-6) }}</span>
        <button @click="initSession" class="new-session-btn">
          新会话
        </button>
      </div>
    </header>

    <!-- 消息区域 -->
    <main class="message-area">
      <!-- 引导提示 -->
      <div v-if="messages.length === 0" class="welcome-guide">
        <div class="guide-icon">💡</div>
        <h3>欢迎使用编程思维助手</h3>
        <p>输入编程问题获取解决思路，例如：</p>
        <ul class="example-list">
          <li>"如何优化这段循环代码？"</li>
          <li>"这个错误提示是什么意思？"</li>
        </ul>
      </div>

      <!-- 消息列表 -->
      <div class="message-list">
        <div 
          v-for="(msg, index) in messages"
          :key="index"
          :class="['message-bubble', msg.role]"
        >
          <div class="bubble-header">
            <span class="role-icon">{{ msg.role === 'user' ? '❓' : '💡' }}</span>
            <span class="role-text">{{ msg.role === 'user' ? '你的问题' : '思路建议' }}</span>
          </div>
          <div class="bubble-content">{{ msg.content }}</div>
          <div class="bubble-time">{{ new Date(msg.timestamp).toLocaleTimeString() }}</div>
        </div>
      </div>

      <!-- 加载状态 -->
      <div v-if="isLoading" class="loading-indicator">
        <div class="loader"></div>
        <span>正在分析问题...</span>
      </div>
    </main>

    <!-- 输入区域 -->
    <div class="input-section">
      <textarea
        v-model="inputMessage"
        placeholder="输入编程问题 (尽量描述具体现象)"
        @keydown.enter.prevent="handleSubmit"
        :disabled="isLoading"
        class="input-text"
      ></textarea>
      <button 
        class="send-button"
        @click="handleSubmit"
        :disabled="isLoading"
      >
        {{ isLoading ? '思考中...' : '发送' }}
      </button>
    </div>

    <!-- 简洁错误提示 -->
    <div v-if="error" class="error-toast">
      <span>{{ error }}</span>
      <button @click="error = null" class="close-btn">&times;</button>
    </div>
  </div>
</div>  
</template>

<style scoped>
/* 现代科技风格优化 */
html, body, #app {
  height: 100%;
  margin: 0;
  padding: 0;
}


.bigfather {
  min-height: 100vh;
  width: 100vw;
  display: flex;
  align-items: center;
  justify-content: center;
  background: none;
}

.chat-container {
  width: 100%;
  max-width: 480px;           /* 限制最大宽度 */
  min-height: 80vh;
  margin: 40px auto;           /* 垂直居中并有上下间距 */
  border-radius: 18px;
  background: rgba(10,26,61,0.92);
  box-shadow: 0 8px 32px 0 rgba(31,38,135,0.37);
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  padding: 2rem 1.5rem 1.5rem 1.5rem; /* 内边距 */
}

/* 头部样式优化 */
.app-header {
  padding: 1rem;
  border-bottom: 1px solid rgba(45,109,246,0.3);
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.logo {
  font-size: 1.8rem;
  background: linear-gradient(135deg, #2d6df6 0%, #8ad4ff 100%);
  background-clip: text;           /* 标准写法 */
  -webkit-background-clip: text;   /* Webkit 浏览器 */
  -webkit-text-fill-color: transparent;
  text-shadow: 0 0 15px rgba(45,109,246,0.3);
  position: relative;
  display: inline-block;
  padding-bottom: 0.5rem;
}

.logo::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 0;
  width: 50px;
  height: 2px;
  background: #72f58f;
}

.session-info {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.new-session-btn {
  background: none;
  border: none;
  color: #8ad4ff;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 0.3rem;
  padding: 0.5rem;
  border-radius: 5px;
  transition: background 0.3s;
}

.new-session-btn:hover {
  background: rgba(45,109,246,0.1);
}

/* 消息气泡优化 */
.message-bubble {
  backdrop-filter: blur(5px);
  border-radius: 10px;
  margin: 0.8rem 0;
  padding: 1rem;
  box-shadow: 0 2px 6px rgba(0,0,0,0.1);
  transition: transform 0.3s;
  position: relative;
  overflow: hidden;
}

.message-bubble::before {
  content: '';
  position: absolute;
  left: 0;
  top: 0;
  bottom: 0;
  width: 4px;
  background: #2d6df6;
}

.message-bubble.assistant::before {
  background: #72f58f;
}

.message-bubble:hover {
  transform: translateY(-3px);
  box-shadow: 0 4px 15px rgba(45,109,246,0.2);
}

.bubble-header {
  padding-bottom: 0.5rem;
  margin-bottom: 0.8rem;
  border-bottom: 1px dashed rgba(255,255,255,0.1);
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.role-icon {
  width: 24px;
  height: 24px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  background: rgba(45,109,246,0.2);
}

/* 输入区域优化 */
.input-section {
  border-top: 1px solid rgba(45,109,246,0.2);
  display: flex;
  padding: 0.5rem 0 0 0;
  background: transparent;
}

.input-text {
  flex: 1;
  background: rgba(10,26,61,0.5);
  border: 1px solid rgba(45,109,246,0.3);
  border-radius: 5px;
  padding: 0.8rem;
  transition: all 0.3s;
  color: white;
  font-size: 1rem;
  resize: none;
}

.input-text:focus {
  outline: none;
  border-color: #72f58f;
  box-shadow: 0 0 0 3px rgba(114,245,143,0.2);
}

.send-button {
  background: linear-gradient(135deg, #2d6df6 0%, #72f58f 100%);
  font-weight: 500;
  border: none;
  border-radius: 5px;
  padding: 0.8rem 1.2rem;
  margin-left: 0.8rem;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  cursor: pointer;
  color: white;
  font-size: 1rem;
}

.send-button:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(45,109,246,0.3);
}

.send-button:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

/* 加载动画优化 */
.loading-indicator {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 1rem;
  color: #72f58f;
}

.loader {
  border: 2px solid rgba(114,245,143,0.2);
  border-top-color: #72f58f;
  animation: spin 0.8s linear infinite;
  width: 20px;
  height: 20px;
  border-radius: 50%;
  margin-right: 0.5rem;
}

/* 错误提示优化 */
.error-toast {
  background: linear-gradient(90deg, #ff4444 0%, #ff6b6b 100%);
  box-shadow: 0 4px 12px rgba(255,68,68,0.2);
  border: 1px solid rgba(255,255,255,0.1);
  border-radius: 5px;
  padding: 0.5rem;
  margin: 0.5rem;
  text-align: center;
  position: fixed;
  top: 1rem;
  right: 1rem;
  z-index: 100;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.close-btn {
  padding: 0 0.5rem;
  transition: opacity 0.3s;
  cursor: pointer;
  background: none;
  border: none;
  color: white;
}

.close-btn:hover {
  opacity: 0.8;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

/* 移动端优化 */
@media screen and (max-width: 768px) {
  .bigfather {
    border-radius: 0;
    min-height: 100vh;
  }

  .chat-container {
    max-width: 100vw;
    min-height: 100vh;
    margin: 0;
    border-radius: 0;
    padding: 1rem 0.2rem 1rem 0.2rem;
  }

  .app-header {
    padding: 0.8rem;
  }

  .logo {
    font-size: 1.5rem;
  }

  .message-bubble {
    max-width: 95%;
  }

  .input-section {
    padding: 0.8rem;
  }

  .input-text {
    height: 60px;
  }
}

/* 引导提示样式优化 */
.welcome-guide {
  text-align: left;
  color: #b3e5fc;
  margin-bottom: 2rem;
  background: rgba(45,109,246,0.08);
  border-radius: 8px;
  padding: 1rem 1.2rem;
}

.guide-icon {
  font-size: 1.5rem;
  margin-bottom: 0.5rem;
}

.example-list {
  margin: 0.5rem 0 0 1.2rem;
  color: #fff;
  font-size: 0.98rem;
}
</style>