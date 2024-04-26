import React from 'react';
import './Sidebar.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faTerminal, faFolderOpen, faChartBar } from '@fortawesome/free-solid-svg-icons'; // Importa los iconos que necesitas

const Sidebar = ({ onOptionSelect }) => {
    return (
        <div className="sidebar">
            <div className="sidebar-header">
                <h1 className="logo__name no-margin center-text"><span className="logo__bold">LAB</span> MIA</h1>
            </div>
            <ul>
                <li onClick={() => onOptionSelect('Consola')}>
                    <FontAwesomeIcon icon={faTerminal} className="sidebar-icon" />
                    Consola
                </li>
                <li onClick={() => onOptionSelect('Explorador')}>
                    <FontAwesomeIcon icon={faFolderOpen} className="sidebar-icon" />
                    Explorador
                </li>
                <li onClick={() => onOptionSelect('Reportes')}>
                    <FontAwesomeIcon icon={faChartBar} className="sidebar-icon" />
                    Reportes
                </li>
            </ul>
        </div>
    );
}

export default Sidebar;