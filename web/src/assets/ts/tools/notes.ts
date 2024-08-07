import type { Note } from "../data/note"

export interface Note_reply {
    note_id?: number
    error?: string
}

export interface Delete_Note_reply {
    status?: string
    error?: string
}

export interface Update_Note_reply {
    status?: string
    error?: string
}

export async function UpdateNote(jwt: string, title: string, content: string, note_id: number): Promise<Update_Note_reply> {
    const result = await fetch(`/api/notes/update/${note_id}`, {
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
    const data: Update_Note_reply = await result.json()

    return data
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


export interface Note_reply {
    id?: number
    user_id?: number
    title?: string
    contents?: string
    error?: string
}


export async function GetNote(jwt: string, note_id: number): Promise<Note_reply> {
    const result = await fetch(`/api/notes/get/${note_id}`, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${jwt}`,
        },
    })
    const data: Note_reply = await result.json()

    return data
}

export async function DeleteNote(jwt: string, note_id: number): Promise<Delete_Note_reply> {
    const result = await fetch(`/api/notes/delete/${note_id}`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${jwt}`,
        },
    })
    const data: Delete_Note_reply = await result.json()

    return data
}
