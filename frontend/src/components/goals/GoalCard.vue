<script setup lang="ts">
import type { Goal } from '@/types/goal'
defineProps<{ goal: Goal }>()
</script>

<template>
  <div class="bg-white rounded shadow p-4 flex items-center justify-between">
    <span :class="{ 'line-through text-gray-400': goal.completed }">{{ goal.title }}</span>
    <span v-if="goal.target_date" class="text-xs text-gray-400">{{ goal.target_date }}</span>
  </div>
</template>
