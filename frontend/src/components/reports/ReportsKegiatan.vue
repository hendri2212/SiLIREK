<template>
    <div class="mb-5">
        <div class="mb-5">
            <h3 class="fw-bold text-dark mb-1">
                <i class="bi bi-diagram-3 me-2 text-success opacity-75"></i>Laporan Pohon Kegiatan
            </h3>
            <p class="text-muted mb-0">Ringkasan seluruh kegiatan, sub kegiatan, dan rekening belanja.</p>
        </div>

        <div v-if="isLoading" class="text-center my-5">
            <div class="spinner-border text-primary" role="status">
                <span class="visually-hidden">Loading...</span>
            </div>
            <p class="mt-2 text-muted">Memuat data kegiatan...</p>
        </div>

        <div v-else class="accordion" id="accordionActivity">
            <div v-for="(activity, indexAct) in activities" :key="activity.id" class="accordion-item mb-3 border-0 shadow-sm rounded">
                <h2 class="accordion-header" :id="'headingAct' + activity.id">
                    <button class="accordion-button fw-bold collapsed rounded" type="button" data-bs-toggle="collapse"
                        :data-bs-target="'#collapseAct' + activity.id" aria-expanded="false"
                        :aria-controls="'collapseAct' + activity.id">
                        <i class="bi bi-folder-fill text-warning me-2"></i> 
                        {{ activity.code }} - {{ activity.name }}
                    </button>
                </h2>
                <div :id="'collapseAct' + activity.id" class="accordion-collapse collapse"
                    :aria-labelledby="'headingAct' + activity.id" data-bs-parent="#accordionActivity">
                    <div class="accordion-body bg-light p-3">
                        
                        <!-- Sub Activities Accordion -->
                        <div v-if="activity.sub_activities && activity.sub_activities.length > 0" class="accordion" :id="'accordionSub' + activity.id">
                            <div v-for="(sub, indexSub) in activity.sub_activities" :key="sub.id" class="accordion-item mb-2 border-0 shadow-sm rounded">
                                <h2 class="accordion-header" :id="'headingSub' + sub.id">
                                    <button class="accordion-button fw-bold collapsed rounded bg-white" type="button" data-bs-toggle="collapse"
                                        :data-bs-target="'#collapseSub' + sub.id" aria-expanded="false"
                                        :aria-controls="'collapseSub' + sub.id">
                                        <i class="bi bi-folder2-open text-info me-2"></i>
                                        {{ sub.code }} - {{ sub.name }}
                                    </button>
                                </h2>
                                <div :id="'collapseSub' + sub.id" class="accordion-collapse collapse"
                                    :aria-labelledby="'headingSub' + sub.id" :data-bs-parent="'#accordionSub' + activity.id">
                                    <div class="accordion-body bg-white p-3 border-top">
                                        
                                        <!-- Expenditure Accounts List -->
                                        <div v-if="sub.expenditure_accounts && sub.expenditure_accounts.length > 0">
                                            <ul class="list-group list-group-flush">
                                                <li v-for="acc in sub.expenditure_accounts" :key="acc.id" class="list-group-item d-flex justify-content-between align-items-center px-0">
                                                    <div>
                                                        <i class="bi bi-file-earmark-text text-secondary me-2"></i>
                                                        <span class="fw-medium">{{ acc.code }}</span> - {{ acc.description }}
                                                    </div>
                                                    <div>
                                                        <span class="badge bg-success me-3">{{ acc.budget_ceiling?.toLocaleString('id-ID', { style: 'currency', currency: 'IDR' }) }}</span>
                                                        <router-link :to="{ name: 'reports.rekening', params: { id: acc.id } }" class="btn btn-sm btn-outline-primary rounded-pill px-3 hover-lift">
                                                            <i class="bi bi-box-arrow-up-right me-1"></i> Detail
                                                        </router-link>
                                                    </div>
                                                </li>
                                            </ul>
                                        </div>
                                        <div v-else class="text-muted small py-2">
                                            <i class="bi bi-info-circle me-1"></i> Belum ada Rekening Belanja
                                        </div>
                                        
                                    </div>
                                </div>
                            </div>
                        </div>
                        <div v-else class="text-muted small py-2">
                            <i class="bi bi-info-circle me-1"></i> Belum ada Sub Kegiatan
                        </div>

                    </div>
                </div>
            </div>
            
            <div v-if="activities.length === 0" class="text-center text-muted py-5">
                <i class="bi bi-folder-x fs-1"></i>
                <p class="mt-3">Tidak ada data kegiatan.</p>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import api from '@/plugins/axios'

const activities = ref([])
const isLoading = ref(true)

onMounted(async () => {
    try {
        const response = await api.get('/reports/activities-tree');
        activities.value = response.data;
    } catch (error) {
        console.error('API error:', error);
    } finally {
        isLoading.value = false;
    }
})
</script>

<style scoped>
.hover-lift {
    transition: transform 0.2s ease, box-shadow 0.2s ease;
}
.hover-lift:hover {
    transform: translateY(-2px);
    box-shadow: 0 4px 8px rgba(0,0,0,0.1);
}
.accordion-button:not(.collapsed) {
    color: var(--bs-primary);
    background-color: #f8f9fa;
    box-shadow: inset 0 -1px 0 rgba(0,0,0,.125);
}
</style>