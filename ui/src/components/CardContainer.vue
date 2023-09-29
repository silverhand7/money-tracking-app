<template>
    <div class="card bg-base-100 shadow-sm" :class="class">
        <div class="card-body">
            <slot></slot>
        </div>
    </div>
</template>

<script>

export default {
    props: {
        class: ''
    }
}

</script>