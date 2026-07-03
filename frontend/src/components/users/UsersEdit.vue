<template>
    <div class="row justify-content-center form-container">
        <div class="col-md-8 col-lg-6">
            <div class="card shadow-lg rounded-4 border-0 overflow-hidden">
                <div class="bg-warning" style="height: 6px;"></div>
                
                <div class="card-body p-4 p-md-5">
                    <div class="d-flex align-items-center mb-4">
                        <router-link :to="{ name: 'users.list' }" class="btn btn-light rounded-circle me-3 d-flex align-items-center justify-content-center hover-lift" style="width: 45px; height: 45px;" title="Kembali">
                            <i class="bi bi-arrow-left fs-5 text-dark"></i>
                        </router-link>
                        <div>
                            <h4 class="card-title text-dark fw-bold mb-1">Edit Pengguna</h4>
                            <p class="text-muted small mb-0">Perbarui posisi dan jabatan pengguna ini</p>
                        </div>
                    </div>

                    <form @submit.prevent="updateUser" class="mt-4">
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
import { reactive, onBeforeMount, ref } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import api from '@/plugins/axios';

const router = useRouter();
const route = useRoute();

const user = reactive({
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

const getUser = async () => {
    try {
        const response = await api.get(`/users/${route.params.id}`);
        Object.assign(user, response.data);
    } catch (error) {
        console.error('Error fetching user data:', error);
    }
};

const updateUser = async () => {
    try {
        const formData = new FormData();
        formData.append('position_id', user.position_id);
        formData.append('leader_id', user.leader_id);

        await api.put(`/users/${route.params.id}`, formData, {
            headers: {
                'Content-Type': 'multipart/form-data'
            }
        });
        router.push({ name: 'users.list' });
    } catch (error) {
        console.error('Error updating user:', error);
    }
};

onBeforeMount(() => {
    getUser();
    fetchPositions();
    fetchLeaders();
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