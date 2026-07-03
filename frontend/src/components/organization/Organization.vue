<template>
    <div class="card shadow-sm rounded-4 border-0">
        <div class="card-body">
            <div class="mb-4 d-flex align-items-center justify-content-between">
                <h4 class="card-title text-info">Management Organisasi</h4>
                <router-link :to="{ name: 'organization.create' }" class="btn btn-info text-white">
                    <i class="bi bi-plus-lg me-1"></i> Tambah Organisasi
                </router-link>
            </div>
            
            <div class="table-responsive">
                <table class="table">
                    <thead class="text-nowrap">
                        <tr>
                            <th scope="col">Nama Organisasi</th>
                            <th scope="col">Nomor Organisasi</th>
                            <th scope="col">Induk (Parent)</th>
                            <th class="text-center" scope="col">Aksi</th>
                        </tr>
                    </thead>
                    <tbody class="text-nowrap">
                        <tr v-for="org in organizations" :key="org.id">
                            <td>{{ org.name }}</td>
                            <td>{{ org.number }}</td>
                            <td>{{ org.parent ? org.parent.name : '-' }}</td>
                            <td class="text-center">
                                <router-link :to="{ name: 'organization.edit', params: { id: org.id } }"
                                    class="btn btn-warning btn-sm me-2">
                                    <i class="bi bi-pencil-fill"></i> Edit
                                </router-link>
                                <button class="btn btn-danger btn-sm" @click="deleteOrganization(org.id)">
                                    <i class="bi bi-trash-fill"></i> Hapus
                                </button>
                            </td>
                        </tr>
                        <tr v-if="organizations.length === 0">
                            <td colspan="4" class="text-center text-muted">Belum ada data organisasi</td>
                        </tr>
                    </tbody>
                </table>
            </div>
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
