export interface ScheduleSlot {
  id: number
  goal_id: number
  day: string
  start_time: string
  end_time: string
  note: string
  created_at: string
}

export interface CreateSlotRequest {
  goal_id?: number
  day: string
  start_time: string
  end_time: string
  note?: string
}
