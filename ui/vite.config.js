/** @type {import('vite').UserConfig} */
import { defineConfig } from 'vite'
import { tempo } from '@tempots/vite'

export default defineConfig({
  base: '',
  plugins: [tempo({
    mode: "ssg",
    devtools: true,
   })],
  build: {
    rollupOptions: {
      output: {
        entryFileNames: 'assets/[name].js',
        assetFileNames: 'assets/[name].[ext]'
      }
    }
  }
})
