<template>
    <div class="mb-5">
        <div class="d-flex justify-content-between align-items-center mb-4 no-print">
            <h5 class="fw-bold mb-0">Laporan Rincian Item Rekening</h5>
            <div>
                <router-link :to="{ name: 'reports.kegiatan' }" class="btn btn-outline-secondary me-2">
                    <i class="bi bi-arrow-left"></i> Kembali
                </router-link>
                <button class="btn btn-primary" @click="printReport">
                    <i class="bi bi-printer"></i> Print
                </button>
            </div>
        </div>

        <div v-if="isLoading" class="text-center my-5">
            <div class="spinner-border text-primary" role="status">
                <span class="visually-hidden">Loading...</span>
            </div>
        </div>

        <div v-else id="PrintArea" class="bg-white p-4 p-md-5 border rounded shadow-sm">
            <!-- Kop Surat / Header Laporan -->
            <div class="text-center mb-4 border-bottom pb-3">
                <h6 class="fw-bold fs-5 mb-1">DAFTAR RINCIAN ITEM REKENING BELANJA</h6>
                <h6 class="fw-bold mb-1">KECAMATAN PULAULAUT SIGAM KABUPATEN KOTABARU</h6>
                <h6 class="fw-bold mb-0">TAHUN ANGGARAN {{ currentYear }}</h6>
            </div>

            <!-- Informasi Rekening -->
            <div class="mb-4">
                <table class="table-borderless">
                    <tbody>
                        <tr>
                            <td class="fw-bold pe-3" style="width: 150px;">Kode Rekening</td>
                            <td>: {{ account?.code || '-' }}</td>
                        </tr>
                        <tr>
                            <td class="fw-bold pe-3">Uraian Rekening</td>
                            <td>: {{ account?.description || '-' }}</td>
                        </tr>
                        <tr>
                            <td class="fw-bold pe-3">Pagu Anggaran</td>
                            <td>: {{ account?.budget_ceiling?.toLocaleString('id-ID', { style: 'currency', currency: 'IDR' }) }}</td>
                        </tr>
                    </tbody>
                </table>
            </div>

            <!-- Tabel Formal Item -->
            <table class="table table-bordered align-middle">
                <thead class="table-light text-center">
                    <tr>
                        <th width="5%">No</th>
                        <th width="15%">Tanggal</th>
                        <th width="20%">No. Bukti / Kode</th>
                        <th width="40%">Uraian Pembayaran</th>
                        <th width="20%">Jumlah (Rp)</th>
                    </tr>
                </thead>
                <tbody v-if="items.length === 0">
                    <tr>
                        <td colspan="5" class="text-center text-muted py-4">Tidak ada data item untuk rekening ini.</td>
                    </tr>
                </tbody>
                <tbody v-else>
                    <tr v-for="(item, index) in items" :key="item.id">
                        <td class="text-center">{{ index + 1 }}</td>
                        <td class="text-center">{{ formatDate(item.date) }}</td>
                        <td>{{ item.code }}</td>
                        <td>{{ item.description }}</td>
                        <td class="text-end">{{ item.credit?.toLocaleString('id-ID', { minimumFractionDigits: 2 }) }}</td>
                    </tr>
                </tbody>
                <tfoot v-if="items.length > 0" class="table-light fw-bold">
                    <tr>
                        <td colspan="4" class="text-end">TOTAL :</td>
                        <td class="text-end">{{ totalAmount.toLocaleString('id-ID', { style: 'currency', currency: 'IDR' }) }}</td>
                    </tr>
                    <tr>
                        <td colspan="4" class="text-end">SISA PAGU ANGGARAN :</td>
                        <td class="text-end">{{ remainingBudget.toLocaleString('id-ID', { style: 'currency', currency: 'IDR' }) }}</td>
                    </tr>
                </tfoot>
            </table>

            <!-- Tanda Tangan -->
            <div class="d-flex justify-content-end mt-5 pt-3">
                <div class="text-center" style="width: 300px;">
                    <p class="mb-5">Camat Pulaulaut Sigam</p>
                    <br /><br />
                    <p class="fw-bold text-decoration-underline mb-0">PIA WIDYA LAKSMI P, ST</p>
                    <p class="mb-0">NIP. 19840929 201001 2 006</p>
                    <p>Pembina (IV/a)</p>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRoute } from 'vue-router'
import api from '@/plugins/axios'

const route = useRoute()
const items = ref([])
const account = ref(null)
const isLoading = ref(true)

const currentYear = computed(() => {
    return new Date().getFullYear();
})

const totalAmount = computed(() => {
    return items.value.reduce((sum, item) => sum + (item.credit || 0), 0)
})

const remainingBudget = computed(() => {
    if (!account.value) return 0;
    return (account.value.budget_ceiling || 0) - totalAmount.value;
})

const formatDate = (dateString) => {
    if (!dateString) return '-';
    const date = new Date(dateString);
    return date.toLocaleDateString('id-ID', { day: '2-digit', month: 'long', year: 'numeric' });
}

onMounted(async () => {
    const accountId = route.params.id;
    if (!accountId) return;

    try {
        // Ambil info Rekening
        const accResponse = await api.get(`/expenditure-account/${accountId}`);
        account.value = accResponse.data;

        // Ambil daftar Item berdasarkan Rekening
        const itemsResponse = await api.get(`/item?expenditure_account_id=${accountId}`);
        items.value = itemsResponse.data;
    } catch (error) {
        console.error('API error:', error);
    } finally {
        isLoading.value = false;
    }
})

function printReport() {
    const printContents = document.getElementById('PrintArea').innerHTML;
    const printWindow = window.open('', '', 'height=800,width=1000');
    printWindow.document.write('<html><head><title>Print Laporan Item</title>');
    printWindow.document.write('<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css">');
    printWindow.document.write(`
        <style>
            body { font-family: 'Times New Roman', Times, serif; font-size: 12pt; padding: 20px; }
            .table-borderless td, .table-borderless th { border: none; padding: 4px 8px; }
            .table-bordered th, .table-bordered td { border: 1px solid #000 !important; }
            .no-print { display: none !important; }
            @media print {
                @page { margin: 1.5cm; }
            }
        </style>
    `);
    printWindow.document.write('</head><body>');
    printWindow.document.write(printContents);
    printWindow.document.write('</body></html>');
    printWindow.document.close();
    printWindow.focus();
    
    // Memberikan sedikit jeda agar CSS dimuat
    setTimeout(() => {
        printWindow.print();
        printWindow.close();
    }, 500);
}
</script>

<style scoped>
@media print {
    .no-print {
        display: none !important;
    }
}
</style>
