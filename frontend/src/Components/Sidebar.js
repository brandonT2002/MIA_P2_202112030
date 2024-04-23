import React from 'react';
import './Sidebar.css';

const Sidebar = ({ onOptionSelect }) => {
    return (
        <div className="sidebar">
            <div className="sidebar-header">
                <h1 className="logo__name no-margin center-text"><span className="logo__bold">LAB</span> MIA</h1>
            </div>
            <ul>
                <li onClick={() => onOptionSelect('Consola')}>Consola</li>
                <li onClick={() => onOptionSelect('Explorador')}>Explorador</li>
                <li onClick={() => onOptionSelect('Reportes')}>Reportes</li>
            </ul>
        </div>
    );
}

export default Sidebar;