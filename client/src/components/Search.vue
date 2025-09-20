<template>
    <div :class="`wrapper ${isFocused ? 'focused' : ''}`">
        <TextInput :type="InputDataTypes.Text"
                    v-model:data="expression"
                    :required="false"
                    placeholder="Search classroom"
                    :message="false"
                    @focus-change="isFocused = $event" />
    </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import TextInput from './UI/TextInput.vue';
import { InputDataTypes } from '../types/input-data-types';

export default defineComponent({
    name: "Search",
    data() {
        return({
            expression: ref(""),
            isFocused: false,
            InputDataTypes,
        });
    },
    components: {
        TextInput,
    },
});
</script>

<style lang="scss" scoped>
@use "/src/assets/styles/colors" as *;
@use "/src/assets/styles/shadows";

.wrapper {
    width: 180px;
    
    &.focused {
        // width: 100%;
    }
}

</style>