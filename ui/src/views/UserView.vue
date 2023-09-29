<template>
    <div>
        <h1 class="text-2xl mb-5">{{ user.name }}</h1>

        <ul class="w-full menu menu-md bg-white rounded-box">
            <li><a>md item 1</a></li>
            <li><a>md item 2</a></li>
            <li>
                <a
                    @click="logout"
                    class="bg-red-500 text-white hover:bg-red-400 hover:text-white"
                >
                    Logout
                </a>
            </li>
        </ul>
    </div>
</template>

<script>
import router from '../router';

export default {
    data() {
        return {
            user: JSON.parse(localStorage.getItem("user"))
        }
    },
    methods: {
        logout() {
            localStorage.removeItem("user")
            router.push({name: 'login'})
        }
    },
}

</script>