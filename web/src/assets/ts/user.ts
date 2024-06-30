
export function UsernameIsValid(username: string): boolean {
    return !new RegExp("[^A-z0-9]+").test(username)
}
