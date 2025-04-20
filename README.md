# XYZ Multifinance API

## Deskripsi

API ini menyediakan layanan untuk mengelola data konsumen, transaksi pembiayaan, dan proses upload dokumen KTP untuk aplikasi XYZ Multifinance.

## Arsitektur Aplikasi

Aplikasi ini mengadopsi arsitektur **Clean Architecture** atau **Layered Architecture** yang bertujuan untuk memisahkan tanggung jawab dan dependensi antar lapisan. Berikut adalah lapisan utama dalam aplikasi ini:

1.  **Presentation Layer (Controller):**
    * Bertanggung jawab untuk menerima request dari klien (melalui HTTP).
    * Melakukan validasi dasar terhadap request.
    * Meneruskan request ke lapisan Service.
    * Menerima response dari lapisan Service dan mengembalikannya ke klien dalam format JSON.

2.  **Application Layer (Service):**
    * Berisi logika bisnis utama aplikasi.
    * Menerima request dari Controller.
    * Melakukan validasi bisnis dan orkestrasi operasi.
    * Berinteraksi dengan lapisan Domain dan Infrastructure (Repository).
    * Tidak memiliki dependensi langsung pada framework atau database.

3.  **Domain Layer (Entity/Model):**
    * Berisi objek-objek bisnis (entitas) dan aturan bisnis yang terkait.
    * Merepresentasikan data dan perilaku inti aplikasi.
    * Tidak memiliki dependensi pada lapisan lain.

4.  **Infrastructure Layer (Repository, Database):**
    * Bertanggung jawab untuk interaksi dengan sumber data (database).
    * Implementasi dari interface Repository yang didefinisikan di lapisan Domain atau Application.
    * Menggunakan library atau framework khusus database (misalnya, `database/sql` di Go).

5.  **Helper Layer:**
    * Berisi fungsi-fungsi utilitas yang digunakan di berbagai lapisan (misalnya, response formatting, error handling, data validation).

**Alur Request:**

1.  Klien mengirimkan request HTTP ke salah satu endpoint Controller.
2.  Controller menerima request, melakukan validasi dasar, dan meneruskannya ke Service yang sesuai.
3.  Service menerima request, melakukan validasi bisnis, dan berinteraksi dengan Repository untuk mengambil atau menyimpan data.
4.  Repository berkomunikasi dengan database.
5.  Data dikembalikan melalui Repository ke Service.
6.  Service memproses data dan mengembalikannya ke Controller.
7.  Controller memformat response dan mengirimkannya kembali ke klien.

## Struktur Database

Berikut adalah struktur tabel database utama yang digunakan oleh aplikasi ini.

### Tabel `konsumen`

| Kolom           | Tipe Data        | Constraints       | Deskripsi                                  |
| --------------- | ---------------- | ----------------- | ------------------------------------------ |
| `id`            | `INT`            | `PRIMARY KEY`, `AUTO_INCREMENT` | ID unik konsumen                             |
| `nik`           | `VARCHAR(255)`   | `UNIQUE`, `NOT NULL` | Nomor Induk Kependudukan konsumen            |
| `full_name`     | `VARCHAR(255)`   | `NOT NULL`        | Nama lengkap konsumen                        |
| `legal_name`    | `VARCHAR(255)`   |                   | Nama sesuai dokumen hukum konsumen          |
| `tempat_lahir`  | `VARCHAR(255)`   |                   | Tempat lahir konsumen                        |
| `tanggal_lahir` | `DATE`           |                   | Tanggal lahir konsumen                       |
| `gaji`          | `DECIMAL(10, 2)` |                   | Pendapatan bulanan konsumen                |
| `foto_ktp`      | `VARCHAR(255)`   |                   | Nama file foto KTP konsumen                  |
| `foto_selfie`   | `VARCHAR(255)`   |                   | Nama file foto selfie konsumen               |


### Tabel `limit_konsumen`

| Kolom   | Tipe Data    | Constraints       | Deskripsi                                      |
| ------- | ------------ | ----------------- | ---------------------------------------------- |
| `id`    | `INT(11)`    | `PRIMARY KEY`, `AUTO_INCREMENT` | ID unik pengajuan kredit                      |
| `nik`   | `VARCHAR(16)`| `NOT NULL`        | Nomor Induk Kependudukan konsumen (FK ke `konsumen`) |
| `tenor` | `INT(11)`    | `NOT NULL`        | Jangka waktu (tenor) pengajuan kredit (dalam bulan) |
| `jumlah`| `INT(11)`    | `NULL`            | Jumlah dana yang diajukan                     |


