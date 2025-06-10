import axios from 'axios'

const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL,
  timeout: 15000,
  headers: {
    'Content-Type': 'application/json'
  }
})

// 统一响应处理
const handleResponse = (response) => {
  if (response.status !== 200) {
    throw new Error(`请求失败: ${response.statusText}`)
  }
  return response.data
}

export const apiService = {
  // AI对话接口
  sendMessage: async (message, sessionId) => {
    try {
      const response = await api.post('/ai/run', {
        message,
        session_id: sessionId
      })
      return handleResponse(response)
    } catch (error) {
      throw new Error(error.response?.data?.message || 'AI服务暂不可用')
    }
  },

  // 用户注册接口
  register: async (userData) => {
    try {
      const response = await api.post('/user/register', userData)
      return handleResponse(response)
    } catch (error) {
      throw new Error(error.response?.data?.message || '注册失败')
    }
  }
}