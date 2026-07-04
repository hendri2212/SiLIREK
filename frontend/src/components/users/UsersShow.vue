<template>
    <div class="container-fluid px-0 py-4">
        <div class="row justify-content-center">
            <div class="col-lg-10 col-xl-8">
                <div class="card soft-shadow border-0 rounded-5 overflow-hidden glass-profile-card">
                    <!-- Header Banner -->
                    <div class="p-5 position-relative border-bottom border-success border-opacity-10">
                        <div class="position-absolute top-0 end-0 opacity-10 p-4 text-success" style="transform: scale(2) translate(10%, -10%);">
                            <i class="bi bi-person-circle" style="font-size: 8rem;"></i>
                        </div>
                        <h3 class="text-success fw-bold mb-0 position-relative z-1">Profil Pengguna</h3>
                        <p class="text-secondary mb-0 position-relative z-1">Perbarui informasi profil dan kredensial Anda</p>
                    </div>

                    <div class="card-body p-4 p-md-5" style="background: rgba(255, 255, 255, 0.6);">
                        <form @submit.prevent="profileUser" class="row g-4">
                            <!-- Left Column: Avatar/Profile Info -->
                            <div class="col-md-4 text-center d-flex flex-column align-items-center justify-content-center border-end-md pe-md-4 mb-4 mb-md-0">
                                <div class="position-relative mb-4">
                                    <div class="rounded-circle bg-light d-flex align-items-center justify-content-center border border-4 border-white soft-shadow" style="width: 150px; height: 150px; overflow: hidden;">
                                        <!-- Placeholder or Image -->
                                        <i v-if="!imagePreview" class="bi bi-person-fill text-secondary opacity-50" style="font-size: 5rem;"></i>
                                        <img v-else :src="imagePreview" alt="Profile" class="w-100 h-100 object-fit-cover" />
                                    </div>
                                    <!-- Edit Badge -->
                                    <label for="formFile" class="position-absolute bottom-0 end-0 bg-success text-white rounded-circle p-2 shadow-sm hover-scale" style="cursor: pointer; transform: translate(-10%, -10%);">
                                        <i class="bi bi-camera-fill"></i>
                                    </label>
                                </div>
                                <h5 class="fw-bold text-dark mb-1">{{ user.full_name || 'Nama Pengguna' }}</h5>
                                <p class="text-muted small mb-0">{{ user.email || 'email@example.com' }}</p>
                                <span class="badge bg-success bg-opacity-10 text-success border border-success border-opacity-25 mt-2" v-if="user.nip">NIP: {{ user.nip }}</span>
                                
                                <!-- Hidden File Input -->
                                <input @change="handlePhotoUpload" class="d-none" type="file" id="formFile" accept="image/*">
                            </div>

                            <!-- Right Column: Form Fields -->
                            <div class="col-md-8 ps-md-4">
                                <h6 class="fw-bold mb-4 text-success text-uppercase opacity-75 border-bottom pb-2">Informasi Dasar</h6>
                                
                                <div class="form-floating mb-3">
                                    <input v-model="user.full_name" type="text" class="form-control bg-light border-0 soft-input" id="floatingFullName" placeholder="Nama Lengkap" required>
                                    <label for="floatingFullName" class="text-muted"><i class="bi bi-person me-2"></i>Nama Lengkap</label>
                                </div>
                                
                                <div class="form-floating mb-3">
                                    <input v-model="user.email" type="email" class="form-control bg-light border-0 soft-input" id="floatingEmail" placeholder="Email"
                                        pattern="^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$"
                                        title="Masukkan format email yang benar" required>
                                    <label for="floatingEmail" class="text-muted"><i class="bi bi-envelope me-2"></i>Email</label>
                                </div>

                                <div class="form-floating mb-4">
                                    <input v-model="user.nip" type="text" class="form-control bg-light border-0 soft-input" id="floatingNIP" placeholder="NIP">
                                    <label for="floatingNIP" class="text-muted"><i class="bi bi-credit-card-2-front me-2"></i>NIP</label>
                                </div>

                                <h6 class="fw-bold mb-3 text-success text-uppercase opacity-75 border-bottom pb-2 mt-4">Keamanan</h6>

                                <div class="form-floating mb-2">
                                    <input v-model="user.password" type="password" class="form-control bg-light border-0 soft-input" id="floatingPassword"
                                        placeholder="Password Baru">
                                    <label for="floatingPassword" class="text-muted"><i class="bi bi-lock me-2"></i>Password Baru</label>
                                </div>
                                <div class="text-muted small fst-italic mb-4">
                                    <i class="bi bi-info-circle me-1"></i> Biarkan kosong jika tidak ingin mengubah kata sandi Anda.
                                </div>

                                <div class="d-flex justify-content-end mt-4 pt-3 border-top">
                                    <button type="button" class="btn btn-light me-2 px-4 rounded-pill" @click="$router.push({ name: 'dashboard' })">Batal</button>
                                    <button type="submit" class="btn btn-success px-4 fw-medium shadow-sm hover-lift rounded-pill" :disabled="isSaving">
                                        <span v-if="isSaving" class="spinner-border spinner-border-sm me-2" role="status" aria-hidden="true"></span>
                                        <i v-else class="bi bi-check2-circle me-2"></i> Simpan Perubahan
                                    </button>
                                </div>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { reactive, ref, onMounted } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import api from '@/plugins/axios';
