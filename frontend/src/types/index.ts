export interface FileInfo {
  name: string
  path: string
  relativePath: string
  size: number
  sizeHuman: string
  modTime: string
  order: number
}

export interface ScannerInfo {
  device: string
  vendor: string
  model: string
  type: string
  description: string
}

export interface ScanSettings {
  resolution: number
  quality: number
  format: string
  device: string
}

export interface APIResponse<T = any> {
  success: boolean
  message?: string
  data?: T
  error?: string
}

export interface BatchOperation {
  action: 'delete' | 'archive' | 'pdf'
  files: string[]
}

export interface ConvertRequest {
  files: string[]
  outputName?: string
}

export interface RenameRequest {
  oldName: string
  newName: string
}

export interface ProgressInfo {
  total: number
  current: number
  percent: number
  status: string
  fileName?: string
}

export type TabType = 'scans' | 'archives' | 'pdfs' | 'settings'

export interface Toast {
  id: number
  type: 'success' | 'error' | 'info' | 'warning'
  message: string
  duration?: number
}

export interface AppSettings {
  resolution: string
  quality: string
  format: string
  default_device: string
  theme: 'light' | 'dark' | 'system'
}
