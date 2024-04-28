import React from 'react';
import './Explorer.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faArrowLeft } from '@fortawesome/free-solid-svg-icons';
import icono from '../img/Particion.png';

const Partitions = ({ onShowExplorer, partitions }) => {
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
                    <div className="file" key={index}>
                        <img src={icono} alt="Icono de particion" title={partition[0]}/>
                        <p>{partition[1]}</p>
                    </div>
                ))}
            </div>
        </div>
    );
}

export default Partitions;