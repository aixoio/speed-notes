<template>
    <div>
        <DashNavBar></DashNavBar>
        <div class="flex justify-center m-4">
            <h1 class="text-4xl font-mono">Editing note</h1>
        </div>
        <div class="flex justify-center">
            <div class="grid grid-cols-1">
                <span>Title</span>
                <input type="text" placeholder="Title"
                    class="border border-gray-950 rounded-lg m-4 bg-gray-50 p-2 outline-none" v-model="title">
                <span>Note</span>
                <textarea type="text" placeholder="Note"
                    class="border border-gray-950 rounded-lg m-4 bg-gray-50 p-2 outline-none resize-none"
                    v-model="content" rows="5"></textarea>
                <button class="btn" @click.prevent="save">Save</button>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import type { Note } from '@/assets/ts/data/note';
import { GetNote, UpdateNote } from '@/assets/ts/tools/notes';
import DashNavBar from '@/components/DashNavBar.vue';
import router from '@/router';
import { useUserStore } from '@/stores/userstore';
import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute()
const userStore = useUserStore()

const note = ref<Note|null>(null)
const title = ref("")
const content = ref("")

onMounted(async () => {
    if (userStore.jwt == null) {
        router.push({ name: "login" })
    }

    const result = await GetNote(userStore.jwt as string, parseInt(route.params.id as string))
    if (result.error != null) {
        router.push({ name: "dash" })
        return
    }

    note.value = {
        contents: result.contents as string,
        id: result.id as number,
        title: result.title as string,
        user_id: result.user_id as number,
    }

    title.value = note.value.title
    content.value = note.value.contents

})

async function save() {
    const result = await UpdateNote(userStore.jwt as string, title.value, content.value, note.value?.id as number)
    if (result.error != null) {
        alert(result.error)
        return
    }

    router.push({ name: "dash" })
}

</script>
