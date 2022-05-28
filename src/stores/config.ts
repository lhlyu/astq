import { defineStore } from 'pinia'
import CryptoJS from 'crypto-js';

export interface Req {
    iss: string
    kid: string
    bid: string
    pk: string
    aud: string

    id: string
    api: string
    result: string
}

const useConfigStore = defineStore({
    id: 'astq',
    state: () =>
        ({
            iss: '',
            kid: '',
            bid: '',
            pk: '',
            aud: 'appstoreconnect-v1',
            id: '',
            api: 'subscriptions',
            result: ''
        } as Req),
    getters: {
        getReq: (state): string => {
            const val = {
                iss: state.iss,
                kid: state.kid,
                bid: state.bid,
                pk: state.pk,
                aud: state.aud,
                id: state.id,
            }
            const wordArray = CryptoJS.enc.Utf8.parse(JSON.stringify(val))
            return CryptoJS.enc.Base64.stringify(wordArray)
        }
    },
    actions: {
        reset() {
            this.iss = ''
            this.kid = ''
            this.bid = ''
            this.pk = ''
            this.aud = 'appstoreconnect-v1'
            this.id = ''
            this.api = 'subscriptions'
            this.result = ''
        }
    },
    // 启用持久化
    persist: {
        enabled: true,
        storage: window.localStorage
    }
})

export default useConfigStore
