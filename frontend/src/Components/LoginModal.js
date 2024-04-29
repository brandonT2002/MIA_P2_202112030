import React, { useState, useEffect } from 'react';
import './Login.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faUser, faLock } from '@fortawesome/free-solid-svg-icons';
import Alert from '@mui/material/Alert';
import Stack from '@mui/material/Stack';
import { sendMessageToBot } from './Utils';

const LoginModal = ({ onClose, selectedPartitionId, onLoginSuccess }) => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');
    const [serverResponse, setServerResponse] = useState(null);

    useEffect(() => {
        let timer;
        if (error || serverResponse) {
            timer = setTimeout(() => {
                setError('');
                setServerResponse(null);
                onClose(); // Cerrar el modal después de 5 segundos
            }, 5000); // 5000 milisegundos = 5 segundos
        }
        return () => clearTimeout(timer);
    }, [error, serverResponse, onClose]);

    const handleFormSubmit = async (e) => {
        e.preventDefault();
        
        // Verificar si se han completado ambos campos
        if (!username || !password) {
            // Establecer el mensaje de error
            setError('Todos los campos son obligatorios');
            return;
        }

        try {
            // Construir la cadena de inicio de sesión con los valores de usuario, contraseña y ID de partición
            const loginString = `login -user=${username} -pass=${password} -id=${selectedPartitionId}`;
            
            // Enviar la cadena al bot y obtener la respuesta
            const response = await sendMessageToBot(loginString);
            
            // Establecer la respuesta del servidor para mostrarla en un alert
            setServerResponse(response);
            
            console.log("Se ha iniciado sesión correctamente");
            onLoginSuccess(); // Llamar a la función de éxito de inicio de sesión
        } catch (error) {
            console.log("Ocurrió un error:", error);
        }
    };

    const handleModalClick = (e) => {
        e.stopPropagation();
    };

    return (
        <>
            <div className="login-modal" onClick={onClose}>
                <div className="login-modal-content" onClick={handleModalClick}>
                    <h2>Login</h2>
                    <Stack spacing={2}> {/* Utilizar Stack para alinear elementos */}
                        {error && (
                            <Alert severity="warning">
                                {error}
                            </Alert>
                        )}
                        {serverResponse && (
                            serverResponse.message.startsWith("Error") ? (
                                <Alert severity="error">
                                    {serverResponse.message}
                                </Alert>
                            ) : (
                                <Alert severity="success">
                                    {serverResponse.message}
                                </Alert>
                            )
                        )}
                        <form onSubmit={handleFormSubmit}>
                            <label htmlFor="username">
                                <FontAwesomeIcon icon={faUser} />
                                Usuario:
                            </label>
                            <input 
                                type="text" 
                                id="username" 
                                name="username" 
                                value={username} 
                                onChange={(e) => setUsername(e.target.value)} 
                            />
                            <label htmlFor="password">
                                <FontAwesomeIcon icon={faLock} />
                                Contraseña:
                            </label>
                            <input 
                                type="password" 
                                id="password" 
                                name="password" 
                                value={password} 
                                onChange={(e) => setPassword(e.target.value)} 
                            />
                            <button type="submit">Iniciar Sesión</button>
                        </form>
                    </Stack>
                </div>
            </div>
        </>
    );
};

export default LoginModal;
