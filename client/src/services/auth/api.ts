// src/services/auth/api.ts
import axios from "axios";

const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL || "http://127.0.0.1:8002",
});

// Додаємо інтерсептор для автоматичної підстановки токена
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem("token"); // ключ можна змінити на свій
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => Promise.reject(error)
);

export default api;
