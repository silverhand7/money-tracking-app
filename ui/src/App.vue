<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <div class="max-w-xl h-screen mx-auto bg-slate-100 relative">
    <div class="w-100 px-5 py-5">
      <RouterView />
    </div>

    <footer class="bg-slate-200 absolute bottom-0 w-full flex justify-around">
      <RouterLink to="/" class="py-3 max-w-[25px]">
        <img src="@/assets/images/home.png" alt="home">
      </RouterLink>
      <RouterLink to="/about" class="py-3 max-w-[25px]">
        <img src="@/assets/images/wallet.png" alt="wallet">
      </RouterLink>
      <RouterLink to="/" class="py-3 max-w-[25px]">
        <img src="@/assets/images/trend.png" alt="trend">
      </RouterLink>
      <RouterLink to="/about" class="py-3 max-w-[25px]">
        <img src="@/assets/images/user.png" alt="user">
      </RouterLink>
    </footer>
  </div>
</template>

<style scoped>
header {
  line-height: 1.5;
  max-height: 100vh;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

nav {
  width: 100%;
  font-size: 12px;
  text-align: center;
  margin-top: 2rem;
}

nav a.router-link-exact-active {
  color: var(--color-text);
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
}

nav a {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
}

nav a:first-of-type {
  border: 0;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }

  nav {
    text-align: left;
    margin-left: -1rem;
    font-size: 1rem;

    padding: 1rem 0;
    margin-top: 1rem;
  }
}
</style>
