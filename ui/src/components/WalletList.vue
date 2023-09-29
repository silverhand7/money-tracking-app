<template>
    <div class="card card-side bg-base-100 shadow-sm mb-4">
        <figure class="w-1/4">
            <img src="@/assets/images/wallet-image.png" alt="Wallet Icon" class="pl-8"/>
        </figure>
        <div class="card-body">
            <h2 class="card-title">{{ wallet.name }}</h2>
            <p>Balance: {{ wallet.currency + ' ' + convert(wallet.balance) }}</p>
            <div class="card-actions justify-start">
                <RouterLink
                    :to="{
                        name: 'wallets.transactions',
                        params: {
                            walletId: wallet.id
                        }
                    }"
                    class="btn btn-primary btn-sm"
                >
                    See Transactions
                </RouterLink>
            </div>
        </div>
    </div>
</template>

<script>
import { RouterLink } from 'vue-router';

export default {
    props: {
        wallet: Object
    },
    methods: {
        convert(number) {
            return number.toString().replace(/\D/g, "").replace(/\B(?=(\d{3})+(?!\d))/g, ",");
        }
    },
    components: { RouterLink, RouterLink }
}
</script>