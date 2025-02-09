import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './App.css';

function App() {
    const [containers, setContainers] = useState([]);

    useEffect(() => {
        const fetchContainers = async () => {
            try {
                const response = await axios.get('http://localhost:8081/containers');
                setContainers(response.data);
            } catch (error) {
                console.error('Error fetching containers:', error);
            }
        };

        fetchContainers();
        const interval = setInterval(fetchContainers, 30000);
        return () => clearInterval(interval);
    }, []);

    return (
        <div className="App">
            <h1>Docker Containers Monitor</h1>
            <table>
                <thead>
                <tr>
                    <th>ID</th>
                    <th>Name</th>
                    <th>IP</th>
                    <th>Status</th>
                    <th>Last Ping</th>
                </tr>
                </thead>
                <tbody>
                {containers.map(container => (
                    <tr key={container.id}>
                        <td>{container.id}</td>
                        <td>{container.name}</td>
                        <td>{container.ip}</td>
                        <td>{container.status}</td>
                        <td>{new Date(container.ping_time).toLocaleString()}</td>
                    </tr>
                ))}
                </tbody>
            </table>
        </div>
    );
}

export default App;