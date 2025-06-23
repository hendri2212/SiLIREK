<template>
    <tr>
        <td class="text-center">
            <span v-if="hasChildren" @click="isExpanded = !isExpanded" style="cursor: pointer;">
                {{ isExpanded ? '▼' : '▶' }}
            </span>
        </td>
        <td>{{ props.data.User?.full_name || '-' }}</td>
        <td class="text-center">{{ props.data.Quarterly?.quarter }} - {{ props.data.Quarterly?.year }}</td>
        <td :style="{ paddingLeft: `${props.level * 20}px` }">
            {{ props.data.Activity?.name || '-' }}
        </td>
        <td class="text-end">
            {{ new Intl.NumberFormat('id-ID', {
                style: 'currency',
                currency: 'IDR',
                minimumFractionDigits: 0
            }).format(props.data.budget) }}
        </td>
        <td class="text-center">
            <button @click="$emit('edit-budget', props.data)" class="btn btn-sm btn-primary">
                Edit
            </button>
        </td>
    </tr>
    <template v-if="isExpanded">
        <BudgetRow 
            v-for="child in props.data.children" 
            :key="child.id" 
            :data="child" 
            :level="props.level + 1" 
            @edit-budget="$emit('edit-budget', $event)"
        />
    </template>
</template>

<script setup>
import { ref, computed } from 'vue'
import BudgetRow from './BudgetRow.vue'

const props = defineProps({
    data: Object,
    level: Number
})

const isExpanded = ref(true)
const hasChildren = computed(() => props.data.children && props.data.children.length > 0)
</script>