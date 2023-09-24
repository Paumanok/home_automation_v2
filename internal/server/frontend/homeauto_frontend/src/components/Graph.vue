<template>
  <Line
    id="my-chart-id"
    :options="chartOptions"
    :v-if="loaded"
    :data="chartData"
  />
</template>

<script>
import axios from 'axios';
import { Line } from 'vue-chartjs'
import { Chart as ChartJS, Title, Tooltip, Legend, LineElement, CategoryScale, LinearScale, PointElement} from 'chart.js'

ChartJS.register(Title, Tooltip, Legend, LineElement, CategoryScale, LinearScale, PointElement)

export default {
  name: 'LinePlot',
  components: {Line},
  data() {
    return {
      chartData: {
        labels: [], // Labels will now hold the 'createdAt' timestamps
        datasets: [
          {
            label: 'Temperature (°C)',
            data: [],
            borderColor: 'rgba(75, 192, 192, 1)',
            backgroundColor: 'rgba(75, 192, 192, 0.2)',
          },
          {
            label: 'Humidity (%)',
            data: [],
            borderColor: 'rgba(255, 99, 132, 1)',
            backgroundColor: 'rgba(255, 99, 132, 0.2)',
          },
        ],
      },
      chartOptions: { 
        responsive: true 
      }
    };
  },
  async mounted() {
  //this.fetchData();
    this.loaded = false
      try {
        const response = await axios.get('/api/measurements/last/hour');
        const data = response.data;
// Extracting 'temp', 'humidity', and 'createdAt' values
        const labels = data.map(item => item.createdAt);
        const tempValues = data.map(item => item.temp);
        const humidityValues = data.map(item => item.humidity);

        this.chartData.labels = labels;
        this.chartData.datasets[0].data = tempValues;
        this.chartData.datasets[1].data = humidityValues;
        
      } catch (error) {
        console.error('Error fetching data:', error);
      }
  },
}; 
</script>
