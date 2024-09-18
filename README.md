# 🎓 Belajar Golang x Flutter 🚀

Selamat datang di repository **Belajar-Golang-x-Flutter**! 🧑‍💻 Repository ini merupakan proyek pembelajaran yang menghubungkan aplikasi frontend Flutter dengan backend Golang yang terhubung ke database.

## 📚 Deskripsi

Proyek ini terdiri dari dua bagian utama:

- **Backend (Golang):** Kode backend dibangun menggunakan Golang dan terhubung ke database untuk menyediakan API yang dapat diakses oleh aplikasi frontend. Backend ini bertanggung jawab untuk menangani data dan permintaan dari frontend.

- **Frontend (Flutter):** Kode frontend dibangun menggunakan Flutter. Aplikasi ini menampilkan data dari backend dalam format daftar barang tanpa interaksi pengguna tambahan.

## 📂 Struktur Repository

- **`apk/golang/`** - Berisi kode backend menggunakan Golang, termasuk pengaturan koneksi database dan penyediaan API.
- **`flutter/`** - Berisi kode frontend menggunakan Flutter, termasuk tampilan antarmuka pengguna dan pengambilan data dari API backend.

## ⚠️ Catatan

Pada awal pembelajaran, beberapa kode penting telah di-push ke repository. Sayangnya, commit tersebut hilang karena kesalahan dalam penggunaan command Git. 😔 Commit tersebut tidak dapat dipulihkan, dan saya sedang memperbaiki dan memperbarui repository ini agar melanjutkan pembelajaran saya.

## 🚀 Cara Memulai

1. **Clone repository ini:**
    ```bash
    git clone https://github.com/ItsukaChiyogami/Belajar-Golang-x-Flutter.git
    ```

2. **Masuk ke direktori proyek:**
    ```bash
    cd Belajar-Golang-x-Flutter
    ```

3. **Menyiapkan Backend:**
    - Navigasi ke direktori `apk/golang/`.
    - Ikuti petunjuk di `README.md` dalam direktori ini untuk mengatur dan menjalankan server backend.
    - Pastikan untuk mengonfigurasi koneksi database sesuai dengan pengaturan lokal Anda.

4. **Menyiapkan Frontend:**
    - Navigasi ke direktori `flutter/`.
    - Jalankan perintah berikut untuk menyiapkan dependensi dan menjalankan aplikasi Flutter:
      ```bash
      flutter pub get
      flutter run
      ```

5. **Integrasi:**
    - Aplikasi Flutter akan mengambil data dari API backend dan menampilkannya dalam format daftar barang. Endpoint API backend dapat ditemukan di kode backend Golang dan harus disesuaikan dengan aplikasi Flutter.

## 🛠️ Kode Kunci

### Kode Flutter

- **`lib/main.dart`**: Berisi kode utama aplikasi Flutter, termasuk pengambilan data dari API backend dan tampilan antarmuka pengguna.
- **`lib/barang.dart`**: Berisi model data `Barang` yang digunakan untuk mendekode data JSON dari API.

### Kode Golang

- **`apk/golang/`**: Berisi kode backend, pengaturan koneksi database, dan API endpoints.
