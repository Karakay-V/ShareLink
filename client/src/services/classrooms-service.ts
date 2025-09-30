import api from "./api";

export interface Classroom {
    id: number;
    building_id: number;
    name: string;
}

export async function getClassrooms() {
    const response = await api.get<Classroom[]>("/classrooms");
    return response.data;
}
