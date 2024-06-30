<template>
    <div>
        <NavBar></NavBar>
        <div class="flex justify-center m-4">
            <h1 class="text-4xl font-mono">Signup</h1>
        </div>
        <div class="flex justify-center">
            <div class="grid grid-cols-1">
                <span>Username</span>
                <input type="text" placeholder="Username"
                    class="border border-gray-950 rounded-lg m-4 bg-gray-50 p-2 outline-none" maxlength="255"
                    v-model="username">
                <span>Password</span>
                <input type="password" placeholder="Password"
                    class="border border-gray-950 rounded-lg m-4 bg-gray-50 p-2 outline-none" v-model="password">
                <span class="bg-red-200 p-2 m-2 border border-red-700 rounded-lg text-red-950" v-show="error != ''">{{ error }}</span>
                <button class="btn" @click.prevent="signup">Signup</button>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import NavBar from '@/components/NavBar.vue';
import { ref } from 'vue';
import { isEmpty } from "lodash"


let username = ref("");
let password = ref("");
let error = ref("");

function signup() {
    if (isEmpty(username.value.trim()) || isEmpty(password.value.trim())) {
        error.value = "You must enter a username and password"
        return
    }
}

</script>
