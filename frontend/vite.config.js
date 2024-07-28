import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },

  server: {
    host: true,
    port: 5000,
  }
  //server: {
  //  port: 5000,
  //  watch: {usePolling: false},
  //  strictPort: true,
  //  hmr: {
  //          host: "frontend",
  //          protocol: "ws",
  //          port: 80,
  //      },
  //},
})