### Tabel `transaksi`

| Kolom           | Tipe Data        | Constraints       | Deskripsi                                    |
| --------------- | ---------------- | ----------------- | -------------------------------------------- |
| `id`            | `INT`            | `PRIMARY KEY`, `AUTO_INCREMENT` | ID unik transaksi                             |
| `konsumen_id`   | `INT`            | `NOT NULL`, `FOREIGN KEY` (`konsumen_id`) REFERENCES `konsumen`(`id`) |
| `no_kontrak`    | `VARCHAR(255)`   | `UNIQUE`, `NOT NULL` | Nomor kontrak transaksi                      |
| `otr`           | `DECIMAL(12, 2)` | `NOT NULL`        | Harga On The Road aset                       |
| `admin_fee`     | `DECIMAL(10, 2)` |                   | Biaya administrasi transaksi               |
| `jumlah_cicilan`| `INT`            | `NOT NULL`        | Jumlah cicilan transaksi                     |
| `jumlah_bunga`  | `DECIMAL(8, 2)`  | `NOT NULL`        | Total bunga transaksi                        |
| `nama_aset`     | `VARCHAR(255)`   | `NOT NULL`        | Nama aset yang dibiayai                     |
| `created_at`    | `TIMESTAMP`      | `DEFAULT CURRENT_TIMESTAMP` | Waktu pembuatan catatan                     |
| `updated_at`    | `TIMESTAMP`      | `DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP` | Waktu terakhir pembaruan 

**Catatan:**

* Struktur database di atas adalah contoh dasar dan mungkin berbeda tergantung pada kebutuhan aplikasi Anda yang lebih spesifik.
* Anda mungkin memiliki tabel lain untuk data pengguna, konfigurasi, atau data pendukung lainnya.
* Pastikan Anda telah membuat database dan menjalankan migrasi atau skrip SQL untuk membuat tabel-tabel ini.

## Cara Menjalankan Aplikasi

1.  **Prerequisites:**
    * Go (versi terbaru yang disarankan)
    * MySQL atau database lain yang didukung
    * [Optional] Postman atau alat HTTP client lainnya untuk menguji API

2.  **Konfigurasi Database:**
    * Buat database dengan nama yang Anda inginkan.
    * Konfigurasi koneksi database di file konfigurasi aplikasi Anda (misalnya, di dalam fungsi `app.NewDB()` di `app/database.go`). Ini biasanya melibatkan pengaturan host, port, username, password, dan nama database.

3.  **Install Dependencies:**
    ```bash
    go mod tidy
    go mod vendor
    ```

4.  **Jalankan Aplikasi:**
    ```bash
    go run main.go
    ```

    Aplikasi akan berjalan di `http://localhost:3000` (atau port lain yang Anda konfigurasi).

## Endpoint API

Berikut adalah beberapa endpoint API yang tersedia:

### Konsumen

* `GET /api/konsumen`: Mendapatkan daftar semua konsumen.
* `GET /api/konsumen/{konsumenId}`: Mendapatkan detail konsumen berdasarkan ID.
* `POST /api/konsumen`: Membuat konsumen baru.
* `PUT /api/konsumen/{konsumenId}`: Memperbarui informasi konsumen berdasarkan ID.
* `DELETE /api/konsumen/{konsumenId}`: Menghapus konsumen berdasarkan ID.
* `GET /api/konsumenlimits`: Mendapatkan informasi limit konsumen

### Limit Konsumen

* `POST /api/limit-konsumen`: Membuat limit konsumen baru.

### Transaksi

* `POST /api/transaksi`: Membuat transaksi baru.
* `GET /api/transaksikonsumen`: Mendapatkan daftar transaksi beserta informasi konsumen terkait.

### Upload Dokumen

* `POST /api/upload/ktp/{konsumenId}`: Mengupload file foto KTP untuk konsumen dengan ID tertentu.
* `POST /api/upload/fotoselfie/{konsumenId}`: Mengupload file foto selfie untuk konsumen dengan ID tertentu.

---

**XYZ Multifinance Developer Team**