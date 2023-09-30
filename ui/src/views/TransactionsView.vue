<template>
    <CardContainer>
        <div class="text-center">
            <span class="text-xs text-gray-400">{{ wallet.name }}</span>
            <h1 class="text-2xl">{{ wallet.currency + ' ' + wallet.balance }}</h1>
        </div>
    </CardContainer>
    <div>
        <ul v-for="(transactionList, date) in transactions"
            :key="date"
        >
            <li v-for="transaction in transactionList"
                :key="transaction.id"
            >
                {{ date }} <br>
                {{ transaction.nominal }}
            </li>
        </ul>
    </div>
</template>

<script>
import axios from 'axios';
import config from '@/config.js'
import CardContainer from '@/components/CardContainer.vue'
import dayjs from 'dayjs'

export default {
    components: {
        CardContainer
    },
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
            let transactions = response.data.data
            transactions.map((value) => {
                value.date = dayjs(value.date_time).format('YYYY-MM-DD')
                return value
            })
            let grouped = transactions.reduce(function (r, a) {
                r[a.date] = r[a.date] || [];
                r[a.date].push(a);
                return r;
            }, Object.create(null));

            this.transactions = grouped
        })
        .catch((error) => {
            console.error('Error fetching data:', error);
        });
    }
}
</script>