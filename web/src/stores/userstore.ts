import { defineStore } from "pinia";
import { ref } from "vue";

export const useUserStore = defineStore("user", () => {
    const jwt = ref<string | null>(localStorage.getItem("speed_notes_jwt"))

    function setJwt(jwt_str: string) {
        jwt.value = jwt_str
        localStorage.setItem("speed_notes_jwt", jwt_str)
    }

    function logout() {
        jwt.value = null
        localStorage.removeItem("speed_notes_jwt")
    }

    return {
        jwt,
        setJwt,
        logout,
    }
})
