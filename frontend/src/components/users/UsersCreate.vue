<template>
    <div class="card shadow-sm rounded-4 border-0">
        <div class="card-body">
            <div class="mb-4 d-flex align-items-center">
                <h4 class="card-title text-info">Tambah Pengguna</h4>
            </div>
            <form @submit.prevent="createUser">
                <div class="form-floating mb-3">
                    <input v-model="user.full_name" type="text" class="form-control" id="floatingFullName"
                        placeholder="Full Name" required>
                    <label for="floatingFullName">Nama Lengkap</label>
                </div>
                <div class="form-floating mb-3">
                    <input v-model="user.email" type="email" class="form-control" id="floatingEmail" placeholder="Email"
                        pattern="^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$"
                        title="Masukkan format email yang benar" required>

                    <label for="floatingEmail">Email</label>
                </div>
                <div class="form-floating mb-3">
                    <input v-model="user.nip" type="text" class="form-control" id="floatingNIP" placeholder="NIP">
                    <label for="floatingNIP">NIP</label>
                </div>
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
                <button class="w-100 btn btn-lg btn-info text-white" type="submit">Tambah Pengguna</button>
            </form>
        </div>
    </div>
</template>
<script setup>
import { ref, reactive, onBeforeMount } from 'vue';
import { useRouter } from 'vue-router';
import api from '@/plugins/axios';

const router = useRouter();

const user = reactive({
    full_name: '',
    email: '',
    nip: '',
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

onBeforeMount(() => {
    fetchPositions();
    fetchLeaders();
});

const createUser = async () => {
    try {
        await api.post('/users', user);
        router.push({ name: 'users.list' });
    } catch (error) {
        console.error('Error creating user:', error);
    }
};
</script>