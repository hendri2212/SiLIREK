import axios from 'axios';
import router from '@/router'; // pastikan import router-mu
import { useAuthStore } from '@/stores/auth'

// Buat instance axios
const api = axios.create({
    timeout: 10000,
});

// Tambahkan interceptor global
api.interceptors.response.use(
    response => response,
    error => {
        if (error.response) {
            if (error.response.status === 401) {
                console.warn('Unauthorized, redirecting to login...');
                router.push('/login');
            } else {
                console.error('API Error:', error.response.data.message || error.message);
            }
        }
        return Promise.reject(error);
    }
);

api.interceptors.request.use(config => {
    const auth = useAuthStore();
    if (auth.token) {
        config.headers.Authorization = `Bearer ${auth.token}`;
    }
    return config;
});

export default api;