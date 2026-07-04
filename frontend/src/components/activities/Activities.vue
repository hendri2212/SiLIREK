<template>
    <div class="activities-container pb-5">
        <div class="mb-5 d-flex flex-column flex-md-row align-items-md-center justify-content-between gap-3">
            <div>
                <h3 class="fw-bold text-dark mb-1">
                    <i class="bi bi-journal-text me-2 text-success opacity-75"></i>Daftar Kegiatan
                </h3>
                <p class="text-muted mb-0">Manajemen seluruh kegiatan yang terdaftar dalam sistem.</p>
            </div>
            <router-link v-if="isAdminOrSuper" class="btn btn-info text-white rounded-pill px-4 py-2 shadow-sm fw-medium d-flex align-items-center gap-2 hover-lift" :to="{ name: 'activities.create' }">
                <i class="bi bi-plus-circle fs-5"></i>
                <span>Tambah Kegiatan Baru</span>
            </router-link>
        </div>

        <div v-if="isLoading" class="d-flex justify-content-center align-items-center" style="min-height: 300px;">
            <div class="spinner-grow text-info" role="status" style="width: 3rem; height: 3rem;">
                <span class="visually-hidden">Loading...</span>
            </div>
        </div>

        <div v-else-if="activities.length === 0" class="text-center py-5 bg-white rounded-4 shadow-sm border-0 empty-state">
            <div class="display-1 text-info opacity-50 mb-3"><i class="bi bi-inboxes-fill"></i></div>
            <h5 class="fw-bold text-dark">Belum ada kegiatan</h5>
            <p class="text-muted mb-4">Data kegiatan saat ini masih kosong di dalam sistem.</p>
            <router-link v-if="isAdminOrSuper" :to="{ name: 'activities.create' }" class="btn btn-outline-info rounded-pill px-4 py-2 fw-medium hover-lift">
                Buat Kegiatan Pertama
            </router-link>
        </div>

        <div v-else class="row row-cols-1 row-cols-md-2 row-cols-xl-3 g-4">
            <div class="col" v-for="(data, index) in activities" :key="data.id">
                <div class="card h-100 border-0 shadow-sm rounded-4 activity-card position-relative overflow-hidden" style="cursor: pointer;" @click="goToSubActivities(data.id)">
                    <!-- Decorative background element -->
                    <div class="position-absolute top-0 end-0 p-3 text-info opacity-10" style="transform: translate(20%, -20%); pointer-events: none;">
                        <i class="bi bi-activity" style="font-size: 8rem;"></i>
                    </div>
                    
                    <div class="card-body p-4 d-flex flex-column z-1 position-relative">
                        <div class="d-flex justify-content-between align-items-start mb-4">
                            <span class="badge bg-info-subtle text-info fs-6 px-3 py-2 rounded-pill border border-info-subtle shadow-sm">
                                <i class="bi bi-hash me-1"></i> {{ data.code }}
                            </span>
                            <div class="index-indicator text-muted d-flex align-items-center justify-content-center rounded-circle bg-light fw-bold" style="width: 35px; height: 35px;">
                                {{ index + 1 }}
                            </div>
                        </div>
                        
                        <h5 class="card-title fw-bold text-dark mb-4 flex-grow-1" style="line-height: 1.5; font-size: 1.25rem;">
                            {{ data.name }}
                        </h5>
                        
                        <div class="divider mb-3"></div>
                        
                        <div v-if="isAdminOrSuper" class="d-flex gap-2 justify-content-end mt-auto" @click.stop>
                            <router-link :to="{ name: 'activities.edit', params: { id: data.id } }"
                                class="btn btn-light text-warning fw-bold rounded-pill px-3 py-2 shadow-sm btn-action w-50">
                                <i class="bi bi-pencil-fill me-1"></i> Edit
                            </router-link>
                            <button class="btn btn-light text-danger fw-bold rounded-pill px-3 py-2 shadow-sm btn-action w-50" @click.stop="deleteActivity(data.id)">
                                <i class="bi bi-trash-fill me-1"></i> Hapus
                            </button>
                        </div>
                        <div v-else class="d-flex align-items-center justify-content-end text-success mt-auto p-2 bg-success-subtle rounded-pill">
                            <i class="bi bi-check-circle-fill me-2 ms-2"></i>
                            <span class="small fw-bold me-2">Kegiatan Terdaftar</span>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import api from '@/plugins/axios'

const router = useRouter()
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

const goToSubActivities = (id) => {
    router.push({ name: 'subactivities.list', params: { activityId: id } })
}

onMounted(() => {
    fetchActivities()
})
</script>

<style scoped>
.activities-container {
    animation: fadeIn 0.4s ease-out;
}

.activity-card {
    transition: all 0.3s cubic-bezier(0.25, 0.8, 0.25, 1);
    background: linear-gradient(160deg, #ffffff, #fcfcfc);
}

.activity-card:hover {
    transform: translateY(-8px);
    box-shadow: 0 15px 30px rgba(0, 0, 0, 0.08) !important;
}

.divider {
    height: 1px;
    background: linear-gradient(90deg, rgba(0,0,0,0.05) 0%, rgba(0,0,0,0.1) 50%, rgba(0,0,0,0.05) 100%);
    width: 100%;
}

.btn-action {
    transition: all 0.25s ease;
    border: 1px solid transparent;
}

.btn-action.text-warning:hover {
    background-color: #ffc107 !important;
    color: #fff !important;
    border-color: #ffc107;
    transform: translateY(-2px);
}

.btn-action.text-danger:hover {
    background-color: #dc3545 !important;
    color: #fff !important;
    border-color: #dc3545;
    transform: translateY(-2px);
}

.hover-lift {
    transition: transform 0.2s ease, box-shadow 0.2s ease;
}

.hover-lift:hover {
    transform: translateY(-3px);
    box-shadow: 0 8px 15px rgba(13, 202, 240, 0.3) !important;
}

.empty-state {
    background: linear-gradient(to bottom, #ffffff, #f8f9fa) !important;
    border: 1px dashed rgba(0,0,0,0.1) !important;
}

@keyframes fadeIn {
    from { opacity: 0; transform: translateY(10px); }
    to { opacity: 1; transform: translateY(0); }
}
</style>