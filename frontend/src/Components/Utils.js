// Utilidades.js
export const sendMessageToBot = async (message) => {
    try {
        const response = await fetch('http://3.147.86.56:8080/interpreter', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ code: message })
        });

        if (!response.ok) {
            throw new Error('Network response was not ok');
        }

        const data = await response.json();
        return data;
    } catch (error) {
        console.error('Error sending message to bot:', error);
        throw error;
    }
};
