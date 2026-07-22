<template>
    <div class="expenditure-accounts-container pb-5">
        <div class="mb-5 d-flex flex-column flex-md-row align-items-md-center justify-content-between gap-3">
            <div>
                <nav aria-label="breadcrumb">
                    <ol class="breadcrumb mb-1">
                        <li class="breadcrumb-item">
                            <router-link :to="{ name: 'activities.list' }" class="text-decoration-none text-muted">Kegiatan</router-link>
                        </li>
                        <li class="breadcrumb-item">
                            <router-link :to="{ name: 'subactivities.list', params: { activityId: route.params.activityId } }" class="text-decoration-none text-muted">Sub Kegiatan</router-link>
                        </li>
                        <li class="breadcrumb-item active" aria-current="page">Rekening Belanja</li>
                    </ol>
                </nav>
                <h3 class="fw-bold text-dark mb-1">
                    {{ parentSubActivity ? parentSubActivity.name : 'Memuat...' }}
                </h3>
                <p class="text-muted mb-0">Manajemen rekening belanja untuk sub kegiatan ini.</p>
            </div>
            
            <div class="d-flex align-items-center gap-3">
                <router-link class="btn btn-white border-0 bg-white rounded-circle shadow-sm d-flex align-items-center justify-content-center hover-lift" style="width: 45px; height: 45px;" :to="{ name: 'subactivities.list', params: { activityId: route.params.activityId } }" title="Kembali ke Daftar Sub Kegiatan">
                    <i class="bi bi-arrow-left fs-5 text-dark"></i>
                </router-link>
                
                <router-link v-if="isAdminOrSuper" class="btn btn-info text-white rounded-pill px-4 py-2 shadow-sm fw-bold d-flex align-items-center gap-2 hover-lift" :to="{ name: 'expenditureaccounts.create', params: { activityId: route.params.activityId, subActivityId: route.params.subActivityId } }">
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

        <div v-else-if="accounts.length === 0" class="text-center py-5 bg-white rounded-4 shadow-sm border-0 empty-state">
            <div class="display-1 text-info opacity-50 mb-3"><i class="bi bi-cash-stack"></i></div>
            <h5 class="fw-bold text-dark">Belum ada rekening belanja</h5>
            <p class="text-muted mb-4">Belum ada rekening belanja yang terdaftar untuk sub kegiatan ini.</p>
            <router-link v-if="isAdminOrSuper" :to="{ name: 'expenditureaccounts.create', params: { activityId: route.params.activityId, subActivityId: route.params.subActivityId } }" class="btn btn-outline-info rounded-pill px-4 py-2 fw-medium hover-lift">
                Buat Rekening Belanja Pertama
            </router-link>
        </div>

        <div v-else class="d-flex flex-column gap-2">
            <div v-for="(data, index) in accounts" :key="data.id" class="card border-0 shadow-sm rounded-4 activity-card position-relative overflow-hidden w-100" @click="goToItems(data.id)" style="cursor: pointer;">
                <div class="position-absolute top-50 start-0 translate-middle-y ms-2 text-info opacity-10 d-none d-lg-block" style="pointer-events: none;">
                    <i class="bi bi-cash-stack" style="font-size: 6rem;"></i>
                </div>
                
                <div class="card-body p-3 px-4 d-flex flex-column flex-lg-row align-items-lg-center justify-content-between gap-3 z-1 position-relative">
                    
                    <!-- Left: Index, Code, Description -->
                    <div class="d-flex align-items-center gap-3 col-lg-4">
                        <div class="index-indicator text-muted d-flex align-items-center justify-content-center rounded-circle bg-light fw-bold flex-shrink-0" style="width: 35px; height: 35px;">
                            {{ index + 1 }}
                        </div>
                        <div class="min-w-0">
                            <span class="badge bg-info-subtle text-info px-2 py-1 rounded border border-info-subtle mb-1 d-inline-block">
                                <i class="bi bi-hash me-1"></i> {{ data.code }}
                            </span>
                            <h6 class="fw-bold text-dark mb-0 text-line-clamp-2" :title="data.description">
                                {{ data.description }}
                            </h6>
                        </div>
                    </div>

                    <!-- Middle: Budget Info -->
                    <div class="d-flex flex-column flex-md-row gap-3 gap-md-4 col-lg-5 justify-content-lg-center bg-light p-2 rounded-3 text-nowrap">
                        <div class="text-md-end flex-grow-1">
                            <div class="text-muted small fw-medium">Pagu Anggaran</div>
                            <div class="fw-bold text-success">{{ formatCurrency(data.budget_ceiling) }}</div>
                        </div>
                        <div class="divider-vertical d-none d-md-block"></div>
                        <div class="text-md-end flex-grow-1">
                            <div class="text-muted small fw-medium">Total Realisasi</div>
                            <div class="fw-bold text-danger">{{ formatCurrency(data.total_credit || 0) }}</div>
                        </div>
                        <div class="divider-vertical d-none d-md-block"></div>
                        <div class="text-md-end flex-grow-1">
                            <div class="text-dark small fw-bold">Sisa Anggaran</div>
                            <div class="fw-bold" :class="(data.remaining_budget >= 0 || data.remaining_budget == null) ? 'text-primary' : 'text-danger'">
                                {{ formatCurrency(data.remaining_budget !== undefined ? data.remaining_budget : data.budget_ceiling) }}
                            </div>
                        </div>
                    </div>

                    <!-- Right: Actions -->
                    <div class="col-lg-auto d-flex justify-content-lg-end align-items-center pt-3 pt-lg-0 mt-2 mt-lg-0">
                        <div v-if="isAdminOrSuper" class="d-flex gap-2 w-100 justify-content-end">
                            <router-link :to="{ name: 'expenditureaccounts.edit', params: { activityId: route.params.activityId, subActivityId: route.params.subActivityId, accountId: data.id } }"
                                @click.stop
                                class="btn btn-light text-warning fw-bold rounded-circle shadow-sm btn-action d-flex align-items-center justify-content-center" style="width: 36px; height: 36px;" title="Edit">
                                <i class="bi bi-pencil-fill"></i>
                            </router-link>
                            <button class="btn btn-light text-danger fw-bold rounded-circle shadow-sm btn-action d-flex align-items-center justify-content-center" style="width: 36px; height: 36px;" @click.stop="deleteAccount(data.id)" title="Hapus">
                                <i class="bi bi-trash-fill"></i>
                            </button>
                        </div>
                        <div v-else class="d-inline-flex align-items-center justify-content-center text-success p-2 bg-success-subtle rounded-pill w-100" style="height: 36px;">
                            <i class="bi bi-check-circle-fill me-2"></i>
                            <span class="small fw-bold">Aktif</span>
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
const accounts = ref([])
const parentSubActivity = ref(null)

