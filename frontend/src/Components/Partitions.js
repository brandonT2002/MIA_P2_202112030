// Partitions.js

import React, { useState } from 'react';
import './Explorer.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faArrowLeft } from '@fortawesome/free-solid-svg-icons';
import icono from '../img/Particion.png';

const Partitions = ({ onShowExplorer, partitions }) => {
    const [selectedPartition, setSelectedPartition] = useState(null);
    const [modalPosition, setModalPosition] = useState({ x: 0, y: 0 });

    const handlePartitionRightClick = (partition, event) => {
        event.preventDefault();
        setSelectedPartition(partition);
        setModalPosition({ x: event.clientX, y: event.clientY });
    };

    const handleCloseModal = () => {
        setSelectedPartition(null);
    };

    const handleUnmount = () => {
        // Lógica para desmontar la partición
        handleCloseModal();
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
                    onClick={handleCloseModal}
                >
                    <div
                        className="modal"
                        style={{
                            position: 'fixed',
                            top: modalPosition.y,
                            left: modalPosition.x + 50, // Ajuste para posicionar al lado derecho del icono
                        }}
                        onClick={(e) => e.stopPropagation()}
                    >
                        <button className="modal-button-full" onClick={handleUnmount}>Desmontar</button>
                    </div>
                </div>
            )}
        </div>
    );
};

export default Partitions;
