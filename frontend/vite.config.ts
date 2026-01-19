import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig(({ mode }) => {
  // Load env from parent directory (project root) and current directory
  const rootEnv = loadEnv(mode, '../', '')
  const localEnv = loadEnv(mode, './', '')
  
  // Merge envs, local takes priority
  const env = { ...rootEnv, ...localEnv }
  
  const serverPort = parseInt(env.VITE_PORT || env.SERVER_PORT || '3000')
  const backendPort = parseInt(env.VITE_BACKEND_PORT || env.SERVER_PORT || '8080')
  const backendHost = env.VITE_BACKEND_HOST || env.SERVER_HOST || 'localhost'
  
  const backendUrl = `http://${backendHost}:${backendPort}`
  
  return {
    plugins: [vue()],
    server: {
      port: serverPort,
      host: env.VITE_HOST || '0.0.0.0',
      proxy: {
        '/api': {
          target: backendUrl,
          changeOrigin: true
        },
        '/files': {
          target: backendUrl,
          changeOrigin: true
        }
      }
    },
    build: {
      outDir: '../cmd/server/dist',
      emptyOutDir: true
    }
  }
})
