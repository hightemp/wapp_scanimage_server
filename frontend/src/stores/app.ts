import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { FileInfo, ScannerInfo, TabType, Toast, AppSettings } from '../types'
import { scannerApi, scansApi, archivesApi, pdfsApi, settingsApi } from '../api'

export const useAppStore = defineStore('app', () => {
  // State
  const scans = ref<FileInfo[]>([])
  const archives = ref<FileInfo[]>([])
  const pdfs = ref<FileInfo[]>([])
  const scanners = ref<ScannerInfo[]>([])
  const scannersRaw = ref<string>('')
  const settings = ref<AppSettings>({
    resolution: '300',
    quality: '80',
    format: 'jpeg',
    default_device: '',
    theme: 'light'
  })
  
  const activeTab = ref<TabType>('scans')
  const selectedFiles = ref<Set<string>>(new Set())
  const isLoading = ref(false)
  const isScanning = ref(false)
  const toasts = ref<Toast[]>([])
  const imagePreview = ref<string | null>(null)
  
  let toastId = 0

  // Computed
  const selectedCount = computed(() => selectedFiles.value.size)
  const hasSelection = computed(() => selectedFiles.value.size > 0)
  const allSelected = computed(() => {
    const currentFiles = activeTab.value === 'scans' ? scans.value : 
                        activeTab.value === 'archives' ? archives.value : pdfs.value
    return currentFiles.length > 0 && currentFiles.every(f => selectedFiles.value.has(f.name))
  })

  const isDarkMode = computed(() => {
    if (settings.value.theme === 'system') {
      return window.matchMedia('(prefers-color-scheme: dark)').matches
    }
    return settings.value.theme === 'dark'
  })

  // Actions
  const addToast = (type: Toast['type'], message: string, duration = 3000) => {
    const id = ++toastId
    toasts.value.push({ id, type, message, duration })
    if (duration > 0) {
      setTimeout(() => removeToast(id), duration)
    }
    return id
  }

  const removeToast = (id: number) => {
    const index = toasts.value.findIndex(t => t.id === id)
    if (index !== -1) {
      toasts.value.splice(index, 1)
    }
  }

  const fetchScanners = async () => {
    try {
      const [response, rawResponse] = await Promise.all([
        scannerApi.getScanners(),
        scannerApi.getScannersRaw()
      ])
      if (response.data.success && response.data.data) {
        scanners.value = response.data.data
      }
      if (rawResponse.data.success && rawResponse.data.data) {
        scannersRaw.value = rawResponse.data.data
      }
    } catch (error) {
      addToast('error', 'Failed to get scanner list')
    }
  }

  const fetchScans = async () => {
    try {
      isLoading.value = true
      const response = await scansApi.getAll()
      if (response.data.success && response.data.data) {
        scans.value = response.data.data
      }
    } catch (error) {
      addToast('error', 'Failed to load scans')
    } finally {
      isLoading.value = false
    }
  }

  const fetchArchives = async () => {
    try {
      isLoading.value = true
      const response = await archivesApi.getAll()
      if (response.data.success && response.data.data) {
        archives.value = response.data.data
      }
    } catch (error) {
      addToast('error', 'Failed to load archives')
    } finally {
      isLoading.value = false
    }
  }

  const fetchPdfs = async () => {
    try {
      isLoading.value = true
      const response = await pdfsApi.getAll()
      if (response.data.success && response.data.data) {
        pdfs.value = response.data.data
      }
    } catch (error) {
      addToast('error', 'Failed to load PDF files')
    } finally {
      isLoading.value = false
    }
  }

  const fetchSettings = async () => {
    try {
      const response = await settingsApi.getAll()
      if (response.data.success && response.data.data) {
        settings.value = { ...settings.value, ...response.data.data }
        applyTheme()
      }
    } catch (error) {
      console.error('Failed to fetch settings:', error)
    }
  }

  const updateSettings = async (newSettings: Partial<AppSettings>) => {
    try {
      await settingsApi.update(newSettings)
      settings.value = { ...settings.value, ...newSettings }
      applyTheme()
      addToast('success', 'Settings saved')
    } catch (error) {
      addToast('error', 'Failed to save settings')
    }
  }

  const applyTheme = () => {
    const isDark = settings.value.theme === 'dark' || 
      (settings.value.theme === 'system' && window.matchMedia('(prefers-color-scheme: dark)').matches)
    
    document.documentElement.classList.toggle('dark', isDark)
  }

  const scan = async () => {
    try {
      isScanning.value = true
      const response = await scannerApi.scan({
        resolution: parseInt(settings.value.resolution),
        quality: parseInt(settings.value.quality),
        format: settings.value.format,
        device: settings.value.default_device
      })
      if (response.data.success) {
        addToast('success', 'Scan completed')
        await fetchScans()
      } else {
        addToast('error', response.data.error || 'Scan failed')
      }
    } catch (error: any) {
      addToast('error', error.response?.data?.error || 'Scan failed')
    } finally {
      isScanning.value = false
    }
  }

  const toggleSelect = (fileName: string) => {
    if (selectedFiles.value.has(fileName)) {
      selectedFiles.value.delete(fileName)
    } else {
      selectedFiles.value.add(fileName)
    }
  }

  const selectAll = () => {
    const currentFiles = activeTab.value === 'scans' ? scans.value : 
                        activeTab.value === 'archives' ? archives.value : pdfs.value
    currentFiles.forEach(f => selectedFiles.value.add(f.name))
  }

  const deselectAll = () => {
    selectedFiles.value.clear()
  }

  const toggleSelectAll = () => {
    if (allSelected.value) {
      deselectAll()
    } else {
      selectAll()
    }
  }

  const deleteSelected = async () => {
    if (!hasSelection.value) return
    
    try {
      isLoading.value = true
      const files = Array.from(selectedFiles.value)
      
      if (activeTab.value === 'scans') {
        await scansApi.batch({ action: 'delete', files })
        await fetchScans()
      } else if (activeTab.value === 'archives') {
        for (const file of files) {
          await archivesApi.delete(file)
        }
        await fetchArchives()
      } else if (activeTab.value === 'pdfs') {
        for (const file of files) {
          await pdfsApi.delete(file)
        }
        await fetchPdfs()
      }
      
      selectedFiles.value.clear()
      addToast('success', 'Files deleted')
    } catch (error) {
      addToast('error', 'Failed to delete files')
    } finally {
      isLoading.value = false
    }
  }

  const deleteAll = async () => {
    try {
      isLoading.value = true
      
      if (activeTab.value === 'scans') {
        await scansApi.deleteAll()
        await fetchScans()
      } else if (activeTab.value === 'archives') {
        await archivesApi.deleteAll()
        await fetchArchives()
      } else if (activeTab.value === 'pdfs') {
        await pdfsApi.deleteAll()
        await fetchPdfs()
      }
      
      selectedFiles.value.clear()
      addToast('success', 'All files deleted')
    } catch (error) {
      addToast('error', 'Failed to delete files')
    } finally {
      isLoading.value = false
    }
  }

  const archiveSelected = async () => {
    if (!hasSelection.value) return
    
    try {
      isLoading.value = true
      const files = Array.from(selectedFiles.value)
      const response = await archivesApi.create(files)
      if (response.data.success) {
        addToast('success', 'Archive created')
        await fetchArchives()
      }
    } catch (error) {
      addToast('error', 'Failed to create archive')
    } finally {
      isLoading.value = false
    }
  }

  const archiveAll = async () => {
    try {
      isLoading.value = true
      const response = await archivesApi.createAll()
      if (response.data.success) {
        addToast('success', 'Archive created')
        await fetchArchives()
      }
    } catch (error) {
      addToast('error', 'Failed to create archive')
    } finally {
      isLoading.value = false
    }
  }

  const convertToPdfSelected = async () => {
    if (!hasSelection.value) return
    
    try {
      isLoading.value = true
      const files = Array.from(selectedFiles.value)
      const response = await pdfsApi.create({ files })
      if (response.data.success) {
        addToast('success', 'PDF created')
        await fetchPdfs()
      }
    } catch (error) {
      addToast('error', 'Failed to create PDF')
    } finally {
      isLoading.value = false
    }
  }

  const convertToPdfAll = async () => {
    try {
      isLoading.value = true
      const response = await pdfsApi.createAll()
      if (response.data.success) {
        addToast('success', 'PDF created')
        await fetchPdfs()
      }
    } catch (error) {
      addToast('error', 'Failed to create PDF')
    } finally {
      isLoading.value = false
    }
  }

  const updateFileOrder = async (newOrder: FileInfo[]) => {
    const orders: Record<string, number> = {}
    newOrder.forEach((file, index) => {
      orders[file.name] = index
    })
    
    try {
      await scansApi.updateOrder(orders)
      scans.value = newOrder
    } catch (error) {
      addToast('error', 'Failed to save order')
    }
  }

  const setImagePreview = (path: string | null) => {
    imagePreview.value = path
  }

  const refresh = async () => {
    selectedFiles.value.clear()
    await Promise.all([
      fetchScans(),
      fetchArchives(),
      fetchPdfs()
    ])
  }

  // Initialize
  const init = async () => {
    await Promise.all([
      fetchScanners(),
      fetchScans(),
      fetchArchives(),
      fetchPdfs(),
      fetchSettings()
    ])
  }

  return {
    // State
    scans,
    archives,
    pdfs,
    scanners,
    scannersRaw,
    settings,
    activeTab,
    selectedFiles,
    isLoading,
    isScanning,
    toasts,
    imagePreview,
    
    // Computed
    selectedCount,
    hasSelection,
    allSelected,
    isDarkMode,
    
    // Actions
    addToast,
    removeToast,
    fetchScanners,
    fetchScans,
    fetchArchives,
    fetchPdfs,
    fetchSettings,
    updateSettings,
    scan,
    toggleSelect,
    selectAll,
    deselectAll,
    toggleSelectAll,
    deleteSelected,
    deleteAll,
    archiveSelected,
    archiveAll,
    convertToPdfSelected,
    convertToPdfAll,
    updateFileOrder,
    setImagePreview,
    refresh,
    init
  }
})
