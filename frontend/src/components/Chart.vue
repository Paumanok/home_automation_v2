<template>
  <div>
    <canvas ref="lineChart" width="400" height="400"></canvas>
  </div>
</template>

<script setup>
import { ref, onMounted, computed, watch } from 'vue';
import Chart from 'chart.js/auto';
import * as Measurements from '../services/measurements.js'
import 'chartjs-adapter-luxon';

const model = defineModel()
const props = defineProps(['measurementType'])
const lineChart = ref(null);
const chartInstance = ref(null);
//ref({label: "na", data: [{ x: "2024-06-07T01:17:48.809027Z", y: 0}]})

onMounted(async () => {
  fetchData()
})

watch( model, async () => {
  console.log("fetchinnnn in chart")
  fetchData()
})

function fetchData() {
      try {
        //const measurements = await Measurements.get_measurements();
        console.log("model:")
        console.log(model.value.measurements)
        const measurements = model.value.measurements
        const dataset = []
        
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
        if (lineChart.value) {
          const ctx = lineChart.value.getContext('2d');

          if (chartInstance.value) {
            chartInstance.value.destroy(); // Destroy the existing chart instance
          }

          chartInstance.value = new Chart(ctx, {
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
        }
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    
}

function construct_chart_data(measurements) {
  //const data = await Measurements.get_measurements()
  const outdata = []
  
  for(const i in measurements) {
    const d = measurements[i];
    const m = d["measurements"];
    outdata.push(  {
        label: d["deviceInfo"]["nickname"],
        data: m.map( 
          (m) => ({
            x: m["createdAt"],
            y: m["temp"],
        })),
      }
    );
  }
  return outdata
}


</script>

<style>
#app {
  min-width:100%;
}
</style>
