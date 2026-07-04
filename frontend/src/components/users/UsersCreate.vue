<template>
    <div class="row justify-content-center form-container">
        <div class="col-md-10 col-lg-8">
            <div class="card shadow-lg rounded-4 border-0 overflow-hidden">
                <div class="bg-info" style="height: 6px;"></div>
                
                <div class="card-body p-4 p-md-5">
                    <div class="d-flex align-items-center mb-4">
                        <router-link :to="{ name: 'users.list' }" class="btn btn-light rounded-circle me-3 d-flex align-items-center justify-content-center hover-lift" style="width: 45px; height: 45px;" title="Kembali">
                            <i class="bi bi-arrow-left fs-5 text-dark"></i>
                        </router-link>
                        <div>
                            <h4 class="card-title text-dark fw-bold mb-1">Tambah Pengguna</h4>
                            <p class="text-muted small mb-0">Daftarkan akun pengguna baru ke dalam sistem</p>
                        </div>
                    </div>

                    <form @submit.prevent="createUser" class="mt-4">
                        <div class="form-floating mb-4">
                            <input v-model="user.full_name" type="text" class="form-control custom-input" id="floatingFullName"
                                placeholder="Full Name" required>
                            <label for="floatingFullName" class="text-muted"><i class="bi bi-person me-1"></i> Nama Lengkap</label>
                        </div>
                        
                        <div class="form-floating mb-4">
                            <input v-model="user.email" type="email" class="form-control custom-input" id="floatingEmail" placeholder="Email"
                                pattern="^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$"
                                title="Masukkan format email yang benar" required>
                            <label for="floatingEmail" class="text-muted"><i class="bi bi-envelope me-1"></i> Email</label>
                        </div>
                        
                        <div class="form-floating mb-4">
                            <input v-model="user.nip" type="text" class="form-control custom-input" id="floatingNIP" placeholder="NIP">
                            <label for="floatingNIP" class="text-muted"><i class="bi bi-credit-card-2-front me-1"></i> NIP</label>
                        </div>
                        
                        <div class="form-floating mb-4">
                            <select v-model.number="user.position_id" name="position" class="form-select custom-input" id="floatingPosition" required>
                                <option value="" selected disabled>-- Pilih Jabatan --</option>
                                <option v-for="position in positions" :key="position.id" :value="position.id">
                                    {{ position.name }}
                                </option>
                            </select>
                            <label for="floatingPosition" class="text-muted"><i class="bi bi-briefcase me-1"></i> Jabatan Kedinasan</label>
                        </div>
                        
                        <div class="form-floating mb-5">
                            <select v-model.number="user.leader_id" name="leader" class="form-select custom-input" id="floatingLeader" required>
                                <option value="" selected disabled>-- Pilih Kegiatan --</option>
                                <option v-for="leader in leaders" :key="leader.id" :value="leader.id">
                                    {{ leader.name }}
                                </option>
                            </select>
                            <label for="floatingLeader" class="text-muted"><i class="bi bi-diagram-2 me-1"></i> Jabatan Kegiatan</label>
                        </div>

                        <div class="d-flex justify-content-end gap-3 mt-4 pt-3 border-top">
                            <router-link :to="{ name: 'users.list' }" class="btn btn-light px-4 py-2 rounded-pill fw-medium hover-lift">
                                Batal
                            </router-link>
                            <button type="submit" class="btn btn-info text-white px-4 py-2 rounded-pill fw-bold d-flex align-items-center gap-2 hover-lift shadow-sm">
                                <i class="bi bi-person-plus-fill fs-5"></i>
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

const user = reactive({
    full_name: '',
    email: '',
    nip: '',
    position_id: '',
    leader_id: ''
});

const positions = ref([]);
const leaders = ref([]);

const fetchPositions = async () => {
    try {
        const response = await api.get('/positions');
        positions.value = response.data;
    } catch (error) {
        console.error('Error fetching positions:', error);
    }
};

const fetchLeaders = async () => {
    try {
        const response = await api.get('/leaders');
        leaders.value = response.data;
    } catch (error) {
        console.error('Error fetching leaders:', error);
    }
};

onMounted(() => {
    // Jalankan secara paralel untuk performa terbaik
    Promise.all([fetchPositions(), fetchLeaders()]);
});

const createUser = async () => {
    try {
        await api.post('/users', user);
        router.push({ name: 'users.list' });
    } catch (error) {
        console.error('Error creating user:', error);
    }
};
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