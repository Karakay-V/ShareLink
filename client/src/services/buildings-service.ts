// src/services/buildings-service.ts
import api from "./api";

export interface Building {
    id: number;
    name: string;
}

export async function getBuildings() {
    const response = await api.get<Building[]>("/buildings");
    return response.data;
}
