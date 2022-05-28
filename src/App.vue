<template>
    <main>
        <el-card class="box-card">
            <template #header>
                <div class="card-header">
                    <span>{{ apiMap[store.api] }}</span>
                </div>
            </template>
            <el-row :gutter="30">
                <el-col :xs="24" :sm="24" :md="12" :lg="12" :xl="12">
                    <el-form label-position="top" label-width="50px" style="width: 100%">
                        <el-form-item label="ISS - 发行人">
                            <el-input v-model="store.iss" placeholder='您在 App Store Connect 中的密钥页面中的发行者 ID（例如：" 57246542-96fe-1a63-e053-0824d011072a"）' />
                        </el-form-item>
                        <el-form-item label="KID - 秘钥">
                            <el-input v-model="store.kid" placeholder="您在 App Store Connect 中的私钥 ID（例如2X9R4HXF34：）" />
                        </el-form-item>
                        <el-form-item label="BID - 应用的BundleID">
                            <el-input v-model="store.bid" placeholder="应用的BundleID（例如：“com.example.testbundleid2021”)" />
                        </el-form-item>
                        <el-form-item label="PK - 签名的秘钥">
                            <el-input v-model="store.pk" :rows="5" type="textarea" :placeholder="pk" />
                        </el-form-item>
                        <el-form-item label="AUD">
                            <el-input v-model="store.aud" placeholder="appstoreconnect-v1" />
                        </el-form-item>
                        <el-form-item>
                            <el-button plain @click="store.reset()">重置默认值</el-button>
                        </el-form-item>
                    </el-form>
                </el-col>
                <el-col :xs="24" :sm="24" :md="12" :lg="12" :xl="12">
                    <el-form label-position="top" label-width="50px" style="width: 100%; overflow: auto">
                        <el-form-item label="查询接口">
                            <el-select placeholder="查询接口" v-model="store.api" @change="changeApi">
                                <el-option v-for="item in apis" :key="item.value" :label="item.label" :value="item.value" />
                            </el-select>
                        </el-form-item>
                        <el-form-item label="ID">
                            <el-input v-model="store.id" :placeholder="apiIdMap[store.api]" />
                        </el-form-item>
                        <el-form-item>
                            <el-button plain @click="query">查询</el-button>
                        </el-form-item>
                        <el-form-item label="结果">
                            <pre
                                class="line-numbers"
                                lang="zh-Hans-CN"
                                data-prismjs-copy="复制文本"
                                data-prismjs-copy-error="按Ctrl+C复制"
                                data-prismjs-copy-success="文本已复制！"
                            ><code class="language-json" v-html="store.result"></code></pre>
                        </el-form-item>
                    </el-form>
                </el-col>
            </el-row>
        </el-card>
    </main>
</template>

<script lang="ts" setup>
import { ElMessage } from 'element-plus'
import Prism from 'prismjs'
import 'prismjs/components/prism-json'
import 'prismjs/plugins/line-numbers/prism-line-numbers'
import 'prismjs/plugins/toolbar/prism-toolbar'
import 'prismjs/plugins/copy-to-clipboard/prism-copy-to-clipboard'
import useConfigStore from './stores/config'
import { GetSubscriptions, GetLookupOrder, GetHistory, GetRefund } from './api'

const store = useConfigStore()

const pk = `-----BEGIN PRIVATE KEY-----
    MIGTAg23kjjh2h3uhuhfduhJHAKJ23JASjhaskjj234hjHKJ
    hJHAKJ23JASjhaskjj234hjHKJHS31hkjjMIGTAg23kjjh2h
    HS31hkjjwdwqrq24h23jkh
-----END PRIVATE KEY-----`

const apiMap = {
    subscriptions: '获取所有的订阅状态',
    lookup: '查找订单',
    history: '获取历史交易记录',
    refund: '获取退款历史'
}

const apiIdMap = {
    subscriptions: 'transactionId',
    lookup: 'orderId',
    history: 'transactionId',
    refund: 'transactionId'
}

const apis = [
    {
        label: apiMap.subscriptions,
        value: 'subscriptions'
    },
    {
        label: apiMap.lookup,
        value: 'lookup'
    },
    {
        label: apiMap.history,
        value: 'history'
    },
    {
        label: apiMap.refund,
        value: 'refund'
    }
]

const changeApi = (v: string) => {
    store.result = ''
    store.id = ''
}

const errtip = (message: string) => {
    ElMessage({
        showClose: true,
        message: message,
        type: 'error',
        center: true
    })
}

const query = async () => {
    if (!store.iss.trim().length) {
        errtip('ISS不能为空')
        return
    }
    if (!store.kid.trim().length) {
        errtip('KID不能为空')
        return
    }
    if (!store.bid.trim().length) {
        errtip('BID不能为空')
        return
    }
    if (!store.pk.trim().length) {
        errtip('PK不能为空')
        return
    }
    if (!store.aud.trim().length) {
        errtip('AUD不能为空')
        return
    }
    if (!store.id.trim().length) {
        errtip('ID不能为空')
        return
    }
    let resp
    switch (store.api) {
        case 'subscriptions':
            resp = await GetSubscriptions()
            break
        case 'lookup':
            resp = await GetLookupOrder()
            break
        case 'history':
            resp = await GetHistory()
            break
        case 'refund':
            resp = await GetRefund()
            break
        default:
            errtip('未知请求')
            return
    }
    if (resp.code) {
        errtip(resp.msg)
        store.result = ''
        return
    }
    store.result = Prism.highlight(JSON.stringify(JSON.parse(resp.data), null, '    '), Prism.languages.json, 'json')
}
</script>

<style lang="scss">
@import "element-plus/dist/index.css";
@import 'prismjs/plugins/toolbar/prism-toolbar.min.css';
@import 'prismjs/plugins/line-numbers/prism-line-numbers.min.css';
@import 'prismjs/themes/prism.css';
#app {
    display: flex;
    justify-content: center;
    align-items: center;
    min-height: 100vh;
    main {
        max-width: 1200px;
        width: 100%;
    }
}
</style>
