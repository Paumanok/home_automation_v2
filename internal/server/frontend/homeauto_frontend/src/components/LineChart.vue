<template>
  <div>
    <canvas ref="lineChart" width="800" height="400"></canvas>
  </div>
</template>

<script>
import { ref, watch, onMounted } from 'vue';
import Chart from 'chart.js/auto';

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
        const response = await fetch(props.apiEndpoint);
        const data = await response.json();

        const labels = data.map((entry) => entry);
        const measurementData = data.map((entry) => entry[props.measurementType]);

        if (lineChart.value) {
          const ctx = lineChart.value.getContext('2d');

          if (chartInstance.value) {
            chartInstance.value.destroy(); // Destroy the existing chart instance
          }

          chartInstance.value = new Chart(ctx, {
            type: 'line',
            data: {
              labels: labels,
              datasets: [
                {
                  label: props.measurementType,
                  data: measurementData,
                  borderColor: 'rgba(75, 192, 192, 1)',
                  borderWidth: 1,
                  fill: false,
                },
              ],
            },
            options: {
              responsive: true,
              maintainAspectRatio: false,
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

