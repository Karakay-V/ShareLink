import api from "./api";

export interface Lesson {
    id: number;
    pair_number: number;
    start_time: string; // HH:mm
    end_time: string;   // HH:mm
}

export async function getLessons() {
    const response = await api.get<Lesson[]>("/lessons");
    return response.data;
}
