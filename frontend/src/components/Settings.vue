<script setup lang="ts">
import { ref, watch } from 'vue'
import { useAppStore } from '../stores/app'

const store = useAppStore()

const resolution = ref(store.settings.resolution)
const quality = ref(store.settings.quality)
const theme = ref(store.settings.theme)

watch(() => store.settings, (newSettings) => {
  resolution.value = newSettings.resolution
  quality.value = newSettings.quality
  theme.value = newSettings.theme
}, { deep: true })

const saveSettings = () => {
  store.updateSettings({
    resolution: resolution.value,
    quality: quality.value,
    theme: theme.value
  })
}
</script>

<template>
  <div class="max-w-2xl mx-auto">
    <div class="card p-6">
      <h2 class="text-xl font-semibold text-gray-800 dark:text-white mb-6">Scan Settings</h2>
      
      <div class="space-y-6">
        <!-- Resolution -->
        <div>
          <label for="resolution" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
            Resolution (DPI)
          </label>
          <select
            id="resolution"
            v-model="resolution"
            class="select"
          >
            <option value="75">75 DPI (Fast)</option>
            <option value="150">150 DPI (Medium)</option>
            <option value="300">300 DPI (High)</option>
            <option value="600">600 DPI (Very High)</option>
          </select>
          <p class="mt-1 text-sm text-gray-500 dark:text-gray-400">
            Higher resolution = better quality, but larger file size
          </p>
        </div>
        
        <!-- Quality -->
        <div>
          <label for="quality" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
            JPEG Quality: {{ quality }}%
          </label>
          <input
            id="quality"
            type="range"
            v-model="quality"
            min="10"
            max="100"
            step="5"
            class="w-full h-2 bg-gray-200 dark:bg-gray-700 rounded-lg appearance-none cursor-pointer"
          />
          <div class="flex justify-between text-xs text-gray-500 dark:text-gray-400 mt-1">
            <span>Low (10%)</span>
            <span>High (100%)</span>
          </div>
        </div>
        
        <!-- Theme -->
        <div>
          <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
            Theme
          </label>
          <div class="grid grid-cols-3 gap-3">
            <button
              @click="theme = 'light'"
              class="p-4 rounded-lg border-2 transition-colors flex flex-col items-center gap-2"
              :class="[
                theme === 'light' 
                  ? 'border-primary-500 bg-primary-50 dark:bg-primary-900' 
                  : 'border-gray-200 dark:border-gray-700 hover:border-gray-300 dark:hover:border-gray-600'
              ]"
            >
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z" />
              </svg>
              <span class="text-sm">Light</span>
            </button>
            
            <button
              @click="theme = 'dark'"
              class="p-4 rounded-lg border-2 transition-colors flex flex-col items-center gap-2"
              :class="[
                theme === 'dark' 
                  ? 'border-primary-500 bg-primary-50 dark:bg-primary-900' 
                  : 'border-gray-200 dark:border-gray-700 hover:border-gray-300 dark:hover:border-gray-600'
              ]"
            >
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z" />
              </svg>
              <span class="text-sm">Dark</span>
            </button>
            
            <button
              @click="theme = 'system'"
              class="p-4 rounded-lg border-2 transition-colors flex flex-col items-center gap-2"
              :class="[
                theme === 'system' 
                  ? 'border-primary-500 bg-primary-50 dark:bg-primary-900' 
                  : 'border-gray-200 dark:border-gray-700 hover:border-gray-300 dark:hover:border-gray-600'
              ]"
            >
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.75 17L9 20l-1 1h8l-1-1-.75-3M3 13h18M5 17h14a2 2 0 002-2V5a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
              </svg>
              <span class="text-sm">System</span>
            </button>
          </div>
        </div>
        
        <!-- Scanner Info -->
        <div>
          <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
            Scanner Information
          </label>
          <div class="bg-gray-50 dark:bg-gray-900 rounded-lg p-4 text-sm font-mono text-gray-600 dark:text-gray-400 whitespace-pre-wrap">
            <template v-if="store.scannersRaw">{{ store.scannersRaw }}</template>
            <template v-else>No scanners found</template>
          </div>
        </div>
        
        <!-- Save Button -->
        <div class="pt-4">
          <button
            @click="saveSettings"
            class="btn btn-primary w-full sm:w-auto"
          >
            Save Settings
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
