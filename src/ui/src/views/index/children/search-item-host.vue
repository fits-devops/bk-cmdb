<template>
    <div class="search-item-host clearfix" @click="handleHostClick" :title="getHostTitle(host)">
        <div class="host-ip fl">{{host['host']['host_innerip']}}</div>
        <div class="host-biz fr">{{host['biz'][0]['biz_name']}}</div>
    </div>
</template>

<script>
    import { mapGetters } from 'vuex'
    export default {
        props: {
            host: {
                type: Object,
                required: true
            }
        },
        computed: {
            ...mapGetters('objectBiz', ['authorizedBusiness'])
        },
        methods: {
            handleHostClick () {
                const bizId = this.host['biz'][0]['biz_id']
                this.$router.push({
                    name: 'resource',
                    query: {
                        business: bizId,
                        ip: this.host['host']['host_innerip'],
                        outer: false,
                        inner: true,
                        exact: 1,
                        assigned: true
                    }
                })
            },
            checkoutBizAuth (bizId) {
                return this.authorizedBusiness.some(biz => biz['biz_id'] === bizId)
            },
            getHostTitle (host) {
                return `${host['host']['host_innerip']}—${host['biz'][0]['biz_name']}`
            }
        }
    }
</script>

<style lang="scss" scoped>
.search-item-host{
    padding: 0 4px;
}
.search-item-host:hover{
    .host-ip{
        color: #3c96ff;
    }
}
.host-ip {
    width: 60%;
    @include ellipsis;
}
.host-biz{
    width: 35%;
    text-align: right;
    color: #c3cdd7;
    @include ellipsis;
}
</style>
