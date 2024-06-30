<template>
    <div>
        <DashNavBar></DashNavBar>
        <div class="flex justify-center m-4">
            <h1 class="text-4xl font-mono">New note</h1>
        </div>
        <div class="flex justify-center">
            <div class="grid grid-cols-1">
                <span>Title</span>
                <input type="text" placeholder="Title"
                    class="border border-gray-950 rounded-lg m-4 bg-gray-50 p-2 outline-none"
                    v-model="title">
                <span>Note</span>
                <textarea type="text" placeholder="Note"
                    class="border border-gray-950 rounded-lg m-4 bg-gray-50 p-2 outline-none resize-none" v-model="content" rows="5"></textarea>
                <span class="bg-red-200 p-2 m-2 border border-red-700 rounded-lg text-red-950" v-show="error != ''">{{
                    error }}</span>
                <button class="btn" @click.prevent="create">Create</button>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { NewNote } from '@/assets/ts/tools/notes';
import DashNavBar from '@/components/DashNavBar.vue';
import router from '@/router';
import { useNotesStore } from '@/stores/notesstore';
import { useUserStore } from '@/stores/userstore';
import { isEmpty } from 'lodash';
import { onMounted, ref } from 'vue';

const title = ref("")
const content = ref("")
const error = ref("")

const userStore = useUserStore()
const notesStore = useNotesStore()

onMounted(async () => {
    if (userStore.jwt == null) {
        router.push({ name: "login" })
    }
})


async function create() {
    if (isEmpty(title.value.trim()) || isEmpty(content.value.trim())) {
        error.value = "You must enter a title and content"
        return
    }

    const res = await NewNote(userStore.jwt as string, title.value, content.value)
    if (res.error != null) {
        error.value = res.error as string
        return
    }

    

    error.value = ""
}

</script>
