<script setup lang="ts">
import { useAppStore } from '../stores/app'

const store = useAppStore()

const downloadFile = (relativePath: string) => {
  window.open(relativePath, '_blank')
}
</script>

<template>
  <div class="h-full">
    <!-- Loading State -->
    <div v-if="store.isLoading && store.archives.length === 0" class="flex items-center justify-center h-64">
      <div class="spinner"></div>
    </div>
    
    <!-- Empty State -->
    <div v-else-if="store.archives.length === 0" class="flex flex-col items-center justify-center h-64 text-gray-500 dark:text-gray-400">
      <svg class="w-16 h-16 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 8h14M5 8a2 2 0 110-4h14a2 2 0 110 4M5 8v10a2 2 0 002 2h10a2 2 0 002-2V8m-9 4h4" />
      </svg>
      <p class="text-lg">Нет архивов</p>
      <p class="text-sm mt-2">Архивируйте сканы для скачивания</p>
    </div>
    
    <!-- List -->
    <div v-else class="space-y-2">
      <div
        v-for="file in store.archives"
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
        <div class="w-12 h-12 bg-blue-100 dark:bg-blue-900 rounded-lg flex items-center justify-center flex-shrink-0">
          <svg class="w-6 h-6 text-blue-600 dark:text-blue-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 8h14M5 8a2 2 0 110-4h14a2 2 0 110 4M5 8v10a2 2 0 002 2h10a2 2 0 002-2V8m-9 4h4" />
          </svg>
        </div>
        
        <!-- Info -->
        <div class="flex-1 min-w-0">
          <p class="font-medium text-gray-800 dark:text-gray-200 truncate">{{ file.name }}</p>
          <p class="text-sm text-gray-500 dark:text-gray-400">{{ file.sizeHuman }}</p>
        </div>
        
        <!-- Actions -->
        <button
          @click="downloadFile(file.relativePath)"
          class="btn btn-primary flex items-center gap-2"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4" />
          </svg>
          <span class="hidden sm:inline">Скачать</span>
        </button>
      </div>
    </div>
  </div>
</template>
