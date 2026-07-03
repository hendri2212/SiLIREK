<template>
    <div class="card shadow-sm rounded-4 border-0">
        <div class="card-body">
            <div class="mb-4 d-flex align-items-center justify-content-between">
                <h4 class="card-title text-info">Management Program</h4>
                <router-link :to="{ name: 'program.create' }" class="btn btn-info text-white">
                    <i class="bi bi-plus-lg me-1"></i> Tambah Program
                </router-link>
            </div>
            
            <div class="table-responsive">
                <table class="table">
                    <thead class="text-nowrap">
                        <tr>
                            <th scope="col">Kode Program</th>
                            <th scope="col">Nama Program</th>
                            <th scope="col">Organisasi</th>
                            <th class="text-center" scope="col">Aksi</th>
                        </tr>
                    </thead>
                    <tbody class="text-nowrap">
                        <tr v-for="prog in programs" :key="prog.id">
                            <td>{{ prog.code }}</td>
                            <td>{{ prog.name }}</td>
                            <td>{{ prog.organization ? prog.organization.name : '-' }}</td>
                            <td class="text-center">
                                <router-link :to="{ name: 'program.edit', params: { id: prog.id } }"
                                    class="btn btn-warning btn-sm me-2">
                                    <i class="bi bi-pencil-fill"></i> Edit
                                </router-link>
                                <button class="btn btn-danger btn-sm" @click="deleteProgram(prog.id)">
                                    <i class="bi bi-trash-fill"></i> Hapus
                                </button>
                            </td>
                        </tr>
                        <tr v-if="programs.length === 0">
                            <td colspan="4" class="text-center text-muted">Belum ada data program</td>
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

const programs = ref([])

const fetchPrograms = async () => {
    try {
        const response = await api.get('/program')
        programs.value = response.data
    } catch (error) {
        console.error('Error fetching programs:', error)
        alert('Gagal mengambil data program')
    }
}

const deleteProgram = async (id) => {
    if (confirm('Apakah Anda yakin ingin menghapus program ini?')) {
        try {
            await api.delete(`/program/${id}`)
            alert('Program berhasil dihapus')
            await fetchPrograms()
        } catch (error) {
            console.error('Error deleting program:', error)
            alert('Gagal menghapus program.')
        }
    }
}

onMounted(() => {
    fetchPrograms()
})
</script>
