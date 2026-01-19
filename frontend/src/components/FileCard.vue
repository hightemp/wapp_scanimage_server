<script setup lang="ts">
import type { FileInfo } from '../types'

defineProps<{
  file: FileInfo
  selected: boolean
}>()

const emit = defineEmits<{
  toggleSelect: []
  preview: []
}>()
</script>

<template>
  <div
    class="card cursor-pointer transition-all duration-200 hover:shadow-lg"
    :class="{ 'ring-2 ring-primary-500': selected }"
  >
    <!-- Image Preview -->
    <div class="relative aspect-[3/4] bg-gray-100 dark:bg-gray-700 overflow-hidden" @click="emit('preview')">
      <img
        :src="file.relativePath"
        :alt="file.name"
        class="w-full h-full object-cover"
        loading="lazy"
      />
      
      <!-- Checkbox Overlay -->
      <div
        class="absolute top-2 left-2"
        @click.stop="emit('toggleSelect')"
      >
        <div
          class="w-6 h-6 rounded border-2 flex items-center justify-center transition-colors"
          :class="[
            selected 
              ? 'bg-primary-600 border-primary-600' 
              : 'bg-white/80 dark:bg-gray-800/80 border-gray-300 dark:border-gray-600 hover:border-primary-500'
          ]"
        >
          <svg v-if="selected" class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M5 13l4 4L19 7" />
          </svg>
        </div>
      </div>
      
      <!-- Drag Handle -->
      <div class="absolute top-2 right-2 cursor-move opacity-50 hover:opacity-100">
        <svg class="w-5 h-5 text-white drop-shadow" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 8h16M4 16h16" />
        </svg>
      </div>
    </div>
    
    <!-- Info -->
    <div class="p-3">
      <p class="text-sm font-medium text-gray-800 dark:text-gray-200 truncate" :title="file.name">
        {{ file.name }}
      </p>
      <p class="text-xs text-gray-500 dark:text-gray-400 mt-1">
        {{ file.sizeHuman }}
      </p>
    </div>
  </div>
</template>
