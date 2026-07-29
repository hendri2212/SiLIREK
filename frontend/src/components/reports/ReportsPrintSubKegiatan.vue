<template>
    <div>
        <!-- Toolbar (tidak ikut tercetak) -->
        <div class="no-print toolbar-bar d-flex align-items-center justify-content-between px-4 py-3 bg-white border-bottom shadow-sm">
            <div class="d-flex align-items-center gap-2">
                <button class="btn btn-outline-secondary btn-sm" @click="$router.back()">
                    <i class="bi bi-arrow-left me-1"></i> Kembali
                </button>
                <span class="text-muted small ms-2">Preview Laporan Rekening Belanja</span>
            </div>
            <button class="btn btn-primary btn-sm px-4" @click="doPrint" :disabled="isLoading">
                <i class="bi bi-printer me-1"></i> Cetak
            </button>
        </div>

        <!-- Loading -->
        <div v-if="isLoading" class="text-center py-5 no-print">
            <div class="spinner-border text-primary" role="status"></div>
            <p class="mt-3 text-muted">Memuat data laporan...</p>
        </div>

        <!-- Error -->
        <div v-else-if="hasError" class="text-center py-5 no-print">
            <i class="bi bi-exclamation-triangle fs-1 text-danger"></i>
            <p class="mt-3 text-danger">Gagal memuat data. Silakan kembali dan coba lagi.</p>
        </div>

        <!-- Area Cetak -->
        <div v-else id="print-area" class="print-area mx-auto bg-white p-5">

            <!-- Kop Surat -->
            <div class="text-center border-bottom border-2 border-dark pb-3 mb-4">
                <h5 class="fw-bold mb-1" style="font-size: 13pt;">DAFTAR REKENING BELANJA</h5>
                <h6 class="fw-bold mb-1" style="font-size: 11pt;">KECAMATAN PULAULAUT SIGAM KABUPATEN KOTABARU</h6>
                <h6 class="fw-bold mb-0" style="font-size: 11pt;">TAHUN ANGGARAN {{ currentYear }}</h6>
            </div>

            <!-- Info Kegiatan & Sub Kegiatan -->
            <table class="mb-4" style="border: none; font-size: 11pt;">
                <tbody>
                    <tr>
                        <td class="fw-bold pe-2" style="width: 140px; border: none; padding: 2px 0;">Kegiatan</td>
                        <td style="border: none; padding: 2px 0;">: {{ activityCode }} - {{ activityName }}</td>
                    </tr>
                    <tr>
                        <td class="fw-bold pe-2" style="border: none; padding: 2px 0;">Sub Kegiatan</td>
                        <td style="border: none; padding: 2px 0;">: {{ subCode }} - {{ subName }}</td>
                    </tr>
                </tbody>
            </table>

            <!-- Tabel Rekening -->
            <table class="report-table w-100">
                <thead>
                    <tr>
                        <th style="width: 4%; text-align: center;">No</th>
                        <th style="width: 18%;">Kode Rekening</th>
                        <th>Uraian Rekening</th>
                        <th style="width: 17%; text-align: right;">Pagu Anggaran (Rp)</th>
                        <th style="width: 17%; text-align: right;">Penggunaan (Rp)</th>
                        <th style="width: 17%; text-align: right;">Sisa Anggaran (Rp)</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-if="accounts.length === 0">
                        <td colspan="6" class="text-center text-muted py-4">Belum ada data rekening belanja.</td>
                    </tr>
                    <tr v-else v-for="(acc, idx) in accounts" :key="acc.id">
                        <td style="text-align: center;">{{ idx + 1 }}</td>
                        <td>{{ acc.code || '-' }}</td>
                        <td>{{ acc.description || '-' }}</td>
                        <td style="text-align: right;">{{ formatRupiah(acc.budget_ceiling) }}</td>
                        <td style="text-align: right;">{{ formatRupiah(acc.total_credit) }}</td>
                        <td style="text-align: right;">{{ formatRupiah(acc.remaining_budget) }}</td>
                    </tr>
                </tbody>
                <tfoot v-if="accounts.length > 0">
                    <tr class="total-row">
                        <td colspan="3" style="text-align: right; font-weight: bold;">TOTAL</td>
                        <td style="text-align: right; font-weight: bold;">{{ formatRupiah(totalPagu) }}</td>
                        <td style="text-align: right; font-weight: bold;">{{ formatRupiah(totalPenggunaan) }}</td>
                        <td style="text-align: right; font-weight: bold;">{{ formatRupiah(totalSisa) }}</td>
                    </tr>
                </tfoot>
            </table>

            <!-- Tanda Tangan -->
            <div class="d-flex justify-content-end mt-5 pt-3">
                <div class="text-center" style="width: 280px;">
                    <p class="mb-0" style="font-size: 11pt;">Camat Pulaulaut Sigam</p>
                    <div style="height: 70px;"></div>
                    <p class="fw-bold text-decoration-underline mb-0" style="font-size: 11pt;">PIA WIDYA LAKSMI P, ST</p>
                    <p class="mb-0" style="font-size: 10pt;">NIP. 19840929 201001 2 006</p>
                    <p style="font-size: 10pt;">Pembina (IV/a)</p>
                </div>
            </div>

        </div>
    </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import api from '@/plugins/axios'

