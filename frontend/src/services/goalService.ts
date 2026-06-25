import { api } from './api'
import type { Goal, CreateGoalRequest } from '@/types/goal'

export const goalService = {
  getAll: () => api.get<{ goals: Goal[] }>('/api/goals'),
  getOne: (id: number) => api.get<{ goal: Goal }>(`/api/goals/${id}`),
  create: (payload: CreateGoalRequest) => api.post<{ goal: Goal }>('/api/goals', payload),
  update: (id: number, payload: Partial<CreateGoalRequest>) =>
    api.put<{ goal: Goal }>(`/api/goals/${id}`, payload),
  remove: (id: number) => api.delete<void>(`/api/goals/${id}`),
}
