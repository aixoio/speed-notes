
<template>
    <div>
        <DashNavBar></DashNavBar>
        <div class="m-4">
            <h1 class="text-2xl">Notes</h1>
        </div>
        <div class="grid grid-cols-4 gap-4">
            <div v-for="note in notesStore.notes">
                <div class="border border-gray-950 p-2 rounded bg-slate-100 shadow cursor-pointer">
                    <span class="font-extrabold text-lg">{{ note.title }}</span>
                    <p>{{ note.contents }}</p>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { Notes } from '@/assets/ts/tools/notes';
import DashNavBar from '@/components/DashNavBar.vue';
import router from '@/router';
import { useNotesStore } from '@/stores/notesstore';
import { useUserStore } from '@/stores/userstore';
import { onMounted } from 'vue';

const notesStore = useNotesStore()
const userStore = useUserStore()

onMounted(async () => {
    if (userStore.jwt == null) {
        router.push({ name: "login" })
    }
    notesStore.setNotes(await Notes(userStore.jwt as string))
})

</script>
