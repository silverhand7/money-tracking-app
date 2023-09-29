<template>
    <!-- <CategoryForm class="category-form mb-4" @on-submit="addCategory" /> -->
    <div class="text-right mb-2">
        <RouterLink :to="{ name: 'wallets.create' }" class="btn btn-secondary mb-2">Add New</RouterLink>
    </div>
    <CardContainer class="overflow-x-auto h-fit">
        <PageTitle title="List Wallets" />
        <div class="overflow-x-auto">
            <table class="table">
                <thead>
                    <tr>
                        <th>#</th>
                        <th>Name</th>
                        <th>Currency</th>
                        <th>Balance</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="wallet in wallets">
                        <td>{{ wallet.id }}</td>
                        <td>{{ wallet.name }}</td>
                        <td>{{ wallet.currency }}</td>
                        <td>{{ wallet.balance }}</td>
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
            wallets: []
        }
    },
    created() {
        axios.get(config.basePath + "/api/wallets", {
            headers: {
                'Authorization': 'Bearer bfe61cc1e05322127f7d5a0288b98f3701d627cefdd8659d58f5354b7e7d7d9d',
            }
        })
        .then((response) => {
            this.wallets = response.data.data
        })
        .catch((error) => {
            console.error('Error fetching data:', error);
        });
    },
}

</script>