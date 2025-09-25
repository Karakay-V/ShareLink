// src/services/upload-service.ts
import api from "./api";

export interface UploadPayload {
    email: string;
    signature: string;
    lesson: string;
    classroom: string;
    file?: File;
}

export async function uploadFile(data: UploadPayload) {
    const formData = new FormData();
    formData.append("email", data.email);
    formData.append("signature", data.signature);
    formData.append("lesson", data.lesson);
    formData.append("classroom", data.classroom);

    if (data.file) {
        formData.append("file", data.file);
    }

    return api.post("/presentations", formData, {
        headers: { "Content-Type": "multipart/form-data" },
    });
}
