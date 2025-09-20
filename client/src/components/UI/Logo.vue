<template>
    <div class="logo-container">
        <div class="logo-img-wrapper">
            <img class="logo-img" :src="img" alt="ShareLink" />
        </div>
        
        <h2 class="app_name">{{ name }}</h2>
    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import LogoPic from '../../assets/logo.png';

export default defineComponent({
    name: "Logo",
    data() {
        return({
            LogoPic,
        });
    },
    props: {
        img: {
            type: String,
            required: false,
            default: LogoPic,
        },
        name: {
            type: String,
            required: false,
            default: "ShareLink"
        },
    },
});
</script>

<style lang="scss" scoped>
@use "/src/assets/styles/colors" as *;
@use "/src/assets/styles/image";
@use "/src/assets/styles/fonts";

.logo-container {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
    gap: 8px;
    width: max-content;
    height: max-content;

    .logo-img-wrapper {
        width: 20px;
        max-width: 20px;
        display: flex;
        align-items: center;
        justify-content: center;
        @include image.prevent-manipulations;
        
        .logo-img {
            width: 100%;
        }
    }

    .app_name {
        @include fonts.noto-font(800);
        @include fonts.responsive-font(20, 20, 1440);
        @include fonts.prevent-selecting;
        color: $text-black;
    }
}
</style>