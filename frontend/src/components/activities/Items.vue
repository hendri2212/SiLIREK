<template>
    <div class="items-container pb-5">
        <div class="mb-5 d-flex flex-column flex-md-row align-items-md-center justify-content-between gap-3">
            <div>
                <nav aria-label="breadcrumb">
                    <ol class="breadcrumb mb-1">
                        <li class="breadcrumb-item">
                            <router-link :to="{ name: 'activities.list' }"
                                class="text-decoration-none text-muted">Kegiatan</router-link>
                        </li>
                        <li class="breadcrumb-item">
                            <router-link
                                :to="{ name: 'subactivities.list', params: { activityId: route.params.activityId } }"
                                class="text-decoration-none text-muted">Sub Kegiatan</router-link>
                        </li>
                        <li class="breadcrumb-item">
                            <router-link
                                :to="{ name: 'expenditureaccounts.list', params: { activityId: route.params.activityId, subActivityId: route.params.subActivityId } }"
                                class="text-decoration-none text-muted">Rekening Belanja</router-link>
                        </li>
                        <li class="breadcrumb-item active" aria-current="page">Item Belanja</li>
                    </ol>
                </nav>
                <h3 class="fw-bold text-dark mb-1">
                    {{ parentAccount ? parentAccount.description : 'Memuat...' }}
                </h3>
                <p class="text-muted mb-0">Manajemen item belanja untuk rekening ini.</p>
            </div>

            <div class="d-flex align-items-center gap-3">
                <router-link
                    class="btn btn-white border-0 bg-white rounded-circle shadow-sm d-flex align-items-center justify-content-center hover-lift"
                    style="width: 45px; height: 45px;"
                    :to="{ name: 'expenditureaccounts.list', params: { activityId: route.params.activityId, subActivityId: route.params.subActivityId } }"
                    title="Kembali">
                    <i class="bi bi-arrow-left fs-5 text-dark"></i>
                </router-link>

                <router-link v-if="isAdminOrSuper"
                    class="btn btn-info text-white rounded-pill px-4 py-2 shadow-sm fw-bold d-flex align-items-center gap-2 hover-lift"
                    :to="{ name: 'items.create', params: { activityId: route.params.activityId, subActivityId: route.params.subActivityId, accountId: route.params.accountId } }">
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

        <div v-else-if="items.length === 0" class="text-center py-5 bg-white rounded-4 shadow-sm border-0 empty-state">
            <div class="display-1 text-info opacity-50 mb-3"><i class="bi bi-list-check"></i></div>
            <h5 class="fw-bold text-dark">Belum ada item belanja</h5>
            <p class="text-muted mb-4">Belum ada item belanja yang terdaftar untuk rekening ini.</p>
            <router-link v-if="isAdminOrSuper"
                :to="{ name: 'items.create', params: { activityId: route.params.activityId, subActivityId: route.params.subActivityId, accountId: route.params.accountId } }"
                class="btn btn-outline-info rounded-pill px-4 py-2 fw-medium hover-lift">
                Buat Item Belanja Pertama
            </router-link>
        </div>

        <div v-else class="d-flex flex-column gap-3">
            <!-- Budget Summary Card -->
            <div class="card border-0 shadow-sm rounded-4 overflow-hidden mb-2 bg-gradient-light">
                <div class="card-body p-4 d-flex flex-column flex-md-row justify-content-between align-items-md-center gap-3">
                    <div class="d-flex align-items-center gap-3">
                        <div class="bg-white p-3 rounded-circle shadow-sm text-info">
                            <i class="bi bi-wallet2 fs-3"></i>
                        </div>
                        <div>
                            <p class="text-muted mb-1 fw-medium">Pagu Anggaran</p>
                            <h4 class="fw-bold text-dark mb-0">{{ formatCurrency(parentAccount?.budget_ceiling || 0) }}</h4>
                        </div>
                    </div>
                    
                    <div class="d-flex flex-column flex-md-row gap-4">
                        <div class="text-md-end">
                            <p class="text-muted mb-1 small fw-medium">Total Kredit</p>
                            <h5 class="fw-bold text-danger mb-0">{{ formatCurrency(totalCredit) }}</h5>
                        </div>
                        <div class="divider-vertical d-none d-md-block"></div>
                        <div class="text-md-end">
                            <p class="text-muted mb-1 small fw-medium">Sisa Anggaran</p>
                            <h5 class="fw-bold mb-0" :class="remainingBudget >= 0 ? 'text-success' : 'text-danger'">
                                {{ formatCurrency(remainingBudget) }}
                            </h5>
                        </div>
                    </div>
                </div>
            </div>

            <!-- Items List -->
            <div v-for="(item, index) in items" :key="item.id" class="card border-0 shadow-sm rounded-4 item-card mb-1">
                <div class="card-body px-3 py-2">
                    <div class="row align-items-center w-100 m-0">
                        <!-- Left: Index, Code, Date -->
                        <div class="col-md-5 mb-2 mb-md-0 d-flex align-items-center gap-2 px-0">
                            <div class="index-indicator text-muted d-flex align-items-center justify-content-center rounded-circle bg-light fw-bold flex-shrink-0" style="width: 28px; height: 28px; font-size: 0.85rem;">
                                {{ index + 1 }}
                            </div>
                            <div class="d-flex flex-column flex-md-row align-items-md-center gap-1 gap-md-2">
                                <span class="badge bg-info-subtle text-info px-2 py-0 border border-info-subtle">
                                    <i class="bi bi-hash"></i>{{ item.code }}
                                </span>
                                <div class="text-dark small fw-medium text-nowrap" style="font-size: 0.8rem;"><i class="bi bi-calendar-date text-muted"></i> {{ formatDate(item.date) }}</div>
                            </div>
                        </div>
                        
                        <!-- Right: Kredit -->
                        <div class="col-md-4 mb-2 mb-md-0 text-md-end px-0 pe-md-3">
                            <div class="text-muted small d-md-none" style="font-size: 0.75rem;">Kredit:</div>
                            <h6 class="fw-bold text-danger mb-0 fs-6">{{ formatCurrency(item.credit) }}</h6>
                        </div>
                        
                        <!-- Far Right: Actions -->
                        <div v-if="isAdminOrSuper" class="col-md-3 text-md-end d-flex justify-content-md-end align-items-center gap-2 mt-2 mt-md-0 px-0">
                            <button class="btn btn-sm btn-white bg-light text-muted rounded-circle shadow-sm border p-0 d-flex align-items-center justify-content-center toggle-arrow hover-lift" style="width: 30px; height: 30px;" @click="toggleExpand(item.id)" title="Lihat Uraian">
                                <i class="bi" :class="isExpanded(item.id) ? 'bi-chevron-up' : 'bi-chevron-down'" style="font-size: 0.9rem;"></i>
                            </button>
                            <div class="divider-vertical" style="height: 20px;"></div>
                            <router-link :to="{ name: 'items.edit', params: { activityId: route.params.activityId, subActivityId: route.params.subActivityId, accountId: route.params.accountId, itemId: item.id } }"
                                class="btn btn-sm btn-light text-warning fw-bold rounded-circle action-btn" style="width: 30px; height: 30px;" title="Edit">
                                <i class="bi bi-pencil-fill" style="font-size: 0.8rem;"></i>
                            </router-link>
                            <button class="btn btn-sm btn-light text-danger fw-bold rounded-circle action-btn" style="width: 30px; height: 30px;" @click="deleteItem(item.id)" title="Hapus">
                                <i class="bi bi-trash-fill" style="font-size: 0.8rem;"></i>
                            </button>
                        </div>
                        <div v-else class="col-md-3 text-md-end d-flex justify-content-md-end align-items-center gap-2 mt-2 mt-md-0 px-0">
                            <button class="btn btn-sm btn-white bg-light text-muted rounded-circle shadow-sm border p-0 d-flex align-items-center justify-content-center toggle-arrow hover-lift" style="width: 30px; height: 30px;" @click="toggleExpand(item.id)" title="Lihat Uraian">
                                <i class="bi" :class="isExpanded(item.id) ? 'bi-chevron-up' : 'bi-chevron-down'" style="font-size: 0.9rem;"></i>
                            </button>
                             <div class="d-inline-flex align-items-center justify-content-center text-success px-2 py-1 bg-success-subtle rounded-pill">
                                <i class="bi bi-check-circle-fill me-1" style="font-size: 0.8rem;"></i>
                                <span class="fw-bold" style="font-size: 0.75rem;">Aktif</span>
                            </div>
                        </div>
                    </div>

                    <!-- Collapsible Description -->
                    <div class="collapse-description bg-light rounded-3 overflow-hidden mt-2" :class="{ 'show': isExpanded(item.id) }">
                        <div class="p-2 border-start border-3 border-info">
                            <p class="text-muted mb-0 small" style="white-space: pre-line;">{{ item.description }}</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRoute } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import api from '@/plugins/axios'

