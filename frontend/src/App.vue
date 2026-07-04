<template>
    <div class="container py-3">
        <header>
            <Navbar v-if="$route.name != 'login'" />
        </header>
        <RouterView />
    </div>
</template>

<script setup>
import { watch } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import Navbar from './components/Navbar.vue';
import { useAuthStore } from '@/stores/auth'
import { isTokenValid } from '@/utils/cek_token'

const router = useRouter();
const route = useRoute();
const auth = useAuthStore();

// Pantau setiap perubahan route, cek token saat itu juga
// Ini lebih andal daripada onBeforeMount yang hanya berjalan sekali
watch(
    () => route.name,
    (routeName) => {
        if (routeName && routeName !== 'login') {
            if (!isTokenValid()) {
                auth.logout();
                router.push({ name: 'login' });
            }
        }
    },
    { immediate: true }
);
</script>