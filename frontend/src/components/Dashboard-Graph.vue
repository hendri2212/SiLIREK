<template>
    <div class="card shadow-sm rounded-4 border-0">
        <div class="card-body">
            <div class="mb-4 d-flex align-items-center">
                <h4 class="card-title text-info">Tambah Pengguna</h4>
            </div>
            <apexchart type="bar" height="300" :options="chartOptions" :series="series" class="mb-4 w-100" />
        </div>
    </div>
</template>
<script setup>
import { ref, computed, onMounted } from 'vue'
import api from '@/plugins/axios';
import ApexCharts from 'vue3-apexcharts'
const apexchart = ApexCharts

const budgets = ref([])

onMounted(async () => {
    try {
        const response = await api.get('/budgets')
        budgets.value = response.data
    } catch (error) {
        console.error('Gagal mengambil data budgets:', error)
    }
})

const series = computed(() => [
    {
        name: 'Pagu',
        data: budgets.value.map(b => b.budget)
    },
    {
        name: 'Terpakai',
        data: budgets.value.map(b => b.used_budget)
    }
])

const chartOptions = computed(() => ({
    chart: {
        id: 'budget-bar'
    },
    xaxis: {
        categories: budgets.value.map(b => b.Activity?.name || 'Tanpa Nama')
    },
    tooltip: {
        y: {
            formatter: val => `Rp${val.toLocaleString('id-ID')}`
        }
    }
}))

defineExpose({ apexchart })
</script>