<template>
    <div class="view-wrapper">
        <div class="info-container">
            <h2 class="title">
                Easily share your study materials with classmates
            </h2>

            <div class="quired_info">
               <img class="image" :src="CoverImage" alt="Easily share your study materials with classmates" />
            </div>
        </div>

        <Section>
            
            <!-- <div class="select-wrapper">
                <label for="select_buildings" class="select-label">
                    Building (Corpus)
                </label>
                <Select v-model="selectedBuilding"
                        :items="buildings.map(b => b.name)"
                        labelKey="select_buildings"
                        valueKey="id"
                        placeholder="Building ..." />
            </div> -->

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
            
            <Button label="Next"
                    class="internal-button"
                    @click="handleNext" />

        </Section>
    </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted, computed } from 'vue';
import { useRouter } from 'vue-router';
import Section from './UI/Section.vue';
import Button from './UI/Button.vue';
import Select from './UI/Select.vue';
import CoverImage from '../assets/covers/cover-image.png';
import { getBuildings, type Building } from '../services/buildings-service';
import { getClassrooms, type Classroom } from '../services/classrooms-service';
import { getLessons, type Lesson } from '../services/lessons-service';

export default defineComponent({
    name: "SelectView",
    components: {
        Section,
        Select,
        Button,
    },
    setup() {
        const router = useRouter();

        const CoverImageRef = ref(CoverImage);
        const selectedBuilding = ref<string | null>(null);
        const selectedClassroom = ref<string | null>(null);
        const selectedLesson = ref<string | null>(null);

        const buildings = ref<Building[]>([]);
        const classrooms = ref<Classroom[]>([]);
        const lessons = ref<Lesson[]>([]);

        onMounted(async () => {
            buildings.value = await fetchWithCache<Building>('buildings', getBuildings);
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
        }
        
        const lessonLabels = computed(() =>
            lessons.value.map(l => `${l.pair_number} (${l.start_time} - ${l.end_time})`)
        );

        const handleNext = () => {
            console.log({ building: selectedBuilding.value, classroom: selectedClassroom.value, lesson: selectedLesson.value, });

            if (!selectedClassroom.value || !selectedLesson.value) {
                alert("Please select all fields!");
                return;
            }

            // редірект з параметрами
            router.push({
                name: 'lesson', // <- ім’я маршруту, не шлях
                params: {
                    lessonId: selectedLesson.value.trim().charAt(0),
                },
                query: {
                    classroom: selectedClassroom.value,
                }
            });
        };

        return {
            CoverImage: CoverImageRef,
            selectedBuilding,
            selectedClassroom,
            selectedLesson,
            buildings,
            classrooms,
            lessons,
            lessonLabels,
            handleNext,
        };
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
        justify-content: center;
        align-items: center;
        width: 100%;

        .image {
            width: 100%;
        }
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
        @include fonts.responsive-font(16, 14, 1440);
        text-align: start;
    }
}

.internal-button {
    margin-top: 10px;
}
</style>