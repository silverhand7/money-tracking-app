<template>
    <CardContainer>
        <PageTitle title="Create New Wallet" />
        <FormInput
            :model-value="name"
            @update:model-value="newValue => name = newValue"
            label="Name"
            placeholder="Wallet Name"
        />
        <ErrorField
            v-if="nameError"
            :message="nameError"
        />

        <FormSelectOption
            model-value=""
            @update:model-value="newValue => currency = newValue"
            :options="currencyOptions"
            label="Currency"
        />
        <ErrorField
            v-if="currencyError"
            :message="currencyError"
        />

        <FormNumber
            :model-value="balance"
            @update:model-value="newValue => balance = newValue"
            label="Balance"
            placeholder="Initial Balance"
        />
        <ErrorField
            v-if="balanceError"
            :message="balanceError"
        />

        <Button
            @click="addWallet"
            class="btn-neutral mt-4"
            text="Save"
            type="submit"
        />
    </CardContainer>
</template>

<script>
import CardContainer from '@/components/CardContainer.vue';
import PageTitle from '@/components/PageTitle.vue';
import FormInput from '@/components/FormInput.vue';
import FormNumber from '@/components/FormNumber.vue';
import Button from '@/components/Button.vue';
import ErrorField from '@/components/ErrorField.vue';
import FormSelectOption from '@/components/FormSelectOption.vue';
import config from '@/config.js'
import axios from 'axios'
import router from '../router';

export default {
    components: {
        CardContainer,
        PageTitle,
        FormInput,
        FormNumber,
        Button,
        ErrorField,
        FormSelectOption
    },
    data() {
        return {
            name: '',
            nameError: '',
            currency: '',
            currencyError: '',
            currencyOptions: [
                {
                    name: 'USD',
                    value: 'USD'
                },
                {
                    name: 'IDR',
                    value: 'IDR'
                },
                {
                    name: 'SGD',
                    value: 'SGD'
                },
            ],
            balance: "",
            balanceError: "",
        }
    },
    methods: {
        addWallet() {
            axios.post(
                config.basePath + "/api/wallets",
                {
                    name: this.name,
                    currency: this.currency,
                    balance: parseInt(this.balance.replace(/[^0-9]/g, '')),
                },
                {
                    headers: {
                        'Content-Type': 'text/plain; charset=utf-8',
                        'Authorization': 'Bearer ' + JSON.parse(localStorage.getItem('user')).api_key
                    }
                }
            )
            .then((response) => {
                router.push({name: 'wallets'})
            })
        }
    }
}

</script>