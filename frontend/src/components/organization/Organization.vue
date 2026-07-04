<template>
    <div class="pb-5">
        <!-- Page Header -->
        <div class="mb-5 d-flex align-items-center justify-content-between gap-3">
            <div>
                <h3 class="fw-bold text-dark mb-1">
                    <i class="bi bi-diagram-3 me-2 text-success opacity-75"></i>Manajemen Organisasi
                </h3>
                <p class="text-muted mb-0">Kelola data induk organisasi dalam sistem.</p>
            </div>
            <router-link :to="{ name: 'organization.create' }" class="btn btn-info text-white rounded-pill px-4 py-2 shadow-sm fw-medium d-flex align-items-center gap-2 flex-shrink-0 hover-lift">
                <i class="bi bi-plus-lg"></i>
                <span class="d-none d-md-inline">Tambah Organisasi</span>
            </router-link>
        </div>

        <!-- Section label -->
        <div class="section-label mb-3">
            <span class="text-success text-uppercase fw-bold small" style="letter-spacing: 0.08em;">Daftar Organisasi</span>
            <div class="section-line mt-1"></div>
        </div>

        <!-- Empty state -->
        <div v-if="organizations.length === 0" class="text-center py-5 text-muted">
            <i class="bi bi-diagram-3 display-3 opacity-25 d-block mb-3"></i>
            <p class="mb-0">Belum ada data organisasi.</p>
        </div>

        <!-- Table -->
        <div v-else class="table-responsive">
            <table class="table flat-table align-middle">
                <thead>
                    <tr>
                        <th>Nama Organisasi</th>
                        <th>Nomor</th>
                        <th>Induk (Parent)</th>
                        <th class="text-end">Aksi</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="org in organizations" :key="org.id" class="flat-row">
                        <td class="fw-medium text-dark">{{ org.name }}</td>
                        <td class="text-muted">{{ org.number }}</td>
                        <td>
                            <span v-if="org.parent" class="badge bg-success bg-opacity-10 text-success border border-success border-opacity-25 rounded-pill px-3">
                                {{ org.parent.name }}
                            </span>
                            <span v-else class="text-muted small">—</span>
                        </td>
                        <td class="text-end">
                            <router-link :to="{ name: 'organization.edit', params: { id: org.id } }"
                                class="btn btn-sm btn-light text-warning rounded-circle action-btn me-1" title="Edit">
                                <i class="bi bi-pencil-fill"></i>
                            </router-link>
                            <button class="btn btn-sm btn-light text-danger rounded-circle action-btn" @click="deleteOrganization(org.id)" title="Hapus">
                                <i class="bi bi-trash-fill"></i>
                            </button>
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

const organizations = ref([])

const fetchOrganizations = async () => {
    try {
        const response = await api.get('/organization')
        organizations.value = response.data
    } catch (error) {
        console.error('Error fetching organizations:', error)
        alert('Gagal mengambil data organisasi')
    }
}

const deleteOrganization = async (id) => {
    if (confirm('Apakah Anda yakin ingin menghapus organisasi ini?')) {
        try {
            await api.delete(`/organization/${id}`)
            alert('Organisasi berhasil dihapus')
            await fetchOrganizations()
        } catch (error) {
            console.error('Error deleting organization:', error)
            alert('Gagal menghapus organisasi. Pastikan organisasi ini tidak memiliki sub-organisasi (children).')
        }
    }
}

onMounted(() => {
    fetchOrganizations()
})
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
    padding: 0.9rem 0.75rem;
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
