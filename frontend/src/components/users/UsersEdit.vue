<template>
    <div class="card shadow-sm rounded-4 border-0">
        <div class="card-body">
            <div class="mb-4 d-flex align-items-center">
                <h4 class="card-title text-info">Edit Pengguna</h4>
            </div>
            <form @submit.prevent="updateUser">
                <div class="form-floating mb-3">
                    <select v-model.number="user.position_id" name="position" class="form-control" id="floatingPosition"
                        required>
                        <option value="" selected disabled>-- Pilih Jabatan --</option>
                        <option v-for="position in positions" :key="position.id" :value="position.id">
                            {{ position.name }}
                        </option>
                    </select>
                    <label for="floatingPosition">Jabatan Kedinasan</label>
                </div>
                <div class="form-floating mb-3">
                    <select v-model.number="user.leader_id" name="leader" class="form-control" id="floatingLeader"
                        required>
                        <option value="" selected disabled>-- Pilih Kegiatan --</option>
                        <option v-for="leader in leaders" :key="leader.id" :value="leader.id">
                            {{ leader.name }}
                        </option>
                    </select>
                    <label for="floatingLeader">Jabatan Kegiatan</label>
                </div>
                <button class="w-100 btn btn-lg btn-info text-white" type="submit">Update Pengguna</button>
            </form>
        </div>
    </div>
</template>
<script setup>
import { reactive, onBeforeMount, ref } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import api from '@/plugins/axios';

const router = useRouter();
const route = useRoute();

const user = reactive({
    position_id: '',
    leader_id: ''
});

const positions = ref([]);
const leaders = ref([]);

const fetchPositions = async () => {
    try {
        const response = await api.get('/positions');
        positions.value = response.data;
    } catch (error) {
        console.error('Error fetching positions:', error);
    }
};

const fetchLeaders = async () => {
    try {
        const response = await api.get('/leaders');
        leaders.value = response.data;
    } catch (error) {
        console.error('Error fetching leaders:', error);
    }
};

const getUser = async () => {
    try {
        const response = await api.get(`/users/${route.params.id}`);
        Object.assign(user, response.data);
    } catch (error) {
        console.error('Error fetching user data:', error);
    }
};

const updateUser = async () => {
    try {
        const formData = new FormData();
        formData.append('position_id', user.position_id);
        formData.append('leader_id', user.leader_id);

        await api.put(`/users/${route.params.id}`, formData, {
            headers: {
                'Content-Type': 'multipart/form-data'
            }
        });
        router.push({ name: 'users.list' });
    } catch (error) {
        console.error('Error updating user:', error);
    }
};

onBeforeMount(() => {
    getUser();
    fetchPositions();
    fetchLeaders();
});
</script>