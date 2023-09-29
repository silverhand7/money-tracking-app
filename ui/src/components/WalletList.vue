<template>
    <div class="card card-side bg-base-100 shadow-sm mb-4">
        <figure class="w-1/4">
            <img src="@/assets/images/wallet-image.png" alt="Wallet Icon" class="pl-8"/>
        </figure>
        <div class="card-body">
            <h2 class="card-title">{{ wallet.name }}</h2>
            <p>Balance: {{ wallet.currency + ' ' + wallet.balance }}</p>
            <div class="card-actions justify-start">
                <button class="btn btn-primary btn-sm">See Transactions</button>
            </div>
        </div>
    </div>
</template>

<script>

export default {
    props: {
        wallet: Object
    }
}
</script>