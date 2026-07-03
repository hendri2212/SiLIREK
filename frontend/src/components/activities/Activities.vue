<template>
    <div class="card shadow-sm rounded-4 border-0">
        <div class="card-body">
            <div class="mb-4 d-flex align-items-center justify-content-between">
                <h4 class="card-title text-info">Data Kegiatan</h4>
                <router-link v-if="isAdminOrSuper" class="btn btn-info text-white" :to="{ name: 'activities.create' }">
                    <i class="bi bi-person-fill-add me-1"></i>
                    Kegiatan
                </router-link>
            </div>
            <div v-if="isLoading" class="text-center my-4">
                Loading...
            </div>
            <div v-else class="table-responsive">
                <table class="table table-border-bottom-0">
                    <thead>
                        <tr>
                            <th class="text-center">No</th>
                            <th>Kode Kegiatan</th>
                            <th>Nama Kegiatan</th>
                            <th v-if="isAdminOrSuper" class="text-center" scope="col">Aksi</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="(data, index) in activities" :key="data.id" class="align-middle">
                            <td class="text-center">{{ index + 1 }}</td>
                            <td>{{ data.code }}</td>
                            <td>{{ data.name }}</td>
                            <td v-if="isAdminOrSuper" class="text-center">
                                <router-link :to="{ name: 'activities.edit', params: { id: data.id } }"
                                    class="btn btn-warning btn-sm me-2">
                                    <i class="bi bi-pencil-fill"></i> Edit
                                </router-link>
                                <button class="btn btn-danger btn-sm" @click="deleteActivity(data.id)">
                                    <i class="bi bi-trash-fill"></i> Hapus
                                </button>
                            </td>
                        </tr>
                        <tr v-if="activities.length === 0">
                            <td colspan="4" class="text-center text-muted">Belum ada data kegiatan</td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, onBeforeMount, computed } from 'vue'
import { useAuthStore } from '@/stores/auth'
import api from '@/plugins/axios'

const authStore = useAuthStore()
const isAdminOrSuper = computed(() => authStore.user?.role === 'superadmin' || authStore.user?.role === 'admin')

const isLoading = ref(true)
const activities = ref([])

const fetchActivities = async () => {
    try {
        const response = await api.get('/activity')
        activities.value = response.data
    } catch (error) {
        console.error('Error fetching activities:', error)
    } finally {
        isLoading.value = false
    }
}

const deleteActivity = async (id) => {
    if (confirm('Apakah Anda yakin ingin menghapus kegiatan ini?')) {
        try {
            await api.delete(`/activity/${id}`)
            alert('Kegiatan berhasil dihapus')
            await fetchActivities()
        } catch (error) {
            console.error('Error deleting activity:', error)
            alert('Gagal menghapus kegiatan.')
        }
    }
}

onBeforeMount(() => {
    fetchActivities()
})
</script>