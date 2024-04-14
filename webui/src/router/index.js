import {createRouter, createWebHashHistory} from 'vue-router'
import Login from '../views/Login.vue'
import ProfileView from '../views/ProfileView.vue'
import UserProfileView from '../views/UserProfileView.vue'
import SearchView from '../views/SearchView.vue'
import MyPhotosView from '../views/MyPhotosView.vue'
import ClientError from '../views/400.vue' 
import ServerError from '../views/500.vue'
import NotFound from '../views/404.vue'
import Unauthorized from '../views/401.vue'
import Forbidden from '../views/403.vue'
import { onBeforeUpdate } from 'vue'
const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', redirect: '/session'},
		{path: '/error/500', component: ServerError},
		{path: '/error/400', component: ClientError},
		{path:  '/error/404', component: NotFound},
		{path:  '/error/401', component: Unauthorized},
		{path:  '/error/403', component: Forbidden},
		
		{path: '/session', component: Login},
		{path: '/users', component: SearchView},
		{path: '/users/:id', component: ProfileView},
		{path: '/users/:id/myPhotos/', component: MyPhotosView},
		{path:  '/users/:id/profile', component: UserProfileView},
		
	]
})

export default router
