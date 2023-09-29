<template>
    <div class="form-control w-full">
        <label class="label">
            <span class="label-text">{{ label }}</span>
        </label>
        <input
            :value="modelValue"
            @input="$emit('update:modelValue', $event.target.value)"
            :type="type"
            :placeholder="placeholder"
            class="input input-bordered"
        />
    </div>
</template>

<script>

export default {
    props:{
        modelValue: {
            type: String,
            default: "",
        },
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
    },
    data(){
        return {
            input: ""
        }
    }
}

</script>