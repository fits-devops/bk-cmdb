export default {
    props: {
        properties: {
            type: Array,
            required: true
        },
        propertyGroups: {
            type: Array,
            required: true
        },
        objectUnique: {
            type: Array,
            default () {
                return []
            }
        }
    },
    computed: {
        $sortedGroups () {
            const publicGroups = []
            const metadataGroups = []
            this.propertyGroups.forEach(group => {
                if (this.$tools.getMetadataBiz(group)) {
                    metadataGroups.push(group)
                } else {
                    publicGroups.push(group)
                }
            })
            const sortKey = 'group_index'
            publicGroups.sort((groupA, groupB) => groupA[sortKey] - groupB[sortKey])
            metadataGroups.sort((groupA, groupB) => groupA[sortKey] - groupB[sortKey])
            const allGroups = [
                ...publicGroups,
                ...metadataGroups,
                {
                    'group_id': 'none',
                    'group_name': this.$t('Common["更多属性"]')
                }
            ]
            allGroups.forEach((group, index) => {
                group['group_index'] = index
            })
            return allGroups
        },
        $sortedProperties () {
            const unique = this.objectUnique.find(unique => unique.must_check) || {}
            const uniqueKeys = unique.keys || []
            const sortKey = 'property_index'
            const properties = this.properties.filter(property => {
                return !property['isapi']
                    && !uniqueKeys.some(key => key.key_id === property.id)
            })
            return properties.sort((propertyA, propertyB) => propertyA[sortKey] - propertyB[sortKey])
        },
        $groupedProperties () {
            return this.$sortedGroups.map(group => {
                return this.$sortedProperties.filter(property => {
                    const inGroup = property['property_group'] === group['group_id']
                    const isAsst = ['singleasst', 'multiasst'].includes(property['property_type'])
                    return inGroup && !isAsst
                })
            })
        }
    }
}