import { useAuthStore } from '@/stores/auth'

const router = useRouter();
const route = useRoute();
const authStore = useAuthStore()

const isSaving = ref(false);
const imagePreview = ref(null);

const user = reactive({
    full_name: '',
    email: '',
    nip: '',
    password: '',
    photo: null,
});

const getUser = async () => {
    try {
        const response = await api.get(`/users/${route.params.id}`);
        Object.assign(user, response.data);
        user.password = '';
        
        // If backend returns a photo path, we could set imagePreview here
        // Assuming response.data.photo contains the URL path like '/uploads/abc.jpg'
        if (response.data.photo) {
            // Check if it's already a full URL or a local path
            if (response.data.photo.startsWith('http')) {
                imagePreview.value = response.data.photo;
            } else {
                imagePreview.value = import.meta.env.VITE_API_URL.replace('/api', '') + '/' + response.data.photo.replace(/^\//, '');
            }
        }
    } catch (error) {
        console.error('Error fetching user data:', error);
    }
};

const handlePhotoUpload = (e) => {
    const file = e.target.files[0];
    if (file) {
        user.photo = file;
        // Create preview URL
        imagePreview.value = URL.createObjectURL(file);
    }
};

const profileUser = async () => {
    isSaving.value = true;
    try {
        const formData = new FormData();
        formData.append('full_name', user.full_name);
        formData.append('email', user.email);
        formData.append('nip', user.nip || '');

        if (user.password) {
            formData.append('password', user.password);
        }
        if (user.photo && typeof user.photo !== 'string') {
            formData.append('photo', user.photo);
        }

        await api.put(`/users/${route.params.id}`, formData, {
            headers: { 'Content-Type': 'multipart/form-data' }
        });

        await authStore.fetchUser() // agar user.value diperbarui

        // Optional: show a success toast here
        
        // If user is just editing themselves from dashboard, maybe go back to dashboard. 
        // If they are superadmin editing others, go to users list.
        // For safety, we just go back to the previous page or dashboard
        if (window.history.length > 2) {
            router.back();
        } else {
            router.push({ name: 'dashboard' });
        }
    } catch (error) {
        console.error('Error updating user:', error);
    } finally {
        isSaving.value = false;
    }
};

onMounted(() => {
    getUser();
});
</script>

<style scoped>
.glass-profile-card {
    background: rgba(255, 255, 255, 0.45);
    backdrop-filter: blur(20px);
    -webkit-backdrop-filter: blur(20px);
    border: 1px solid rgba(255, 255, 255, 0.6) !important;
}
.soft-shadow {
    box-shadow: 0 10px 40px rgba(0, 0, 0, 0.05) !important;
}
.soft-input {
    background-color: #f8f9fa !important;
    border-radius: 0.75rem !important;
    transition: all 0.2s ease;
}
.soft-input:focus {
    background-color: #fff !important;
    box-shadow: 0 0 0 0.25rem rgba(25, 135, 84, 0.15) !important;
}
.hover-scale {
    transition: transform 0.2s ease, box-shadow 0.2s ease;
}
.hover-scale:hover {
    transform: translate(-10%, -10%) scale(1.1) !important;
}
.border-end-md {
    border-right: none;
}
@media (min-width: 768px) {
    .border-end-md {
        border-right: 1px dashed #e9ecef;
    }
}
.hover-lift {
    transition: transform 0.2s ease, box-shadow 0.2s ease;
}
.hover-lift:hover:not(:disabled) {
    transform: translateY(-2px);
    box-shadow: 0 .5rem 1rem rgba(25, 135, 84, 0.25)!important;
}
.form-floating > .form-control:focus ~ label::after, .form-floating > .form-control:not(:placeholder-shown) ~ label::after {
    background-color: transparent !important;
}
</style>