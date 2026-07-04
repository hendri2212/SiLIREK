<template>
    <div class="row justify-content-center form-container">
        <div class="col-md-10 col-lg-8">
            <div class="card shadow-lg rounded-4 border-0 overflow-hidden">
                <div class="bg-info" style="height: 6px;"></div>
                
                <div class="card-body p-4 p-md-5">
                    <div class="d-flex align-items-center mb-4">
                        <router-link :to="{ name: 'organization.list' }" class="btn btn-light rounded-circle me-3 d-flex align-items-center justify-content-center hover-lift" style="width: 45px; height: 45px;" title="Kembali">
                            <i class="bi bi-arrow-left fs-5 text-dark"></i>
                        </router-link>
                        <div>
                            <h4 class="card-title text-dark fw-bold mb-1">Tambah Organisasi</h4>
                            <p class="text-muted small mb-0">Masukkan detail organisasi baru ke dalam sistem</p>
                        </div>
                    </div>

                    <form @submit.prevent="createOrganization" class="mt-4">
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
                                <option v-for="org in organizations" :key="org.id" :value="org.id">
                                    {{ org.number }} - {{ org.name }}
                                </option>
                            </select>
                            <label for="parent_id" class="text-muted"><i class="bi bi-diagram-3 me-1"></i> Induk Organisasi (Opsional)</label>
                        </div>

                        <div class="d-flex justify-content-end gap-3 mt-4 pt-3 border-top">
                            <router-link :to="{ name: 'organization.list' }" class="btn btn-light px-4 py-2 rounded-pill fw-medium hover-lift">
                                Batal
                            </router-link>
                            <button type="submit" class="btn btn-info text-white px-4 py-2 rounded-pill fw-bold d-flex align-items-center gap-2 hover-lift shadow-sm">
                                <i class="bi bi-cloud-arrow-up-fill fs-5"></i>
                                <span>Simpan Data</span>
                            </button>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import api from '@/plugins/axios'

const router = useRouter()
const organizations = ref([])

const form = ref({
    name: '',
    number: '',
    parent_id: null
})

const fetchOrganizations = async () => {
    try {
        const response = await api.get('/organization')
        organizations.value = response.data
    } catch (error) {
        console.error('Error fetching organizations:', error)
    }
}

const createOrganization = async () => {
    try {
        await api.post('/organization', form.value)
        alert('Organisasi berhasil ditambahkan')
        router.push({ name: 'organization.list' })
    } catch (error) {
        console.error('Error creating organization:', error)
        if (error.response && error.response.data && error.response.data.error) {
            alert(error.response.data.error)
        } else {
            alert('Terjadi kesalahan saat menyimpan organisasi')
        }
    }
}

onMounted(() => {
    fetchOrganizations()
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
    border-color: #0dcaf0;
    box-shadow: 0 0 0 0.25rem rgba(13, 202, 240, 0.15);
    background-color: #ffffff;
}

.form-floating > label {
    padding-left: 1.25rem;
}

.form-floating > .form-control:focus ~ label,
.form-floating > .form-control:not(:placeholder-shown) ~ label,
.form-floating > .form-select ~ label {
    color: #0dcaf0;
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

.btn-info.hover-lift:hover {
    box-shadow: 0 6px 12px rgba(13, 202, 240, 0.3) !important;
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
