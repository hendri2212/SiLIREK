<template>
    <div id="PrintArea">
        <button class="btn btn-primary mb-3 float-end no-print" @click="printReport">Print</button>
        <h6 class="text-center fw-bold">DAFTAR NAMA PEJABAT PELAKSANA TEKNIS KEGIATAN</h6>
        <h6 class="text-center fw-bold">KECAMATAN PULAULAUT SIGAM KABUPATEN KOTABARU</h6>
        <h6 class="text-center fw-bold">TAHUN ANGGARAN {{ activities[0]?.Quarterly.year }}</h6>

        <table class="table table-border-bottom-0 align-middle mt-3">
            <thead>
                <tr class="align-middle">
                    <th class="text-center">No</th>
                    <th>Nama Lengkap / NIP</th>
                    <th>Jabatan Kedinasan</th>
                    <th>Jabatan Dalam Kegiatan</th>
                    <th>Program / Kegiatan / Sub Kegiatan</th>
                    <th>Pagu Anggaran</th>
                </tr>
            </thead>
            <tbody v-if="isLoading" class="text-center my-4">
                Loading...
            </tbody>
            <tbody v-else>
                <tr v-for="(data, index) in activities" :key="index">
                    <td class="text-center">{{ index + 1 }}</td>
                    <td class="text-nowrap">
                        {{ data.User?.full_name }} <br />
                        NIP. {{ data.User?.nip }}
                    </td>
                    <td>{{ data.User.position?.name }}</td>
                    <td>{{ data.User.leader?.name }}</td>
                    <td>{{ data.Activity?.name }}</td>
                    <td class="text-end">{{ data.budget?.toLocaleString('id-ID', { style: 'currency', currency: 'IDR' }) }}</td>
                </tr>
            </tbody>
        </table>
        <div class="d-flex justify-content-end">
            <div class="mt-5">
                <p>Camat Pulaulaut Sigam</p>
                <br />
                <br />
                <br />
                <p class="fw-bold mb-0">PIA WIDYA LAKSMI P, ST</p>
                <p class="mb-0">NIP. 19840929 201001 2 006</p>
                <p>Pembina (IV/a)</p>
            </div>
        </div>
    </div>
</template>
<script setup>
import { ref, onMounted } from 'vue'
import api from '@/plugins/axios'

const activities = ref([])
const isLoading = ref(true)

onMounted(async () => {
    try {
        const response = await api.get('/budgets');
        activities.value = response.data;
    } catch (error) {
        console.error('API error:', error);
    }
    isLoading.value = false;
})

function printReport() {
    const printContents = document.getElementById('PrintArea').innerHTML;
    const printWindow = window.open('', '', 'height=600,width=800');
    printWindow.document.write('<html><head><title>Print Report</title>');
    printWindow.document.write('<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css">');
    printWindow.document.write('<style>.no-print { display: none !important; }</style>');
    printWindow.document.write('</head><body>');
    printWindow.document.write(printContents);
    printWindow.document.write('</body></html>');
    printWindow.document.close();
    printWindow.focus();
    printWindow.onload = function () {
        printWindow.print();
        printWindow.close();
    };
}
</script>