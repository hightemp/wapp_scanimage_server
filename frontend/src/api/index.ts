import axios from 'axios'
import type { 
  APIResponse, 
  FileInfo, 
  ScannerInfo, 
  ScanSettings, 
  BatchOperation, 
  ConvertRequest,
  AppSettings 
} from '../types'

const api = axios.create({
  baseURL: '/api',
  headers: {
    'Content-Type': 'application/json'
  }
})

// Scanner API
export const scannerApi = {
  getScanners: () => 
    api.get<APIResponse<ScannerInfo[]>>('/scanners'),
  
  getScannersRaw: () => 
    api.get<APIResponse<string>>('/scanners/raw'),
  
  scan: (settings?: Partial<ScanSettings>) => 
    api.post<APIResponse<{ fileName: string }>>('/scan', settings || {})
}

// Scans API
export const scansApi = {
  getAll: () => 
    api.get<APIResponse<FileInfo[]>>('/scans'),
  
  delete: (name: string) => 
    api.delete<APIResponse>(`/scans/${encodeURIComponent(name)}`),
  
  deleteAll: () => 
    api.delete<APIResponse>('/scans'),
  
  rename: (oldName: string, newName: string) => 
    api.put<APIResponse>(`/scans/${encodeURIComponent(oldName)}`, { newName }),
  
  batch: (operation: BatchOperation) => 
    api.post<APIResponse<{ fileName?: string }>>('/scans/batch', operation),
  
  updateOrder: (orders: Record<string, number>) => 
    api.put<APIResponse>('/scans/order', orders)
}

// Archives API
export const archivesApi = {
  getAll: () => 
    api.get<APIResponse<FileInfo[]>>('/archives'),
  
  create: (files: string[], outputName?: string) => 
    api.post<APIResponse<{ fileName: string }>>('/archives', { files, outputName }),
  
  createAll: () => 
    api.post<APIResponse<{ fileName: string }>>('/archives/all'),
  
  delete: (name: string) => 
    api.delete<APIResponse>(`/archives/${encodeURIComponent(name)}`),
  
  deleteAll: () => 
    api.delete<APIResponse>('/archives')
}

// PDFs API
export const pdfsApi = {
  getAll: () => 
    api.get<APIResponse<FileInfo[]>>('/pdfs'),
  
  create: (request: ConvertRequest) => 
    api.post<APIResponse<{ fileName: string }>>('/pdfs', request),
  
  createAll: () => 
    api.post<APIResponse<{ fileName: string }>>('/pdfs/all'),
  
  delete: (name: string) => 
    api.delete<APIResponse>(`/pdfs/${encodeURIComponent(name)}`),
  
  deleteAll: () => 
    api.delete<APIResponse>('/pdfs')
}

// Settings API
export const settingsApi = {
  getAll: () => 
    api.get<APIResponse<AppSettings>>('/settings'),
  
  update: (settings: Partial<AppSettings>) => 
    api.put<APIResponse>('/settings', settings),
  
  get: (key: string) => 
    api.get<APIResponse<Record<string, string>>>(`/settings/${key}`),
  
  set: (key: string, value: string) => 
    api.put<APIResponse>(`/settings/${key}`, { value })
}

export default api
