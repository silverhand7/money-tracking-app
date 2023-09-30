<template>
    <CardContainer>
        <PageTitle title="Add New Transaction" />
        <FormSelectOption
            v-if="walletOptions"
            model-value=""
            @update:model-value="newValue => wallet_id = newValue"
            :options="walletOptions"
            label="Wallet"
        />

        <FormSelectOption
            v-if="categoryOptions"
            model-value=""
            @update:model-value="newValue => category_id = newValue"
            :options="categoryOptions"
            label="Category"
        />

        <FormNumber
            :model-value="nominal"
            @update:model-value="newValue => nominal = newValue"
            label="Nominal"
            placeholder="Nominal"
        />

        <FormInput
            :model-value="note"
            @update:model-value="newValue => note = newValue"
            label="Note"
            placeholder="Note"
        />

        <Button
            @click="addTransaction"
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
import axios from 'axios'
import config from '@/config.js'

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
            nominal: '',
            category_id: '',
            wallet_id: '',
            note: '',
            walletOptions: '',
            categoryOptions: '',
        }
    },
    created() {
        axios.get(config.basePath + "/api/categories")
        .then((response) => {
            this.categoryOptions = response.data.data.map((value) => {
                return {
                    name: value.name,
                    value: value.id,
                }
            })
            console.log(this.categoryOptions)
        })
        .catch((error) => {
            console.error('Error fetching data:', error);
        });

        axios.get(config.basePath + "/api/wallets", {
            headers: {
                'Authorization': 'Bearer ' + JSON.parse(localStorage.getItem('user')).api_key
            }
        })
        .then((response) => {
            this.walletOptions = response.data.data.map((value) => {
                return {
                    name: value.name,
                    value: value.id,
                }
            })
        })
        .catch((error) => {
            console.error('Error fetching data:', error);
        });

    }
}

</script>