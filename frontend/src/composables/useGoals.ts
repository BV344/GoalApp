import { onMounted } from 'vue'
import { storeToRefs } from 'pinia'
import { useGoalStore } from '@/stores/goalStore'
import { goalService } from '@/services/goalService'

export function useGoals() {
  const store = useGoalStore()
  const { goals, loading, error } = storeToRefs(store)

  async function fetchGoals() {
    loading.value = true
    error.value = null
    try {
      const res = await goalService.getAll()
      store.setGoals(res.goals)
    } catch (e) {
      error.value = e instanceof Error ? e.message : 'Failed to load goals'
    } finally {
      loading.value = false
    }
  }

  onMounted(fetchGoals)

  return { goals, loading, error, fetchGoals }
}
