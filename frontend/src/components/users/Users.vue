<template>
    <div class="pb-5">
        <!-- Page Header -->
        <div class="mb-5 d-flex align-items-center justify-content-between gap-3">
            <div>
                <h3 class="fw-bold text-dark mb-1">
                    <i class="bi bi-people me-2 text-success opacity-75"></i>Data Pengguna
                </h3>
                <p class="text-muted mb-0">Kelola seluruh akun pengguna yang terdaftar dalam sistem.</p>
            </div>
            <router-link :to="{ name: 'users.create' }" class="btn btn-success text-white rounded-pill px-4 py-2 shadow-sm fw-medium d-flex align-items-center gap-2 flex-shrink-0">
                <i class="bi bi-person-fill-add"></i>
                <span class="d-none d-md-inline">Tambah Pengguna</span>
            </router-link>
        </div>

        <!-- Section label -->
        <div class="section-label mb-3">
            <span class="text-success text-uppercase fw-bold small" style="letter-spacing: 0.08em;">Daftar Pengguna</span>
            <div class="section-line mt-1"></div>
        </div>

        <!-- Loading -->
        <div v-if="isLoading" class="text-center py-5 text-muted">
            <div class="spinner-grow spinner-grow-sm text-success me-2" role="status"></div>
            Memuat data pengguna...
        </div>

        <!-- Empty state -->
        <div v-else-if="users.length === 0" class="text-center py-5 text-muted">
            <i class="bi bi-people display-3 opacity-25 d-block mb-3"></i>
            <p class="mb-0">Belum ada pengguna yang terdaftar.</p>
        </div>

        <!-- Table -->
        <div v-else class="table-responsive">
            <table class="table flat-table align-middle">
                <thead>
                    <tr>
                        <th class="text-center" style="width: 48px;">No</th>
                        <th>Nama Lengkap</th>
                        <th>Email</th>
                        <th>Jabatan Kedinasan</th>
                        <th>Jabatan Kegiatan</th>
                        <th class="text-end">Aksi</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="(user, index) in users" :key="user.id" class="flat-row">
                        <td class="text-center text-muted small">{{ index + 1 }}</td>
                        <td>
                            <div class="d-flex align-items-center gap-2">
                                <div class="avatar-sm rounded-circle bg-success bg-opacity-10 text-success d-flex align-items-center justify-content-center fw-bold flex-shrink-0" style="width: 36px; height: 36px; font-size: 0.85rem;">
                                    {{ user.full_name?.charAt(0)?.toUpperCase() || '?' }}
                                </div>
                                <span class="fw-medium text-dark">{{ user.full_name }}</span>
                            </div>
                        </td>
                        <td class="text-muted small">{{ user.email }}</td>
                        <td>
                            <span v-if="user.position?.name" class="badge bg-success bg-opacity-10 text-success border border-success border-opacity-25 rounded-pill px-3">
                                {{ user.position.name }}
                            </span>
                            <span v-else class="text-muted small">—</span>
                        </td>
                        <td class="text-muted small">{{ user.leader?.name || '—' }}</td>
                        <td class="text-end">
                            <router-link
                                :to="{ name: 'users.edit', params: { id: user.id } }"
                                class="btn btn-sm btn-light text-warning rounded-circle action-btn me-1" title="Edit">
                                <i class="bi bi-pencil-fill"></i>
                            </router-link>
                            <!-- <button class="btn btn-sm btn-light text-danger rounded-circle action-btn" @click="deleteUser(user.id)" title="Hapus">
                                <i class="bi bi-trash-fill"></i>
                            </button> -->
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import api from '@/plugins/axios'

const users = ref([])
const isLoading = ref(true)

onMounted(async () => {
    try {
        const response = await api.get('/users')
        users.value = response.data
    } catch (error) {
        console.error('API error:', error)
    } finally {
        isLoading.value = false
    }
})

const deleteUser = async (id) => {
    if (!confirm('Apakah Anda yakin ingin menghapus user ini?')) return
    try {
        await api.delete(`/users/${id}`)
        users.value = users.value.filter(user => user.id !== id)
    } catch (error) {
        console.error('Gagal menghapus user:', error)
    }
}
</script>

<style scoped>
.section-label .section-line {
    height: 2px;
    background: linear-gradient(90deg, rgba(25, 135, 84, 0.3), transparent);
    border-radius: 2px;
}
.flat-table {
    border-collapse: separate;
    border-spacing: 0 4px;
}
.flat-table thead th {
    background: transparent;
    border: none;
    color: #6c757d;
    font-size: 0.78rem;
    text-transform: uppercase;
    letter-spacing: 0.06em;
    padding-bottom: 0.75rem;
    border-bottom: 2px solid rgba(25, 135, 84, 0.1);
}
.flat-table tbody .flat-row td {
    background: rgba(255, 255, 255, 0.5);
    border: none;
    border-top: 1px solid rgba(0,0,0,0.04);
    border-bottom: 1px solid rgba(0,0,0,0.04);
    padding: 0.85rem 0.75rem;
    transition: background 0.2s ease;
}
.flat-table tbody .flat-row td:first-child {
    border-left: 1px solid rgba(0,0,0,0.04);
    border-radius: 0.5rem 0 0 0.5rem;
}
.flat-table tbody .flat-row td:last-child {
    border-right: 1px solid rgba(0,0,0,0.04);
    border-radius: 0 0.5rem 0.5rem 0;
}
.flat-row:hover td {
    background: rgba(255, 255, 255, 0.85) !important;
}
.action-btn {
    width: 32px;
    height: 32px;
    display: inline-flex;
    align-items: center;
    justify-content: center;
    transition: all 0.2s ease;
    border: 1px solid transparent;
}
.action-btn:hover {
    transform: translateY(-1px);
    box-shadow: 0 4px 8px rgba(0,0,0,0.1);
}
</style>