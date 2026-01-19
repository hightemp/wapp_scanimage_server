<script setup lang="ts">
import { ref, watch } from 'vue'
import { useAppStore } from '../stores/app'

const store = useAppStore()
const scale = ref(1)
const position = ref({ x: 0, y: 0 })
const isDragging = ref(false)
const dragStart = ref({ x: 0, y: 0 })

watch(() => store.imagePreview, () => {
  // Reset zoom and position when image changes
  scale.value = 1
  position.value = { x: 0, y: 0 }
})

const close = () => {
  store.setImagePreview(null)
}

const zoomIn = () => {
  scale.value = Math.min(scale.value * 1.25, 5)
}

const zoomOut = () => {
  scale.value = Math.max(scale.value / 1.25, 0.5)
}

const resetZoom = () => {
  scale.value = 1
  position.value = { x: 0, y: 0 }
}

const handleWheel = (e: WheelEvent) => {
  e.preventDefault()
  if (e.deltaY < 0) {
    zoomIn()
  } else {
    zoomOut()
  }
}

const handleMouseDown = (e: MouseEvent) => {
  if (scale.value > 1) {
    isDragging.value = true
    dragStart.value = { x: e.clientX - position.value.x, y: e.clientY - position.value.y }
  }
}

const handleMouseMove = (e: MouseEvent) => {
  if (isDragging.value) {
    position.value = {
      x: e.clientX - dragStart.value.x,
      y: e.clientY - dragStart.value.y
    }
  }
}

const handleMouseUp = () => {
  isDragging.value = false
}

const handleKeyDown = (e: KeyboardEvent) => {
  if (e.key === 'Escape') {
    close()
  } else if (e.key === '+' || e.key === '=') {
    zoomIn()
  } else if (e.key === '-') {
    zoomOut()
  } else if (e.key === '0') {
    resetZoom()
  }
}
</script>

<template>
  <Teleport to="body">
    <div
      v-if="store.imagePreview"
      class="fixed inset-0 z-50 bg-black/90 flex items-center justify-center"
      @click.self="close"
      @keydown="handleKeyDown"
      @wheel="handleWheel"
      tabindex="0"
      ref="viewer"
    >
      <!-- Toolbar -->
      <div class="absolute top-4 left-1/2 -translate-x-1/2 flex items-center gap-2 bg-white/10 backdrop-blur rounded-lg p-2 z-10">
        <button
          @click="zoomOut"
          class="p-2 rounded hover:bg-white/20 text-white transition-colors"
          title="Уменьшить (−)"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 12H4" />
          </svg>
        </button>
        
        <span class="text-white text-sm px-2 min-w-[60px] text-center">
          {{ Math.round(scale * 100) }}%
        </span>
        
        <button
          @click="zoomIn"
          class="p-2 rounded hover:bg-white/20 text-white transition-colors"
          title="Увеличить (+)"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
          </svg>
        </button>
        
        <div class="w-px h-6 bg-white/30 mx-1"></div>
        
        <button
          @click="resetZoom"
          class="p-2 rounded hover:bg-white/20 text-white transition-colors"
          title="Сбросить (0)"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 8V4m0 0h4M4 4l5 5m11-1V4m0 0h-4m4 0l-5 5M4 16v4m0 0h4m-4 0l5-5m11 5l-5-5m5 5v-4m0 4h-4" />
          </svg>
        </button>
      </div>
      
      <!-- Close Button -->
      <button
        @click="close"
        class="absolute top-4 right-4 p-2 rounded-lg bg-white/10 hover:bg-white/20 text-white transition-colors z-10"
        title="Закрыть (Esc)"
      >
        <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
        </svg>
      </button>
      
      <!-- Image Container -->
      <div
        class="overflow-hidden w-full h-full flex items-center justify-center"
        :class="{ 'cursor-grab': scale > 1, 'cursor-grabbing': isDragging }"
        @mousedown="handleMouseDown"
        @mousemove="handleMouseMove"
        @mouseup="handleMouseUp"
        @mouseleave="handleMouseUp"
      >
        <img
          :src="store.imagePreview"
          alt="Preview"
          class="max-w-none select-none"
          :style="{
            transform: `scale(${scale}) translate(${position.x / scale}px, ${position.y / scale}px)`,
            transition: isDragging ? 'none' : 'transform 0.1s ease-out'
          }"
          draggable="false"
        />
      </div>
      
      <!-- Instructions -->
      <div class="absolute bottom-4 left-1/2 -translate-x-1/2 text-white/60 text-sm">
        Колесо мыши для масштабирования • Перетаскивание для перемещения • Esc для закрытия
      </div>
    </div>
  </Teleport>
</template>
