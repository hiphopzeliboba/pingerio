const API_URL = process.env.REACT_APP_API_URL || 'http://localhost:8081';

export const getContainers = async () => {
    const response = await fetch(`${API_URL}/containers`);
    if (!response.ok) {
        throw new Error('Failed to fetch containers');
    }
    return response.json();
};