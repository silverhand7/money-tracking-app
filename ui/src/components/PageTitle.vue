<template>
    <h1 class="font-bold text-lg">{{ title }}</h1>
</template>

<script>
export default {
    props:{
        title: {
            type: String,
            default: ""
        },
    }
}

</script>