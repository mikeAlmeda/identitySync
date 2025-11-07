import { useEffect, useState } from 'react';
import { fetchStudents, syncStudents } from '../api/students';
import type { Student, SyncResult } from '../api/students';

export default function StudentList() {
    const [students, setStudents] = useState<Student[]>([]);
    const [loading, setLoading] = useState(false);
    const [syncResult, setSyncResult] = useState<SyncResult | null>(null);
    const [error, setError] = useState<string | null>(null);

    const loadStudents = async () => {
        setLoading(true);
        setError(null);
        try {
            const data = await fetchStudents();
            setStudents(data);
        } catch (err) {
            setError('Failed to load students. Is the API running?');
            console.error(err);
        } finally {
            setLoading(false);
        }
    };

    const handleSync = async () => {
        setLoading(true);
        setError(null);
        setSyncResult(null);
        try {
            const result = await syncStudents();
            setSyncResult(result);
            // Reload students after sync
            await loadStudents();
        } catch (err) {
            setError('Sync failed. Check the API.');
            console.error(err);
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        loadStudents();
    }, []);

    return (
    <div style={{ padding: '2rem', maxWidth: '1200px', margin: '0 auto' }}>
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '2rem' }}>
        <h1>RosterBridge - Student Directory</h1>
        <button
          onClick={handleSync}
          disabled={loading}
          style={{
            padding: '0.75rem 1.5rem',
            fontSize: '1rem',
            backgroundColor: '#2563eb',
            color: 'white',
            border: 'none',
            borderRadius: '0.5rem',
            cursor: loading ? 'not-allowed' : 'pointer',
            opacity: loading ? 0.6 : 1,
          }}
        >
          {loading ? 'Syncing...' : 'Sync Students'}
        </button>
      </div>

      {error && (
        <div style={{ padding: '1rem', backgroundColor: '#fee', color: '#c00', borderRadius: '0.5rem', marginBottom: '1rem' }}>
          {error}
        </div>
      )}

      {syncResult && (
        <div style={{ padding: '1rem', backgroundColor: '#efe', color: '#060', borderRadius: '0.5rem', marginBottom: '1rem' }}>
          <strong>Sync Complete:</strong> {syncResult.synced} students processed 
          ({syncResult.created} created, {syncResult.updated} updated, {syncResult.errors} errors)
        </div>
      )}

      {loading && !syncResult ? (
        <p>Loading...</p>
      ) : students.length === 0 ? (
        <p>No students found. Click "Sync Students" to populate.</p>
      ) : (
        <table style={{ width: '100%', borderCollapse: 'collapse', backgroundColor: 'white', boxShadow: '0 1px 3px rgba(0,0,0,0.1)' }}>
          <thead>
            <tr style={{ backgroundColor: '#f3f4f6', borderBottom: '2px solid #e5e7eb' }}>
              <th style={{ padding: '0.75rem', textAlign: 'left' }}>Source ID</th>
              <th style={{ padding: '0.75rem', textAlign: 'left' }}>First Name</th>
              <th style={{ padding: '0.75rem', textAlign: 'left' }}>Last Name</th>
              <th style={{ padding: '0.75rem', textAlign: 'left' }}>Grade</th>
              <th style={{ padding: '0.75rem', textAlign: 'left' }}>School ID</th>
              <th style={{ padding: '0.75rem', textAlign: 'left' }}>Last Updated</th>
            </tr>
          </thead>
          <tbody>
            {students.map((student) => (
              <tr key={student.id} style={{ borderBottom: '1px solid #e5e7eb' }}>
                <td style={{ padding: '0.75rem' }}>{student.source_id}</td>
                <td style={{ padding: '0.75rem' }}>{student.first_name}</td>
                <td style={{ padding: '0.75rem' }}>{student.last_name}</td>
                <td style={{ padding: '0.75rem' }}>{student.grade_level}</td>
                <td style={{ padding: '0.75rem' }}>{student.school_id}</td>
                <td style={{ padding: '0.75rem' }}>{new Date(student.updated_at).toLocaleString()}</td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
}
