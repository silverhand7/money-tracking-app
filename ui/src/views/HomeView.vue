<template>
  <main>
      <TheWelcome />
  </main>
</template>

<script>
import TheWelcome from '../components/TheWelcome.vue'
import config from '@/config.js'
import axios from 'axios'

export default{
  components: {
    TheWelcome
  },
  created(){
    axios.get(config.basePath + "/api/categories")
      .then((response) => {
        console.log(response)
      })
      .catch((error) => {
        console.error('Error fetching data:', error);
      });
  }
}

</script>