<script setup lang="ts">
import { useSchedule } from '@/composables/useSchedule'

const { slots, loading, error } = useSchedule()
</script>

<template>
  <main class="p-8">
    <h1 class="text-3xl font-bold mb-6">Weekly Schedule</h1>
    <p v-if="loading" class="text-gray-400">Loading...</p>
    <p v-else-if="error" class="text-red-500">{{ error }}</p>
    <p v-else-if="slots.length === 0" class="text-gray-400">No scheduled slots yet.</p>
  </main>
</template>
