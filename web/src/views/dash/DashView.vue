
<template>
    <div>
        <DashNavBar></DashNavBar>
        <div class="m-4">
            <h1 class="text-2xl">Notes</h1>
        </div>
        <div class="grid grid-cols-4 gap-4 mr-60 ml-60">
            <div v-for="note in notesStore.notes">
                <DashNote :note="note"></DashNote>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { Notes } from '@/assets/ts/tools/notes';
import DashNavBar from '@/components/DashNavBar.vue';
import DashNote from '@/components/DashNote.vue';
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