const route = useRoute()
const authStore = useAuthStore()
const isAdminOrSuper = computed(() => authStore.user?.role === 'superadmin' || authStore.user?.role === 'admin')

const isLoading = ref(true)
const items = ref([])
const parentAccount = ref(null)

const expandedIds = ref(new Set())

const toggleExpand = (id) => {
    if (expandedIds.value.has(id)) {
        expandedIds.value.delete(id)
    } else {
        expandedIds.value.add(id)
    }
}

const isExpanded = (id) => expandedIds.value.has(id)

const totalCredit = computed(() => {
    return items.value.reduce((sum, item) => sum + Number(item.credit), 0)
})

const remainingBudget = computed(() => {
    if (!parentAccount.value) return 0
    return Number(parentAccount.value.budget_ceiling) - totalCredit.value
})

const formatCurrency = (value) => {
    return new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR' }).format(value);
}

const formatDate = (dateString) => {
    const options = { year: 'numeric', month: 'long', day: 'numeric' }
    return new Date(dateString).toLocaleDateString('id-ID', options)
}

const fetchParentAccount = async () => {
    try {
        const response = await api.get(`/expenditure-account/${route.params.accountId}`)
        parentAccount.value = response.data
    } catch (error) {
        console.error('Error fetching parent account:', error)
    }
}

