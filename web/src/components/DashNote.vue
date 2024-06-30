<template>
    <div class="border border-gray-950 p-2 rounded bg-slate-100 shadow cursor-pointer">
        <div class="flex gap-2 justify-between">
            <span class="font-extrabold text-lg text-wrap truncate">{{ props.note.title }}</span>
            <button
                class="p-2 m-1 border border-gray-950 rounded-2xl bg-gray-100 text-xl cursor-pointer hover:bg-transparent shadow hover:shadow-lg">
                <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px"
                    class="fill-gray-950">
                    <path
                        d="M280-120q-33 0-56.5-23.5T200-200v-520h-40v-80h200v-40h240v40h200v80h-40v520q0 33-23.5 56.5T680-120H280Zm400-600H280v520h400v-520ZM360-280h80v-360h-80v360Zm160 0h80v-360h-80v360ZM280-720v520-520Z" />
                </svg>
            </button>
        </div>
        <p class="truncate">{{ props.note.contents }}</p>
    </div>
</template>

<script lang="ts" setup>
import type { Note } from '@/assets/ts/data/note';


const props = defineProps<{
    note: Note
}>()

</script>
