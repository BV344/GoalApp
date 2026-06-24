export interface Goal {
  id: number
  title: string
  description: string
  target_date: string
  completed: boolean
  created_at: string
  updated_at: string
}

export interface CreateGoalRequest {
  title: string
  description?: string
  target_date?: string
}
