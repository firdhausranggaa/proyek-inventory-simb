<script setup>
import { ref } from 'vue';
import api from '../api';

const emit = defineEmits(['loginSuccess']);
const username = ref('');
const password = ref('');
const error = ref('');

const handleLogin = async () => {
    try {
        const res = await api.post('/login', {
            username: username.value,
            password: password.value
        });
        localStorage.setItem('token', res.data.token);
        localStorage.setItem('role', res.data.role);
        emit('loginSuccess');
    } catch (err) {
        error.value = err.response?.data?.error || 'Login gagal';
    }
};
</script>

<template>
    <div style="max-width: 300px; margin: 50px auto; padding: 20px; border: 1px solid #ccc; border-radius: 8px;">
        <h2 style="text-align: center;">Login SIMB</h2>
        <p v-if="error" style="color: red; text-align: center;">{{ error }}</p>
        <input v-model="username" placeholder="Username"
            style="display: block; margin-bottom: 15px; width: 100%; padding: 8px; box-sizing: border-box;" />
        <input v-model="password" type="password" placeholder="Password"
            style="display: block; margin-bottom: 15px; width: 100%; padding: 8px; box-sizing: border-box;" />
        <button @click="handleLogin"
            style="width: 100%; padding: 10px; background-color: #4CAF50; color: white; border: none; border-radius: 4px; cursor: pointer;">Masuk</button>
    </div>
</template>