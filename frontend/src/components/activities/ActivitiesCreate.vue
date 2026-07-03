<template>
    <div class="row justify-content-center form-container">
        <div class="col-md-8 col-lg-6">
            <div class="card shadow-lg rounded-4 border-0 overflow-hidden">
                <!-- Decorative top border -->
                <div class="bg-info" style="height: 6px;"></div>
                
                <div class="card-body p-4 p-md-5">
                    <div class="d-flex align-items-center mb-4">
                        <router-link :to="{ name: 'activities.list' }" class="btn btn-light rounded-circle me-3 d-flex align-items-center justify-content-center hover-lift" style="width: 45px; height: 45px;" title="Kembali">
                            <i class="bi bi-arrow-left fs-5 text-dark"></i>
                        </router-link>
                        <div>
                            <h4 class="card-title text-dark fw-bold mb-1">Tambah Kegiatan</h4>
                            <p class="text-muted small mb-0">Masukkan detail kegiatan baru ke dalam sistem</p>
                        </div>
                    </div>

                    <form @submit.prevent="createActivity" class="mt-4">
                        <!-- Code Input -->
                        <div class="form-floating mb-4">
                            <input v-model="activity.code" type="text" class="form-control custom-input" id="floatingActivity"
                                placeholder="Kode Kegiatan" required>
                            <label for="floatingActivity" class="text-muted"><i class="bi bi-hash me-1"></i> Kode Kegiatan</label>
                        </div>
                        
                        <!-- Name Input -->
                        <div class="form-floating mb-5">
                            <input v-model="activity.name" type="text" class="form-control custom-input" id="floatingNameActivity" 
                                placeholder="Nama Kegiatan" required>
                            <label for="floatingNameActivity" class="text-muted"><i class="bi bi-card-text me-1"></i> Nama Kegiatan</label>
                        </div>
                        
                        <!-- Actions -->
                        <div class="form-floating mb-5">
                            <select v-model.number="activity.program_id" class="form-select custom-input" id="program_id" required>
                                <option value="" disabled selected>-- Pilih Program --</option>
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
import { ref, reactive, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import api from '@/plugins/axios';

const router = useRouter();
const programs = ref([]);

const activity = reactive({
    code: '',
    name: '',
    program_id: ''
});

const fetchPrograms = async () => {
    try {
        const response = await api.get('/program');
        programs.value = response.data;
    } catch (error) {
        console.error('Error fetching programs:', error);
    }
};

const createActivity = async () => {
    try {
        await api.post('/activity', activity);
        router.push({ name: 'activities.list' });
    } catch (error) {
        console.error('Error creating activities:', error);
        if (error.response && error.response.data && error.response.data.error) {
            alert(error.response.data.error);
        } else {
            alert('Terjadi kesalahan saat menambahkan kegiatan.');
        }
    }
};

onMounted(() => {
    fetchPrograms();
});
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
.form-floating > .form-control:not(:placeholder-shown) ~ label {
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