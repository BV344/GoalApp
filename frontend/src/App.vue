<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import { useAuthStore } from '@/stores/authStore'
import { useRouter } from 'vue-router'

const auth = useAuthStore()
const router = useRouter()

function logout() {
  auth.logout()
  router.push('/login')
}
</script>

<template>
  <div class="min-h-screen bg-gray-50">
    <nav v-if="auth.isAuthenticated" class="bg-white shadow px-6 py-3 flex gap-6 items-center">
      <RouterLink to="/" class="font-semibold text-blue-600 hover:underline">Home</RouterLink>
      <RouterLink to="/goals" class="hover:underline">Goals</RouterLink>
      <RouterLink to="/schedule" class="hover:underline">Schedule</RouterLink>
      <button class="ml-auto text-sm text-gray-500 hover:text-red-500" @click="logout">
        Logout
      </button>
    </nav>
    <RouterView />
  </div>
</template>
