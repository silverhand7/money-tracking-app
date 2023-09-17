<template>
    <CategoryForm class="category-form mb-4" @on-submit="addCategory" />
</template>

<script>
import CategoryForm from '@/components/Forms/CategoryForm.vue';
import axios from 'axios'
import config from '@/config.js'
import router from '../router';

export default {
    components: {
        CategoryForm,
    },
    methods:{
        addCategory(input) {
            axios.post(
                config.basePath + "/api/categories",
                {
                    name: input.name,
                    icon: input.icon,
                    type: input.type
                },
                {
                    headers: {
                        'Content-Type': 'text/plain; charset=utf-8'
                    }
                }
            )
            .then((response) => {
                router.push({name: 'categories'})
            })
        }
    }
}

</script>