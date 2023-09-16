<template>
    <div class="form-control w-full">
        <label class="label">
            <span class="label-text">{{ label }}</span>
        </label>
        <input :type="type" :placeholder="placeholder" class="input input-bordered" />
    </div>
</template>

<script>

export default {
    props:{
        label: {
            type: String,
            default: ""
        },
        type: {
            type: String,
            default: "text"
        },
        placeholder: {
            type: String,
            default: "Input here"
        },
    }
}

</script>