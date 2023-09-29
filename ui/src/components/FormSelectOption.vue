<template>
    <div class="w-full">
        <label class="label">
            <span class="label-text">{{ label }}</span>
        </label>

        <select
            :value="selected"
            @input="$emit('update:modelValue', $event.target.value)"
            class="select select-bordered w-full"
        >
            <option v-for="option in options"
                :key="option.value"
                :value="option.value"
            >
                {{ option.name }}
            </option>

        </select>
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
        options: {
            type: Object,
            default: [
                {
                    name: 'Default',
                    value: 0
                },
                {
                    name: 'Value',
                    value: 0,
                }
            ],
        }
    },
    data() {
        return {
            selected: this.$props.modelValue == '' ? this.$props.options[0].value : this.$props.modelValue
        }
    }
}
</script>