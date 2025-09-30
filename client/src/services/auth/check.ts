// src/services/auth/check.ts
import api from "./api";

export async function checkAuth(): Promise<boolean> {
    try {
        await api.get("/auth/check");
        return true; // ✅ токен живий
    } catch {
        return false;
    }
}