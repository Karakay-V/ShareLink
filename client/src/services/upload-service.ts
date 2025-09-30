// src/services/upload-service.ts
import api from "./api";

export interface PresentationUploadPayload {
    lesson_id: string;     // саме snake_case як у Go
    classroom_id: string;  // те саме
    email: string;
    description: string;
    file: File;
}

export async function uploadFile(data: PresentationUploadPayload) {
    const formData = new FormData();
    formData.append("lesson_id", data.lesson_id);
    formData.append("classroom_id", data.classroom_id);
    formData.append("email", data.email);
    formData.append("description", data.description);
    formData.append("file", data.file);

    return api.post("/presentations/", formData, {
        headers: { "Content-Type": "multipart/form-data" },
    });
}
