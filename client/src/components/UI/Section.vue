<template>
    <div class="wrapper"
        :style="{ width: width, maxWidth: maxWidth, }">
        <slot />
    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
    name: "Section",
    props: {
        width: {
            type: String,
            required: false,
            default: '100%',
        },
        maxWidth: {
            type: String,
            required: false,
            default: '500px',
        },
    }
});
</script>

<style lang="scss" scoped>
@use "/src/assets/styles/colors" as *;
@use "/src/assets/styles/image";
@use "/src/assets/styles/fonts";
@use "/src/assets/styles/shadows";

.wrapper {
    background-color: $background-white;
    border-radius: 20px;
    @include shadows.box-shadow-medium;

    padding: 24px;
    display: flex;
    gap: 16px 30px;
    flex-direction: column;
    justify-content: space-between;
    align-items: center;
}
</style>