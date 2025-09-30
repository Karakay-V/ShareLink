// src/services/upload-service.ts
import api from "./api";

export interface PresentationUploadPayload {
    title: string;
    description: string;
    lessonId: string;
    classroomId: string;
    file: File; // файл тепер обов’язковий
}

export interface Presentation {
    id: string;
    title: string;
    description: string;
    fileUrl: string;
    lessonId: string;
    classroomId: string;
    createdAt: string;
    updatedAt: string;
}

export async function uploadFile(data: PresentationUploadPayload) {
    const formData = new FormData();
    formData.append("title", data.title);
    formData.append("description", data.description);
    formData.append("lessonId", data.lessonId);
    formData.append("classroomId", data.classroomId);
    formData.append("file", data.file);

    return api.post<Presentation>("/presentations", formData, {
        headers: { "Content-Type": "multipart/form-data" },
    });
}
