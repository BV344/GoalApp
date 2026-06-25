import { api } from './api'
import type { ScheduleSlot, CreateSlotRequest } from '@/types/schedule'

export const scheduleService = {
  getAll: () => api.get<{ slots: ScheduleSlot[] }>('/api/schedule'),
  create: (payload: CreateSlotRequest) =>
    api.post<{ slot: ScheduleSlot }>('/api/schedule', payload),
  remove: (id: number) => api.delete<void>(`/api/schedule/${id}`),
}
