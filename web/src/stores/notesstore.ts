import type { Note } from "@/assets/ts/data/note";
import { defineStore } from "pinia";
import { ref } from "vue";

export const useNotesStore = defineStore("notes", () => {
    const notes = ref<Note[]>([])

    function setNotes(notes_dat: Note[]) {
        notes.value = notes_dat
    }

    function addNote(note: Note) {
        notes.value.push(note)
    }

    return {
        notes,
        setNotes,
        addNote,
    }
})
