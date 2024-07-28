<script setup>
import HelloWorld from './components/HelloWorld.vue'
import TheWelcome from './components/TheWelcome.vue'
import NextCounter from './components/NextCounter.vue'
import Homepage from './views/Homepage.vue'
import Sidebar from 'primevue/sidebar';
import Button from 'primevue/button';
import * as Kasa from './services/kasa.js';
import { ref } from 'vue';

const visible = ref(false);
function handleClick() {
  Kasa.toggle_lamp();

}
</script>

<template>
    
    <v-layout class="rounded rounded-md">
      <v-app-bar color="black">
        <v-img
          class="mx-2"
          src="src/assets/logo-house.svg"
          max-height="40"
          max-width="40"
          contain
        ></v-img>
        <v-app-bar-title> Homeauto </v-app-bar-title>
        <v-spacer></v-spacer>
        <v-btn color="primary" @click="handleClick">
         Lamp 
      </v-btn>
      </v-app-bar>
      <v-navigation-drawer
          expand-on-hover 
          rail 
          color="black"
          mobile-breakpoint='xs'
          >
          <v-list density="compact" nav> 
            <v-list-item prepend-icon="pi pi-chart-line" title="Environment" to="/environment"></v-list-item>
            <v-list-item prepend-icon="pi pi-stopwatch" title="Pomodoro" to="/pomodoro"></v-list-item>
            <v-list-item prepend-icon="pi pi-download" title="Deluge" href="http://deluge.datapaddock.lan"></v-list-item>
          </v-list>
      </v-navigation-drawer>
    </v-layout>

    <router-view />
</template>

<style scoped>
header {
  line-height: 1.5;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}
  #app {
    min-width: 90%;
  }
@media (min-width: 100%) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }
}
</style>
