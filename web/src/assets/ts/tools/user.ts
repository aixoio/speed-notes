
export interface UserReply {
    status?: string
    error?: string
}

export async function signupUser(username: string, password: string): Promise<UserReply> {
    const result = await fetch("/api/signup")
    const data: UserReply = await result.json()

    return data
}
