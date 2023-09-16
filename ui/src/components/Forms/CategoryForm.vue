<template>
    <CardContainer>
        <PageTitle title="Create New Category" />
        <FormInput
            label="Name"
            type="text"
            placeholder="Input..."
        />
        <FormInput
            label="Icon Name"
            type="text"
            placeholder="Input..."
            class="mb-3"
        />

        <Button
            class="btn-neutral"
            text="Save"
        />
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
    }
}

</script>