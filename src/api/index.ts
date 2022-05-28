import instance from './request'
import useConfigStore from '../stores/config'

// 查询所有的订阅情况
export const GetSubscriptions = async () => {
    const store = useConfigStore()
    return instance.post('/api/subscriptions', store.getReq)
}

// 查询订单信息
export const GetLookupOrder = async () => {
    const store = useConfigStore()
    return instance.post('/api/lookup', store.getReq)
}

// 查询退款信息
export const GetRefund = async () => {
    const store = useConfigStore()
    return instance.post('/api/refund', store.getReq)
}

// 查询历史交易信息
export const GetHistory = async () => {
    const store = useConfigStore()
    return instance.post('/api/history', store.getReq)
}
