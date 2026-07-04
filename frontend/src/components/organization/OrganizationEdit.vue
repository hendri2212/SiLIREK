<template>
    <div class="row justify-content-center form-container">
        <div class="col-md-10 col-lg-8">
            <div class="card shadow-lg rounded-4 border-0 overflow-hidden">
                <div class="bg-warning" style="height: 6px;"></div>
                
                <div class="card-body p-4 p-md-5">
                    <div class="d-flex align-items-center mb-4">
                        <router-link :to="{ name: 'organization.list' }" class="btn btn-light rounded-circle me-3 d-flex align-items-center justify-content-center hover-lift" style="width: 45px; height: 45px;" title="Kembali">
                            <i class="bi bi-arrow-left fs-5 text-dark"></i>
                        </router-link>
                        <div>
                            <h4 class="card-title text-dark fw-bold mb-1">Edit Organisasi</h4>
                            <p class="text-muted small mb-0">Perbarui detail data organisasi ini</p>
                        </div>
                    </div>

                    <div v-if="isLoading" class="text-center my-5">
                        <div class="spinner-border text-warning" role="status">
                            <span class="visually-hidden">Loading...</span>
                        </div>
                    </div>

                    <form v-else @submit.prevent="updateOrganization" class="mt-4">
                        <div class="form-floating mb-4">
                            <input type="text" v-model="form.name" class="form-control custom-input" id="name" required
                                placeholder="Nama Organisasi">
                            <label for="name" class="text-muted"><i class="bi bi-building me-1"></i> Nama Organisasi</label>
                        </div>

                        <div class="form-floating mb-4">
                            <input type="text" v-model="form.number" class="form-control custom-input" id="number" required
                                placeholder="Nomor Organisasi">
                            <label for="number" class="text-muted"><i class="bi bi-123 me-1"></i> Nomor Organisasi</label>
                        </div>

                        <div class="form-floating mb-5">
                            <select v-model.number="form.parent_id" class="form-select custom-input" id="parent_id">
                                <option :value="null">-- Tidak Ada (Sebagai Induk Utama) --</option>
                                <option v-for="org in availableParents" :key="org.id" :value="org.id">
                                    {{ org.number }} - {{ org.name }}
                                </option>
                            </select>
                            <label for="parent_id" class="text-muted"><i class="bi bi-diagram-3 me-1"></i> Induk Organisasi (Opsional)</label>
                            
                            <div class="form-text text-danger mt-2" v-if="form.parent_id === orgId">
                                <i class="bi bi-exclamation-triangle-fill"></i> Organisasi tidak bisa menjadi induk bagi dirinya sendiri.
                            </div>
                        </div>

                        <div class="d-flex justify-content-end gap-3 mt-4 pt-3 border-top">
                            <router-link :to="{ name: 'organization.list' }" class="btn btn-light px-4 py-2 rounded-pill fw-medium hover-lift">
                                Batal
                            </router-link>
                            <button type="submit" class="btn btn-warning text-dark px-4 py-2 rounded-pill fw-bold d-flex align-items-center gap-2 hover-lift shadow-sm" :disabled="form.parent_id === orgId">
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
import { ref, onMounted, computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import api from '@/plugins/axios'

const router = useRouter()
const route = useRoute()
const orgId = parseInt(route.params.id)

const organizations = ref([])
const isLoading = ref(true)

const form = ref({
    name: '',
    number: '',
    parent_id: null
})

const availableParents = computed(() => {
    return organizations.value.filter(org => org.id !== orgId)
})

const fetchData = async () => {
    try {
        const orgsResponse = await api.get('/organization')
        organizations.value = orgsResponse.data

        const response = await api.get(`/organization/${orgId}`)
        const data = response.data
        
        form.value.name = data.name
        form.value.number = data.number
        form.value.parent_id = data.parent_id

    } catch (error) {
        console.error('Error fetching data:', error)
        alert('Gagal mengambil data')
        router.push({ name: 'organization.list' })
    } finally {
        isLoading.value = false
    }
}

const updateOrganization = async () => {
    if (form.value.parent_id === orgId) {
        alert('Organisasi tidak dapat menjadi induk bagi dirinya sendiri.')
        return
    }

    try {
        await api.put(`/organization/${orgId}`, form.value)
        alert('Organisasi berhasil diperbarui')
        router.push({ name: 'organization.list' })
    } catch (error) {
        console.error('Error updating organization:', error)
        if (error.response && error.response.data && error.response.data.error) {
            alert(error.response.data.error)
        } else {
            alert('Terjadi kesalahan saat memperbarui organisasi')
        }
    }
}

onMounted(() => {
    fetchData()
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
