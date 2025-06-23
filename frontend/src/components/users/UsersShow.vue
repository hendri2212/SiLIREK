<template>
    <div class="card shadow-sm rounded-4 border-0">
        <div class="card-body">
            <div class="mb-4 d-flex align-items-center">
                <h4 class="card-title text-info">Profil Pengguna</h4>
            </div>
            <form @submit.prevent="profileUser">
                <div class="form-floating mb-3">
                    <input v-model="user.full_name" type="text" class="form-control" id="floatingFullName"
                        placeholder="Full Name" required>
                    <label for="floatingFullName">Full Name</label>
                </div>
                <div class="form-floating mb-3">
                    <input v-model="user.email" type="email" class="form-control" id="floatingEmail" placeholder="Email"
                        pattern="^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$"
                        title="Masukkan format email yang benar" required>

                    <label for="floatingEmail">Email</label>
                </div>
                <div class="form-floating mb-3">
                    <input v-model="user.password" type="password" class="form-control" id="floatingPassword"
                        placeholder="Password (biarkan kosong jika tidak diubah)">
                    <label for="floatingPassword">Password (opsional)</label>
                    <span class="text-body-secondary fw-light fst-italic small">* Kosongkan jika tidak ingin mengubah password</span>
                </div>
                <div class="form-floating mb-3">
                    <input v-model="user.nip" type="text" class="form-control" id="floatingNIP" placeholder="NIP">
                    <label for="floatingNIP">NIP</label>
                </div>
                <div class="mb-3">
                    <input @change="handlePhotoUpload" class="form-control" type="file" id="formFile" accept="image/*">
                    <span class="text-body-secondary fw-light fst-italic small">* Foto Profil</span>
                </div>
                <button class="w-100 btn btn-lg btn-info text-white" type="submit">Update Profile</button>
            </form>
        </div>
    </div>
</template>
<script setup>
import { reactive, onMounted } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import api from '@/plugins/axios';
import { useAuthStore } from '@/stores/auth'

const router = useRouter();
const route = useRoute();
const authStore = useAuthStore()

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
    } catch (error) {
        console.error('Error fetching user data:', error);
    }
};

const handlePhotoUpload = (e) => {
    user.photo = e.target.files[0];
};

const profileUser = async () => {
    try {
        const formData = new FormData();
        formData.append('full_name', user.full_name);
        formData.append('email', user.email);
        formData.append('nip', user.nip || '');

        if (user.password) {
            formData.append('password', user.password);
        }
        if (user.photo) {
            formData.append('photo', user.photo);
        }

        await api.put(`/users/${route.params.id}`, formData, {
            headers: { 'Content-Type': 'multipart/form-data' }
        });

        await authStore.fetchUser() // agar user.value diperbarui

        router.push({ name: 'users.list' });
    } catch (error) {
        console.error('Error updating user:', error);
    }
};

onMounted(() => {
    getUser();
});
</script>