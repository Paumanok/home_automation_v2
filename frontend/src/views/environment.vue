<template>
  <h2> environment </h2>
  <!--<v-container class="bg-surface-variant">
    <v-row no-gutters>
      <v-col dense>
        <Chart2 measurementType="temp" update="update" /> 
      </v-col>
      <v-col dense>
        <Chart2 measurementType="humidity" update="update" /> 
      </v-col>
    </v-row>

    <v-row>
      <v-col>
        <Chart v-model="measModel" measurementType="temp"/>
      </v-col>
    </v-row>
  </v-container>-->

<v-container fluid class="bg-surface-variant pa-10">
    <v-row no-gutters>
      <v-col cols="6">
        <v-sheet class="pa-1 ma-2">
          Temperature
          <Chart v-model="measModel" measurementType="temp"/>
        </v-sheet>
      </v-col>
      <v-col cols="6">
        <v-sheet class="pa-1 ma-2">
          Humidity
          <Chart v-model="measModel" measurementType="humidity"/>
        </v-sheet>
      </v-col>

      <v-responsive width="100%"></v-responsive>

      <v-col cols="6">
        <v-sheet class="pa-1 ma-2">
          Pressure
          <Chart v-model="measModel" measurementType="pressure"/>
        </v-sheet>
      </v-col>

      <v-col cols="6">
        <v-sheet class="pa-1 ma-2">
          pm25
          <Chart v-model="measModel" measurementType="pm25"/>
        </v-sheet>
      </v-col>
    </v-row>
  </v-container>
</template>


<script setup>
import { ref, onMounted, watch } from 'vue';
import NextCounter from '../components/NextCounter.vue';
import * as Measurements from '../services/measurements.js';
import Chart from '../components/Chart.vue';
//import Chart2 from '../components/chart2.vue';


const update = ref(false);

const measModel = ref(null);
const next = ref(0);

onMounted(async () => {
  Measurements.get_next(next)

  measModel.value = {"measurements": await Measurements.get_measurements()}
  console.log(measModel.value)

  //const measurements = await Measurements.get_measurements();
  //measModel.value = {"measurements": measurements};
  console.log("hello")
})

watch( update, async () => {
    console.log("watch hit")
    if(update.value == true) {
      update.value = false
      measModel.value = {"measurements": await Measurements.get_measurements()}
    }
})

setInterval(decrement, 1000)

function decrement() {
  next.value--
  if(next.value <= 0) {
    console.log("getting next")
    Measurements.get_next(next)
    update.value=true
  }
}



</script>

<style>

#app {
    min-width: 100%;
}
</style>
