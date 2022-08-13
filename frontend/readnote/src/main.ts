import { createApp } from 'vue'
import App from './App.vue'

import './assets/main.css'

const app = createApp(App)
app.config.globalProperties.backendURL = 'http://127.0.0.1:8090'

app.mount('#app')
