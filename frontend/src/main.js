import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
//import { createMemoryHistory, createRouter } from 'vue-router'
import PrimeVue from 'primevue/config';
import App from './App.vue'
import router from './router'

import Sidebar from 'primevue/sidebar';
import Button from 'primevue/button';

import 'primeicons/primeicons.css';
import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import { aliases, mdi } from "vuetify/iconsets/mdi-svg"
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'



import Chart from 'chart.js/auto';
import { Line } from 'vue-chartjs';
//custom icons
import logo from '@/customIcons/logo.vue'
//const aliasesCustom = {
//  ...aliases,
//  logo,
//}
//
const vuetify = createVuetify({
  components,
  directives,
})
//  icons: {
//    defaultSet: 'pi',
//    aliases: {
//      ...aliasesCustom
//    },
//    sets: {
//      mdi,
//    },
//  },
//})


const app = createApp(App)

app.use(createPinia())
app.use(PrimeVue);
app.use(router)
app.use(vuetify)
app.mount('#app')

app.component('sidebar', Sidebar)
app.component('chart', Chart)
