<template>
    <div class="card shadow-sm rounded-4 border-0">
        <div class="card-body">
            <div class="mb-4 d-flex align-items-center">
                <h4 class="card-title text-info">Edit Kegiatan</h4>
            </div>
            
            <form @submit.prevent="updateActivity" v-if="!isLoading">
                <div class="form-floating mb-3">
                    <input v-model="activity.code" type="text" class="form-control" id="floatingActivity"
                        placeholder="Kode Kegiatan" required>
                    <label for="floatingActivity">Kode Kegiatan</label>
                </div>
                <div class="form-floating mb-3">
                    <input v-model="activity.name" type="text" class="form-control" id="floatingNameActivity" placeholder="Nama Kegiatan"
                        required>
                    <label for="floatingNameActivity">Nama Kegiatan</label>
                </div>
                <div class="d-flex justify-content-end">
                    <router-link :to="{ name: 'activities.list' }" class="btn btn-secondary me-2">
                        Batal
                    </router-link>
                    <button type="submit" class="btn btn-info text-white">
                        <i class="bi bi-floppy-fill me-1"></i> Update
                    </button>
                </div>
            </form>
            
            <div v-else class="text-center py-4">
                <div class="spinner-border text-info" role="status">
                    <span class="visually-hidden">Loading...</span>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import api from '@/plugins/axios'

const router = useRouter()
const route = useRoute()
const activityId = route.params.id

const isLoading = ref(true)
const activity = ref({
    code: '',
    name: '',
})

const fetchActivity = async () => {
    try {
        const response = await api.get(`/activity/${activityId}`)
        activity.value.code = response.data.code
        activity.value.name = response.data.name
    } catch (error) {
        console.error('Error fetching activity:', error)
        alert('Gagal mengambil data kegiatan')
        router.push({ name: 'activities.list' })
    } finally {
        isLoading.value = false
    }
}

const updateActivity = async () => {
    try {
        await api.put(`/activity/${activityId}`, activity.value)
        alert('Kegiatan berhasil diperbarui')
        router.push({ name: 'activities.list' })
    } catch (error) {
        console.error('Error updating activity:', error)
        if (error.response && error.response.data && error.response.data.error) {
            alert(error.response.data.error)
        } else {
            alert('Terjadi kesalahan saat memperbarui kegiatan.')
        }
    }
}

onMounted(() => {
    fetchActivity()
})
</script>