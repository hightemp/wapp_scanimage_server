<script setup lang="ts">
import { computed } from 'vue'
import { useAppStore } from '../stores/app'
import draggable from 'vuedraggable'
import FileCard from './FileCard.vue'

const store = useAppStore()

const scans = computed({
  get: () => store.scans,
  set: (value) => store.updateFileOrder(value)
})

const handlePreview = (path: string) => {
  store.setImagePreview(path)
}
</script>

<template>
  <div class="h-full">
    <!-- Loading State -->
    <div v-if="store.isLoading && scans.length === 0" class="flex items-center justify-center h-64">
      <div class="spinner"></div>
    </div>
    
    <!-- Empty State -->
    <div v-else-if="scans.length === 0" class="flex flex-col items-center justify-center h-64 text-gray-500 dark:text-gray-400">
      <svg class="w-16 h-16 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
      </svg>
      <p class="text-lg">No scanned images</p>
      <p class="text-sm mt-2">Click "Scan" to start</p>
    </div>
    
    <!-- Grid -->
    <draggable
      v-else
      v-model="scans"
      item-key="name"
      class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 2xl:grid-cols-6 gap-4"
      ghost-class="sortable-ghost"
      chosen-class="sortable-chosen"
      animation="200"
    >
      <template #item="{ element }">
        <FileCard
          :file="element"
          :selected="store.selectedFiles.has(element.name)"
          @toggle-select="store.toggleSelect(element.name)"
          @preview="handlePreview(element.relativePath)"
        />
      </template>
    </draggable>
  </div>
</template>
