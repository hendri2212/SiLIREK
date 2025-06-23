<template>
    <div class="container py-3">
        <header>
            <Navbar v-if="$route.name != 'login'" />
        </header>

        <RouterView />
    </div>
</template>
<script>
import { onBeforeMount } from 'vue';
import { useRouter } from 'vue-router';
import Navbar from './components/Navbar.vue';
import { useAuthStore } from '@/stores/auth'

export default {
    name: 'App',
    components: {
        Navbar
    },
    setup() {
        const router = useRouter();
        const auth = useAuthStore();

        const isTokenExpired = (token) => {
            try {
                const payload = JSON.parse(atob(token.split('.')[1]));
                const exp = payload.exp;
                if (!exp) return true;
                return Date.now() >= exp * 1000;
            } catch (e) {
                return true;
            }
        };

        onBeforeMount(() => {
            const token = auth.token;
            if (!token || isTokenExpired(token)) {
                auth.logout();
                router.push({ name: 'login' });
            }
        });

        return {};
    }
};
</script>