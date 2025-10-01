<template>
    <div class="files-inner">
        <Section>
            <div class="select-wrapper">
                <label for="select_classrooms" class="select-label">
                    Classroom (Auditorium)
                </label>
                <Select v-model="selectedClassroom"
                        :items="classrooms.map(c => c.name)"
                        labelKey="select_classrooms"
                        valueKey="id"
                        placeholder="Classroom ..." />
            </div>


            <div class="select-wrapper">
                <label for="select_lessons" class="select-label">
                    Class / Lesson Time Slot
                </label>
                <Select v-model="selectedLesson"
                        :items="lessonLabels"
                        labelKey="select_lessons"
                        valueKey="id"
                        placeholder="Lesson ..." />
            </div>
            
            <Button label="Show Presentations"
                    class="internal-button"
                    @click="handleNext" />

        </Section>

        <Section maxWidth="100%" >
            
            <div v-if="presentations.length > 0" class="list-wrapper">
                <File   v-for="p in presentations"
                        :key="p.id"
                        :name="p.file_name"
                        :description="p.description"
                        @click="handleDownload(p.file_id, p.file_name)" />
            </div>


            <div v-else class="heading-wrapper">
                <h2 class="heading-text">
                    <span class="heading-bold">Empty.</span>
                    <br/>
                    Not found any uploaded presentations for this Lesson.
                </h2>
            </div>
        
        </Section>
    </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import { checkAuth } from '../../services/auth/check';
import Section from '../UI/Section.vue';
import Button from '../UI/Button.vue';
import Select from '../UI/Select.vue';
import File from '../UI/File.vue';
import { getClassrooms, type Classroom } from '../../services/classrooms-service';
import { getLessons, type Lesson } from '../../services/lessons-service';
import { fetchWithCache } from '../../services/selects-service';
import { getPresentationsByLesson, downloadPresentation, type Presentation } from '../../services/auth/presentations-service';

export default defineComponent({
    name: "FilesView",
    components: {
        Section,
        Select,
        Button,
        File,
    },
    props: {
        classroomProp: {
            type: String,
            required: false,
            default: '',
        }
    },
    setup(props) {
        const router = useRouter();

        const selectedClassroom = ref<string | null>(null);
        const selectedLesson = ref<string | null>(null);

        const classrooms = ref<Classroom[]>([]);
        const lessons = ref<Lesson[]>([]);

        const presentations = ref<Presentation[]>([]);

        onMounted(async () => {
            const isValid = await checkAuth();
            if (!isValid) {
                router.push({ name: "auth" });
            }

            classrooms.value = await fetchWithCache<Classroom>('classrooms', getClassrooms);
            lessons.value = await fetchWithCache<Lesson>('lessons', getLessons);

            if (props.classroomProp !== '') {
                const found = classrooms.value.find(c => c.name === props.classroomProp);
                if (found) {
                    selectedClassroom.value = found.name;
                }
            }
        });

        const lessonLabels = computed(() =>
            lessons.value.map(l => `${l.pair_number} (${l.start_time} - ${l.end_time})`)
        );

        const handleNext = async () => {
            if (!selectedClassroom.value || !selectedLesson.value) {
                alert("Please select all fields!");
                return;
            }

            const pairNumber = selectedLesson.value.trim().charAt(0);

            presentations.value = await getPresentationsByLesson(pairNumber, selectedClassroom.value);
        };

        const handleDownload = async (file_id: number, file_name: string, file_extension?: string) => {
            try {
                await downloadPresentation(file_id, file_name, file_extension); // сервіс вже завантажує файл
                console.log(`File ${file_id} downloaded successfully`);
            } catch (err) {
                console.error(`Failed to download file ${file_id}:`, err);
                alert("Failed to download file ❌");
            }
        };

        return {
            selectedClassroom,
            selectedLesson,
            classrooms,
            lessons,
            lessonLabels,
            handleNext,
            handleDownload,
            presentations,
        };
    },
});
</script>

<style lang="scss" scoped>
@use "/src/assets/styles/colors" as *;
@use "/src/assets/styles/image";
@use "/src/assets/styles/fonts";
@use "/src/assets/styles/shadows";

.files-inner {
    width: calc(100% - 30px - 30px);
    padding: 30px;
    gap: 30px;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: flex-start;

    @media (max-width: 900px) {
        flex-wrap: wrap;
        justify-content: center;
    }
}

.select-wrapper {
    width: 100%;
    display: flex;
    flex-direction: column;
    gap: 4px;
    justify-content: space-between;
    align-items: flex-start;

    .select-label {
        @include fonts.noto-font(500);
        @include fonts.responsive-font(16, 16, 1440);
        text-align: start;
    }
}

.internal-button {
    margin-top: 10px;
}

.list-wrapper {
    width: 100%;
    display: flex;
    gap: 20px 10px;
    flex-direction: row;
    flex-wrap: wrap;
    justify-content: flex-start;
}

.heading-wrapper {
    padding: 70px 20px;

    .heading-text {
        text-align: center;
        @include fonts.noto-font(500);
        @include fonts.responsive-font(20, 18, 1440);
        
        .heading-bold {
            @include fonts.noto-font(800);
            @include fonts.responsive-font(24, 20, 1440);
        }
    }
}
</style>