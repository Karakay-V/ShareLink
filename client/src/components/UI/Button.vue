<template>
    <button :class="`wrapper ${ isFocused ? 'focused' : '' }`" 
        :style="{ width: width, height: height, }"
        @mouseenter="isFocused = true"
        @mouseleave="isFocused = false"
    >
        <div>
            <p class="label">{{ label }}</p>
        </div>
    </button>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
    name: "Button",
    data() {
        return({
            isFocused: false,
        });
    },
    props: {
        width: {
            type: String,
            required: false,
            default: '100%',
        },
        height: {
            type: String,
            required: false,
            default: '44px',
        },
        label: {
            type: String,
            required: false,
            default: '',
        },
    },
});
</script>

<style lang="scss" scoped>
@use "/src/assets/styles/colors" as *;
@use "/src/assets/styles/image";
@use "/src/assets/styles/fonts";
@use "/src/assets/styles/shadows";

.wrapper {
    background-color: $background-orange;
    border: 2px solid transparent;
    transition: 0.3s ease background-color, 0.3s ease box-shadow, 0.3s ease color, 0.3s ease border-color;
    border-radius: 50px;
    cursor: pointer;
    display: flex;
    justify-content: center;
    align-items: center;

    &.focused {
        @include shadows.box-shadow-medium;

        .label {
            color: $text-gray;
        }
    }

    .label {
        @include fonts.noto-font(600);
        @include fonts.responsive-font(16, 16, 1440);
        @include fonts.prevent-selecting;
        transition: 0.3s ease color;
        text-align: center;
        color: $text-white;
    }
}

</style>