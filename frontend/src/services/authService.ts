import { api } from './api'
import type { AuthResponse, LoginRequest } from '@/types/auth'

export const authService = {
  login: (payload: LoginRequest) => api.post<AuthResponse>('/api/auth/login', payload),
}
