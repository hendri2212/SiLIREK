<template>
    <div class="card shadow-sm rounded-4 border-0">
        <div class="card-body">
            <h4 class="card-title text-info mb-4">Tambah Organisasi</h4>

            <form @submit.prevent="createOrganization">
                <div class="mb-3">
                    <label for="name" class="form-label fw-bold">Nama Organisasi</label>
                    <input type="text" v-model="form.name" class="form-control" id="name" required
                        placeholder="Contoh: Dinas Pendidikan">
                </div>

                <div class="mb-3">
                    <label for="number" class="form-label fw-bold">Nomor Organisasi</label>
                    <input type="text" v-model="form.number" class="form-control" id="number" required
                        placeholder="Contoh: 1.01.0.00.0.00.01.0000">
                </div>

                <div class="mb-4">
                    <label for="parent_id" class="form-label fw-bold">Induk Organisasi (Opsional)</label>
                    <select v-model.number="form.parent_id" class="form-select" id="parent_id">
                        <option :value="null">-- Tidak Ada (Sebagai Induk Utama) --</option>
                        <option v-for="org in organizations" :key="org.id" :value="org.id">
                            {{ org.number }} - {{ org.name }}
                        </option>
                    </select>
                </div>

                <div class="d-flex justify-content-end">
                    <router-link :to="{ name: 'organization.list' }" class="btn btn-secondary me-2">
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
    name: '',
    number: '',
    parent_id: null
})

const fetchOrganizations = async () => {
    try {
        const response = await api.get('/organization')
        organizations.value = response.data
    } catch (error) {
        console.error('Error fetching organizations:', error)
    }
}

const createOrganization = async () => {
    try {
        await api.post('/organization', form.value)
        alert('Organisasi berhasil ditambahkan')
        router.push({ name: 'organization.list' })
    } catch (error) {
        console.error('Error creating organization:', error)
        if (error.response && error.response.data && error.response.data.error) {
            alert(error.response.data.error)
        } else {
            alert('Terjadi kesalahan saat menyimpan organisasi')
        }
    }
}

onMounted(() => {
    fetchOrganizations()
})
</script>
