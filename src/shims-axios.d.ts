import { AxiosRequestConfig } from 'axios'

export namespace ReqRes {
    export interface ResponseResult {
        code: number
        msg: string
        data: string
    }
}

declare module 'axios' {
    export interface AxiosInstance {
        request<R = ReqRes.ResponseResult>(config: AxiosRequestConfig): Promise<R>
        get<R = ReqRes.ResponseResult>(url: string, config?: AxiosRequestConfig): Promise<R>
        delete<R = ReqRes.ResponseResult>(url: string, config?: AxiosRequestConfig): Promise<R>
        head<R = ReqRes.ResponseResult>(url: string, config?: AxiosRequestConfig): Promise<R>
        options<R = ReqRes.ResponseResult>(url: string, config?: AxiosRequestConfig): Promise<R>
        post<R = ReqRes.ResponseResult>(url: string, data?: any, config?: AxiosRequestConfig): Promise<R>
        put<R = ReqRes.ResponseResult>(url: string, data?: any, config?: AxiosRequestConfig): Promise<R>
        patch<R = ReqRes.ResponseResult>(url: string, data?: any, config?: AxiosRequestConfig): Promise<R>
    }
}
