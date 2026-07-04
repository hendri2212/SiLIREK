<template>
    <div class="d-flex flex-column flex-md-row align-items-center mb-4 border-bottom">
        <a href="/" class="d-flex align-items-center link-body-emphasis text-decoration-none">
            <span class="fs-3 fw-bold text-info">SiLIREK</span>
        </a>

        <nav class="d-inline-flex mt-2 mt-md-0 ms-md-auto align-items-center">
            <router-link v-if="isAdminOrSuper" :to="{ name: 'program.list' }"
                class="fw-bold me-3 py-2 link-body-emphasis text-decoration-none">Program</router-link>
            <router-link :to="{ name: 'activities.list' }"
                class="fw-bold me-3 py-2 link-body-emphasis text-decoration-none">Kegiatan</router-link>
            <router-link v-if="isAdminOrSuper" :to="{ name: 'reports.kegiatan' }"
                class="fw-bold me-3 py-2 link-body-emphasis text-decoration-none">Laporan</router-link>
            <router-link v-if="isSuperadmin" :to="{ name: 'organization.list' }"
                class="fw-bold me-3 py-2 link-body-emphasis text-decoration-none">Organisasi</router-link>
            <router-link v-if="isSuperadmin" :to="{ name: 'users.list' }"
                class="fw-bold me-3 py-2 link-body-emphasis text-decoration-none">Pengguna</router-link>
            <div v-if="$route.name != 'login'" class="dropdown">
                <a class="btn p-0" href="#" role="button" data-bs-toggle="dropdown" aria-expanded="false">
                    <img v-if="auth.user?.photo" :src="`${baseApiUrl}/uploads/photos/${auth.user.photo}`" alt="Profile"
                        class="rounded-circle" width="32" height="32" />
                    <i v-else class="bi bi-person-circle fs-1"></i>
                </a>
                <ul class="dropdown-menu dropdown-menu-end p-2">
                    <li>
                        <router-link v-if="auth.user?.id" class="dropdown-item"
                            :to="{ name: 'users.show', params: { id: auth.user.id } }">
                            Profile
                        </router-link>
                    </li>
                    <li>
                        <hr class="dropdown-divider" />
                    </li>
                    <li>
                        <button class="dropdown-item" @click="logout">Logout</button>
                    </li>
                </ul>
            </div>
        </nav>
    </div>
</template>
<script setup>
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

// const baseApiUrl = import.meta.env.MODE === 'development'
//     ? import.meta.env.VITE_API_URL_DEV
//     : import.meta.env.VITE_API_URL_PROD

const baseApiUrl = import.meta.env.MODE === 'development'
    ? import.meta.env.VITE_API_URL_DEV.replace(/\/api$/, '')
    : import.meta.env.VITE_API_URL_PROD.replace(/\/api$/, '')

const router = useRouter()
const auth = useAuthStore()

const role = computed(() => auth.user?.role)

const isAdminOrSuper = computed(() =>
    role.value === 'superadmin' || role.value === 'admin'
)

const isUserOrSuper = computed(() =>
    role.value === 'user' || role.value === 'superadmin'
)

const isSuperadmin = computed(() =>
    role.value === 'superadmin'
)

const logout = () => {
    auth.logout()
    window.location.reload()
    router.push({ name: 'login' })
}
</script>