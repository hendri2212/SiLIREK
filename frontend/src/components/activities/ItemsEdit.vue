<template>
    <div class="row justify-content-center form-container">
        <div class="col-md-10 col-lg-8">
            <div class="card shadow-lg rounded-4 border-0 overflow-hidden">
                <div class="bg-warning" style="height: 6px;"></div>
                
                <div class="card-body p-4 p-md-5">
                    <div class="d-flex align-items-center mb-4">
                        <router-link :to="{ name: 'items.list', params: { activityId: route.params.activityId, subActivityId: route.params.subActivityId, accountId: route.params.accountId } }" class="btn btn-light rounded-circle me-3 d-flex align-items-center justify-content-center hover-lift" style="width: 45px; height: 45px;" title="Kembali">
                            <i class="bi bi-arrow-left fs-5 text-dark"></i>
                        </router-link>
                        <div>
                            <h4 class="card-title text-dark fw-bold mb-1">Edit Item Belanja</h4>
                            <p class="text-muted small mb-0">Perbarui detail pengeluaran ini</p>
                        </div>
                    </div>

                    <div v-if="isLoading" class="text-center my-5">
                        <div class="spinner-border text-warning" role="status">
                            <span class="visually-hidden">Loading...</span>
                        </div>
                    </div>

                    <form v-else @submit.prevent="updateItem" class="mt-4">
                        <div class="form-floating mb-4">
                            <input type="text" class="form-control custom-input" id="code" v-model="item.code"
                                placeholder="Kode Item" required>
                            <label for="code" class="text-muted"><i class="bi bi-hash me-1"></i> Kode Item</label>
                        </div>
                        
                        <div class="form-floating mb-4">
                            <input type="date" class="form-control custom-input" id="date" v-model="item.date"
                                placeholder="Tanggal" required>
                            <label for="date" class="text-muted"><i class="bi bi-calendar-date me-1"></i> Tanggal</label>
                        </div>
                        
                        <div class="form-floating mb-4">
                            <input type="text" class="form-control custom-input" id="description" v-model="item.description"
                                placeholder="Uraian Belanja" required>
                            <label for="description" class="text-muted"><i class="bi bi-card-text me-1"></i> Uraian / Keterangan Belanja</label>
                        </div>

                        <div class="form-floating mb-5">
                            <input type="text" inputmode="numeric" class="form-control custom-input" id="credit" v-model="formattedCredit"
                                @focus="isCreditFocused = true" @blur="isCreditFocused = false"
                                placeholder="Nominal Kredit (Rp)" required>
                            <label for="credit" class="text-muted"><i class="bi bi-cash-coin me-1"></i> Nominal Kredit (Rp)</label>
                        </div>
                        
                        <div class="d-flex justify-content-end gap-3 mt-4 pt-3 border-top">
                            <router-link :to="{ name: 'items.list', params: { activityId: route.params.activityId, subActivityId: route.params.subActivityId, accountId: route.params.accountId } }"
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
import { ref, reactive, computed, onBeforeMount } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import api from '@/plugins/axios'

const router = useRouter()
const route = useRoute()

const isLoading = ref(true)

const item = reactive({
    code: '',
    date: '',
    description: '',
    credit: 0,
    expenditure_account_id: Number(route.params.accountId)
})

const isCreditFocused = ref(false)
const formattedCredit = computed({
    get() {
        if (isCreditFocused.value) return item.credit || '';
        if (!item.credit) return '';
        return new Intl.NumberFormat('id-ID').format(item.credit);
    },
    set(val) {
        const parsed = parseInt(String(val).replace(/\D/g, ''), 10);
        item.credit = isNaN(parsed) ? 0 : parsed;
    }
})

const fetchItem = async () => {
    try {
        const response = await api.get(`/item/${route.params.itemId}`)
        item.code = response.data.code
        // Extract YYYY-MM-DD for date input
        item.date = response.data.date ? response.data.date.split('T')[0] : ''
        item.description = response.data.description
        item.credit = response.data.credit
        item.expenditure_account_id = response.data.expenditure_account_id
    } catch (error) {
        console.error('Error fetching item:', error)
        alert('Gagal mengambil data item belanja.')
        router.push({ name: 'items.list', params: { activityId: route.params.activityId, subActivityId: route.params.subActivityId, accountId: route.params.accountId } })
    } finally {
        isLoading.value = false
    }
}

const updateItem = async () => {
    try {
        let finalDate = item.date;
        if (finalDate && !finalDate.includes('T')) {
            finalDate = finalDate + 'T00:00:00Z';
        }
        const payload = { 
            ...item, 
            date: finalDate,
            credit: parseFloat(item.credit) || 0,
            expenditure_account_id: parseInt(route.params.accountId)
        }
        await api.put(`/item/${route.params.itemId}`, payload)
        alert('Item Belanja berhasil diupdate')
        router.push({ name: 'items.list', params: { activityId: route.params.activityId, subActivityId: route.params.subActivityId, accountId: route.params.accountId } })
    } catch (error) {
        console.error('Error updating item:', error)
        if (error.response && error.response.data && error.response.data.error) {
            alert(error.response.data.error)
        } else {
            alert('Gagal mengupdate item belanja.')
        }
    }
}

onBeforeMount(() => {
    fetchItem()
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
