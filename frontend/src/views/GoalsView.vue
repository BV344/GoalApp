<script setup lang="ts">
import { useGoals } from '@/composables/useGoals'

const { goals, loading, error } = useGoals()
</script>

<template>
  <main class="p-8">
    <h1 class="text-3xl font-bold mb-6">Goals</h1>
    <p v-if="loading" class="text-gray-400">Loading...</p>
    <p v-else-if="error" class="text-red-500">{{ error }}</p>
    <ul v-else class="space-y-2">
      <li
        v-for="goal in goals"
        :key="goal.id"
        class="bg-white p-4 rounded shadow flex items-center justify-between"
      >
        <span :class="{ 'line-through text-gray-400': goal.completed }">{{ goal.title }}</span>
      </li>
      <li v-if="goals.length === 0" class="text-gray-400">No goals yet.</li>
    </ul>
  </main>
</template>
