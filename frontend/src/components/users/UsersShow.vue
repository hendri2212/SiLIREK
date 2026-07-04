<template>
    <div class="container-fluid px-0 py-4">
        <div class="row justify-content-center">
            <div class="col-lg-10 col-xl-8">

                <!-- Page Header -->
                <div class="mb-5">
                    <h3 class="text-success fw-bold mb-1">
                        <i class="bi bi-person-circle me-2 opacity-75"></i>Profil Pengguna
                    </h3>
                    <p class="text-muted mb-0">Perbarui informasi profil dan kredensial Anda</p>
                </div>

                <form @submit.prevent="profileUser">
                    <div class="row g-5">
                        <!-- Left Column: Avatar -->
                        <div class="col-md-4 text-center d-flex flex-column align-items-center">
                            <!-- Avatar -->
                            <div class="position-relative mb-3">
                                <div class="avatar-ring rounded-circle d-flex align-items-center justify-content-center" style="width: 140px; height: 140px; overflow: hidden;">
                                    <i v-if="!imagePreview" class="bi bi-person-fill text-secondary opacity-50" style="font-size: 5rem;"></i>
                                    <img v-else :src="imagePreview" alt="Profile" class="w-100 h-100 object-fit-cover" />
                                </div>
                                <label for="formFile" class="avatar-edit-btn position-absolute bottom-0 end-0 bg-success text-white rounded-circle d-flex align-items-center justify-content-center shadow" style="width: 36px; height: 36px; cursor: pointer; transform: translate(-5%, -5%);">
                                    <i class="bi bi-camera-fill small"></i>
                                </label>
                            </div>
                            <h5 class="fw-bold text-dark mb-1">{{ user.full_name || 'Nama Pengguna' }}</h5>
                            <p class="text-muted small mb-2">{{ user.email || 'email@example.com' }}</p>
                            <span v-if="user.nip" class="badge bg-success bg-opacity-10 text-success border border-success border-opacity-25 px-3 py-1 rounded-pill">
                                NIP {{ user.nip }}
                            </span>
                            <input @change="handlePhotoUpload" class="d-none" type="file" id="formFile" accept="image/*">
                        </div>

                        <!-- Right Column: Form -->
                        <div class="col-md-8">
                            <!-- Section: Informasi Dasar -->
                            <div class="section-label mb-4">
                                <span class="text-success text-uppercase fw-bold small letter-spacing">Informasi Dasar</span>
                                <div class="section-line mt-1"></div>
                            </div>

                            <div class="form-floating mb-3">
                                <input v-model="user.full_name" type="text" class="form-control flat-input" id="floatingFullName" placeholder="Nama Lengkap" required>
                                <label for="floatingFullName" class="text-muted"><i class="bi bi-person me-2"></i>Nama Lengkap</label>
                            </div>

                            <div class="form-floating mb-3">
                                <input v-model="user.email" type="email" class="form-control flat-input" id="floatingEmail" placeholder="Email"
                                    pattern="^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$"
                                    title="Masukkan format email yang benar" required>
                                <label for="floatingEmail" class="text-muted"><i class="bi bi-envelope me-2"></i>Email</label>
                            </div>

                            <div class="form-floating mb-5">
                                <input v-model="user.nip" type="text" class="form-control flat-input" id="floatingNIP" placeholder="NIP">
                                <label for="floatingNIP" class="text-muted"><i class="bi bi-credit-card-2-front me-2"></i>NIP</label>
                            </div>

                            <!-- Section: Keamanan -->
                            <div class="section-label mb-4">
                                <span class="text-success text-uppercase fw-bold small letter-spacing">Keamanan</span>
                                <div class="section-line mt-1"></div>
                            </div>

                            <div class="form-floating mb-1">
                                <input v-model="user.password" type="password" class="form-control flat-input" id="floatingPassword" placeholder="Password Baru">
                                <label for="floatingPassword" class="text-muted"><i class="bi bi-lock me-2"></i>Password Baru</label>
                            </div>
                            <p class="text-muted small fst-italic mt-2 mb-5">
                                <i class="bi bi-info-circle me-1"></i> Biarkan kosong jika tidak ingin mengubah kata sandi.
                            </p>

                            <!-- Action Buttons -->
                            <div class="d-flex justify-content-end gap-2 pt-4 border-top border-success border-opacity-10">
                                <button type="button" class="btn btn-light px-4 rounded-pill fw-medium" @click="$router.push({ name: 'dashboard' })">
                                    Batal
                                </button>
                                <button type="submit" class="btn btn-success px-4 rounded-pill fw-medium hover-lift shadow-sm" :disabled="isSaving">
                                    <span v-if="isSaving" class="spinner-border spinner-border-sm me-2" role="status" aria-hidden="true"></span>
                                    <i v-else class="bi bi-check2-circle me-2"></i>Simpan Perubahan
                                </button>
                            </div>
                        </div>
                    </div>
                </form>

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

const getPhotoUrl = (photo) => {
    if (!photo) return '';
    if (photo.startsWith('http')) return photo;
    const baseApiUrl = api.defaults.baseURL.replace(/\/api$/, '');
    const path = photo.startsWith('uploads/') ? photo : `uploads/photos/${photo}`;
    return `${baseApiUrl}/${path}`;
};

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
        
        if (response.data.photo) {
            imagePreview.value = getPhotoUrl(response.data.photo);
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
/* Avatar */
.avatar-ring {
    background-color: rgba(255, 255, 255, 0.7);
    border: 3px solid rgba(25, 135, 84, 0.15);
    box-shadow: 0 4px 20px rgba(25, 135, 84, 0.1);
    transition: box-shadow 0.3s ease;
}
.avatar-ring:hover {
    box-shadow: 0 6px 24px rgba(25, 135, 84, 0.2);
}
.avatar-edit-btn {
    transition: transform 0.2s ease, box-shadow 0.2s ease;
}
.avatar-edit-btn:hover {
    transform: translate(-5%, -5%) scale(1.12);
    box-shadow: 0 4px 12px rgba(25, 135, 84, 0.3) !important;
}

/* Section divider label */
.section-label .section-line {
    height: 2px;
    background: linear-gradient(90deg, rgba(25, 135, 84, 0.3), transparent);
    border-radius: 2px;
}
.letter-spacing {
    letter-spacing: 0.08em;
}

/* Flat inputs - no border box, just bottom line */
.flat-input {
    background-color: rgba(255, 255, 255, 0.55) !important;
    border: none !important;
    border-bottom: 2px solid rgba(0, 0, 0, 0.08) !important;
    border-radius: 0 !important;
    transition: all 0.25s ease;
}
.flat-input:focus {
    background-color: rgba(255, 255, 255, 0.85) !important;
    border-bottom: 2px solid rgba(25, 135, 84, 0.5) !important;
    box-shadow: none !important;
}

/* Floating label fix for flat inputs */
.form-floating > .form-control:focus ~ label,
.form-floating > .form-control:not(:placeholder-shown) ~ label {
    color: #198754;
}
.form-floating > .form-control:focus ~ label::after,
.form-floating > .form-control:not(:placeholder-shown) ~ label::after {
    background-color: transparent !important;
}

/* Button */
.hover-lift {
    transition: transform 0.2s ease, box-shadow 0.2s ease;
}
.hover-lift:hover:not(:disabled) {
    transform: translateY(-2px);
    box-shadow: 0 0.5rem 1rem rgba(25, 135, 84, 0.25) !important;
}
</style>