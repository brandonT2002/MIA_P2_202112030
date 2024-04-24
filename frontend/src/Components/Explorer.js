import React from 'react';
import "./Explorer.css";
import icono from '../img/disco-duro.png'

const Explorer = () => {
    return (
        <div className="explorer">
            <h1>Explorador de archivos</h1>
            <div className="file-grid">
                <div className="file">
                    <img src={icono} alt="Icono de disco" />
                    <p>Disco 1</p>
                </div>
                <div className="file">
                    <img src={icono} alt="Icono de disco" />
                    <p>Disco 2</p>
                </div>
                <div className="file">
                    <img src={icono} alt="Icono de disco" />
                    <p>Disco 3</p>
                </div>
                <div className="file">
                    <img src={icono} alt="Icono de disco" />
                    <p>Disco 3</p>
                </div>
                <div className="file">
                    <img src={icono} alt="Icono de disco" />
                    <p>Disco 3</p>
                </div>
                <div className="file">
                    <img src={icono} alt="Icono de disco" />
                    <p>Disco 3</p>
                </div>
                <div className="file">
                    <img src={icono} alt="Icono de disco" />
                    <p>Disco 3</p>
                </div>
                <div className="file">
                    <img src={icono} alt="Icono de disco" />
                    <p>Disco 3</p>
                </div>
                <div className="file">
                    <img src={icono} alt="Icono de disco" />
                    <p>Disco 3</p>
                </div>
                <div className="file">
                    <img src={icono} alt="Icono de disco" />
                    <p>Disco 3</p>
                </div>
                <div className="file">
                    <img src={icono} alt="Icono de disco" />
                    <p>Disco 3</p>
                </div>
                {/* Agrega más elementos de archivo según sea necesario */}
            </div>
        </div>
    );
}

export default Explorer;