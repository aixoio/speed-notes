<template>
    <div>
        <div class="border border-gray-950 p-2 rounded bg-gray-100 shadow cursor-pointer">
            <div class="flex gap-2 justify-between">
                <span class="font-extrabold text-lg text-wrap truncate">{{ props.note.title }}</span>
                <Popover v-slot="{ open }">
                    <PopoverButton
                        class="outline-none p-2 m-1 border border-gray-950 rounded-2xl bg-gray-100 text-xl cursor-pointer hover:bg-transparent shadow hover:shadow-lg">
                        <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px"
                            class="fill-gray-950">
                            <path
                                d="M280-120q-33 0-56.5-23.5T200-200v-520h-40v-80h200v-40h240v40h200v80h-40v520q0 33-23.5 56.5T680-120H280Zm400-600H280v520h400v-520ZM360-280h80v-360h-80v360Zm160 0h80v-360h-80v360ZM280-720v520-520Z" />
                        </svg>
                    </PopoverButton>

                    <PopoverPanel v-slot="{ close }"
                        class="z-50 bg-slate-50 border border-gray-950 rounded p-2 m-2 absolute cursor-default shadow-xl">
                        <span class="font-bold">Are you sure you want to delete this note?</span>
                        <div class="flex gap-2 justify-end">
                            <button
                                class="p-2 border border-gray-950 rounded-2xl bg-gray-100 text-lg cursor-pointer hover:bg-transparent shadow hover:shadow-md"
                                @click="deleteNote(close)">Yes</button>
                            <button
                                class="p-2 border border-gray-950 rounded-2xl bg-gray-100 text-lg cursor-pointer hover:bg-transparent shadow hover:shadow-md"
                                @click="close">No</button>
                        </div>
                    </PopoverPanel>
                </Popover>

            </div>
            <p class="truncate">{{ props.note.contents }}</p>
        </div>
    </div>
</template>

<script lang="ts" setup>
import type { Note } from '@/assets/ts/data/note';
import { DeleteNote } from '@/assets/ts/tools/notes';
import { useNotesStore } from '@/stores/notesstore';
import { useUserStore } from '@/stores/userstore';
import { Popover, PopoverButton, PopoverPanel } from '@headlessui/vue'

const props = defineProps<{
    note: Note
}>()

const userStore = useUserStore()
const notesStore = useNotesStore();

async function deleteNote(close: any) {
    close()

    const res = await DeleteNote(userStore.jwt as string, props.note.id)
    if (res.error != null) {
        alert("Cannot delete not")
        return
    }

    notesStore.removeNote(props.note)
}

</script>
