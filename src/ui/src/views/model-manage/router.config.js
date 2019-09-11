import Meta from '@/router/meta'
import { getMetadataBiz } from '@/utils/tools'
import { NAV_MODEL_MANAGEMENT } from '@/dictionary/menu'
import {
    C_MODEL_GROUP,
    U_MODEL_GROUP,
    D_MODEL_GROUP,
    C_MODEL,
    U_MODEL,
    D_MODEL,
    GET_AUTH_META
} from '@/dictionary/auth'

export const OPERATION = {
    C_MODEL_GROUP,
    U_MODEL_GROUP,
    D_MODEL_GROUP,
    C_MODEL,
    U_MODEL,
    D_MODEL
}

const modelPath = '/model'

export default [{
    name: 'model',
    path: modelPath,
    component: () => import('./index.vue'),
    meta: new Meta({
        menu: {
            id: 'model',
            i18n: 'Nav["模型"]',
            path: modelPath,
            order: 1,
            parent: NAV_MODEL_MANAGEMENT
        },
        auth: {
            operation: Object.values(OPERATION),
            setAuthScope (to, from, app) {
                const isAdminView = app.$store.getters.isAdminView
                this.authScope = isAdminView ? 'global' : 'business'
            }
        }
    })
}, {
    name: 'modelDetails',
    path: '/model/details/:modelId',
    component: () => import('./children/index.vue'),
    meta: new Meta({
        auth: {
            operation: [
                OPERATION.U_MODEL,
                OPERATION.D_MODEL
            ],
            setAuthScope (to, from, app) {
                const modelId = to.params.modelId
                const model = app.$store.getters['objectModelClassify/getModelById'](modelId)
                const bizId = getMetadataBiz(model)
                this.authScope = bizId ? 'business' : 'global'
            },
            setDynamicMeta: (to, from, app) => {
                const modelId = to.params.modelId
                const model = app.$store.getters['objectModelClassify/getModelById'](modelId)
                const bizId = getMetadataBiz(model)
                const resourceMeta = [{
                    ...GET_AUTH_META(OPERATION.U_MODEL),
                    resource_id: model.id
                }, {
                    ...GET_AUTH_META(OPERATION.D_MODEL),
                    resource_id: model.id
                }]
                if (bizId) {
                    resourceMeta.forEach(meta => {
                        meta.biz_id = parseInt(bizId)
                    })
                }
                app.$store.commit('auth/setResourceMeta', resourceMeta)
            }
        },
        checkAvailable: (to, from, app) => {
            const modelId = to.params.modelId
            const model = app.$store.getters['objectModelClassify/getModelById'](modelId)
            return !!model
        }
    })
}]
