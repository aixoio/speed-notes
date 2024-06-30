
<template>
    <div>
        <DashNavBar></DashNavBar>
        <div class="m-4">
            <h1 class="text-2xl">Notes</h1>
        </div>
        <div class="grid grid-cols-4 gap-4">
            <div class="border border-gray-950 p-2 rounded bg-slate-100 shadow cursor-pointer">
                <span class="font-extrabold text-lg">[NOTE TITLE]</span>
                <p>NOTE DATA</p>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import DashNavBar from '@/components/DashNavBar.vue';

</script>
