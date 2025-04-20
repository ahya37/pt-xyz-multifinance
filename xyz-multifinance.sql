-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Waktu pembuatan: 20 Apr 2025 pada 05.59
-- Versi server: 10.4.28-MariaDB
-- Versi PHP: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `xyz-multifinance`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `konsumen`
--

CREATE TABLE `konsumen` (
  `id` int(11) NOT NULL,
  `nik` varchar(16) DEFAULT NULL,
  `full_name` varchar(255) NOT NULL,
  `legal_name` varchar(255) NOT NULL,
  `tempat_lahir` varchar(100) DEFAULT NULL,
  `tanggal_lahir` date DEFAULT NULL,
  `gaji` int(11) DEFAULT NULL,
  `foto_ktp` varchar(255) DEFAULT NULL,
  `foto_selfie` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `konsumen`
--

INSERT INTO `konsumen` (`id`, `nik`, `full_name`, `legal_name`, `tempat_lahir`, `tanggal_lahir`, `gaji`, `foto_ktp`, `foto_selfie`) VALUES
(1, '1234511890123422', 'Budi', 'Budi', 'Jakarta', '1996-01-01', 5000000, 'foto_selfie_1_1745120406804636000.jpg', 'https://example.com/budi_selfie.jpg'),
(2, '1234511890123421', 'Annisa', 'Annisa', 'Bandung', '1999-01-01', 6000000, 'https://example.com/anisa_ktp.jpg', 'https://example.com/anisa_selfie.jpg'),
(4, '1234511890123420', 'Annisa', 'Annisa', 'Bandung', '1999-01-01', 6000000, '', '');

-- --------------------------------------------------------

--
-- Struktur dari tabel `limit_konsumen`
--

CREATE TABLE `limit_konsumen` (
  `id` int(11) NOT NULL,
  `nik` varchar(16) NOT NULL,
  `tenor` int(11) NOT NULL,
  `jumlah` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `limit_konsumen`
--

INSERT INTO `limit_konsumen` (`id`, `nik`, `tenor`, `jumlah`) VALUES
(2, '1234511890123422', 1, 100000),
(3, '1234511890123422', 2, 200000),
(4, '1234511890123422', 3, 500000),
(5, '1234511890123422', 4, 700000),
(6, '1234511890123421', 4, 2000000),
(7, '1234511890123421', 3, 1500000),
(8, '1234511890123421', 2, 1200000),
(9, '1234511890123421', 1, 1000000);

-- --------------------------------------------------------

--
-- Struktur dari tabel `transaksi`
--

CREATE TABLE `transaksi` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `no_kontrak` varchar(255) NOT NULL,
  `otr` int(11) NOT NULL,
  `admin_fee` int(11) NOT NULL,
  `jumlah_cicilan` int(11) NOT NULL,
  `jumlah_bunga` int(11) NOT NULL,
  `nama_aset` varchar(255) NOT NULL,
  `konsumen_id` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `transaksi`
--

INSERT INTO `transaksi` (`id`, `no_kontrak`, `otr`, `admin_fee`, `jumlah_cicilan`, `jumlah_bunga`, `nama_aset`, `konsumen_id`, `created_at`, `updated_at`) VALUES
(1, 'KTR-20250419-001', 25000000, 500000, 12, 3000000, 'Motor Honda Beat', 1, '2025-04-18 20:50:30', '2025-04-18 20:50:30'),
(3, 'KTR-20250419-002', 25000000, 500000, 6, 3000000, 'Motor Honda Vario', 2, '2025-04-19 00:04:00', '2025-04-19 00:04:00'),
(5, 'KTR-20250419-012', 25000000, 500000, 6, 3000000, 'Motor Honda Vario', 2, '2025-04-19 00:29:50', '2025-04-19 00:29:50');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `konsumen`
--
ALTER TABLE `konsumen`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `limit_konsumen`
--
ALTER TABLE `limit_konsumen`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `transaksi`
--
ALTER TABLE `transaksi`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `no_kontrak` (`no_kontrak`),
  ADD KEY `fk_konsumen` (`konsumen_id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `konsumen`
--
ALTER TABLE `konsumen`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT untuk tabel `limit_konsumen`
--
ALTER TABLE `limit_konsumen`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT untuk tabel `transaksi`
--
ALTER TABLE `transaksi`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `transaksi`
--
ALTER TABLE `transaksi`
  ADD CONSTRAINT `fk_konsumen` FOREIGN KEY (`konsumen_id`) REFERENCES `konsumen` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
