<template>
    <div class="card shadow-sm rounded-4 border-0">
        <div class="card-body">
            <div class="mb-4 d-flex align-items-center justify-content-between">
                <h4 class="card-title text-info">Penggunaan Anggaran</h4>
                <div class="col-sm-4">
                    <input type="text" class="form-control" v-model="searchTerm" placeholder="Cari kegiatan...">
                </div>
            </div>
            <div class="table-responsive">
                <table class="table table-border-bottom-0 align-middle">
                    <thead class="text-nowrap">
                        <tr>
                            <th class="text-center">No</th>
                            <th>Kegiatan</th>
                            <th class="text-end">Pagu</th>
                            <th class="text-end">Penggunakan</th>
                            <th class="text-end">Sisa</th>
                            <th></th>
                        </tr>
                    </thead>
                    <tbody>
                        <!-- <tr v-for="(data, index) in budgets.filter(b => b.budget !== 0)" :key="data.id"> -->
                        <tr v-for="(data, index) in filteredBudgets" :key="data.id">
                            <td class="text-center">{{ index + 1 }}</td>
                            <td>{{ data.activity?.name }}</td>
                            <td class="text-end">{{ data.budget === 0 ? '' : data.budget?.toLocaleString('id-ID', { style: 'currency', currency: 'IDR' }) }}</td>
                            <td class="text-end">{{ data.budget === 0 ? '' : data.used_budget?.toLocaleString('id-ID', { style: 'currency', currency: 'IDR' }) }}</td>
                            <td class="text-end">{{ data.budget === 0 ? '' : data.remaining_budget?.toLocaleString('id-ID', { style: 'currency', currency: 'IDR' }) }}</td>
                            <td>
                                <div class="btn-group" role="group" aria-label="Basic example" v-if="data.budget != 0">
                                    <router-link
                                      class="btn btn-outline-warning d-flex align-items-center gap-2 px-3 py-1 rounded-pill shadow-sm"
                                      :to="{ name: 'budgeting.edit', params: { id: data.id } }"
                                    >
                                      <i class="bi bi-pencil-square fs-6"></i>
                                      <span>Edit</span>
                                    </router-link>
                                </div>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import api from '@/plugins/axios'
import { useAuthStore } from '@/stores/auth'

const auth = useAuthStore()
const budgets = ref([])

const searchTerm = ref('')

const filteredBudgets = computed(() => {
    if (!searchTerm.value) return budgets.value
    return budgets.value.filter(b => 
        b.activity?.name?.toLowerCase().includes(searchTerm.value.toLowerCase())
    )
})

onMounted(async () => {
    try {
        const response = await api.get(`/budgets/user/${auth.user.id}`)
        budgets.value = response.data
    } catch (error) {
        console.error('Gagal memuat data anggaran:', error)
    }
})
</script>