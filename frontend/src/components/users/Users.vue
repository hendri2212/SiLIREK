<template>
    <div class="card shadow-sm rounded-4 border-0">
        <div class="card-body">
            <div class="mb-4 d-flex align-items-center justify-content-between">
                <h4 class="card-title text-info">Data Pengguna</h4>
                <router-link class="btn btn-info text-white" :to="{ name: 'users.create' }">
                    <i class="bi bi-person-fill-add me-1"></i>
                    User
                </router-link>
            </div>
            <div v-if="isLoading" class="text-center my-4">
                Loading...
            </div>
            <div v-else class="table-responsive">
                <table class="table table-border-bottom-0">
                    <thead>
                        <tr>
                            <th class="text-center">No</th>
                            <th>Nama Lengkap</th>
                            <th>Email</th>
                            <th>Jabatan Kedinasan</th>
                            <th colspan="2">Jabatan Kegiatan</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="(user, index) in users" :key="index" class="align-middle">
                            <td class="text-center">{{ index + 1 }}</td>
                            <td>{{ user.full_name }}</td>
                            <td>{{ user.email }}</td>
                            <td>{{ user.position?.name || '-' }}</td>
                            <td>{{ user.leader?.name || '-' }}</td>
                            <td class="d-flex justify-content-end">
                                <div class="btn-group" role="group" aria-label="Basic example">
                                    <router-link class="btn btn-warning text-white"
                                        :to="{ name: 'users.edit', params: { id: user.id } }">
                                        <i class="bi bi-pencil-fill me-1"></i>
                                    </router-link>
                                    <!-- <button class="btn btn-danger text-white" @click="deleteUser(user.id)">
                                        <i class="bi bi-trash3-fill me-1"></i>
                                    </button> -->
                                </div>
                            </td>
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

const users = ref([])
const isLoading = ref(true)

onMounted(async () => {
    try {
        const response = await api.get('/users');
        users.value = response.data;
    } catch (error) {
        console.error('API error:', error);
    }
    isLoading.value = false;
})

const deleteUser = async (id) => {
    if (!confirm('Apakah Anda yakin ingin menghapus user ini?')) return;
    try {
        await api.delete(`/users/${id}`);
        users.value = users.value.filter(user => user.id !== id);
    } catch (error) {
        console.error('Gagal menghapus user:', error);
    }
};
</script>