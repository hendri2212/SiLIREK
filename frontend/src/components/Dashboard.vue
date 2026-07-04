<template>
    <div class="mb-5">
        <!-- Welcome Banner -->
        <div class="card shadow-sm rounded-4 border-0 bg-primary text-white mb-4 overflow-hidden position-relative">
            <div class="position-absolute top-0 end-0 opacity-10 p-4" style="transform: scale(2.5) translate(-10%, 10%);">
                <i class="bi bi-wallet2" style="font-size: 15rem;"></i>
            </div>
            <div class="card-body p-4 p-md-5 position-relative z-1">
                <div class="row align-items-center">
                    <div class="col-md-8">
                        <h3 class="fw-bold mb-2">Selamat Datang di SiLIREK</h3>
                        <p class="fs-5 text-white-50 mb-4">Sistem Pengendalian Realisasi Keuangan Kecamatan Pulaulaut Sigam</p>
                        <p class="mb-0">Kelola anggaran, catat transaksi pengeluaran, dan hasilkan laporan realisasi dengan mudah dan transparan.</p>
                    </div>
                </div>
            </div>
        </div>

        <div class="row g-4 mb-4">
            <!-- Program -->
            <div class="col-md-4">
                <router-link :to="{ name: 'program.list' }" class="text-decoration-none">
                    <div class="card shadow-sm h-100 border-0 rounded-4 hover-lift">
                        <div class="card-body p-4 text-center">
                            <div class="bg-primary bg-opacity-10 text-primary rounded-circle d-inline-flex p-3 mb-3">
                                <i class="bi bi-folder-fill fs-2"></i>
                            </div>
                            <h5 class="fw-bold text-dark">Program</h5>
                            <p class="text-muted small mb-0">Kelola data induk program tingkat atas.</p>
                        </div>
                    </div>
                </router-link>
            </div>

            <!-- Kegiatan & Transaksi -->
            <div class="col-md-4">
                <router-link :to="{ name: 'activities.list' }" class="text-decoration-none">
                    <div class="card shadow-sm h-100 border-0 rounded-4 hover-lift">
                        <div class="card-body p-4 text-center">
                            <div class="bg-info bg-opacity-10 text-info rounded-circle d-inline-flex p-3 mb-3">
                                <i class="bi bi-file-earmark-text-fill fs-2"></i>
                            </div>
                            <h5 class="fw-bold text-dark">Kegiatan & Transaksi</h5>
                            <p class="text-muted small mb-0">Kelola hierarki kegiatan hingga pencatatan item belanja.</p>
                        </div>
                    </div>
                </router-link>
            </div>

            <!-- Laporan -->
            <div class="col-md-4">
                <router-link :to="{ name: 'reports.kegiatan' }" class="text-decoration-none">
                    <div class="card shadow-sm h-100 border-0 rounded-4 hover-lift">
                        <div class="card-body p-4 text-center">
                            <div class="bg-success bg-opacity-10 text-success rounded-circle d-inline-flex p-3 mb-3">
                                <i class="bi bi-printer-fill fs-2"></i>
                            </div>
                            <h5 class="fw-bold text-dark">Laporan Realisasi</h5>
                            <p class="text-muted small mb-0">Hasilkan laporan pohon kegiatan dan rincian transaksi.</p>
                        </div>
                    </div>
                </router-link>
            </div>
        </div>

        <div class="row g-4 mb-4">
            <!-- Tech Stack Update -->
            <div class="col-md-12">
                <div class="card shadow-sm h-100 border-0 rounded-4 bg-light">
                    <div class="card-body p-4">
                        <h6 class="fw-bold text-uppercase text-muted mb-3 fs-7">Sistem Informasi</h6>
                        <ul class="list-unstyled mb-0">
                            <li class="mb-2"><i class="bi bi-check-circle-fill text-success me-2"></i> <span class="fw-medium">Frontend:</span> Vue 3 & Bootstrap 5</li>
                            <li class="mb-2"><i class="bi bi-check-circle-fill text-success me-2"></i> <span class="fw-medium">Backend:</span> Golang (Gin Framework)</li>
                            <li class="mb-2"><i class="bi bi-check-circle-fill text-success me-2"></i> <span class="fw-medium">Database:</span> MySQL (GORM)</li>
                            <li><i class="bi bi-check-circle-fill text-success me-2"></i> <span class="fw-medium">Hak Akses:</span> Role-based (Superadmin & Admin)</li>
                        </ul>
                    </div>
                </div>
            </div>
        </div>

        <div class="text-center mt-5 text-muted small pb-3">
            <p>Dikembangkan oleh <span class="fw-bold text-dark">Hendri Arifin</span> | Versi 2.0.0</p>
        </div>
    </div>
</template>

<script setup>
</script>

<style scoped>
.hover-lift {
    transition: transform 0.2s ease, box-shadow 0.2s ease;
}
.hover-lift:hover {
    transform: translateY(-5px);
    box-shadow: 0 .5rem 1rem rgba(0,0,0,.15)!important;
}
.fs-7 {
    font-size: 0.85rem;
}
</style>