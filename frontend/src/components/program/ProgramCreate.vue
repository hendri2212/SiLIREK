<template>
    <div class="card shadow-sm rounded-4 border-0">
        <div class="card-body">
            <h4 class="card-title text-info mb-4">Tambah Program</h4>

            <form @submit.prevent="createProgram">
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
                        <option value="" disabled selected>-- Pilih Organisasi --</option>
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
                        <i class="bi bi-floppy-fill me-1"></i> Simpan
                    </button>
                </div>
            </form>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import api from '@/plugins/axios'

const router = useRouter()
const organizations = ref([])

const form = ref({
    code: '',
    name: '',
    organization_id: ''
})

const fetchOrganizations = async () => {
    try {
        const response = await api.get('/organization')
        organizations.value = response.data
    } catch (error) {
        console.error('Error fetching organizations:', error)
    }
}

const createProgram = async () => {
    try {
        await api.post('/program', form.value)
        alert('Program berhasil ditambahkan')
        router.push({ name: 'program.list' })
    } catch (error) {
        console.error('Error creating program:', error)
        if (error.response && error.response.data && error.response.data.error) {
            alert(error.response.data.error)
        } else {
            alert('Terjadi kesalahan saat menyimpan program')
        }
    }
}

onMounted(() => {
    fetchOrganizations()
})
</script>
