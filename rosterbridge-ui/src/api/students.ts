const API_BASE = 'http://localhost:8080';

export interface Student {
    id: string;
    source_id: string;
    first_name: string;
    last_name: string;
    grade_level: number;
    school_id: string;
    created_at: string;
    updated_at: string;
}

export interface SyncResult {
    synced: number;
    created: number;
    updated: number;
    errors: number;
}

export async function fetchStudents(): Promise<Student []> {
    const response = await fetch(`${API_BASE}/students`);
    if (!response.ok) {
        throw new Error(`Error fetching students: ${response.statusText}`);
    }
    return response.json();
}

export async function syncStudents(): Promise<SyncResult> {
    const response = await fetch(`${API_BASE}/sync`, {
        method: 'POST',
    });
    if (!response.ok) {
        throw new Error(`Sync failed`);
    }
    return response.json();
}