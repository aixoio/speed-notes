<template>
  <div class="flex justify-between border-b border-b-gray-950">
    <div class="flex items-center align-middle">
      <h1 class="text-3xl cursor-pointer" @click="home">Speed notes</h1>
    </div>
    <div class="flex gap-2">
      <button class="btn" @click="signup">Signup</button>
      <button class="btn" @click="login">Login</button>
    </div>
  </div>
</template>

<script lang="ts" setup>
import router from '@/router';

function home() {
  router.push({ name: "home" })
}

function signup() {
  router.push({ name: "signup" })
}

function login() {
  router.push({ name: "login" })
}

</script>
