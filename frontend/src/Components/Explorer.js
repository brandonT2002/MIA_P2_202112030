import React, { useState, useEffect } from 'react';
import "./Explorer.css";
import icono from '../img/disco-duro1.png';
import Partitions from './Partitions'; // Importa el componente Partitions desde su archivo

const Explorer = () => {
    const [files, setFiles] = useState([]);
    const [showPartitions, setShowPartitions] = useState(false);
    const [selectedDisk, setSelectedDisk] = useState(null);
    const [explorerVisible, setExplorerVisible] = useState(true); // Nuevo estado para controlar la visibilidad de Explorer

    useEffect(() => {
        fetchFiles();
    }, []);

    const fetchFiles = async () => {
        try {
            const response = await fetch('http://localhost:8080/files');
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            const data = await response.json();
            setFiles(data.files || []); // Si data.files es null, se establece un arreglo vacío
        } catch (error) {
            console.error('There was a problem with the fetch operation:', error);
        }
    };

    const handleDiskClick = (disk) => {
        setSelectedDisk(disk);
        setShowPartitions(true);
        setExplorerVisible(false); // Oculta Explorer cuando se muestra Partitions
    };

    const handleShowExplorer = () => {
        setExplorerVisible(true); // Muestra Explorer cuando se hace clic en "Mostrar Explorer"
        setShowPartitions(false); // Oculta Partitions
    };

    return (
        <div>
            {explorerVisible && (
                <div className="explorer">
                    <h2>Discos</h2>
                    <div className="file-grid">
                        {files.map((fileName, index) => (
                            <div className="file" key={index}>
                                <img src={icono} alt="Icono de disco" onClick={() => handleDiskClick(fileName)}/>
                                <p>{fileName}</p>
                            </div>
                        ))}
                    </div>
                </div>
            )}
            {showPartitions && <Partitions onShowExplorer={handleShowExplorer} />}
        </div>
    );
};

export default Explorer;