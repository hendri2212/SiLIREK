import axios from 'axios';
import router from '@/router';
import { useAuthStore } from '@/stores/auth'

const baseURL = import.meta.env.DEV
    ? import.meta.env.VITE_API_URL_DEV
    : import.meta.env.VITE_API_URL_PROD

// Buat instance axios dengan baseURL yang eksplisit
const api = axios.create({
    baseURL,
    timeout: 10000,
});

// Request interceptor: tambahkan token ke setiap request
api.interceptors.request.use(config => {
    const auth = useAuthStore();
    if (auth.token) {
        config.headers.Authorization = `Bearer ${auth.token}`;
    }
    return config;
}, error => Promise.reject(error));

// Response interceptor: tangani error global
api.interceptors.response.use(
    response => response,
    error => {
        if (error.response) {
            if (error.response.status === 401) {
                console.warn('Unauthorized, redirecting to login...');
                const auth = useAuthStore();
                auth.logout();
                router.push({ name: 'login' });
            } else {
                console.error('API Error:', error.response.data?.message || error.message);
            }
        } else if (error.request) {
            // Request dibuat tapi tidak ada response (network error)
            console.error('Network Error: No response received', error.request);
        } else {
            console.error('Request Setup Error:', error.message);
        }
        return Promise.reject(error);
    }
);

export default api;