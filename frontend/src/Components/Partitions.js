import React, { useState, useEffect } from 'react';
import './Explorer.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faArrowLeft } from '@fortawesome/free-solid-svg-icons';
import icono from '../img/Particion.png';
import LoginModal from './LoginModal';
import Alert from '@mui/material/Alert'; // Importa el componente Alert
import { sendMessageToBot } from './Utils';

const Partitions = ({ onShowExplorer, partitions }) => {
    const [selectedPartition, setSelectedPartition] = useState(null);
    const [showLoginModal, setShowLoginModal] = useState(false);
    const [modalPosition, setModalPosition] = useState({ x: 0, y: 0 });
    const [isLoggedIn, setIsLoggedIn] = useState(false); // Estado de bandera para el inicio de sesión
    const [serverResponse, setServerResponse] = useState(null); // Estado para almacenar la respuesta del servidor

    // Efecto para limpiar la respuesta del servidor después de 5 segundos
    useEffect(() => {
        let timer;
        if (serverResponse) {
            timer = setTimeout(() => {
                setServerResponse(null);
            }, 5000); // 5000 milisegundos = 5 segundos
        }
        return () => clearTimeout(timer);
    }, [serverResponse]);

    const handlePartitionRightClick = (partition, event) => {
        event.preventDefault();
        setSelectedPartition(partition);
        setModalPosition({ x: event.clientX, y: event.clientY });
    };

    const handleUnmount = () => {
        setSelectedPartition(null);
    };

    const handleLoginClick = () => {
        setShowLoginModal(true);
    };

    const handleCloseLoginModal = () => {
        setShowLoginModal(false);
        setSelectedPartition(null);
    };

    const handleLoginSuccess = () => {
        setIsLoggedIn(true); // Actualiza el estado después de iniciar sesión correctamente
    };

    const handleLogout = async () => {
        try {
            await sendMessageToBot("logout")
            const response = { message: "Logout exitoso" };
            console.log("Se ha cerrado sesión correctamente");
            setSelectedPartition(null);
            setIsLoggedIn(false); // Actualiza el estado después de cerrar sesión
            setServerResponse(response); // Guarda la respuesta del servidor
        } catch (error) {
            console.log("Ocurrió un error al cerrar sesión:", error);
        }
    };

    return (
        <div className="explorer">
            <div className="partitions-header">
                <button className="back-button" onClick={onShowExplorer}>
                    <FontAwesomeIcon icon={faArrowLeft} />
                </button>
                <h2>Particiones</h2>
            </div>
            <div className="file-grid">
                {partitions.map((partition, index) => (
                    <div
                        className="file"
                        key={index}
                        onContextMenu={(event) => handlePartitionRightClick(partition, event)}
                        style={{ position: 'relative' }}
                    >
                        <img src={icono} alt="Icono de particion" title={partition[0]} />
                        <p>{partition[1]}</p>
                    </div>
                ))}
            </div>
            {selectedPartition && (
                <div
                    className="modal-background"
                    onClick={() => setSelectedPartition(null)}
                >
                    <div
                        className="modal"
                        style={{
                            position: 'fixed',
                            top: modalPosition.y,
                            left: modalPosition.x + 50,
                        }}
                        onClick={(e) => e.stopPropagation()}
                    >
                        <button className="modal-button-full" onClick={handleUnmount}>Desmontar</button>
                        {isLoggedIn ? ( // Verifica el estado para mostrar "Login" o "Logout"
                            <button className="modal-button-full" onClick={handleLogout}>Logout</button>
                        ) : (
                            <button className="modal-button-full" onClick={handleLoginClick}>Login</button>
                        )}
                    </div>
                </div>
            )}
            {showLoginModal && (
                <div className="modal-background" onClick={handleCloseLoginModal}>
                    <div className="modal">
                        <LoginModal onClose={handleCloseLoginModal} selectedPartitionId={selectedPartition ? selectedPartition[0] : null} onLoginSuccess={handleLoginSuccess} />
                    </div>
                </div>
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
        </div>
    );
};

export default Partitions;
