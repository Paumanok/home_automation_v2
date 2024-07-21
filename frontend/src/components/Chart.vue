<template>
  <div>
    <canvas ref="lineChart" width="400" height="400"></canvas>
  </div>
</template>

<script setup>
import { ref, onMounted, computed, watch, shallowRef } from 'vue';
import Chart from 'chart.js/auto';
import * as Measurements from '../services/measurements.js'
import 'chartjs-adapter-luxon';

const model = defineModel()
const props = defineProps(['measurementType'])
const lineChart = ref(null);
const lcContext = ref(null);
//important! 
//https://stackoverflow.com/questions/77700265/chart-js-with-vue3-update-fails-gives-infinite-recursion-errors
const chartInstance = shallowRef(null);

onMounted(async () => {
  lcContext.value = lineChart.value.getContext("2d")
  createChart()
})

watch( model, async (newValue) => {
  try {
    if (chartInstance.value != null){
      updateChart(newValue)
      }
  } catch (error) {
    console.log(error)
  }
})

function collectData(newModel) {
      const dataset = []
      if( model.value != null) {

        const measurements = newModel.measurements
        
        for(const i in measurements) {
          const d = measurements[i];
          const m = d["measurements"];
          dataset.push(  {
              label: d["deviceInfo"]["nickname"],
              data: m.map( 
                (m) => ({
                  x: m["createdAt"],
                  y: m[props.measurementType],
              })),
                fill: false,
                boarderWidth: 1,
            }
          );
        }
      }
      return dataset
}


function updateChart(newModel) {
    const dataset = collectData(newModel)
    setTimeout(() => 0, 100) //this sleep is needed to avoid glitchy charts, not extremely interested in discovering why.
    chartInstance.value.data.datasets = dataset

    chartInstance.value.update()
}

function createChart() {
    try {
    const dataset = collectData()
    chartInstance.value = new Chart(lcContext.value, {
      type: 'line',
      data: {
        datasets: dataset
      },
      options: {
        responsive: true,
        maintainAspectRatio: false,
        pointRadius: 0,
        scales: {
            x: {
            type: 'time',
            time: {
              tooltipFormat: 'YYYY-MM-DDTHH:mm:ss.SSSZ', // Format for the tooltip
              displayFormats: {
                millisecond: 'YYYY-MM-DDTHH:mm:ss.SSSZ', // Format for milliseconds
                second: 'LLL dd h:mm', // Format for seconds
                minute: 'LLL dd h:mm', // Format for minutes
                hour: 'YYYY-MM-DDTHH:mm:ss.SSSZ', // Format for hours
                day: 'YYYY-MM-DDTHH:mm:ss.SSSZ', // Format for days
                week: 'YYYY-MM-DDTHH:mm:ss.SSSZ', // Format for weeks
                month: 'YYYY-MM-DDTHH:mm:ss.SSSZ', // Format for months
                quarter: 'YYYY-MM-DDTHH:mm:ss.SSSZ', // Format for quarters
                year: 'YYYY-MM-DDTHH:mm:ss.SSSZ', // Format for years
              },
              unit: 'minute',
            },
          }
        }
      },
    });
  } catch (error) {
      console.error('Error fetching data:', error);
  }

}


</script>

<style>
#app {
  min-width:100%;
}
</style>
