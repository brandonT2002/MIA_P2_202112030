import React, { useState, useEffect } from 'react';
import "./Explorer.css";
import icono from '../img/disco-duro1.png';
import Partitions from './Partitions';

const Explorer = () => {
    const [files, setFiles] = useState([]);
    const [showPartitions, setShowPartitions] = useState(false);
    const [selectedDisk, setSelectedDisk] = useState(null);
    const [explorerVisible, setExplorerVisible] = useState(true);
    const [partitions, setPartitions] = useState([]); // Define la variable partitions en el estado

    useEffect(() => {
        fetchFiles();
    }, []);

    const fetchFiles = async () => {
        try {
            const response = await fetch('http://3.147.86.56:8080/files');
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            const data = await response.json();
            setFiles(data.files || []);
        } catch (error) {
            console.error('There was a problem with the fetch operation:', error);
        }
    };

    const handleDiskClick = async (diskName) => {
        setSelectedDisk(diskName);
        console.log(selectedDisk)
        setShowPartitions(true);
        setExplorerVisible(false);

        try {
            const response = await fetch('http://3.147.86.56:8080/partitions', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ disk: diskName.split('.')[0] })
            });
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            const data = await response.json();
            setPartitions(data.partitions); // Guarda las particiones en el estado
        } catch (error) {
            console.error('There was a problem with the fetch operation:', error);
        }
    };

    const handleShowExplorer = () => {
        setExplorerVisible(true);
        setShowPartitions(false);
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
                                <p data-disk={fileName}>{fileName}</p>
                            </div>
                        ))}
                    </div>
                </div>
            )}
            {showPartitions && <Partitions onShowExplorer={handleShowExplorer} partitions={partitions} />} {/* Pasa partitions como propiedad */}
        </div>
    );
};

export default Explorer;
