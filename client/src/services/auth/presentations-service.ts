// src/services/auth/presentations-service.ts
import api from "./api";

export interface Presentation {
    id: number;
    pair_number: string;
    classroom_name: string;
    email: string;
    description: string;
    file_id: number;
    file_name: string;
    file_extension: string,
    uploaded_at: string;
    is_outdated: boolean;
}
// ==================== GET ALL ====================
export async function getAllPresentations(showHidden = false) {
    const res = await api.get<Presentation[]>(
        `/presentations/all?show_hidden=${showHidden}`
    );
    return res.data;
}

// ==================== GET BY LESSON ====================
export async function getPresentationsByLesson(
    pairNumber: string,
    classroomName: string
) {
    const res = await api.get<Presentation[]>(
        `/presentations/get/lesson/${pairNumber}?classroom=${encodeURIComponent(
            classroomName
        )}`
    );
    return res.data;
}

// ==================== DOWNLOAD ====================
export async function downloadPresentation(fileId: number, fileName?: string, fileExt?: string) {
    const res = await api.get(`/presentations/download/${fileId}`, {
        responseType: "blob",
    });

    // Формуємо повну назву файлу
    let fullName = `file_${fileId}`;
    if (fileName) fullName = fileName;
    if (fileExt) fullName += `.${fileExt}`;

    // Автоматично зберігаємо файл у браузері
    const url = window.URL.createObjectURL(res.data);
    const link = document.createElement("a");
    link.href = url;
    link.setAttribute("download", fullName);
    document.body.appendChild(link);
    link.click();
    link.remove();
}

