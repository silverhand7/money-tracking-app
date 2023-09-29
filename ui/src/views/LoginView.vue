<template>
    <CardContainer>
        <FormInput
            :model-value="email"
            @update:model-value="newValue => email = newValue"
            label="Email"
            placeholder="Email Address"
        />
        <ErrorField
            v-if="emailError"
            :message="emailError"
        />

        <FormInput
            :model-value="password"
            @update:model-value="newValue => password = newValue"
            label="Password"
            placeholder="Password"
            type="password"
        />

        <Button
            @click="login"
            class="btn-neutral mt-5"
            text="Login"
            type="button"
        />
    </CardContainer>
</template>

<script>
import CardContainer from '@/components/CardContainer.vue';
import FormInput from '@/components/FormInput.vue';
import Button from '@/components/Button.vue';
import ErrorField from '@/components/ErrorField.vue';
import axios from 'axios'
import config from '@/config.js'
import router from '../router';

export default {
    data(){
        return {
            email: 'dummy@email.com',
            emailError:'',
            password: 'password',
        }
    },
    components: {
        CardContainer,
        FormInput,
        Button,
        ErrorField,
    },
    methods: {
        login() {
            axios.post(
                config.basePath + "/api/login",
                {
                    email: this.email,
                    password: this.password,
                },
                {
                    headers: {
                        'Content-Type': 'text/plain; charset=utf-8'
                    }
                }
            )
            .then((response) => {
                localStorage.setItem("user", JSON.stringify(response.data.data))
                router.push({name: 'home'})
            }).catch(err => {
                if (err.response.status == 404) {
                    this.emailError = err.response.data.data
                }
            })
        }
    }
}

</script>