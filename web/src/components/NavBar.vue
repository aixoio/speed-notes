<template>
    <div class="flex justify-between border-b border-b-gray-950">
      <div class="flex items-center align-middle">
        <h1 class="text-3xl">Speed notes</h1>
      </div>
      <div class="flex gap-2">
        <button class="btn" @click="signup">Signup</button>
        <button class="btn">Login</button>
      </div>
    </div>
</template>

<script lang="ts" setup>
import router from '@/router';


function signup() {
  router.push("signup")
}

</script>
