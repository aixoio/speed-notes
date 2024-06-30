import type { Note } from "../data/note"

export interface Note_reply {
    note_id?: number
    error?: string
}

export async function NewNote(jwt: string, title: string, content: string): Promise<Note_reply> {
    const result = await fetch("/api/notes/new", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${jwt}`,
        },
        body: JSON.stringify({
            title,
            content,
        }),
    })
    const data: Note_reply = await result.json()

    return data
}

export async function Notes(jwt: string): Promise<Note[]> {
    const result = await fetch("/api/notes", {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${jwt}`,
        },
    })
    const data: Note[] = await result.json()

    return data
}
