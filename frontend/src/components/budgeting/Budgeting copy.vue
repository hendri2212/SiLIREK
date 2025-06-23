<template>
    <div class="card shadow-sm rounded-4 border-0">
        <div class="card-body">
            <div class="mb-4 d-flex align-items-center justify-content-between">
                <h4 class="card-title text-info">Penganggaran</h4>
            </div>
            <form @submit.prevent="createBudget">
                <div class="d-flex justify-content-between gap-3">
                    <div>
                        <div class="form-floating mb-3">
                            <select v-model.number="form.user_id" name="user" class="form-control" id="floatingUser">
                                <option value="0" selected disabled>-- Pilih Pengguna --</option>
                                <option v-for="data in users" :key="data.id" :value="data.id">
                                    {{ data.full_name }}
                                </option>
                            </select>
                            <label for="floatingQuarter">Nama Pengguna</label>
                        </div>
                        <div class="form-floating mb-3">
                            <select v-model.number="form.quarterly_id" name="quarter" class="form-control" id="floatingQuarter"
                                required>
                                <option value="" selected disabled>-- Pilih Quarter --</option>
                                <option v-for="data in quarterlies" :key="data.id" :value="data.id">
                                    {{ data.quarter }} - {{ data.year }}
                                </option>
                            </select>
                            <label for="floatingQuarter">Triwulan</label>
                        </div>
                        <div class="form-floating mb-3">
                            <input v-model="computedBudget" type="text" class="form-control" id="floatingBudget"
                                placeholder="Jumlah Anggaran">
                            <label for="floatingBudget">Jumlah Anggaran</label>
                        </div>
                    </div>
                    <div>
                        <div class="form-floating mb-3">
                            <select v-model.number="form.activity_id" name="activity" class="form-control" id="floatingActivity"
                                required>
                                <option value="" selected disabled>-- Pilih Kegiatan --</option>
                                <option v-for="data in activities" :key="data.id" :value="data.id">
                                    {{ data.code }} - {{ data.name }}
                                </option>
                            </select>
                            <label for="floatingPosition">Nama Kegiatan</label>
                        </div>
                        <div class="form-floating mb-3">
                            <select v-model.number="form.sub_activity_id" name="sub_activity" class="form-control" id="floatingActivity">
                                <option value="" selected disabled>-- Pilih Sub Kegiatan --</option>
                                <option v-for="data in activities" :key="data.id" :value="data.id">
                                    {{ data.code }} - {{ data.name }}
                                </option>
                            </select>
                            <label for="floatingPosition">Nama Sub Kegiatan</label>
                        </div>
                        <div class="d-flex justify-content-end align-items-bottom">
                            <button class="btn btn-info text-white" type="submit">
                                <i class="bi bi-floppy me-1"></i>
                                Anggaran
                            </button>
                        </div>
                    </div>
                </div>
            </form>
        </div>
    </div>
    <div class="card shadow-sm rounded-4 border-0">
        <div class="card-body">
            <!-- <div class="mb-4 d-flex align-items-center justify-content-between">
                <h4 class="card-title text-info">Penganggaran</h4>
            </div> -->
            <div class="table-responsive">
                <table class="table">
                    <thead class="text-nowrap">
                        <tr>
                            <th class="text-center" scope="col">No</th>
                            <th scope="col">Nama Pengguna</th>
                            <th scope="col">Triwulan</th>
                            <th scope="col">Kegiatan</th>
                            <th scope="col">Sub Kegiatan</th>
                            <th scope="col">Jumlah Anggaran</th>
                        </tr>
                    </thead>
                    <tbody class="text-nowrap">
                        <tr v-for="(data, index) in budgets" :key="index">
                            <td class="text-center">{{ index + 1 }}</td>
                            <td>{{ data.User?.full_name || '-' }}</td>
                            <td class="text-center">{{ data.Quarterly?.quarter }} - {{ data.Quarterly?.year }}</td>
                            <td>{{ data.Activity?.name || '-' }}</td>
                            <td>{{ data.sub_activity_id || '-' }}</td>
                            <td class="text-end">{{ new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR', minimumFractionDigits: 0 }).format(data.budget) }}</td>
                        </tr>
                    </tbody>
                    <tfoot class="border-bottom-0">
                        <tr class="fw-bold">
                            <td colspan="5" class="text-end">Total Anggaran</td>
                            <td class="text-end">{{ new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR', minimumFractionDigits: 0 }).format(budgets.reduce((acc, data) => acc + data.budget, 0)) }}</td>
                        </tr>
                    </tfoot>
                </table>
            </div>
        </div>
    </div>
</template>
<script setup>
import { ref, onMounted, watch, computed } from 'vue'
import api from '@/plugins/axios'

const users = ref([])
const quarterlies = ref([])
const activities = ref([])
const budgets = ref([])

const form = ref({
    user_id: null,
    quarterly_id: null,
    activity_id: null,
    sub_activity_id: null,
    budget: null
})

const computedBudget = computed({
  get() {
    return form.value.budget === null ? '' : form.value.budget.toLocaleString('id-ID')
  },
  set(value) {
    const numeric = parseInt(value.replace(/\./g, ''), 10)
    form.value.budget = isNaN(numeric) ? null : numeric
  }
})

onMounted(async () => {
    try {
        const resUser = await api.get('/users')
        users.value = resUser.data

        const resQuarter = await api.get('/quarterlies')
        quarterlies.value = resQuarter.data

        const resActivity = await api.get('/activity')
        activities.value = resActivity.data
    } catch (error) {
        console.error('Error fetching data:', error)
    }
})

watch([() => form.value.user_id, () => form.value.quarterly_id], async ([userId, quarterlyId]) => {
    if (quarterlyId) {
        const userParam = userId ? userId : 'null'
        const resBudget = await api.get(`/budgets/user/${userParam}/quarterly/${quarterlyId}`)
        budgets.value = resBudget.data
    }
})

const createBudget = async () => {
    try {
        await api.post('/budgets', form.value)
        alert('Anggaran berhasil ditambahkan')
        const userParam = form.value.user_id ? form.value.user_id : 'null'
        const resBudget = await api.get(`/budgets/user/${userParam}/quarterly/${form.value.quarterly_id}`)
        budgets.value = resBudget.data
    } catch (error) {
        console.error('Error creating budget:', error)
        if (error.response && error.response.data && error.response.data.error) {
            alert(error.response.data.error)
        } else {
            alert('Terjadi kesalahan saat menambahkan anggaran.')
        }
    }
}
</script>