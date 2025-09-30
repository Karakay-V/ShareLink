<template>
    <div class="view-wrapper">
        <div class="info-container">
            <h2 class="title">
                Easily share your study materials with classmates
            </h2>

            <div class="quired_info">
                <p class="classroom">Classroom: {{ classroom }}</p>
                <p class="lesson">Lesson: {{ lesson }} ({{ selectedLesson?.start_time }} - {{ selectedLesson?.end_time }})</p>
            </div>
        </div>

        <Section>
            <div class="input-wrapper">
                <label class="input-label">
                    Leave your email address
                </label>
                <TextInput  v-model:data="email" 
                            :type="InputDataTypes.Email"
                            placeholder="andrii@example.com"
                            :required="true" />
            </div>

            <div class="input-wrapper">
                <label class="input-label">
                    And message for your teacher
                </label>
            
                <TextInput  v-model:data="description"
                            :type="InputDataTypes.Text"
                            placeholder="Description (e.g., Andrii, Mykola BIP2-22)"
                            :required="true" />
            </div>

            <DragAndDropArea v-model:files="files" />
            
            <Button label="Submit File"
                    class="internal-button"
                    @click="handleSubmit" />

        </Section>
    </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted, computed } from 'vue';
import { useRouter } from 'vue-router';
import Section from './UI/Section.vue';
import DragAndDropArea from './UI/DragAndDropArea.vue';
import Button from './UI/Button.vue';
import TextInput from './UI/TextInput.vue';
import { InputDataTypes } from '../types/input-data-types';
import { uploadFile } from "../services/upload-service";
import { getClassrooms, type Classroom } from '../services/classrooms-service';
import { getLessons, type Lesson } from '../services/lessons-service';

export default defineComponent({
    name: "UploadView",
    data() {
        return({
            InputDataTypes,
            email: ref(''),
            description: ref(''),
            files: [] as File[],
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
            if (!this.files[0]) {
                alert("Please select a file before submitting.");
                return;
            }

              if (!this.selectedLesson || !this.selectedClassroom) {
                alert("Invalid classroom or lesson.");
                return;
            }

            try {
                const res = await uploadFile({
                    email: this.email,
                    description: this.description,
                    lesson_id: this.selectedLesson?.id.toString(),
                    classroom_id: this.selectedClassroom?.id.toString(),
                    file: this.files[0],
                });
                console.log("Upload success:", res.data);
                alert("File submitted successfully ✅\n\nThe teacher will have access to them during lesson.");
                this.router.push({
                    name: 'select',
                });
            } catch (err) {
                console.error("Upload error:", err);
                alert("Failed to submit file ❌\n\nPlease try again later.");
            }
        },
    },
    setup(props) {
        const router = useRouter();
        
        const classrooms = ref<Classroom[]>([]);
        const lessons = ref<Lesson[]>([]);

        const selectedClassroom = computed(() =>
            classrooms.value.find(c => String(c.name) === props.classroom)
        );

        const selectedLesson = computed(() =>
            lessons.value.find(l => String(l.pair_number) === props.lesson)
        );

        onMounted(async () => {
            classrooms.value = await fetchWithCache<Classroom>('classrooms', getClassrooms);
            lessons.value = await fetchWithCache<Lesson>('lessons', getLessons);
        });

        async function fetchWithCache<T>(key: string, apiFn: () => Promise<T[]>): Promise<T[]> {
            const localData = localStorage.getItem(key);
            if (localData) {
                return JSON.parse(localData) as T[];
            } else {
                try {
                    const data = await apiFn();
                    localStorage.setItem(key, JSON.stringify(data));
                    return data;
                } catch (error) {
                    console.error(`Failed to fetch ${key}`, error);
                    return [];
                }
            }
        };

        return {
            router,
            classrooms,
            lessons,
            selectedClassroom,
            selectedLesson
        };
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

.input-wrapper {
    width: 100%;
    display: flex;
    flex-direction: column;
    gap: 4px;
    justify-content: space-between;
    align-items: flex-start;

    .input-label {
        @include fonts.noto-font(500);
        @include fonts.responsive-font(16, 16, 1440);
        text-align: start;
    }
}
</style>