<script setup>
import { ref } from 'vue';
import Login from './components/Login.vue';
import BookList from './components/BookList.vue';

const isAuthenticated = ref(!!localStorage.getItem('token'));

const onLoginSuccess = () => {
  isAuthenticated.value = true;
};

const handleLogout = () => {
  localStorage.removeItem('token');
  localStorage.removeItem('role');
  isAuthenticated.value = false;
};
</script>

<template>
  <div>
    <header v-if="isAuthenticated"
      style="padding: 15px 20px; background: #2c3e50; text-align: right; box-shadow: 0 2px 4px rgba(0,0,0,0.1);">
      <button @click="handleLogout"
        style="padding: 8px 15px; cursor: pointer; background-color: #e74c3c; color: white; border: none; border-radius: 4px; font-weight: bold;">Logout</button>
    </header>

    <Login v-if="!isAuthenticated" @login-success="onLoginSuccess" />
    <BookList v-else />
  </div>
</template>