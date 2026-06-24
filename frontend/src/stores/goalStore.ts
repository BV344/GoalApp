import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { Goal } from '@/types/goal'

export const useGoalStore = defineStore('goals', () => {
  const goals = ref<Goal[]>([])
  const loading = ref(false)
  const error = ref<string | null>(null)

  function setGoals(data: Goal[]) {
    goals.value = data
  }

  return { goals, loading, error, setGoals }
})
