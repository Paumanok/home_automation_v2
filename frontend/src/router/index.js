import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Homepage.vue'
import Environment from '../views/environment.vue'
import Pomodoro from '../views/pomodoro.vue'

const router = createRouter({
	history: createWebHistory(),
	routes: [
		{
			path: '/',
			component: Environment
		},
		{
			path: '/environment',
			component: Environment,
		},
		{
			path: '/pomodoro',
			component: Pomodoro,
		},
	],
})

export default router
