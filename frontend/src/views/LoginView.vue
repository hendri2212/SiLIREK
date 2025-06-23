<template>
    <div class="pricing-header p-3 pb-md-4 mx-auto text-center">
        <h1 class="display-4 fw-normal fw-bold text-info">SiLIREK</h1>
        <p class="fs-5 text-body-secondary">Sistem Pengendalian Realisasi Keuangan</p>

        <form @submit.prevent="handleLogin">
            <div class="form-floating mb-3">
                <input type="text" class="form-control" id="email" placeholder="Email" v-model="email" required>
                <label for="email">Email</label>
            </div>
            <div class="form-floating mb-3">
                <input type="password" class="form-control" id="password" placeholder="Password" v-model="password" required>
                <label for="password">Password</label>
            </div>
            <button class="w-100 btn btn-lg btn-info text-white" type="submit">Sign in</button>
        </form>
    </div>
</template>

<script setup>
import { ref } from 'vue';
import api from '@/plugins/axios';
import { useAuthStore } from '@/stores/auth';
import { useRouter } from 'vue-router';

const email = ref('');
const password = ref('');
const auth = useAuthStore();
const router = useRouter();

const handleLogin = async () => {
    try {
        const response = await api.post('/login', {
            email: email.value,
            password: password.value,
        });
        const token = response.data.token;
        auth.login(token);
        await auth.fetchUser();

        alert('Login berhasil!');
        router.replace({ name: 'dashboard' })
    } catch (error) {
        console.error('Login error:', error);
        alert('Login gagal, periksa email atau password.');
    }
};
</script>