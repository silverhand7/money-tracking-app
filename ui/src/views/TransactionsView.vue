<template>
    <div>
        Transactions {{ wallet.name }}
        <ul>
            <li v-for="transaction in transactions">
                {{ transaction.nominal }} {{ transaction.date_time }} {{ transaction.type }}
                {{ transaction }}
            </li>
        </ul>
    </div>
</template>

<script>
import axios from 'axios';
import config from '@/config.js'

export default {
    data() {
        return {
            wallet: '',
            transactions: '',
        }
    },
    created() {
        axios.get(config.basePath + "/api/wallets/" + this.$route.params.walletId, {
            headers: {
                'Authorization': 'Bearer ' + JSON.parse(localStorage.getItem('user')).api_key
            }
        })
        .then((response) => {
            this.wallet = response.data.data
        })
        .catch((error) => {
            console.error('Error fetching data:', error);
        });

        axios.get(config.basePath + "/api/wallets/" + this.$route.params.walletId + '/transactions', {
            headers: {
                'Authorization': 'Bearer ' + JSON.parse(localStorage.getItem('user')).api_key
            }
        })
        .then((response) => {
            this.transactions = response.data.data
        })
        .catch((error) => {
            console.error('Error fetching data:', error);
        });
    }
}
</script>