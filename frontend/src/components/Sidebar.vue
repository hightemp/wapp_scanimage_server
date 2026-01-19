<script setup lang="ts">
import { useAppStore } from '../stores/app'
import type { TabType } from '../types'

const store = useAppStore()

const tabs: { id: TabType; label: string; icon: string }[] = [
  { id: 'scans', label: 'Сканы', icon: 'M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z' },
  { id: 'archives', label: 'Архивы', icon: 'M5 8h14M5 8a2 2 0 110-4h14a2 2 0 110 4M5 8v10a2 2 0 002 2h10a2 2 0 002-2V8m-9 4h4' },
  { id: 'pdfs', label: 'PDF', icon: 'M7 21h10a2 2 0 002-2V9.414a1 1 0 00-.293-.707l-5.414-5.414A1 1 0 0012.586 3H7a2 2 0 00-2 2v14a2 2 0 002 2z' },
  { id: 'settings', label: 'Настройки', icon: 'M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z M15 12a3 3 0 11-6 0 3 3 0 016 0z' }
]

const setTab = (tab: TabType) => {
  store.activeTab = tab
  store.deselectAll()
}
</script>

<template>
  <aside class="w-full lg:w-64 bg-white dark:bg-gray-800 shadow-lg flex flex-col">
    <!-- Header -->
    <div class="p-4 border-b border-gray-200 dark:border-gray-700">
      <h1 class="text-xl font-bold text-gray-800 dark:text-white flex items-center gap-2">
        <svg class="w-6 h-6 text-primary-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
        </svg>
        <span class="hidden lg:inline">Сканирование</span>
      </h1>
    </div>
    
    <!-- Scanner Info -->
    <div class="p-3 border-b border-gray-200 dark:border-gray-700 text-xs text-gray-500 dark:text-gray-400 hidden lg:block">
      <div v-if="store.scanners.length > 0" class="space-y-1">
        <div v-for="scanner in store.scanners" :key="scanner.device" class="truncate" :title="scanner.description">
          {{ scanner.description || scanner.device }}
        </div>
      </div>
      <div v-else class="text-gray-400">
        Сканеры не найдены
      </div>
    </div>
    
    <!-- Navigation -->
    <nav class="flex-1 p-2">
      <div class="flex lg:flex-col gap-1">
        <button
          v-for="tab in tabs"
          :key="tab.id"
          @click="setTab(tab.id)"
          class="flex items-center gap-3 px-4 py-3 rounded-lg transition-colors flex-1 lg:flex-none"
          :class="[
            store.activeTab === tab.id 
              ? 'bg-primary-100 dark:bg-primary-900 text-primary-700 dark:text-primary-300' 
              : 'text-gray-600 dark:text-gray-400 hover:bg-gray-100 dark:hover:bg-gray-700'
          ]"
        >
          <svg class="w-5 h-5 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="tab.icon" />
          </svg>
          <span class="hidden lg:inline font-medium">{{ tab.label }}</span>
          
          <!-- Badge -->
          <span
            v-if="tab.id === 'scans' && store.scans.length > 0"
            class="hidden lg:inline ml-auto bg-primary-600 text-white text-xs px-2 py-0.5 rounded-full"
          >
            {{ store.scans.length }}
          </span>
          <span
            v-else-if="tab.id === 'archives' && store.archives.length > 0"
            class="hidden lg:inline ml-auto bg-gray-500 text-white text-xs px-2 py-0.5 rounded-full"
          >
            {{ store.archives.length }}
          </span>
          <span
            v-else-if="tab.id === 'pdfs' && store.pdfs.length > 0"
            class="hidden lg:inline ml-auto bg-red-500 text-white text-xs px-2 py-0.5 rounded-full"
          >
            {{ store.pdfs.length }}
          </span>
        </button>
      </div>
    </nav>
    
    <!-- Theme Toggle -->
    <div class="p-4 border-t border-gray-200 dark:border-gray-700 hidden lg:block">
      <button
        @click="store.updateSettings({ theme: store.isDarkMode ? 'light' : 'dark' })"
        class="w-full flex items-center justify-center gap-2 px-4 py-2 rounded-lg bg-gray-100 dark:bg-gray-700 text-gray-600 dark:text-gray-300 hover:bg-gray-200 dark:hover:bg-gray-600 transition-colors"
      >
        <svg v-if="store.isDarkMode" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z" />
        </svg>
        <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z" />
        </svg>
        <span>{{ store.isDarkMode ? 'Светлая тема' : 'Темная тема' }}</span>
      </button>
    </div>
  </aside>
</template>
