<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <div class="max-w-xl min-h-screen h-full mx-auto bg-slate-100 relative">
    <div class="w-100 px-5 pt-5 pb-16">
      <RouterView v-slot="{ Component }">
          <Component :is="Component" />
      </RouterView>
    </div>

    <footer class="bg-slate-200 fixed bottom-0 w-full flex justify-around max-w-xl">
      <RouterLink to="/" class="py-3 max-w-[25px]">
        <img src="@/assets/images/home.png" alt="home">
      </RouterLink>
      <RouterLink to="/wallets" class="py-3 max-w-[25px]">
        <img src="@/assets/images/wallet.png" alt="wallet">
      </RouterLink>
      <RouterLink :to="{ name: 'transactions.create' }" class="py-3 max-w-[25px]">
        <div class="bg-blue-500 text-3xl block pl-4 pr-4 text-center relative w-9 h-10 rounded-full border-4 border-slate-100">
          <span class="absolute left-[7.5px] -top-1 text-white">+</span>
        </div>
      </RouterLink>
      <RouterLink to="/" class="py-3 max-w-[25px]">
        <img src="@/assets/images/trend.png" alt="trend">
      </RouterLink>
      <RouterLink to="/user" class="py-3 max-w-[25px]">
        <img src="@/assets/images/user.png" alt="user">
      </RouterLink>
    </footer>
  </div>
</template>

<style scoped>
.fade-enter-from,
.fade-leave-to {
	opacity: 0;
}

.fade-enter-active,
.fade-leave-active {
	transition: opacity 0.3s ease-out;
}
</style>
