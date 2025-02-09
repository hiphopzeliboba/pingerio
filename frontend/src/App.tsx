import React, { useEffect, useState } from 'react';
import { getContainers } from './api';
import './App.css';

interface Container {
    id: string;
    name: string;
    ip: string;
    status: string;
    ping_time: string;
}

function App() {
    const [containers, setContainers] = useState<Container[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    const fetchContainers = async () => {
        try {
            setLoading(true);
            const data = await getContainers();
            setContainers(data);
            setError(null);
        } catch (err) {
            setError('Failed to fetch containers');
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        fetchContainers();
        const interval = setInterval(fetchContainers, 30000);
        return () => clearInterval(interval);
    }, []);

    if (loading) {
        return <div className="loading">Loading...</div>;
    }

    if (error) {
        return (
            <div className="error">
                <p>{error}</p>
                <button onClick={fetchContainers}>Retry</button>
            </div>
        );
    }

    return (
        <div className="app">
            <header className="header">
                <h1>Docker Container Monitor</h1>
            </header>

            <main className="container-list">
                {containers.length === 0 ? (
                    <p className="no-containers">No containers found</p>
                ) : (
                    containers.map(container => (
                        <div key={container.id} className="container-card">
                            <div className="container-header">
                                <span className={`status ${container.status.toLowerCase()}`} />
                                <h3>{container.name}</h3>
                            </div>
                            <div className="container-info">
                                <p><strong>ID:</strong> {container.id}</p>
                                <p><strong>IP:</strong> {container.ip || 'N/A'}</p>
                                <p><strong>Status:</strong> {container.status}</p>
                                <p><strong>Last Ping:</strong> {new Date(container.ping_time).toLocaleString()}</p>
                            </div>
                        </div>
                    ))
                )}
            </main>
        </div>
    );
}

export default App;