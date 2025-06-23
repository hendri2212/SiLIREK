<template>
    <div class="card shadow-sm rounded-4 border-0">
        <div class="card-body">
            <div class="mb-4 d-flex align-items-center justify-content-between">
                <h4 class="card-title text-info">Penggunaan Anggaran</h4>
            </div>
            <form @submit.prevent="updateBudget">
                <div class="row mb-3">
                    <label class="col-sm-2 col-form-label">Nama Kegiatan</label>
                    <div class="col-sm-10">
                        <input type="text" class="form-control-plaintext" :value="data.activity?.name" readonly>
                    </div>
                </div>
                <div class="row mb-3">
                    <label class="col-sm-2 col-form-label">Pagu Anggaran</label>
                    <div class="col-sm-10">
                        <input type="text" class="form-control-plaintext" :value="data.budget?.toLocaleString('id-ID', { style: 'currency', currency: 'IDR' })" readonly>
                    </div>
                </div>
                <div class="mb-3 row">
                    <label class="col-sm-2 col-form-label">Sisa Anggaran</label>
                    <div class="col-sm-10">
                        <input type="text" readonly class="form-control-plaintext" :value="data.remaining_budget?.toLocaleString('id-ID', { style: 'currency', currency: 'IDR' })">
                    </div>
                </div>
                <div class="mb-3 row">
                    <label class="col-sm-2 col-form-label">Penggunaan</label>
                    <div class="col-sm-10">
                        <input type="text" ref="usedBudgetInput" :value="formattedUsedBudget" @input="onUsedBudgetInput" class="form-control" required>
                    </div>
                </div>
                <div class="d-flex justify-content-end">
                    <button type="submit" class="btn btn-info text-white">
                        <i class="bi bi-floppy me-1"></i>
                        Update</button>
                </div>
            </form>
        </div>
    </div>
</template>
<script setup>
import { ref, onMounted, computed, nextTick } from 'vue'
import api from '@/plugins/axios'
import { useRoute, useRouter } from 'vue-router'

const usedBudgetInput = ref(null)

const router = useRouter()

const route = useRoute()
const data = ref({})

onMounted(async () => {
    try {
        const response = await api.get(`/budgets/${route.params.id}`)
        data.value = response.data
        nextTick(() => usedBudgetInput.value?.focus())
    } catch (error) {
        console.error('Gagal mengambil data anggaran:', error)
    }
})

async function updateBudget() {
    try {
        await api.put(`/budgets/${route.params.id}`, {
            used_budget: data.value.used_budget,
        })
        router.push({ name: 'budgeting.using' })
    } catch (error) {
        console.error('Gagal memperbarui data anggaran:', error)
        alert('Terjadi kesalahan saat memperbarui data.')
    }
}

const formattedUsedBudget = computed(() => {
    if (data.value.used_budget === null || data.value.used_budget === undefined) return ''
    return Number(data.value.used_budget).toLocaleString('id-ID')
})

function onUsedBudgetInput(event) {
    const raw = event.target.value.replace(/\./g, '').replace(/,/g, '')
    data.value.used_budget = Number(raw)
}
</script>