const formatCurrency = (value) => {
    return new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR' }).format(value);
}

const fetchParentSubActivity = async () => {
    try {
        const response = await api.get(`/sub-activity/${route.params.subActivityId}`)
        parentSubActivity.value = response.data
    } catch (error) {
        console.error('Error fetching parent sub activity:', error)
    }
}

const fetchAccounts = async () => {
    try {
        const response = await api.get(`/expenditure-account?sub_activity_id=${route.params.subActivityId}`)
        accounts.value = response.data || []
    } catch (error) {
        console.error('Error fetching expenditure accounts:', error)
    } finally {
        isLoading.value = false
    }
}

const deleteAccount = async (id) => {
    if (confirm('Apakah Anda yakin ingin menghapus rekening belanja ini?')) {
        try {
            await api.delete(`/expenditure-account/${id}`)
            alert('Rekening belanja berhasil dihapus')
            await fetchAccounts()
        } catch (error) {
            console.error('Error deleting account:', error)
            alert('Gagal menghapus rekening belanja.')
        }
    }
}

const goToItems = (accountId) => {
    router.push({ name: 'items.list', params: { activityId: route.params.activityId, subActivityId: route.params.subActivityId, accountId: accountId } })
}

onMounted(async () => {
    // Jalankan secara paralel untuk performa terbaik
    await Promise.all([fetchParentSubActivity(), fetchAccounts()])
})
</script>

<style scoped>
.expenditure-accounts-container {
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

.divider-vertical {
    width: 1px;
    background-color: rgba(0,0,0,0.05);
    height: auto;
}

.min-w-0 {
    min-width: 0;
}

.text-line-clamp-2 {
    display: -webkit-box;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
    overflow: hidden;
    white-space: normal;
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
