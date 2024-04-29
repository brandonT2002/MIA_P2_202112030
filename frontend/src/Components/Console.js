import React, { useState, useEffect, useRef } from 'react';
import "./Console.css";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faPaperPlane } from '@fortawesome/free-solid-svg-icons';

const Console = () => {
    const [messages, setMessages] = useState([]);
    const [inputValue, setInputValue] = useState('');
    const [textareaRows, setTextareaRows] = useState(1);
    const textareaRef = useRef(null);
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

    useEffect(() => {
        adjustTextareaHeight();
    }, [inputValue]);

    const scrollToBottom = () => {
        messagesEndRef.current?.scrollIntoView({ behavior: "auto" });
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
        if (event.key === 'Enter' && event.shiftKey) {
            event.preventDefault();
            setInputValue(prev => prev + '\n');
        } else if (event.key === 'Enter' && !event.shiftKey) {
            event.preventDefault();
            const message = inputValue.trim();
            if (message !== '') {
                setInputValue('');
                await sendMessageToBot(message);
            }
        }
    };

    const adjustTextareaHeight = () => {
        const textarea = textareaRef.current;
        if (textarea) {
            const rows = textarea.value.split('\n').length;
            setTextareaRows(rows < 7 ? rows : 7);
        }
    };

    return (
        <div className="console">
            <div className="messages">
                {messages.map((conversation, index) => (
                    <div key={index} className="message-container">
                        <div className="message user-message">
                            {conversation[0].split('\n').map((line, i) => (
                                <React.Fragment key={i}>
                                    {line}
                                    {i < conversation[0].split('\n').length - 1 && <br />}
                                </React.Fragment>
                            ))}
                        </div>
                        <div className="message bot-message">
                            {conversation[1].split('\n').map((line, i) => (
                                <React.Fragment key={i}>
                                    {line}
                                    {i < conversation[1].split('\n').length - 1 && <br />}
                                </React.Fragment>
                            ))}
                        </div>
                    </div>
                ))}
                <div ref={messagesEndRef} />
            </div>
            <div className="input-container">
                <textarea
                    ref={textareaRef}
                    value={inputValue}
                    onChange={(e) => setInputValue(e.target.value)}
                    onKeyDown={handleKeyPress}
                    placeholder="Escribe un mensaje..."
                    rows={textareaRows}
                />
            </div>
        </div>
    );
}

export default Console;
