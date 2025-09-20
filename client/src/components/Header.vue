<script setup lang="ts">
import { ref } from 'vue';
import { useFixedHeader } from 'vue-use-fixed-header';

const headerRef = ref(null);

const { styles } = useFixedHeader(headerRef)
</script>

<template>
    <div class="header_indent"></div>

    <header id="header" ref="headerRef" :style="styles" class="header">
        <Logo />

        <Search />
    </header>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Logo from './UI/Logo.vue';
import Search from './Search.vue';

export default defineComponent({
    name: "Header",
    components: {
        Logo,
        Search,
    },
});
</script>

<style lang="scss" scoped>
@use "/src/assets/styles/colors" as *;
@use "/src/assets/styles/shadows";

.header_indent {
    height: 72px;
}

.header {
    position: fixed;
    top: 0;
    left: 0;
    height: 72px;
    width: calc(100% - 20px - 15px);
    background-color: $background-gray;
    @include shadows.box-shadow-medium;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    padding: 0 20px 0 15px;
}

</style>