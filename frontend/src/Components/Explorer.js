import React, { useState, useEffect } from 'react';
import "./Explorer.css";
import icono from '../img/disco-duro1.png'

const Explorer = () => {
    const [files, setFiles] = useState([]);

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

    return (
        <div className="explorer">
            <h2>Explorador de archivos</h2>
            <div className="file-grid">
                {files.map((fileName, index) => (
                    <div className="file" key={index}>
                        <img src={icono} alt="Icono de disco" />
                        <p>{fileName}</p>
                    </div>
                ))}
            </div>
        </div>
    );
}

export default Explorer;