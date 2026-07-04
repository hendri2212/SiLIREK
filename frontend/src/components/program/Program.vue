<template>
    <div class="mb-5">
        <!-- Header Section -->
        <div class="d-flex flex-column flex-md-row justify-content-between align-items-md-center mb-4 gap-3">
            <div>
                <h3 class="fw-bold text-dark mb-1">Manajemen Program</h3>
                <p class="text-muted small mb-0">Kelola daftar program beserta organisasinya</p>
            </div>
            <router-link :to="{ name: 'program.create' }" class="btn btn-info text-white rounded-pill px-4 shadow-sm hover-lift d-flex align-items-center gap-2">
                <i class="bi bi-plus-lg fs-5"></i>
                <span class="fw-bold">Tambah Program</span>
            </router-link>
        </div>
        
        <!-- Cards List -->
        <div class="d-flex flex-column gap-3">
            <div v-for="prog in programs" :key="prog.id" class="card border-0 shadow-sm rounded-4 hover-card overflow-hidden">
                <div class="card-body p-0">
                    <div class="row g-0 align-items-center p-3 p-md-4">
                        <!-- Icon & Code -->
                        <div class="col-12 col-md-3 mb-3 mb-md-0 d-flex align-items-center gap-3">
                            <div class="bg-info bg-opacity-10 text-info rounded-circle d-flex align-items-center justify-content-center flex-shrink-0" style="width: 48px; height: 48px;">
                                <i class="bi bi-folder2-open fs-4"></i>
                            </div>
                            <div>
                                <span class="badge bg-light text-dark border px-2 py-1">{{ prog.code }}</span>
                            </div>
                        </div>
                        
                        <!-- Details -->
                        <div class="col-12 col-md-5 mb-3 mb-md-0">
                            <h6 class="mb-1 fw-bold text-dark">{{ prog.name }}</h6>
                            <div class="text-muted small d-flex align-items-center">
                                <i class="bi bi-building me-2"></i> 
                                <span>
                                    {{ prog.organization ? prog.organization.name : 'Tidak ada organisasi' }}
                                </span>
                            </div>
                        </div>
                        
                        <!-- Actions -->
                        <div class="col-12 col-md-4 text-md-end">
                            <div class="d-flex justify-content-md-end gap-2">
                                <router-link :to="{ name: 'program.edit', params: { id: prog.id } }" class="btn btn-light btn-sm text-warning rounded-pill px-3 hover-lift fw-medium">
                                    <i class="bi bi-pencil-fill me-1"></i> Edit
                                </router-link>
                                <button @click="deleteProgram(prog.id)" class="btn btn-light btn-sm text-danger rounded-pill px-3 hover-lift fw-medium">
                                    <i class="bi bi-trash-fill me-1"></i> Hapus
                                </button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            
            <!-- Empty State -->
            <div v-if="programs.length === 0" class="text-center py-5 bg-white rounded-4 shadow-sm">
                <i class="bi bi-inbox text-muted" style="font-size: 4rem; opacity: 0.5;"></i>
                <h5 class="text-muted fw-bold mt-3">Belum Ada Program</h5>
                <p class="text-muted small mb-4">Tambahkan program baru untuk memulai pengelolaan data.</p>
                <router-link :to="{ name: 'program.create' }" class="btn btn-info text-white rounded-pill px-4 hover-lift">
                    <i class="bi bi-plus-lg me-1"></i> Tambah Program Sekarang
                </router-link>
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

<style scoped>
.hover-card {
    transition: transform 0.2s ease, box-shadow 0.2s ease;
}

.hover-card:hover {
    transform: translateY(-3px);
    box-shadow: 0 8px 15px rgba(0, 0, 0, 0.08) !important;
}

.hover-lift {
    transition: transform 0.2s ease, background-color 0.2s ease;
}

.hover-lift:hover {
    transform: translateY(-2px);
}
</style>
