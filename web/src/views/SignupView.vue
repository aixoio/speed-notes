<template>
    <div>
        <NavBar></NavBar>
        <div class="flex justify-center m-4">
            <h1 class="text-4xl font-mono">Signup</h1>
        </div>
        <div class="flex justify-center">
            <div class="grid grid-cols-1">
                <span>Username</span>
                <input type="text" placeholder="Username" class="border border-gray-950 rounded-lg m-4 bg-gray-50 p-2 outline-none" maxlength="255">
                <span>Password</span>
                <input type="password" placeholder="Password" class="border border-gray-950 rounded-lg m-4 bg-gray-50 p-2 outline-none">
                <button class="btn">Signup</button>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import NavBar from '@/components/NavBar.vue';

</script>
