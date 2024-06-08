<template>
  <div>
    <canvas ref="lineChart" width="800" height="400"></canvas>
  </div>
</template>

<script>
import { ref, watch, onMounted } from 'vue';
import Chart from 'chart.js/auto';
import * as Measurements from '../services/measurements.js'
import 'chartjs-adapter-luxon';

export default {
  props: {
    apiEndpoint: String, // API endpoint to fetch data
    measurementType: String, // 'temp' or 'humidity'
  },
  setup(props) {
    const lineChart = ref(null);
    const chartInstance = ref(null);

    const fetchData = async () => {
      try {
        //const response = await fetch(props.apiEndpoint);
        //const data = await response.json();
        const measurements = await Measurements.get_measurements();
        //const labels = data.map((entry) => entry);
        //const measurementData = data.map((entry) => entry[props.measurementType]);
        
        const dataset = []
        
        for(const i in measurements) {
          const d = measurements[i];
          const m = d["measurements"];
          dataset.push(  {
              label: d["deviceInfo"]["nickname"],
              data: m.map( 
                (m) => ({
                  x: m["createdAt"],
                  y: m["temp"],
              })),
                fill: false,
                boarderWidth: 1,
            }
          );
        }
        console.log(dataset)
        if (lineChart.value) {
          const ctx = lineChart.value.getContext('2d');

          if (chartInstance.value) {
            chartInstance.value.destroy(); // Destroy the existing chart instance
          }

          chartInstance.value = new Chart(ctx, {
            type: 'line',
            data: {
              //labels: labels,
              datasets: dataset
              //[
              //  {
              //    label: props.measurementType,
              //    data: measurementData,
              //    borderColor: 'rgba(75, 192, 192, 1)',
              //    borderWidth: 1,
              //    fill: false,
              //  },
              //],
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
                      second: 'hh:mm:ss - DD', // Format for seconds
                      minute: 'YYYY-MM-DDTHH:mm:ss.SSSZ', // Format for minutes
                      hour: 'YYYY-MM-DDTHH:mm:ss.SSSZ', // Format for hours
                      day: 'YYYY-MM-DDTHH:mm:ss.SSSZ', // Format for days
                      week: 'YYYY-MM-DDTHH:mm:ss.SSSZ', // Format for weeks
                      month: 'YYYY-MM-DDTHH:mm:ss.SSSZ', // Format for months
                      quarter: 'YYYY-MM-DDTHH:mm:ss.SSSZ', // Format for quarters
                      year: 'YYYY-MM-DDTHH:mm:ss.SSSZ', // Format for years
                    },
                    unit: 'second',
                  },
                  //time: {
                  //  //parser: 'YY:MM:DDTHH:mm:ss',
                  //  parser: "yyyy-MM-ddTHH:mm:ss.SSSSZ",
                  //  unit: 'hour',
                  //  //displayFormats: {
                  //  //  hour: 'ddTHH:mm'
                  //  //},
                  //  //tooltipFormat: 'D MMM YYYY - HH:mm:ss'
                  //}
                }
              }
            },
          });
        }
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    };

    // Fetch data when the component is mounted and re-fetch when the API endpoint changes
    onMounted(() => {
      fetchData();
    });

    watch(() => props.apiEndpoint, () => {
      fetchData();
    });

    return {
      lineChart,
    };
  },
};
</script>

<style scoped>
/* Add any custom styles for your chart here */
</style>

