<template>
    <div>
        <DashNavBar></DashNavBar>
        <div class="m-4">
            <h1 class="text-2xl">Editing note - {{ $route.params.id }}</h1>
        </div>
    </div>
</template>

<script lang="ts" setup>
import DashNavBar from '@/components/DashNavBar.vue';

</script>
