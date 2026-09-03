<script setup>
import { ref, onMounted } from 'vue';
import api from '../api';

const books = ref([]);
const borrowings = ref([]);
const searchQuery = ref('');
const returnTransactionId = ref('');

const fetchBooks = async () => {
    try {
        const res = await api.get(`/books?search=${searchQuery.value}`);
        books.value = res.data.data;
    } catch (err) {
        console.error("Gagal mengambil data buku:", err);
    }
};

// Fungsi baru untuk menarik data riwayat peminjaman dari endpoint
const fetchBorrowings = async () => {
    try {
        const res = await api.get('/borrowings/me');
        borrowings.value = res.data.data;
    } catch (err) {
        console.error("Gagal mengambil riwayat peminjaman:", err);
    }
};

const borrowBook = async (bookId) => {
    try {
        const res = await api.post('/borrow', { book_id: bookId });
        alert(`${res.data.message}\nID Transaksi Anda: ${res.data.data.id}`);
        fetchBooks();
        fetchBorrowings();
    } catch (err) {
        alert(err.response?.data?.error || 'Gagal meminjam buku');
    }
};

const returnBook = async () => {
    if (!returnTransactionId.value) {
        alert("Masukkan ID Transaksi terlebih dahulu!");
        return;
    }
    try {
        const res = await api.post(`/return/${returnTransactionId.value}`);
        alert(res.data.message);
        returnTransactionId.value = '';
        fetchBooks();
        fetchBorrowings();
    } catch (err) {
        alert(err.response?.data?.error || 'Gagal mengembalikan buku. Pastikan ID Transaksi benar.');
    }
};

// Memanggil kedua API secara bersamaan saat halaman pertama kali dimuat
onMounted(() => {
    fetchBooks();
    fetchBorrowings();
});
</script>

<template>
    <div style="padding: 20px;">
        <h2>Daftar Inventaris Buku</h2>

        <!-- Fitur Pencarian -->
        <div style="margin-bottom: 20px;">
            <input v-model="searchQuery" @keyup.enter="fetchBooks" placeholder="Cari judul buku atau penulis..."
                style="padding: 8px; width: 300px; margin-right: 10px; border: 1px solid #ccc; border-radius: 4px;" />
            <button @click="fetchBooks"
                style="padding: 8px 15px; cursor: pointer; background-color: #2c3e50; color: white; border: none; border-radius: 4px;">Cari</button>
            <button @click="searchQuery = ''; fetchBooks()"
                style="padding: 8px 15px; cursor: pointer; margin-left: 5px; border: 1px solid #ccc; background: #fff; border-radius: 4px;">Reset</button>
        </div>

        <!-- Tabel Daftar Buku -->
        <table border="1" cellPadding="8"
            style="width: 100%; border-collapse: collapse; text-align: left; margin-bottom: 30px;">
            <thead style="background-color: #2c3e50; color: white;">
                <tr>
                    <th>ID</th>
                    <th>Judul</th>
                    <th>Penulis</th>
                    <th>Stok</th>
                    <th>Aksi</th>
                </tr>
            </thead>
            <tbody>
                <tr v-if="books.length === 0">
                    <td colspan="5" style="text-align: center; color: red;">Buku tidak ditemukan</td>
                </tr>
                <tr v-for="book in books" :key="book.id">
                    <td>{{ book.id }}</td>
                    <td>{{ book.title }}</td>
                    <td>{{ book.author }}</td>
                    <td>{{ book.stock }}</td>
                    <td>
                        <button @click="borrowBook(book.id)" :disabled="book.stock <= 0"
                            style="padding: 5px 10px; cursor: pointer; background-color: #3498db; color: white; border: none; border-radius: 4px;">
                            {{ book.stock > 0 ? 'Pinjam' : 'Habis' }}
                        </button>
                    </td>
                </tr>
            </tbody>
        </table>

        <!-- Layout Grid untuk Loket & Riwayat -->
        <div style="display: grid; grid-template-columns: 1fr 2fr; gap: 20px; align-items: start;">

            <!-- Loket Pengembalian -->
            <div style="padding: 20px; border: 1px solid #4CAF50; border-radius: 8px; background-color: #f9fff9;">
                <h3 style="margin-top: 0; color: #2e7d32;">Loket Pengembalian</h3>
                <p style="font-size: 14px; margin-bottom: 10px;">Masukkan <strong>ID Transaksi Peminjaman</strong> Anda.
                </p>
                <div style="display: flex; gap: 10px;">
                    <input v-model="returnTransactionId" type="number" placeholder="Contoh: 1"
                        style="padding: 8px; width: 100px; border: 1px solid #ccc; border-radius: 4px;" />
                    <button @click="returnBook"
                        style="padding: 8px 15px; cursor: pointer; background-color: #4CAF50; color: white; border: none; border-radius: 4px; flex-grow: 1;">Kembalikan</button>
                </div>
            </div>

            <!-- Tabel Riwayat Peminjaman -->
            <div>
                <h3 style="margin-top: 0;">Riwayat Peminjaman Saya</h3>
                <table border="1" cellPadding="8"
                    style="width: 100%; border-collapse: collapse; text-align: left; font-size: 14px;">
                    <thead style="background-color: #ecf0f1; color: #333;">
                        <tr>
                            <th>ID Transaksi</th>
                            <th>ID Buku</th>
                            <th>Tgl Pinjam</th>
                            <th>Tgl Kembali</th>
                            <th>Status</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-if="borrowings.length === 0">
                            <td colspan="5" style="text-align: center; color: gray;">Belum ada riwayat transaksi.</td>
                        </tr>
                        <tr v-for="b in borrowings" :key="b.id">
                            <td><strong>{{ b.id }}</strong></td>
                            <td>{{ b.book_id }}</td>
                            <td>{{ new Date(b.borrow_date).toLocaleDateString('id-ID') }}</td>
                            <td>{{ b.return_date ? new Date(b.return_date).toLocaleDateString('id-ID') : '-' }}</td>
                            <td>
                                <span :style="{
                                    padding: '3px 8px',
                                    borderRadius: '12px',
                                    fontSize: '12px',
                                    backgroundColor: b.status === 'BORROWED' ? '#fff3cd' : '#d4edda',
                                    color: b.status === 'BORROWED' ? '#856404' : '#155724'
                                }">
                                    {{ b.status === 'BORROWED' ? 'Dipinjam' : 'Dikembalikan' }}
                                </span>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>

        </div>
    </div>
</template>