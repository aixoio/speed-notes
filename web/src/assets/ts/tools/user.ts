
export interface UserReply {
    status?: string
    error?: string
}

export async function signupUser(username: string, password: string): Promise<UserReply> {
    const result = await fetch("/api/signup", {
        method: "POST",
        body: JSON.stringify({
            username,
            password
        }),
        headers: {
            "Content-Type": "application/json"
        }
    })
    const data: UserReply = await result.json()

    return data
}
