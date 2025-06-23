import { useAuthStore } from '@/stores/auth'

export function isTokenValid() {
    const auth = useAuthStore()
    const token = auth.token
    if (!token) return false;

    try {
        const payload = JSON.parse(atob(token.split('.')[1]));
        const exp = payload.exp;
        if (!exp) return false;
        return Date.now() < exp * 1000;
    } catch {
        return false;
    }
}