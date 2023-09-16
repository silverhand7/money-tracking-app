<template>
    <CardContainer>
        <PageTitle title="Create New Category" />

        <form @submit.prevent="$emit('on-submit', form)" method="post">
            <div class="form-control w-full">
                <label class="label">
                    <span class="label-text">Name</span>
                </label>
                <input v-model="form.name" type="text" placeholder="Input..." class="input input-bordered" />
            </div>

            <div class="form-control w-full">
                <label class="label">
                    <span class="label-text">Icon</span>
                </label>
                <input v-model="form.icon" type="text" placeholder="Input..." class="input input-bordered" />
            </div>

            <div class="form-control w-full mb-3">
                <label class="label">
                    <span class="label-text">Type</span>
                </label>
                <select
                    v-model="form.type"
                    class="select select-bordered"
                >
                    <option value="E" selected>Expense</option>
                    <option value="I">Income</option>
                </select>
            </div>

            <Button
                class="btn-neutral"
                text="Save"
                type="submit"
            />
        </form>

      </CardContainer>
</template>


<script>
import CardContainer from '@/components/CardContainer.vue';
import FormInput from '@/components/FormInput.vue';
import PageTitle from '@/components/PageTitle.vue';
import Button from '@/components/Button.vue';

export default {
    components: {
        CardContainer,
        FormInput,
        PageTitle,
        Button
    },
    data() {
        return {
            form: {
                name: "",
                icon: "",
                type: "E",
            }
        }
    },
    methods:{

    }
}

</script>