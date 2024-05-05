import {createApp, reactive} from 'vue'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js';
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'
import Post from './components/Post.vue'
import UserEntry from './components/UserEntry.vue'
import DeletablePost from './components/DeletablePost.vue'
import './assets/dashboard.css'
import './assets/main.css'

const app = createApp(App)
app.config.globalProperties.$axios = axios;
app.component("ErrorMsg", ErrorMsg);
app.component("Post", Post);
app.component("UserEntry", UserEntry);
app.component("LoadingSpinner", LoadingSpinner);
app.component("DeletablePost", DeletablePost)
app.use(router)
app.mount('#app')
