<script setup lang="ts">
import { useAppStore } from '../stores/app'

const store = useAppStore()

const openPdf = (relativePath: string) => {
  window.open(relativePath, '_blank')
}
</script>

<template>
  <div class="h-full">
    <!-- Loading State -->
    <div v-if="store.isLoading && store.pdfs.length === 0" class="flex items-center justify-center h-64">
      <div class="spinner"></div>
    </div>
    
    <!-- Empty State -->
    <div v-else-if="store.pdfs.length === 0" class="flex flex-col items-center justify-center h-64 text-gray-500 dark:text-gray-400">
      <svg class="w-16 h-16 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 21h10a2 2 0 002-2V9.414a1 1 0 00-.293-.707l-5.414-5.414A1 1 0 0012.586 3H7a2 2 0 00-2 2v14a2 2 0 002 2z" />
      </svg>
      <p class="text-lg">No PDF files</p>
      <p class="text-sm mt-2">Convert scans to PDF</p>
    </div>
    
    <!-- List -->
    <div v-else class="space-y-2">
      <div
        v-for="file in store.pdfs"
        :key="file.name"
        class="card p-4 flex items-center gap-4 hover:shadow-md transition-shadow"
      >
        <!-- Checkbox -->
        <div
          class="cursor-pointer"
          @click="store.toggleSelect(file.name)"
        >
          <div
            class="w-6 h-6 rounded border-2 flex items-center justify-center transition-colors"
            :class="[
              store.selectedFiles.has(file.name)
                ? 'bg-primary-600 border-primary-600' 
                : 'border-gray-300 dark:border-gray-600 hover:border-primary-500'
            ]"
          >
            <svg v-if="store.selectedFiles.has(file.name)" class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M5 13l4 4L19 7" />
            </svg>
          </div>
        </div>
        
        <!-- Icon -->
        <div class="w-12 h-12 bg-red-100 dark:bg-red-900 rounded-lg flex items-center justify-center flex-shrink-0">
          <svg class="w-6 h-6 text-red-600 dark:text-red-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21h10a2 2 0 002-2V9.414a1 1 0 00-.293-.707l-5.414-5.414A1 1 0 0012.586 3H7a2 2 0 00-2 2v14a2 2 0 002 2z" />
          </svg>
        </div>
        
        <!-- Info -->
        <div class="flex-1 min-w-0">
          <p class="font-medium text-gray-800 dark:text-gray-200 truncate">{{ file.name }}</p>
          <p class="text-sm text-gray-500 dark:text-gray-400">{{ file.sizeHuman }}</p>
        </div>
        
        <!-- Actions -->
        <button
          @click="openPdf(file.relativePath)"
          class="btn btn-primary flex items-center gap-2"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
          </svg>
          <span class="hidden sm:inline">Open</span>
        </button>
        
        <a
          :href="file.relativePath"
          :download="file.name"
          class="btn btn-secondary flex items-center gap-2"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4" />
          </svg>
          <span class="hidden sm:inline">Download</span>
        </a>
      </div>
    </div>
  </div>
</template>
