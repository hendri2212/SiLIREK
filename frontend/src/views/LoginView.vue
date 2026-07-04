<template>
    <div class="d-flex align-items-center justify-content-center min-vh-100 position-relative">


        <div class="container position-relative z-1">
            <div class="row justify-content-center">
                <div class="col-12 col-md-10 col-lg-7 col-xl-6">
                    <div class="card border-0 rounded-5 glass-card overflow-hidden">

                        <!-- Header Content inside the glass card -->
                        <div class="card-body p-4 p-md-5">
                            <div class="text-center mb-5 mt-3">
                                <div class="icon-box mx-auto mb-3 text-primary d-flex align-items-center justify-content-center bg-primary bg-opacity-10 rounded-circle"
                                    style="width: 80px; height: 80px;">
                                    <i class="bi bi-wallet2 fs-1"></i>
                                </div>
                                <h2 class="fw-bold text-dark mb-1">SiLIREK</h2>
                                <p class="text-muted small fw-medium">Sistem Pengendalian Realisasi Keuangan</p>
                            </div>

                            <div v-if="alertMessage"
                                :class="['alert', `alert-${alertType}`, 'alert-dismissible', 'fade', 'show', 'border-0', 'shadow-sm', 'rounded-4']"
                                role="alert">
                                <i class="bi"
                                    :class="alertType === 'success' ? 'bi-check-circle-fill' : 'bi-exclamation-triangle-fill'"></i>
                                <span class="ms-2 small">{{ alertMessage }}</span>
                                <button type="button" class="btn-close" @click="alertMessage = ''"
                                    aria-label="Close"></button>
                            </div>

                            <form @submit.prevent="handleLogin">
                                <div class="form-floating mb-3">
                                    <input type="email" class="form-control bg-transparent border-bottom-only"
                                        id="email" placeholder="Email" v-model="email" required :disabled="isLoading">
                                    <label for="email" class="text-muted"><i class="bi bi-envelope me-2"></i>Alamat
                                        Email</label>
                                </div>

                                <div class="form-floating mb-5">
                                    <input type="password" class="form-control bg-transparent border-bottom-only"
                                        id="password" placeholder="Password" v-model="password" required
                                        :disabled="isLoading">
                                    <label for="password" class="text-muted"><i class="bi bi-lock me-2"></i>Kata
                                        Sandi</label>
                                </div>

                                <button
                                    class="w-100 btn btn-lg btn-primary fw-bold text-white rounded-pill hover-lift shadow-lg py-3"
                                    type="submit" :disabled="isLoading">
                                    <span v-if="isLoading" class="spinner-border spinner-border-sm me-2" role="status"
                                        aria-hidden="true"></span>
                                    <span v-else>Masuk <i class="bi bi-arrow-right ms-2"></i></span>
                                </button>
                            </form>
                        </div>

                        <!-- Footer -->
                        <div class="text-center pb-4 pt-2 opacity-75">
                            <small class="text-muted" style="font-size: 0.75rem;">
                                &copy; {{ new Date().getFullYear() }} Kec. Pulaulaut Sigam<br />
                                Kabupaten Kotabaru
                            </small>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue';
import api from '@/plugins/axios';
import { useAuthStore } from '@/stores/auth';
import { useRouter } from 'vue-router';

const email = ref('');
const password = ref('');
const alertMessage = ref('');
const alertType = ref('success');
const isLoading = ref(false);
const auth = useAuthStore();
const router = useRouter();

const handleLogin = async () => {
    isLoading.value = true;
    alertMessage.value = '';

    try {
        const response = await api.post('/login', {
            email: email.value,
            password: password.value,
        });
        const token = response.data.token;
        auth.login(token);
        await auth.fetchUser();

        alertMessage.value = 'Berhasil masuk! Mengalihkan...';
        alertType.value = 'success';

        setTimeout(() => {
            router.replace({ name: 'dashboard' });
        }, 1000);
    } catch (error) {
        console.error('Login error:', error);
        isLoading.value = false;
        alertMessage.value = 'Login gagal. Periksa kembali email atau kata sandi Anda.';
        alertType.value = 'danger';
    }
};
</script>

<style scoped>
/* Glassmorphism Card */
.glass-card {
    background: rgba(255, 255, 255, 0.75);
    backdrop-filter: blur(16px);
    -webkit-backdrop-filter: blur(16px);
    border: 1px solid rgba(255, 255, 255, 0.5) !important;
    box-shadow: 0 20px 40px rgba(0, 0, 0, 0.04) !important;
}

/* Elegant Inputs - Just bottom borders */
.border-bottom-only {
    border: none !important;
    border-bottom: 2px solid #e2e8f0 !important;
    border-radius: 0 !important;
    padding-left: 0.5rem;
}

.border-bottom-only:focus {
    box-shadow: none !important;
    border-bottom: 2px solid var(--bs-primary) !important;
    background-color: transparent !important;
}

/* Floating label offset fixing for border-bottom-only */
.form-floating>.form-control:focus~label,
.form-floating>.form-control:not(:placeholder-shown)~label {
    color: var(--bs-primary);
    transform: scale(.85) translateY(-1rem) translateX(0.15rem);
}

.form-floating>label {
    padding-left: 0.5rem;
}

/* Subtle button lift */
.hover-lift {
    transition: transform 0.3s cubic-bezier(0.175, 0.885, 0.32, 1.275), box-shadow 0.3s ease;
}

.hover-lift:hover:not(:disabled) {
    transform: translateY(-3px);
    box-shadow: 0 10px 20px rgba(13, 110, 253, 0.25) !important;
}
</style>