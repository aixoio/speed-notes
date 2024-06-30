<template>
  <div class="flex justify-between border-b border-b-gray-950">
    <div class="flex items-center align-middle">
      <h1 class="text-3xl cursor-pointer">Speed notes</h1>
    </div>
    <div class="flex gap-2">
      <button class="btn" @click="dash">Dashboard</button>
      <button class="btn" @click="">New note</button>
      <button class="btn" @click="">Logout</button>
    </div>
  </div>
</template>

<script lang="ts" setup>
import router from '@/router';

function dash() {
  router.push({ name: "dash" })
}


</script>
