<template>
    <CardContainer class="mb-4">
        <div class="text-center">
            <span class="text-xs text-gray-400">{{ wallet.name }}</span>
            <h1 class="text-2xl">{{ wallet.currency + ' ' + wallet.balance }}</h1>
        </div>
    </CardContainer>

    <CardContainer
        v-for="(transactionList, date) in transactions"
        :key="date"
        class="mb-4"
    >

        <div class="flex gap-3 mb-2">
            <div class="text-4xl font-bold">
                {{ splitDate(date)[1] }}
            </div>
            <div class="font-normal text-base">
                <div>
                    {{ splitDate(date)[0] }}
                </div>
                <div class="text-gray-400">
                    {{ splitDate(date)[2] + ' ' + splitDate(date)[3] }}
                </div>
            </div>
        </div>
        <div
            v-for="transaction in transactionList"
            class="flex gap-3 justify-between mb-4"
        >
            <div class="flex gap-4">
                <div class="w-10">
                    <img src="@/assets/images/wallet-image.png" alt="Icon">
                </div>
                <div>
                    {{ transaction.name }}
                    <p class="text-sm text-gray-400">
                        {{ transaction.note }}
                    </p>
                </div>
            </div>
            <div :class="[transaction.type == 'E' ? 'text-red-500' : 'text-green-500']">
                {{ transaction.nominal }}
            </div>
        </div>
    </CardContainer>
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
                value.date = dayjs(value.date_time).format('dddd D MMM YYYY')
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
    },
    methods: {
        splitDate(date) {
            return date.split(' ') // ["Saturday", "30", "Sep", "2023"]
        }
    }
}
</script>