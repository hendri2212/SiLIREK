import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { createPersistedStatePlugin } from 'pinia-plugin-persistedstate-2'

import App from './App.vue'
import router from './router'
import api from './plugins/axios'; // Import api

// Determine API base URL based on environment
const API_URL = import.meta.env.PROD
  ? import.meta.env.VITE_API_URL_PROD
  : import.meta.env.VITE_API_URL_DEV;

// Set axios baseURL dynamically
api.defaults.baseURL = API_URL;

const app = createApp(App)

// Tambahkan ke global properties
app.config.globalProperties.$axios = api;

const pinia = createPinia()
pinia.use(createPersistedStatePlugin())
app.use(pinia)
app.use(router)

router.isReady().then(() => {
    app.mount('#app')
})
