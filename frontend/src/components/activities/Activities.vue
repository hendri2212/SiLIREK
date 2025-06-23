<template>
    <div class="card shadow-sm rounded-4 border-0">
        <div class="card-body">
            <div class="mb-4 d-flex align-items-center justify-content-between">
                <h4 class="card-title text-info">Data Kegiatan</h4>
                <router-link class="btn btn-info text-white" :to="{ name: 'activities.create' }">
                    <i class="bi bi-person-fill-add me-1"></i>
                    Kegiatan
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
                            <th>Kode Kegiatan</th>
                            <!-- <th colspan="2">Nama Kegiatan</th> -->
                            <th>Nama Kegiatan</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="(data, index) in activities" :key="index" class="align-middle">
                            <td class="text-center">{{ index + 1 }}</td>
                            <td>{{ data.code }}</td>
                            <td>{{ data.name }}</td>
                            <!-- <td class="d-flex justify-content-end">
                                <div class="d-flex" role="group" aria-label="Basic example">
                                    <router-link
                                        class="btn btn-outline-warning d-flex align-items-center gap-2 px-3 py-1 rounded-pill shadow-sm"
                                        :to="{ name: 'activities.edit', params: { id: data.id } }">
                                        <i class="bi bi-pencil-square fs-6"></i>
                                        <span>Edit</span>
                                    </router-link>
                                    <button
                                        class="btn btn-outline-danger d-flex align-items-center gap-2 px-3 py-1 rounded-pill shadow-sm ms-2"
                                        @click="deleteUser(data.id)">
                                        <i class="bi bi-trash fs-6"></i>
                                        <span>Hapus</span>
                                    </button>
                                </div>
                            </td> -->
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</template>
<script setup>
import { ref, onBeforeMount } from 'vue'
import api from '@/plugins/axios'

const isLoading = ref(true)
const activities = ref([])

const fetchActivities = async () => {
    try {
        const response = await api.get('/activity');
        activities.value = response.data;
    } catch (error) {
        console.error('Error fetching activities:', error);
    } finally {
        isLoading.value = false;
    }
};

onBeforeMount(() => {
    fetchActivities();
});
</script>