<script setup lang="ts">
import type { ScheduleSlot } from '@/types/schedule'
defineProps<{ slots: ScheduleSlot[] }>()

const days = ['monday', 'tuesday', 'wednesday', 'thursday', 'friday', 'saturday', 'sunday']
</script>

<template>
  <div class="grid grid-cols-7 gap-2">
    <div v-for="day in days" :key="day" class="text-center">
      <div class="font-semibold capitalize mb-2">{{ day.slice(0, 3) }}</div>
      <div
        v-for="slot in slots.filter(s => s.day === day)"
        :key="slot.id"
        class="bg-blue-100 rounded p-1 text-xs mb-1"
      >
        {{ slot.start_time }} – {{ slot.end_time }}
      </div>
    </div>
  </div>
</template>
