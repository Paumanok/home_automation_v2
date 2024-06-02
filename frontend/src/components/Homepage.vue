<template> 
  <header>
    <img alt="Home logo" class="logo" src="@/assets/logo-house.svg" width="125" height="125" />
  </header>

  <h1> this is my main header </h1>

  <div>
    <v-col dense>
      <v-row>
      <h3> another thing </h3>
      </v-row>
      <v-row>
      <h3> wow </h3>
      </v-row>

    </v-col>
    <v-col dense>
      <v-row>
      <h3> another thing </h3>
      </v-row>
      <v-row>
      <h3> wow </h3>
      </v-row>

    </v-col>
    

  </div>



</template>


<script setup>


</script>

<style scoped>
header {
  line-height: 1.5;
  max-height: 100vh;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}
</style>
