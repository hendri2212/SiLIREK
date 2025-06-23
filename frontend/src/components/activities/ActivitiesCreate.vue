<template>
    <div class="card shadow-sm rounded-4 border-0">
        <div class="card-body">
            <div class="mb-4 d-flex align-items-center">
                <h4 class="card-title text-info">Tambah Kegiatan</h4>
            </div>
            <form @submit.prevent="createActivity">
                <div class="form-floating mb-3">
                    <input v-model="activity.code" type="text" class="form-control" id="floatingActivity"
                        placeholder="Kode Kegiatan" required>
                    <label for="floatingActivity">Kode Kegiatan</label>
                </div>
                <div class="form-floating mb-3">
                    <input v-model="activity.name" type="text" class="form-control" id="floatingNameActivity" placeholder="Nama Kegiatan"
                        required>
                    <label for="floatingNameActivity">Nama Kegiatan</label>
                </div>
                <button class="w-100 btn btn-lg btn-info text-white" type="submit">Tambah Kegiatan</button>
            </form>
        </div>
    </div>
</template>
<script setup>
import { reactive } from 'vue';
import { useRouter } from 'vue-router';
import api from '@/plugins/axios';

const router = useRouter();

const activity = reactive({
    code: '',
    name: '',
});

const createActivity = async () => {
    try {
        await api.post('/activity', activity);
        router.push({ name: 'activities.list' });
    } catch (error) {
        console.error('Error creating activities:', error);
        if (error.response && error.response.data && error.response.data.error) {
            alert(error.response.data.error);
        } else {
            alert('Terjadi kesalahan saat menambahkan kegiatan.');
        }
    }
};
</script>