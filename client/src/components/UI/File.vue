<template>
    <div class="file-wrapper">
        <div class="download-wrapper">
            <img class="download-icon" :src="DownloadIcon" alt="Remove" />
        </div>
        
        <img class="file-icon" :src="icon" :alt="name" />
        
        <span class="file-name">{{ name }}</span>

        <span class="file-description">{{ description }}</span>
    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import FileIcon from '../../assets/icons/file.icon.svg';
import DownloadIcon from '../../assets/icons/download.icon.png';

export default defineComponent({
    name: "File",
    data() {
        return({
            DownloadIcon,
        });
    },
    props: {
        icon: {
            type: String,
            required: false,
            default: FileIcon,
        },
        name: {
            type: String,
            required: false,
            default: 'File',
        },
        description: {
            type: String,
            required: false,
            default: '',
        },

    },
    methods: {

    },
});
</script>

<style lang="scss" scoped>
@use "/src/assets/styles/colors" as *;
@use "/src/assets/styles/image";
@use "/src/assets/styles/fonts";
@use "/src/assets/styles/shadows";

.file-wrapper {
    border: 2px solid $border-gray;
    background: $background-white;
    border-radius: 10px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: space-between;
    padding: 4px 6px;
    cursor: pointer;
    position: relative;
    transition: ease 0.3s all;
    width: 140px;
    height: 150px;

    @media (max-width: 900px) {
        width: 110px;
        height: 120px;
    }

    &:hover {
        background-color: $background-gray;
        border-color: $border-danger;
        transform: scale(1.025);

        .download-wrapper {
            opacity: 1;
        }
    }

    .download-wrapper {
        transition: ease 0.3s opacity;
        opacity: 0;
        width: 100%;
        height: 100%;
        position: absolute;
        top: 0;
        left: 0;
        display: flex;
        justify-content: center;
        align-items: center;
        
        .download-icon {
            width: 50%;
        }
    }

    .file-icon {
        width: 90px;

        @media (max-width: 900px) {
            width: 60px;
        }
    }

    .file-name {
        @include fonts.noto-font(600);
        @include fonts.responsive-font(13, 13, 1440);
        @include fonts.prevent-selecting;
        @include fonts.hide-overflowed-text;
        color: $text-black;
        text-align: center;
    }

    .file-description {
        width: 100%;
        @include fonts.noto-font(600);
        @include fonts.responsive-font(13, 13, 1440);
        @include fonts.prevent-selecting;
        @include fonts.hide-overflowed-text-multiline(2);
        color: $text-gray;
        text-align: center;
    }
}
</style>