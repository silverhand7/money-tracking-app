<template>
    <!-- <CategoryForm class="category-form mb-4" @on-submit="addCategory" /> -->
    <div class="text-right mb-2">
        <RouterLink :to="{ name: 'categories.create' }" class="btn btn-secondary mb-2">Add New</RouterLink>
    </div>
    <CardContainer class="overflow-x-auto h-fit">
        <PageTitle title="List Categories" />
        <div class="overflow-x-auto">
            <table class="table">
                <thead>
                    <tr>
                        <th>#</th>
                        <th>Name</th>
                        <th>Type</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="category in categories">
                        <th>{{ category.id }}</th>
                        <td>{{ category.name }}</td>
                        <td><div class="badge" :class="[category.type == 'E' ? 'badge-secondary' : 'badge-primary']">{{ category.type == 'E' ? 'expense' : 'income' }}</div></td>
                    </tr>
                </tbody>
            </table>
        </div>
    </CardContainer>
</template>

<script>
import CategoryForm from '@/components/Forms/CategoryForm.vue';
import CardContainer from '@/components/CardContainer.vue';
import PageTitle from '@/components/PageTitle.vue';
import axios from 'axios'
import config from '@/config.js'

export default {
    components: {
        CategoryForm,
        CardContainer,
        PageTitle,
    },
    data() {
        return {
            categories: []
        }
    },
    created() {
        axios.get(config.basePath + "/api/categories")
        .then((response) => {
            this.categories = response.data.data
        })
        .catch((error) => {
            console.error('Error fetching data:', error);
        });


    },
}

</script>