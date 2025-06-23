<template>
    <div class="card shadow-sm rounded-4 border-0">
        <div class="card-body">
            <div class="mb-4 d-flex align-items-center justify-content-between">
                <h4 class="card-title text-info">Penganggaran</h4>
            </div>
            <form @submit.prevent="createOrUpdateBudget">
                <div class="row">
                    <div class="col">
                        <div class="form-floating mb-3">
                            <select v-model.number="form.quarterly_id" name="quarter" class="form-control"
                                id="floatingQuarter" required>
                                <option value="" selected disabled>-- Pilih Quarter --</option>
                                <option v-for="data in quarterlies" :key="data.id" :value="data.id">
                                    {{ data.quarter }} - {{ data.year }}
                                </option>
                            </select>
                            <label for="floatingQuarter">Triwulan</label>
                        </div>
                    </div>
                    <div class="col">
                        <div class="form-floating mb-3">
                            <select v-model.number="form.user_id" name="user" class="form-control" id="floatingUser">
                                <option value="" selected disabled>-- Pilih Pengguna --</option>
                                <option v-for="data in users" :key="data.id" :value="data.id">
                                    {{ data.full_name }}
                                </option>
                            </select>
                            <label for="floatingQuarter">Nama Pengguna</label>
                        </div>
                    </div>
                    <div class="col">
                        <div class="form-floating mb-3">
                            <input v-model="computedBudget" type="text" class="form-control" id="floatingBudget"
                                placeholder="Jumlah Anggaran">
                            <label for="floatingBudget">Jumlah Anggaran</label>
                        </div>
                    </div>
                </div>
                <div class="row">
                    <div class="col">
                        <div class="form-floating mb-3">
                            <select v-model.number="form.activity_id" name="activity" class="form-control"
                                id="floatingActivity" required>
                                <option value="" selected disabled>-- Pilih Kegiatan --</option>
                                <option v-for="data in activities" :key="data.id" :value="data.id">
                                    {{ data.code }} - {{ data.name }}
                                </option>
                            </select>
                            <label for="floatingPosition">Nama Kegiatan</label>
                        </div>
                    </div>
                    <div class="col">
                        <div class="form-floating mb-3">
                            <select v-model.number="form.sub_activity_id" name="sub_activity" class="form-control"
                                id="floatingActivity">
                                <option value="" selected disabled>-- Pilih Sub Kegiatan --</option>
                                <option v-for="data in activities" :key="data.id" :value="data.id">
                                    {{ data.code }} - {{ data.name }}
                                </option>
                            </select>
                            <label for="floatingPosition">Nama Sub Kegiatan</label>
                        </div>
                    </div>
                    <div class="col d-flex justify-content-end align-items-center">
                        <button class="btn btn-info text-white" type="submit">
                            <i class="bi bi-floppy me-1"></i>
                            Anggaran
                        </button>
                        <button 
                            v-if="isEditing" 
                            type="button" 
                            class="btn btn-secondary ms-2" 
                            @click="cancelEdit">
                            Cancel
                        </button>
                    </div>
                </div>
            </form>
        </div>
    </div>
    <div class="card shadow-sm rounded-4 border-0 mt-2">
        <div class="card-body">
            <!-- <div class="mb-4 d-flex align-items-center justify-content-between">
                <h4 class="card-title text-info">Penganggaran</h4>
            </div> -->
            <div class="table-responsive">
                <table class="table">
                    <thead class="text-nowrap">
                        <tr>
                            <th class="text-center" scope="col"></th>
                            <th scope="col">Nama Pengguna</th>
                            <th scope="col">Triwulan</th>
                            <th scope="col">Kegiatan</th>
                            <th scope="col">Pagu Anggaran</th>
                            <th class="text-center" scope="col">Actions</th>
                        </tr>
                    </thead>
                    <tbody class="text-nowrap">
                        <template v-for="(data, index) in budgetTree" :key="data.id">
                            <BudgetRow :data="data" :level="0" @edit-budget="startEdit" />
                        </template>
                    </tbody>
                    <tfoot class="border-bottom-0">
                        <tr class="fw-bold">
                            <td colspan="5" class="text-end">Total Anggaran</td>
                            <td class="text-end">{{new Intl.NumberFormat('id-ID', {
                                style: 'currency', currency: 'IDR',
                                minimumFractionDigits: 0
                            }).format(budgets.reduce((acc, data) => acc + data.budget, 0))
                                }}</td>
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
import BudgetRow from './BudgetRow.vue'

const users = ref([])
const quarterlies = ref([])
const activities = ref([])
const budgets = ref([])

// --- Add these new refs for editing state ---
const isEditing = ref(false)
const editingId = ref(null)

const form = ref({
    id: null,
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

const budgetTree = computed(() => {
    const map = new Map()
    const roots = []

    budgets.value.forEach(b => {
        b.children = []
        map.set(b.id, b)
    })

    budgets.value.forEach(b => {
        const parent = map.get(b.parent_id)
        if (parent) {
            parent.children.push(b)
        } else {
            roots.push(b)
        }
    })

    return roots
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

// --- Add startEdit method ---
const startEdit = (rowData) => {
    isEditing.value = true
    editingId.value = rowData.id
    // Populate form fields with the selected row's data
    form.value.user_id = rowData.user_id
    form.value.quarterly_id = rowData.quarterly_id
    form.value.activity_id = rowData.activity_id
    form.value.sub_activity_id = rowData.sub_activity_id
    form.value.budget = rowData.budget
}

// --- Add cancelEdit method ---
const cancelEdit = () => {
    isEditing.value = false
    editingId.value = null
    form.value.user_id = null
    form.value.quarterly_id = null
    form.value.activity_id = null
    form.value.sub_activity_id = null
    form.value.budget = null
}

// Update createBudget to handle both create and update:
const createOrUpdateBudget = async () => {
    try {
        if (isEditing.value && editingId.value) {
            // Update existing budget
            await api.put(`/budgets/${editingId.value}`, form.value)
            alert('Anggaran berhasil diubah')
            isEditing.value = false
            editingId.value = null
        } else {
            // Create new budget
            await api.post('/budgets', form.value)
            alert('Anggaran berhasil ditambahkan')
        }
        // Refresh list after either action
        const userParam = form.value.user_id ? form.value.user_id : 'null'
        const resBudget = await api.get(`/budgets/user/${userParam}/quarterly/${form.value.quarterly_id}`)
        budgets.value = resBudget.data
        // Clear form after update
        if (!isEditing.value) {
            form.value.user_id = null
            form.value.quarterly_id = null
            form.value.activity_id = null
            form.value.sub_activity_id = null
            form.value.budget = null
        }
    } catch (error) {
        console.error('Error saving budget:', error)
        if (error.response && error.response.data && error.response.data.error) {
            alert(error.response.data.error)
        } else {
            alert('Terjadi kesalahan saat menyimpan anggaran.')
        }
    }
}
</script>