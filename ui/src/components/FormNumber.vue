<template>
    <div class="form-control w-full">
        <label class="label">
            <span class="label-text">{{ label }}</span>
        </label>
        <input
            :value="value"
            @input="$emit('update:modelValue', $event.target.value)"
            type="text"
            :placeholder="placeholder"
            class="input input-bordered"
            @keyup="convert()"
        />
    </div>
</template>

<script>
export default {
    props:{
        modelValue: {
            type: String,
            default: 0,
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
            value: this.$props.modelValue
        }
    },
    methods: {
        convert() {
            const result = this.$props.modelValue.replace(/\D/g, "").replace(/\B(?=(\d{3})+(?!\d))/g, ",");
            this.value = result
            // there's a bug when
            // 50.000, becomes 5.0000 in the parent that use this component
        }
    }
}
</script>