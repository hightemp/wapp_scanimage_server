<script setup lang="ts">
import { ref } from 'vue'
import { useAppStore } from '../stores/app'

const store = useAppStore()
const showActions = ref(false)

const handleScan = () => {
  store.scan()
}

const handleRefresh = () => {
  store.refresh()
}
</script>

<template>
  <div class="bg-white dark:bg-gray-800 shadow-sm border-b border-gray-200 dark:border-gray-700 p-4">
    <div class="flex flex-wrap items-center gap-3">
      <!-- Scan Button -->
      <button
        v-if="store.activeTab === 'scans'"
        @click="handleScan"
        :disabled="store.isScanning"
        class="btn btn-success flex items-center gap-2"
      >
        <svg v-if="!store.isScanning" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 9a2 2 0 012-2h.93a2 2 0 001.664-.89l.812-1.22A2 2 0 0110.07 4h3.86a2 2 0 011.664.89l.812 1.22A2 2 0 0018.07 7H19a2 2 0 012 2v9a2 2 0 01-2 2H5a2 2 0 01-2-2V9z" />
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 13a3 3 0 11-6 0 3 3 0 016 0z" />
        </svg>
        <div v-else class="spinner w-5 h-5 border-2"></div>
        <span>{{ store.isScanning ? 'Сканирование...' : 'Сканировать' }}</span>
      </button>
      
      <!-- Refresh Button -->
      <button
        @click="handleRefresh"
        :disabled="store.isLoading"
        class="btn btn-outline flex items-center gap-2"
      >
        <svg class="w-5 h-5" :class="{ 'animate-spin': store.isLoading }" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
        </svg>
        <span class="hidden sm:inline">Обновить</span>
      </button>
      
      <!-- Selection Info -->
      <div v-if="store.hasSelection" class="flex items-center gap-2 text-sm text-gray-600 dark:text-gray-400">
        <span>Выбрано: {{ store.selectedCount }}</span>
      </div>
      
      <!-- Spacer -->
      <div class="flex-1"></div>
      
      <!-- Actions for Scans -->
      <template v-if="store.activeTab === 'scans'">
        <!-- Select All -->
        <button
          @click="store.toggleSelectAll"
          class="btn btn-outline flex items-center gap-2"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-6 9l2 2 4-4" />
          </svg>
          <span class="hidden sm:inline">{{ store.allSelected ? 'Снять' : 'Выбрать все' }}</span>
        </button>
        
        <!-- Actions Dropdown -->
        <div class="relative">
          <button
            @click="showActions = !showActions"
            class="btn btn-secondary flex items-center gap-2"
          >
            <span>Действия</span>
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
            </svg>
          </button>
          
          <div
            v-if="showActions"
            class="absolute right-0 mt-2 w-56 bg-white dark:bg-gray-800 rounded-lg shadow-lg border border-gray-200 dark:border-gray-700 py-1 z-50"
            @click="showActions = false"
          >
            <button
              @click="store.deleteSelected"
              :disabled="!store.hasSelection"
              class="w-full px-4 py-2 text-left text-sm hover:bg-gray-100 dark:hover:bg-gray-700 disabled:opacity-50 disabled:cursor-not-allowed flex items-center gap-2"
            >
              <svg class="w-4 h-4 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
              </svg>
              Удалить выделенное
            </button>
            
            <button
              @click="store.deleteAll"
              class="w-full px-4 py-2 text-left text-sm hover:bg-gray-100 dark:hover:bg-gray-700 flex items-center gap-2 text-red-600"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
              </svg>
              Удалить все
            </button>
            
            <hr class="my-1 border-gray-200 dark:border-gray-700">
            
            <button
              @click="store.archiveSelected"
              :disabled="!store.hasSelection"
              class="w-full px-4 py-2 text-left text-sm hover:bg-gray-100 dark:hover:bg-gray-700 disabled:opacity-50 disabled:cursor-not-allowed flex items-center gap-2"
            >
              <svg class="w-4 h-4 text-blue-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 8h14M5 8a2 2 0 110-4h14a2 2 0 110 4M5 8v10a2 2 0 002 2h10a2 2 0 002-2V8m-9 4h4" />
              </svg>
              Архивировать выделенное
            </button>
            
            <button
              @click="store.archiveAll"
              class="w-full px-4 py-2 text-left text-sm hover:bg-gray-100 dark:hover:bg-gray-700 flex items-center gap-2"
            >
              <svg class="w-4 h-4 text-blue-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 8h14M5 8a2 2 0 110-4h14a2 2 0 110 4M5 8v10a2 2 0 002 2h10a2 2 0 002-2V8m-9 4h4" />
              </svg>
              Архивировать все
            </button>
            
            <hr class="my-1 border-gray-200 dark:border-gray-700">
            
            <button
              @click="store.convertToPdfSelected"
              :disabled="!store.hasSelection"
              class="w-full px-4 py-2 text-left text-sm hover:bg-gray-100 dark:hover:bg-gray-700 disabled:opacity-50 disabled:cursor-not-allowed flex items-center gap-2"
            >
              <svg class="w-4 h-4 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21h10a2 2 0 002-2V9.414a1 1 0 00-.293-.707l-5.414-5.414A1 1 0 0012.586 3H7a2 2 0 00-2 2v14a2 2 0 002 2z" />
              </svg>
              Конвертировать в PDF выделенное
            </button>
            
            <button
              @click="store.convertToPdfAll"
              class="w-full px-4 py-2 text-left text-sm hover:bg-gray-100 dark:hover:bg-gray-700 flex items-center gap-2"
            >
              <svg class="w-4 h-4 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21h10a2 2 0 002-2V9.414a1 1 0 00-.293-.707l-5.414-5.414A1 1 0 0012.586 3H7a2 2 0 00-2 2v14a2 2 0 002 2z" />
              </svg>
              Конвертировать в PDF все
            </button>
          </div>
        </div>
      </template>
      
      <!-- Actions for Archives/PDFs -->
      <template v-else-if="store.activeTab === 'archives' || store.activeTab === 'pdfs'">
        <button
          @click="store.toggleSelectAll"
          class="btn btn-outline flex items-center gap-2"
        >
          <span class="hidden sm:inline">{{ store.allSelected ? 'Снять' : 'Выбрать все' }}</span>
        </button>
        
        <button
          @click="store.deleteSelected"
          :disabled="!store.hasSelection"
          class="btn btn-danger flex items-center gap-2"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
          </svg>
          <span class="hidden sm:inline">Удалить</span>
        </button>
      </template>
    </div>
  </div>
  
  <!-- Click outside to close dropdown -->
  <div v-if="showActions" class="fixed inset-0 z-40" @click="showActions = false"></div>
</template>
