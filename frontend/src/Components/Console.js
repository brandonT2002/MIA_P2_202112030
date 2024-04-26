import React, { useState, useEffect, useRef } from 'react';
import "./Console.css";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faPaperPlane } from '@fortawesome/free-solid-svg-icons';

const Console = () => {
    const [messages, setMessages] = useState([]);
    const [inputValue, setInputValue] = useState('');
    const messagesEndRef = useRef(null);

    useEffect(() => {
        fetchMessages();
    }, []);

    useEffect(() => {
        const intervalId = setInterval(() => {
            fetchMessages();
        }, 5000);

        return () => clearInterval(intervalId);
    }, []);

    useEffect(() => {
        scrollToBottom();
    }, [messages]);

    const scrollToBottom = () => {
        messagesEndRef.current?.scrollIntoView({ behavior: "auto" }); // Cambiado a "auto" para evitar la animación suave
    };

    const fetchMessages = async () => {
        try {
            const response = await fetch('http://127.0.0.1:8080/messages');
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            const data = await response.json();
            setMessages(data.messages);
        } catch (error) {
            console.error('Error fetching messages:', error);
        }
    };

    const sendMessageToBot = async (message) => {
        try {
            const response = await fetch('http://127.0.0.1:8080/interpreter', {
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
            fetchMessages();
        } catch (error) {
            console.error('Error sending message to bot:', error);
        }
    };

    const handleKeyPress = async (event) => {
        if (event.key === 'Enter') {
            const message = event.target.value.trim();
            if (message !== '') {
                setInputValue('');
                await sendMessageToBot(message);
            }
        }
    };

    return (
        <div className="console">
            <div className="messages">
                {messages.map((conversation, index) => (
                    <><div key={index} className="message user-message">{conversation[0]}</div><div key={index} className="message bot-message">{conversation[1]}</div></>
                ))}
                <div ref={messagesEndRef} />
            </div>
            <div className="input-container">
                <input
                    type="text"
                    value={inputValue}
                    onChange={(e) => setInputValue(e.target.value)}
                    onKeyDown={handleKeyPress}
                    placeholder="Escribe un mensaje..."
                />
                <FontAwesomeIcon icon={faPaperPlane} className="icon" onClick={async () => {
                    const message = inputValue.trim();
                    if (message !== '') {
                        setInputValue('');
                        await sendMessageToBot(message);
                    }
                }} />
            </div>
        </div>
    );
}

export default Console;