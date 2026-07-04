<template>
    <div class="row justify-content-center form-container">
        <div class="col-md-10 col-lg-8">
            <div class="card shadow-lg rounded-4 border-0 overflow-hidden">
                <div class="bg-warning" style="height: 6px;"></div>
                
                <div class="card-body p-4 p-md-5">
                    <div class="d-flex align-items-center mb-4">
                        <router-link :to="{ name: 'expenditureaccounts.list', params: { activityId: route.params.activityId, subActivityId: route.params.subActivityId } }" class="btn btn-light rounded-circle me-3 d-flex align-items-center justify-content-center hover-lift" style="width: 45px; height: 45px;" title="Kembali">
                            <i class="bi bi-arrow-left fs-5 text-dark"></i>
                        </router-link>
                        <div>
                            <h4 class="card-title text-dark fw-bold mb-1">Edit Rekening Belanja</h4>
                            <p class="text-muted small mb-0">Perbarui detail anggaran ini</p>
                        </div>
                    </div>

                    <div v-if="isLoading" class="text-center my-5">
                        <div class="spinner-border text-warning" role="status">
                            <span class="visually-hidden">Loading...</span>
                        </div>
                    </div>

                    <form v-else @submit.prevent="updateAccount" class="mt-4">
                        <div class="form-floating mb-4">
                            <input type="text" class="form-control custom-input" id="code" v-model="account.code"
                                placeholder="Kode Rekening" required>
                            <label for="code" class="text-muted"><i class="bi bi-hash me-1"></i> Kode Rekening</label>
                        </div>
                        
                        <div class="form-floating mb-4">
                            <input type="text" class="form-control custom-input" id="description" v-model="account.description"
                                placeholder="Uraian" required>
                            <label for="description" class="text-muted"><i class="bi bi-card-text me-1"></i> Uraian / Deskripsi</label>
                        </div>

                        <div class="form-floating mb-5">
                            <input type="text" inputmode="numeric" class="form-control custom-input" id="budget" v-model="formattedPagu"
                                @focus="isPaguFocused = true" @blur="isPaguFocused = false"
                                placeholder="Pagu Anggaran" required>
                            <label for="budget" class="text-muted"><i class="bi bi-currency-dollar me-1"></i> Pagu Anggaran (Rp)</label>
                        </div>
                        
                        <div class="d-flex justify-content-end gap-3 mt-4 pt-3 border-top">
                            <router-link :to="{ name: 'expenditureaccounts.list', params: { activityId: route.params.activityId, subActivityId: route.params.subActivityId } }"
                                class="btn btn-light px-4 py-2 rounded-pill fw-medium hover-lift">
                                Batal
                            </router-link>
                            <button type="submit" class="btn btn-warning text-dark px-4 py-2 rounded-pill fw-bold d-flex align-items-center gap-2 hover-lift shadow-sm">
                                <i class="bi bi-cloud-arrow-up-fill fs-5"></i>
                                <span>Simpan Perubahan</span>
                            </button>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import api from '@/plugins/axios'

const router = useRouter()
const route = useRoute()

const isLoading = ref(true)

const account = reactive({
    code: '',
    description: '',
    budget_ceiling: 0,
    sub_activity_id: Number(route.params.subActivityId)
})

const isPaguFocused = ref(false)
const formattedPagu = computed({
    get() {
        if (isPaguFocused.value) return account.budget_ceiling || '';
        if (!account.budget_ceiling) return '';
        return new Intl.NumberFormat('id-ID').format(account.budget_ceiling);
    },
    set(val) {
        const parsed = parseInt(String(val).replace(/\D/g, ''), 10);
        account.budget_ceiling = isNaN(parsed) ? 0 : parsed;
    }
})

const fetchAccount = async () => {
    try {
        const response = await api.get(`/expenditure-account/${route.params.accountId}`)
        account.code = response.data.code
        account.description = response.data.description
        account.budget_ceiling = response.data.budget_ceiling
        account.sub_activity_id = response.data.sub_activity_id
    } catch (error) {
        console.error('Error fetching expenditure account:', error)
        alert('Gagal mengambil data rekening belanja.')
        router.push({ name: 'expenditureaccounts.list', params: { activityId: route.params.activityId, subActivityId: route.params.subActivityId } })
    } finally {
        isLoading.value = false
    }
}

const updateAccount = async () => {
    try {
        await api.put(`/expenditure-account/${route.params.accountId}`, account)
        alert('Rekening Belanja berhasil diupdate')
        router.push({ name: 'expenditureaccounts.list', params: { activityId: route.params.activityId, subActivityId: route.params.subActivityId } })
    } catch (error) {
        console.error('Error updating expenditure account:', error)
        if (error.response && error.response.data && error.response.data.error) {
            alert(error.response.data.error)
        } else {
            alert('Gagal mengupdate rekening belanja.')
        }
    }
}

onMounted(() => {
    fetchAccount()
})
</script>

<style scoped>
.form-container {
    animation: slideUpFade 0.5s ease-out;
}

.custom-input {
    border: 2px solid #e9ecef;
    border-radius: 0.75rem;
    transition: all 0.3s ease;
    background-color: #f8f9fa;
}

.custom-input:focus {
    border-color: #ffc107;
    box-shadow: 0 0 0 0.25rem rgba(255, 193, 7, 0.15);
    background-color: #ffffff;
}

.form-floating > label {
    padding-left: 1.25rem;
}

.form-floating > .form-control:focus ~ label,
.form-floating > .form-control:not(:placeholder-shown) ~ label,
.form-floating > .form-select ~ label {
    color: #ffc107;
    font-weight: 600;
    transform: scale(0.85) translateY(-0.75rem) translateX(0.15rem);
}

.hover-lift {
    transition: transform 0.2s ease, box-shadow 0.2s ease;
}

.hover-lift:hover {
    transform: translateY(-2px);
    box-shadow: 0 6px 12px rgba(0, 0, 0, 0.1) !important;
}

.btn-warning.hover-lift:hover {
    box-shadow: 0 6px 12px rgba(255, 193, 7, 0.3) !important;
}

@keyframes slideUpFade {
    from {
        opacity: 0;
        transform: translateY(20px);
    }
    to {
        opacity: 1;
        transform: translateY(0);
    }
}
</style>
