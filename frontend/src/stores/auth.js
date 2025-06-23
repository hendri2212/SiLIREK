import { ref } from 'vue'
import { defineStore } from 'pinia'

export const useAuthStore = defineStore(
    'auth',
    () => {
        const token = ref(null)
        const user = ref(null)

        function login(newToken) {
            token.value = newToken
        }

        async function fetchUser() {
            if (!token.value) {
                user.value = null
                return
            }

            const baseApiUrl = import.meta.env.DEV
                ? import.meta.env.VITE_API_URL_DEV
                : import.meta.env.VITE_API_URL_PROD

            try {
                const response = await fetch(`${baseApiUrl}/me`, {
                    headers: {
                        'Authorization': `Bearer ${token.value}`
                    }
                })
                if (!response.ok) {
                    user.value = null
                    return
                }
                const data = await response.json()
                user.value = data
            } catch {
                user.value = null
            }
        }

        function logout() {
            token.value = null
            user.value = null
        }

        return { token, user, login, fetchUser, logout }
    },
    {
        persist: {
            key: 'auth',
            storage: localStorage,
        }
    }
)
