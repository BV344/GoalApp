<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/authStore'
import { authService } from '@/services/authService'

const username = ref('')
const password = ref('')
const error = ref('')
const router = useRouter()
const auth = useAuthStore()

async function login() {
  error.value = ''
  try {
    const res = await authService.login({ username: username.value, password: password.value })
    auth.setToken(res.token)
    router.push('/')
  } catch {
    error.value = 'Invalid credentials'
  }
}
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50">
    <form class="bg-white p-8 rounded-lg shadow w-full max-w-sm" @submit.prevent="login">
      <h1 class="text-2xl font-bold mb-6">GoalApp</h1>
      <p v-if="error" class="text-red-500 text-sm mb-4">{{ error }}</p>
      <input
        v-model="username"
        type="text"
        placeholder="Username"
        class="w-full border rounded px-3 py-2 mb-4 focus:outline-none focus:ring-2 focus:ring-blue-500"
      />
      <input
        v-model="password"
        type="password"
        placeholder="Password"
        class="w-full border rounded px-3 py-2 mb-6 focus:outline-none focus:ring-2 focus:ring-blue-500"
      />
      <button
        type="submit"
        class="w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700 transition"
      >
        Sign In
      </button>
    </form>
  </div>
</template>
