import { createApp } from 'vue'
import NoteViewer from './NoteViewer.vue'

import './assets/main.css'

const app = createApp(NoteViewer)
app.config.globalProperties.backendURL = 'http://127.0.0.1:8090'

app.mount('#app')
