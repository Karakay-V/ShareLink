<template>
    <div class="view-wrapper">
        <div class="info-container">
            <h2 class="title">
                Easily share your study materials with classmates
            </h2>

            <div class="quired_info">
                <p class="classroom">Classroom: {{ classroom }}</p>
                <p class="lesson">Lesson: {{ lesson }}</p>
            </div>
        </div>

        <Section>
            <TextInput  v-model:data="email" 
                        :type="InputDataTypes.Email"
                        placeholder="andrii@example.com"
                        :required="true" />
            
            <TextInput  v-model:data="signature"
                        :type="InputDataTypes.Text"
                        placeholder="Signature (e.g., Andrii, Mykola BIP2-22)"
                        :required="true" />

            <DragAndDropArea />
            
            <Button label="Submit File"
                    class="internal-button"
                    @click="handleSubmit" />

        </Section>
    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Section from './UI/Section.vue';
import DragAndDropArea from './UI/DragAndDropArea.vue';
import Button from './UI/Button.vue';
import TextInput from './UI/TextInput.vue';
import { InputDataTypes } from '../types/input-data-types';
import { uploadFile } from "../services/upload-service";

export default defineComponent({
    name: "UploadView",
    data() {
        return({
            InputDataTypes,
            email: '',
            signature: '',
            file: null as File | null,
        });
    },
    props: {
        lesson: {
            type: String,
            required: true,
        },
        classroom: {
            type: String,
            required: true,
        },
    },
    methods: {
        async handleSubmit() {
        try {
            const res = await uploadFile({
                email: this.email,
                signature: this.signature,
                lesson: this.lesson,
                classroom: this.classroom,
                file: this.file ?? undefined,
            });
            console.log("Upload success:", res.data);
            alert("File submitted successfully!");
        } catch (err) {
            console.error("Upload error:", err);
            alert("Failed to submit file.");
        }
        },
    },
    components: {
        Section,
        TextInput,
        DragAndDropArea,
        Button,
    },
});
</script>

<style lang="scss" scoped>
@use "/src/assets/styles/colors" as *;
@use "/src/assets/styles/image";
@use "/src/assets/styles/fonts";
@use "/src/assets/styles/shadows";


.view-wrapper {
    margin: 33px;
    display: flex;
    gap: 20px;
    flex-direction: column;
    justify-content: space-between;
    align-items: center;
}

.info-container {
    max-width: 500px;
    
    h2, p {
        margin: 0;
        padding: 0;
    }

    .title {
        @include fonts.noto-font(500);
        @include fonts.responsive-font(24, 20, 1440);
        text-align: center;
    }

    .quired_info {
        margin-top: 20px;
        @include fonts.noto-font(500);
        @include fonts.responsive-font(16, 14, 1440);
        text-align: start;


        gap: 5px;
        flex-direction: column;
        justify-content: space-between;
        align-items: flex-start;
    }
}

.internal-button {
    margin-top: 10px;
}
</style>