<template>
    <button class="btn w-full" :class="class" :type="type">{{ text }}</button>
</template>

<script>
export default {
    props:{
        class: {
            type: String,
            default: "",
        },
        text: {
            type: String,
            default: "Submit"
        },
        type: {
            type: String,
            default: "Submit"
        },
    }
}

</script>