const route = useRoute()

const accounts   = ref([])
const isLoading  = ref(true)
const hasError   = ref(false)

// Info kegiatan & sub kegiatan disimpan dari query params
const activityCode = ref('')
const activityName = ref('')
const subCode      = ref('')
const subName      = ref('')

const currentYear = new Date().getFullYear()

// ── Computed totals ──────────────────────────────────────────
const totalPagu       = computed(() => accounts.value.reduce((s, a) => s + (a.budget_ceiling  || 0), 0))
const totalPenggunaan = computed(() => accounts.value.reduce((s, a) => s + (a.total_credit    || 0), 0))
const totalSisa       = computed(() => accounts.value.reduce((s, a) => s + (a.remaining_budget|| 0), 0))

// ── Helpers ──────────────────────────────────────────────────
function formatRupiah(val) {
    if (val == null) return 'Rp 0'
    return Number(val).toLocaleString('id-ID', { style: 'currency', currency: 'IDR' })
}

function doPrint() {
    // Set nama file print = nama sub kegiatan
    const originalTitle = document.title
    document.title = subName.value
        ? `Rekening - ${subCode.value} ${subName.value}`
        : 'Daftar Rekening Belanja'

    window.print()

    // Restore title setelah dialog print ditutup
    setTimeout(() => {
        document.title = originalTitle
    }, 1000)
}

// ── Fetch data ───────────────────────────────────────────────
onMounted(async () => {
    const subId = route.params.subId

    // Ambil info dari query params (dikirim saat navigasi)
    activityCode.value = route.query.activityCode || ''
    activityName.value = route.query.activityName || ''
    subCode.value      = route.query.subCode      || ''
    subName.value      = route.query.subName      || ''

    try {
        const response = await api.get(`/expenditure-account?sub_activity_id=${subId}`)
        accounts.value = response.data || []
    } catch (err) {
        console.error('Fetch error:', err)
        hasError.value = true
    } finally {
        isLoading.value = false
    }
})
</script>

<style>
/* ── Toolbar ─────────────────────────── */
.toolbar-bar {
    position: sticky;
    top: 0;
    z-index: 100;
}

/* ── Preview area (layar) ───────────────────────── */
.print-area {
    max-width: 1000px;
    font-family: 'Times New Roman', Times, serif;
    font-size: 11pt;
    min-height: 100vh;
}

/* ── Tabel laporan ──────────────────────────────── */
.report-table {
    border-collapse: collapse;
    width: 100%;
}
.report-table th,
.report-table td {
    border: 1px solid #000;
    padding: 6px 10px;
    font-size: 11pt;
}
.report-table thead th {
    background-color: #e8e8e8;
    text-align: center;
    font-weight: bold;
}
.report-table .total-row td {
    background-color: #f0f0f0;
}

/* ── Print media ────────────────────────────────── */
@media print {
    @page {
        size: A4 landscape;
        margin: 1.5cm;
    }

    /* Sembunyikan toolbar halaman ini saat cetak */
    .no-print {
        display: none !important;
    }

    /* Sembunyikan navbar/header aplikasi dari App.vue */
    header,
    nav,
    .navbar {
        display: none !important;
    }

    /* Reset container App.vue agar tidak ada padding/margin ekstra */
    .container {
        max-width: 100% !important;
        padding: 0 !important;
        margin: 0 !important;
    }

    /* Hapus shadow / border halaman preview saat cetak */
    .print-area {
        max-width: 100% !important;
        padding: 0 !important;
        box-shadow: none !important;
        border: none !important;
        min-height: unset !important;
    }

    body {
        background: #fff !important;
    }
}
</style>
