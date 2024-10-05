import { defineConfig } from 'vite'
import tailwindcss from 'tailwindcss'
import react from '@vitejs/plugin-react'
import basicSsl from '@vitejs/plugin-basic-ssl'


export default defineConfig({
  plugins: [
    react(),
    basicSsl()
  ],
  css: {
    postcss: {
      plugins: [tailwindcss()],
    }
  },
})
