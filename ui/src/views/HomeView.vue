<template>
  <main>
    <h1 class="text-3xl font-bold underline">
      Hello world!
    </h1>
    <button class="btn">Button</button>
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