const fetchItems = async () => {
    try {
        const response = await api.get(`/item?expenditure_account_id=${route.params.accountId}`)
        items.value = response.data || []
    } catch (error) {
        console.error('Error fetching items:', error)
    } finally {
        isLoading.value = false
    }
}

const deleteItem = async (id) => {
    if (confirm('Apakah Anda yakin ingin menghapus item belanja ini?')) {
        try {
            await api.delete(`/item/${id}`)
            alert('Item belanja berhasil dihapus')
            await fetchItems()
        } catch (error) {
            console.error('Error deleting item:', error)
            alert('Gagal menghapus item belanja.')
        }
    }
}

onMounted(async () => {
    // Jalankan secara paralel untuk performa terbaik
    await Promise.all([fetchParentAccount(), fetchItems()])
})
</script>

<style scoped>
.items-container {
    animation: fadeIn 0.4s ease-out;
}

.item-card {
    transition: all 0.3s cubic-bezier(0.25, 0.8, 0.25, 1);
    border: 1px solid transparent;
}

.item-card:hover {
    transform: translateY(-4px);
    box-shadow: 0 10px 20px rgba(0, 0, 0, 0.05) !important;
    border-color: rgba(13, 202, 240, 0.1) !important;
}

.collapse-description {
    max-height: 0;
    opacity: 0;
    transition: max-height 0.4s cubic-bezier(0, 1, 0, 1), opacity 0.3s ease, margin 0.3s ease;
    margin-top: 0 !important;
}

.collapse-description.show {
    max-height: 1000px;
    opacity: 1;
    margin-top: 1rem !important;
    transition: max-height 0.5s ease-in-out, opacity 0.4s ease 0.1s, margin 0.3s ease;
}

.bg-gradient-light {
    background: linear-gradient(135deg, #ffffff 0%, #f8f9fa 100%);
}

.divider-vertical {
    width: 1px;
    background-color: rgba(0,0,0,0.1);
    height: auto;
}

.action-btn {
    width: 36px;
    height: 36px;
    display: inline-flex;
    align-items: center;
    justify-content: center;
    transition: all 0.2s ease;
    border: 1px solid transparent;
}

.action-btn.text-warning:hover {
    background-color: #ffc107 !important;
    color: #fff !important;
    transform: translateY(-2px);
    box-shadow: 0 4px 8px rgba(255, 193, 7, 0.2);
}

.action-btn.text-danger:hover {
    background-color: #dc3545 !important;
    color: #fff !important;
    transform: translateY(-2px);
    box-shadow: 0 4px 8px rgba(220, 53, 69, 0.2);
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

.toggle-arrow {
    transition: all 0.3s ease;
}

.toggle-arrow:hover {
    background-color: #0dcaf0 !important;
    color: #fff !important;
    border-color: #0dcaf0 !important;
}

@keyframes fadeIn {
    from { opacity: 0; transform: translateY(10px); }
    to { opacity: 1; transform: translateY(0); }
}
</style>
