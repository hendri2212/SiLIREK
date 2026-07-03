<template>
    <div class="card shadow-sm rounded-4 border-0">
        <div class="card-body">
            <h4 class="card-title text-info mb-4">Edit Organisasi</h4>

            <form @submit.prevent="updateOrganization" v-if="!isLoading">
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
                        <option v-for="org in availableParents" :key="org.id" :value="org.id">
                            {{ org.number }} - {{ org.name }}
                        </option>
                    </select>
                    <div class="form-text text-danger" v-if="form.parent_id === orgId">
                        Organisasi tidak bisa menjadi induk bagi dirinya sendiri.
                    </div>
                </div>

                <div class="d-flex justify-content-end">
                    <router-link :to="{ name: 'organization.list' }" class="btn btn-secondary me-2">
                        Batal
                    </router-link>
                    <button type="submit" class="btn btn-info text-white" :disabled="form.parent_id === orgId">
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
import { ref, onMounted, computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import api from '@/plugins/axios'

const router = useRouter()
const route = useRoute()
const orgId = parseInt(route.params.id)

const organizations = ref([])
const isLoading = ref(true)

const form = ref({
    name: '',
    number: '',
    parent_id: null
})

// Prevent circular reference visually by not letting it pick itself (though they might manually edit HTML, backend should ideally check this too, but for UI it's fine)
const availableParents = computed(() => {
    return organizations.value.filter(org => org.id !== orgId)
})

const fetchData = async () => {
    try {
        // Fetch all organizations for the dropdown
        const orgsResponse = await api.get('/organization')
        organizations.value = orgsResponse.data

        // Fetch the specific organization data
        const response = await api.get(`/organization/${orgId}`)
        const data = response.data
        
        form.value.name = data.name
        form.value.number = data.number
        form.value.parent_id = data.parent_id

    } catch (error) {
        console.error('Error fetching data:', error)
        alert('Gagal mengambil data')
        router.push({ name: 'organization.list' })
    } finally {
        isLoading.value = false
    }
}

const updateOrganization = async () => {
    if (form.value.parent_id === orgId) {
        alert('Organisasi tidak dapat menjadi induk bagi dirinya sendiri.')
        return
    }

    try {
        await api.put(`/organization/${orgId}`, form.value)
        alert('Organisasi berhasil diperbarui')
        router.push({ name: 'organization.list' })
    } catch (error) {
        console.error('Error updating organization:', error)
        if (error.response && error.response.data && error.response.data.error) {
            alert(error.response.data.error)
        } else {
            alert('Terjadi kesalahan saat memperbarui organisasi')
        }
    }
}

onMounted(() => {
    fetchData()
})
</script>
