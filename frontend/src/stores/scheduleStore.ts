import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { ScheduleSlot } from '@/types/schedule'

export const useScheduleStore = defineStore('schedule', () => {
  const slots = ref<ScheduleSlot[]>([])
  const loading = ref(false)
  const error = ref<string | null>(null)

  function setSlots(data: ScheduleSlot[]) {
    slots.value = data
  }

  return { slots, loading, error, setSlots }
})
