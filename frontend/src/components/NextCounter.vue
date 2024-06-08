<template>

  <h2> {{ next }} </h2>

</template>

<script setup> 
import {ref, onMounted} from 'vue';
import * as Measurements from '../services/measurements';

const next = ref(0);

onMounted(() => {
  next.val = Measurements.get_next(next)
  console.log("hello")
})

setInterval(decrement, 1000)

function decrement() {
  next.value--
  if(next.value <= 0) {
    Measurements.get_next(next)
  }
}

//async function get_next() {
//  
//  const response = await fetch("/api/next", {
//    method: "GET",
//    headers: {
//      "Content-Type": "application/json",
//      "Accept": "application/json"
//    },
//  });
//  
//  const resp_json = await response.json()
//  next.value = resp_json["sync_time"]
//}



</script>
