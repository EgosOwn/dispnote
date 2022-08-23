import { createApp } from 'vue'
import NoteCreator from './NoteCreator.vue'

import './assets/main.css'

const app = createApp(NoteCreator)
app.config.globalProperties.backendURL = 'http://127.0.0.1:8090'

app.mount('#app')
