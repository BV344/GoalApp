import { onMounted } from 'vue'
import { storeToRefs } from 'pinia'
import { useScheduleStore } from '@/stores/scheduleStore'
import { scheduleService } from '@/services/scheduleService'

export function useSchedule() {
  const store = useScheduleStore()
  const { slots, loading, error } = storeToRefs(store)

  async function fetchSlots() {
    loading.value = true
    error.value = null
    try {
      const res = await scheduleService.getAll()
      store.setSlots(res.slots)
    } catch (e) {
      error.value = e instanceof Error ? e.message : 'Failed to load schedule'
    } finally {
      loading.value = false
    }
  }

  onMounted(fetchSlots)

  return { slots, loading, error, fetchSlots }
}
