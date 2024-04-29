import React, { useState, useEffect, useRef } from 'react';
import "./Console.css";
import Tree from './Graph'; // Asegúrate de importar correctamente el componente Tree
import { sendMessageToBot } from './Utils';

const Console = () => {
    const [messages, setMessages] = useState([]);
    const [inputValue, setInputValue] = useState('');
    const [textareaRows, setTextareaRows] = useState(1);
    const textareaRef = useRef(null);
    const consoleRef = useRef(null);
    const messagesEndRef = useRef(null);

    useEffect(() => {
        fetchMessages();
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
            const response = await fetch('http://3.147.86.56:8080/messages');
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            const data = await response.json();
            setMessages(data.messages);
        } catch (error) {
            console.error('Error fetching messages:', error);
        }
    };

    const sendMessageToBot_ = async (message) => {
        try {
            await sendMessageToBot(message);
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
                await sendMessageToBot_(message);
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
        <div className="console" ref={consoleRef}>
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
                        {conversation[1].startsWith('digraph') ? (
                            <div className="message bot-message bot-message-graph">
                                <div className="chart-row-1">
                                    <div className="chart-container-tree">
                                        <Tree dotCode={conversation[1]} />
                                    </div>
                                </div>
                            </div>
                        ) : (
                            <div className={`message bot-message ${conversation[1].startsWith('Error') ? 'bot-message-error' : ''}`}>
                                {conversation[1].split('\n').map((line, i) => (
                                    <React.Fragment key={i}>
                                        {line}
                                        {i < conversation[1].split('\n').length - 1 && <br />}
                                    </React.Fragment>
                                ))}
                            </div>
                        )}
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
