<template>
    <div class="sub-activities-container pb-5">
        <div class="mb-5 d-flex flex-column flex-md-row align-items-md-center justify-content-between gap-3">
            <div>
                <nav aria-label="breadcrumb">
                    <ol class="breadcrumb mb-1">
                        <li class="breadcrumb-item">
                            <router-link :to="{ name: 'activities.list' }" class="text-decoration-none text-muted">Kegiatan</router-link>
                        </li>
                        <li class="breadcrumb-item active" aria-current="page">Sub Kegiatan</li>
                    </ol>
                </nav>
                <h3 class="fw-bold text-dark mb-1">
                    {{ parentActivity ? parentActivity.name : 'Memuat...' }}
                </h3>
                <p class="text-muted mb-0">Manajemen sub kegiatan untuk kegiatan ini.</p>
            </div>
            
            <div class="d-flex align-items-center gap-3">
                <router-link class="btn btn-white border-0 bg-white rounded-circle shadow-sm d-flex align-items-center justify-content-center hover-lift" style="width: 45px; height: 45px;" :to="{ name: 'activities.list' }" title="Kembali ke Daftar Kegiatan">
                    <i class="bi bi-arrow-left fs-5 text-dark"></i>
                </router-link>
                
                <router-link v-if="isAdminOrSuper" class="btn btn-info text-white rounded-pill px-4 py-2 shadow-sm fw-bold d-flex align-items-center gap-2 hover-lift" :to="{ name: 'subactivities.create', params: { activityId: route.params.activityId } }">
                    <i class="bi bi-plus-lg fs-6"></i>
                    <span>Tambah</span>
                </router-link>
            </div>
        </div>

        <div v-if="isLoading" class="d-flex justify-content-center align-items-center" style="min-height: 300px;">
            <div class="spinner-grow text-info" role="status" style="width: 3rem; height: 3rem;">
                <span class="visually-hidden">Loading...</span>
            </div>
        </div>

        <div v-else-if="subActivities.length === 0" class="text-center py-5 bg-white rounded-4 shadow-sm border-0 empty-state">
            <div class="display-1 text-info opacity-50 mb-3"><i class="bi bi-node-plus-fill"></i></div>
            <h5 class="fw-bold text-dark">Belum ada sub kegiatan</h5>
            <p class="text-muted mb-4">Belum ada sub kegiatan yang terdaftar untuk kegiatan ini.</p>
            <router-link v-if="isAdminOrSuper" :to="{ name: 'subactivities.create', params: { activityId: route.params.activityId } }" class="btn btn-outline-info rounded-pill px-4 py-2 fw-medium hover-lift">
                Buat Sub Kegiatan Pertama
            </router-link>
        </div>

        <div v-else class="row row-cols-1 row-cols-md-2 row-cols-xl-3 g-4">
            <div class="col" v-for="(data, index) in subActivities" :key="data.id">
                <div class="card h-100 border-0 shadow-sm rounded-4 activity-card position-relative overflow-hidden" @click="goToExpenditureAccounts(data.id)" style="cursor: pointer;">
                    <div class="position-absolute top-0 end-0 p-3 text-info opacity-10" style="transform: translate(20%, -20%); pointer-events: none;">
                        <i class="bi bi-diagram-3-fill" style="font-size: 8rem;"></i>
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
                        
                        <h5 class="card-title fw-bold text-dark mb-4 flex-grow-1" style="line-height: 1.5; font-size: 1.15rem;">
                            {{ data.name }}
                        </h5>
                        
                        <div class="divider mb-3"></div>
                        
                        <div v-if="isAdminOrSuper" class="d-flex gap-2 justify-content-end mt-auto">
                            <router-link :to="{ name: 'subactivities.edit', params: { activityId: route.params.activityId, subActivityId: data.id } }"
                                @click.stop
                                class="btn btn-light text-warning fw-bold rounded-pill px-3 py-2 shadow-sm btn-action w-50">
                                <i class="bi bi-pencil-fill me-1"></i> Edit
                            </router-link>
                            <button class="btn btn-light text-danger fw-bold rounded-pill px-3 py-2 shadow-sm btn-action w-50" @click.stop="deleteSubActivity(data.id)">
                                <i class="bi bi-trash-fill me-1"></i> Hapus
                            </button>
                        </div>
                        <div v-else class="d-flex align-items-center justify-content-end text-success mt-auto p-2 bg-success-subtle rounded-pill">
                            <i class="bi bi-check-circle-fill me-2 ms-2"></i>
                            <span class="small fw-bold me-2">Sub Kegiatan Aktif</span>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import api from '@/plugins/axios'

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()
const isAdminOrSuper = computed(() => authStore.user?.role === 'superadmin' || authStore.user?.role === 'admin')

const isLoading = ref(true)
const subActivities = ref([])
const parentActivity = ref(null)

const fetchParentActivity = async () => {
    try {
        const response = await api.get(`/activity/${route.params.activityId}`)
        parentActivity.value = response.data
    } catch (error) {
        console.error('Error fetching parent activity:', error)
    }
}

const fetchSubActivities = async () => {
    try {
        const response = await api.get(`/sub-activity?activity_id=${route.params.activityId}`)
        subActivities.value = response.data
    } catch (error) {
        console.error('Error fetching sub activities:', error)
    } finally {
        isLoading.value = false
    }
}

const deleteSubActivity = async (id) => {
    if (confirm('Apakah Anda yakin ingin menghapus sub kegiatan ini?')) {
        try {
            await api.delete(`/sub-activity/${id}`)
            alert('Sub Kegiatan berhasil dihapus')
            await fetchSubActivities()
        } catch (error) {
            console.error('Error deleting sub activity:', error)
            alert('Gagal menghapus sub kegiatan.')
        }
    }
}

const goToExpenditureAccounts = (subId) => {
    router.push({ name: 'expenditureaccounts.list', params: { activityId: route.params.activityId, subActivityId: subId } })
}

onMounted(async () => {
    // Jalankan secara paralel untuk performa terbaik
    await Promise.all([fetchParentActivity(), fetchSubActivities()])
})
</script>

<style scoped>
.sub-activities-container {
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
