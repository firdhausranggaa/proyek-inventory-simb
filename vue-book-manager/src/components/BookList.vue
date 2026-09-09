<script setup>
import { ref, onMounted } from 'vue';
import api from '../api';

const books = ref([]);
const borrowings = ref([]);
const returnTransactionId = ref('');

const isLoading = ref(false);
const isActionLoading = ref(false);
const searchQuery = ref('');
const currentPage = ref(1);
const totalPages = ref(1);
let searchTimeout = null;

const onSearchInput = () => {
    clearTimeout(searchTimeout);
    searchTimeout = setTimeout(() => {
        currentPage.value = 1;
        fetchBooks();
    }, 500);
};

const fetchBooks = async () => {
    isLoading.value = true;
    try {
        const res = await api.get(`/books?search=${searchQuery.value}&page=${currentPage.value}&limit=5`);
        books.value = res.data.data;
        totalPages.value = res.data.meta.total_pages || 1;
    } catch (err) {
        console.error("Gagal mengambil data buku:", err);
    } finally {
        isLoading.value = false;
    }
};

const fetchBorrowings = async () => {
    try {
        const res = await api.get('/borrowings/me');
        borrowings.value = res.data.data;
    } catch (err) {
        console.error("Gagal mengambil riwayat:", err);
    }
};

const borrowBook = async (bookId) => {
    isActionLoading.value = true;
    try {
        const res = await api.post('/borrow', { book_id: bookId });
        alert(`Berhasil! ID Transaksi: ${res.data.data.id}`);
        fetchBooks();
        fetchBorrowings();
    } catch (err) {
        alert(err.response?.data?.error || 'Gagal meminjam buku');
    } finally {
        isActionLoading.value = false;
    }
};

const returnBook = async () => {
    if (!returnTransactionId.value) return alert("Masukkan ID Transaksi!");

    isActionLoading.value = true;
    try {
        const res = await api.post(`/return/${returnTransactionId.value}`);
        alert(res.data.message);
        returnTransactionId.value = '';
        fetchBooks();
        fetchBorrowings();
    } catch (err) {
        alert(err.response?.data?.error || 'Gagal mengembalikan buku.');
    } finally {
        isActionLoading.value = false;
    }
};

const changePage = (delta) => {
    currentPage.value += delta;
    fetchBooks();
};

onMounted(() => {
    fetchBooks();
    fetchBorrowings();
});
</script>

