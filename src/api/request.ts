import axios from 'axios'

const instance = axios.create({
    baseURL: 'http://astq-rosy.vercel.app/',
    timeout: 60000
})

// 拦截请求
instance.interceptors.request.use(config => {
    return config
})

const handlerResponce = () => {
    return {
        code: 3,
        msg: '请求异常',
        data: null
    }
}

// 拦截响应
instance.interceptors.response.use(response => {
    if (response.status !== 200 || !response.data) {
        console.error('请求异常: ' + response.status)
        return handlerResponce()
    }
    const result = response.data
    return result
})

export default instance
