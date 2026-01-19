<script setup lang="ts">
import { onMounted } from 'vue'
import { useAppStore } from './stores/app'
import Sidebar from './components/Sidebar.vue'
import ScanGrid from './components/ScanGrid.vue'
import ArchiveList from './components/ArchiveList.vue'
import PdfList from './components/PdfList.vue'
import Settings from './components/Settings.vue'
import ImageViewer from './components/ImageViewer.vue'
import ToastContainer from './components/ToastContainer.vue'
import Toolbar from './components/Toolbar.vue'

const store = useAppStore()

onMounted(() => {
  store.init()
})
</script>

<template>
  <div class="h-full flex flex-col lg:flex-row bg-gray-100 dark:bg-gray-900">
    <!-- Sidebar -->
    <Sidebar />
    
    <!-- Main Content -->
    <main class="flex-1 flex flex-col min-h-0 overflow-hidden">
      <!-- Toolbar -->
      <Toolbar />
      
      <!-- Content Area -->
      <div class="flex-1 overflow-auto p-4">
        <ScanGrid v-if="store.activeTab === 'scans'" />
        <ArchiveList v-else-if="store.activeTab === 'archives'" />
        <PdfList v-else-if="store.activeTab === 'pdfs'" />
        <Settings v-else-if="store.activeTab === 'settings'" />
      </div>
    </main>
    
    <!-- Image Viewer Modal -->
    <ImageViewer />
    
    <!-- Toast Notifications -->
    <ToastContainer />
  </div>
</template>
