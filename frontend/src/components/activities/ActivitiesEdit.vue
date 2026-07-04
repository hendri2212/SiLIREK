<template>
    <div class="row justify-content-center form-container">
        <div class="col-md-10 col-lg-8">
            <div class="card shadow-lg rounded-4 border-0 overflow-hidden">
                <div class="bg-warning" style="height: 6px;"></div>
                
                <div class="card-body p-4 p-md-5">
                    <div class="d-flex align-items-center mb-4">
                        <router-link :to="{ name: 'activities.list' }" class="btn btn-light rounded-circle me-3 d-flex align-items-center justify-content-center hover-lift" style="width: 45px; height: 45px;" title="Kembali">
                            <i class="bi bi-arrow-left fs-5 text-dark"></i>
                        </router-link>
                        <div>
                            <h4 class="card-title text-dark fw-bold mb-1">Edit Kegiatan</h4>
                            <p class="text-muted small mb-0">Perbarui detail data kegiatan ini</p>
                        </div>
                    </div>

                    <div v-if="isLoading" class="text-center my-5">
                        <div class="spinner-border text-warning" role="status">
                            <span class="visually-hidden">Loading...</span>
                        </div>
                    </div>

                    <form v-else @submit.prevent="updateActivity" class="mt-4">
                        <div class="form-floating mb-4">
                            <input v-model="activity.code" type="text" class="form-control custom-input" id="floatingActivity"
                                placeholder="Kode Kegiatan" required>
                            <label for="floatingActivity" class="text-muted"><i class="bi bi-hash me-1"></i> Kode Kegiatan</label>
                        </div>
                        
                        <div class="form-floating mb-5">
                            <input v-model="activity.name" type="text" class="form-control custom-input" id="floatingNameActivity" 
                                placeholder="Nama Kegiatan" required>
                            <label for="floatingNameActivity" class="text-muted"><i class="bi bi-card-text me-1"></i> Nama Kegiatan</label>
                        </div>
                        
                        <div class="form-floating mb-5">
                            <select v-model.number="activity.program_id" class="form-select custom-input" id="program_id" required>
                                <option value="" disabled>-- Pilih Program --</option>
                                <option v-for="program in programs" :key="program.id" :value="program.id">
                                    {{ program.code }} - {{ program.name }}
                                </option>
                            </select>
                            <label for="program_id" class="text-muted"><i class="bi bi-folder-fill me-1"></i> Program Terkait</label>
                        </div>
                        
                        <div class="d-flex justify-content-end gap-3 mt-4 pt-3 border-top">
                            <router-link :to="{ name: 'activities.list' }" class="btn btn-light px-4 py-2 rounded-pill fw-medium hover-lift">
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
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import api from '@/plugins/axios'

const router = useRouter()
const route = useRoute()
const activityId = route.params.id

const isLoading = ref(true)
const programs = ref([])
const activity = ref({
    code: '',
    name: '',
    program_id: ''
})

const fetchData = async () => {
    try {
        const progResponse = await api.get('/program')
        programs.value = progResponse.data

        const response = await api.get(`/activity/${activityId}`)
        activity.value.code = response.data.code
        activity.value.name = response.data.name
        activity.value.program_id = response.data.program_id
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
.form-floating > .form-control:not(:placeholder-shown) ~ label {
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