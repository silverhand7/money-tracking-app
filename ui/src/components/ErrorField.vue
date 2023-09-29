<template>
    <div class="text-red-500">
        {{ message }}
    </div>
</template>

<script>

export default {
    props: {
        message: {
            type: String,
            default: ''
        }
    }
}

</script>