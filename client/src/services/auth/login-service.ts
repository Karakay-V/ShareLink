// src/services/auth/login-service.ts
/***
 * There I use global API service in project
 * In next auth internal services I will use internal API sercvice
 * <bold>services/auth/api.ts</bold>
 */
import api from "../api";

export interface LoginRequest {
    email: string;
    password: string;
}

export interface LoginResponse {
    token: string;
}

export async function loginTeacher(data: LoginRequest): Promise<LoginResponse> {
    const res = await api.post<LoginResponse>("/login", data);
    return res.data;
}
