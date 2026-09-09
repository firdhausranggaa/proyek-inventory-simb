<script setup>
import { ref } from 'vue';
import api from '../api';

const emit = defineEmits(['loginSuccess']);
const username = ref('');
const password = ref('');
const error = ref('');
const isLoading = ref(false);

const handleLogin = async () => {
    if (!username.value || !password.value) return;

    isLoading.value = true;
    error.value = '';

    try {
        const res = await api.post('/login', {
            username: username.value,
            password: password.value
        });
        localStorage.setItem('token', res.data.token);
        localStorage.setItem('role', res.data.role);
        emit('loginSuccess');
    } catch (err) {
        error.value = err.response?.data?.error || 'Koneksi ke server gagal.';
    } finally {
        isLoading.value = false;
    }
};
</script>

<template>
    <div
        style="max-width: 320px; margin: 80px auto; padding: 30px; border-radius: 10px; background: white; box-shadow: 0 4px 6px rgba(0,0,0,0.1);">
        <h2 style="text-align: center; margin-top: 0; color: #2c3e50;">Login SIMB</h2>

        <div v-if="error"
            style="background-color: #f8d7da; color: #721c24; padding: 10px; border-radius: 4px; margin-bottom: 15px; font-size: 14px; text-align: center;">
            {{ error }}
        </div>

        <input v-model="username" @keyup.enter="handleLogin" placeholder="Username"
            style="display: block; margin-bottom: 15px; width: 100%; padding: 10px; box-sizing: border-box; border: 1px solid #ccc; border-radius: 4px;" />

        <input v-model="password" type="password" @keyup.enter="handleLogin" placeholder="Password"
            style="display: block; margin-bottom: 20px; width: 100%; padding: 10px; box-sizing: border-box; border: 1px solid #ccc; border-radius: 4px;" />

        <button @click="handleLogin" :disabled="isLoading"
            style="width: 100%; padding: 12px; background-color: #27ae60; color: white; border: none; border-radius: 4px; cursor: pointer; font-weight: bold;">
            {{ isLoading ? 'Memverifikasi...' : 'Masuk Sistem' }}
        </button>
    </div>
</template>