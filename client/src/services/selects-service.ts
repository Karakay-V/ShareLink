// src/services/selects-service.ts

export async function fetchWithCache<T>(key: string, apiFn: () => Promise<T[]>): Promise<T[]> {
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