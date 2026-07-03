<template>
    <div class="card shadow-sm rounded-4 border-0">
        <div class="card-body">
            <h4 class="card-title text-info mb-4">Edit Program</h4>

            <form @submit.prevent="updateProgram" v-if="!isLoading">
                <div class="mb-3">
                    <label for="code" class="form-label fw-bold">Kode Program</label>
                    <input type="text" v-model="form.code" class="form-control" id="code" required
                        placeholder="Contoh: 1.01.01">
                </div>

                <div class="mb-3">
                    <label for="name" class="form-label fw-bold">Nama Program</label>
                    <input type="text" v-model="form.name" class="form-control" id="name" required
                        placeholder="Contoh: Program Pelayanan Administrasi Perkantoran">
                </div>

                <div class="mb-4">
                    <label for="organization_id" class="form-label fw-bold">Organisasi Terkait</label>
                    <select v-model.number="form.organization_id" class="form-select" id="organization_id" required>
                        <option value="" disabled>-- Pilih Organisasi --</option>
                        <option v-for="org in organizations" :key="org.id" :value="org.id">
                            {{ org.number }} - {{ org.name }}
                        </option>
                    </select>
                </div>

                <div class="d-flex justify-content-end">
                    <router-link :to="{ name: 'program.list' }" class="btn btn-secondary me-2">
                        Batal
                    </router-link>
                    <button type="submit" class="btn btn-info text-white">
                        <i class="bi bi-floppy-fill me-1"></i> Update
                    </button>
                </div>
            </form>
            <div v-else class="text-center py-4">
                <div class="spinner-border text-info" role="status">
                    <span class="visually-hidden">Loading...</span>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import api from '@/plugins/axios'

const router = useRouter()
const route = useRoute()
const programId = parseInt(route.params.id)

const organizations = ref([])
const isLoading = ref(true)

const form = ref({
    code: '',
    name: '',
    organization_id: ''
})

const fetchData = async () => {
    try {
        // Fetch organizations for dropdown
        const orgsResponse = await api.get('/organization')
        organizations.value = orgsResponse.data

        // Fetch program data
        const response = await api.get(`/program/${programId}`)
        const data = response.data
        
        form.value.code = data.code
        form.value.name = data.name
        form.value.organization_id = data.organization_id

    } catch (error) {
        console.error('Error fetching data:', error)
        alert('Gagal mengambil data')
        router.push({ name: 'program.list' })
    } finally {
        isLoading.value = false
    }
}

const updateProgram = async () => {
    try {
        await api.put(`/program/${programId}`, form.value)
        alert('Program berhasil diperbarui')
        router.push({ name: 'program.list' })
    } catch (error) {
        console.error('Error updating program:', error)
        if (error.response && error.response.data && error.response.data.error) {
            alert(error.response.data.error)
        } else {
            alert('Terjadi kesalahan saat memperbarui program')
        }
    }
}

onMounted(() => {
    fetchData()
})
</script>