<template>
    <div style="padding: 20px;">
        <h2>Daftar Inventaris Buku</h2>

        <div style="margin-bottom: 15px;">
            <input v-model="searchQuery" @input="onSearchInput" placeholder="Ketik untuk mencari judul atau penulis..."
                style="padding: 10px; width: 100%; max-width: 400px; border: 1px solid #ccc; border-radius: 4px;" />
        </div>

        <div style="overflow-x: auto;">
            <table border="1" cellPadding="10"
                style="width: 100%; border-collapse: collapse; text-align: left; background: white;">
                <thead style="background-color: #2c3e50; color: white;">
                    <tr>
                        <th width="5%">ID</th>
                        <th width="35%">Judul</th>
                        <th width="30%">Penulis</th>
                        <th width="10%">Stok</th>
                        <th width="20%">Aksi</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-if="isLoading">
                        <td colspan="5" style="text-align: center; color: #666;">Memuat data...</td>
                    </tr>
                    <tr v-else-if="books.length === 0">
                        <td colspan="5" style="text-align: center; color: red;">Buku tidak ditemukan</td>
                    </tr>
                    <tr v-for="book in books" :key="book.id" v-else>
                        <td>{{ book.id }}</td>
                        <td>{{ book.title }}</td>
                        <td>{{ book.author }}</td>
                        <td>
                            <span :style="{ color: book.stock > 0 ? '#27ae60' : '#e74c3c', fontWeight: 'bold' }">
                                {{ book.stock }}
                            </span>
                        </td>
                        <td>
                            <button @click="borrowBook(book.id)" :disabled="book.stock <= 0 || isActionLoading"
                                style="padding: 6px 12px; cursor: pointer; background-color: #3498db; color: white; border: none; border-radius: 4px;">
                                {{ book.stock > 0 ? 'Pinjam Buku' : 'Stok Habis' }}
                            </button>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>

        <div
            style="margin-top: 15px; display: flex; justify-content: flex-end; align-items: center; gap: 15px; margin-bottom: 30px;">
            <button @click="changePage(-1)" :disabled="currentPage === 1 || isLoading"
                style="padding: 8px 15px; border: 1px solid #ccc; background: white; border-radius: 4px; cursor: pointer;">
                « Sebelumnya
            </button>
            <span style="font-size: 14px;">Halaman <strong>{{ currentPage }}</strong> dari {{ totalPages }}</span>
            <button @click="changePage(1)" :disabled="currentPage === totalPages || isLoading"
                style="padding: 8px 15px; border: 1px solid #ccc; background: white; border-radius: 4px; cursor: pointer;">
                Selanjutnya »
            </button>
        </div>

        <div style="display: grid; grid-template-columns: 1fr 2.5fr; gap: 20px; align-items: start;">
            <div style="padding: 20px; border: 1px solid #27ae60; border-radius: 8px; background-color: #f9fff9;">
                <h3 style="margin-top: 0; color: #2e7d32;">Loket Pengembalian</h3>
                <p style="font-size: 14px; margin-bottom: 10px;">Masukkan <strong>ID Transaksi</strong>.</p>
                <div style="display: flex; gap: 10px;">
                    <input v-model="returnTransactionId" @keyup.enter="returnBook" type="number" placeholder="ID"
                        style="padding: 8px; width: 80px; border: 1px solid #ccc; border-radius: 4px;" />
                    <button @click="returnBook" :disabled="isActionLoading"
                        style="padding: 8px 15px; cursor: pointer; background-color: #27ae60; color: white; border: none; border-radius: 4px; flex-grow: 1;">
                        {{ isActionLoading ? 'Memproses...' : 'Kembalikan' }}
                    </button>
                </div>
            </div>

            <div style="background: white; padding: 20px; border: 1px solid #eee; border-radius: 8px;">
                <h3 style="margin-top: 0;">Riwayat Transaksi Saya</h3>
                <table border="1" cellPadding="8"
                    style="width: 100%; border-collapse: collapse; text-align: left; font-size: 14px;">
                    <thead style="background-color: #ecf0f1; color: #333;">
                        <tr>
                            <th width="15%">ID Transaksi</th>
                            <th width="40%">Judul Buku</th>
                            <th width="15%">Tgl Pinjam</th>
                            <th width="15%">Tgl Kembali</th>
                            <th width="15%">Status</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-if="borrowings.length === 0">
                            <td colspan="5" style="text-align: center; color: gray;">Belum ada riwayat transaksi.</td>
                        </tr>
                        <tr v-for="b in borrowings" :key="b.id" v-else>
                            <td><strong>#{{ b.id }}</strong></td>
                            <td>{{ b.book ? b.book.title : 'Data buku dihapus' }}</td>
                            <td>{{ new Date(b.borrow_date).toLocaleDateString('id-ID') }}</td>
                            <td>{{ b.return_date ? new Date(b.return_date).toLocaleDateString('id-ID') : '-' }}</td>
                            <td>
                                <span :style="{
                                    padding: '4px 8px', borderRadius: '4px', fontSize: '12px', fontWeight: 'bold',
                                    backgroundColor: b.status === 'BORROWED' ? '#fff3cd' : '#d4edda',
                                    color: b.status === 'BORROWED' ? '#856404' : '#155724'
                                }">
                                    {{ b.status === 'BORROWED' ? 'Dipinjam' : 'Selesai' }}
                                </span>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</template>