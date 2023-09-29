<template>
    <!-- <CategoryForm class="category-form mb-4" @on-submit="addCategory" /> -->
    <div class="text-right mb-2">
        <RouterLink :to="{ name: 'wallets.create' }" class="btn btn-secondary mb-2">Add New</RouterLink>
    </div>
    <WalletList
        v-for="wallet in wallets"
        :key="wallet.id"
        :wallet="wallet"
    />
</template>

<script>
import CategoryForm from '@/components/Forms/CategoryForm.vue';
import WalletList from '@/components/WalletList.vue';
import PageTitle from '@/components/PageTitle.vue';
import axios from 'axios'
import config from '@/config.js'

export default {
    components: {
        CategoryForm,
        WalletList,
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