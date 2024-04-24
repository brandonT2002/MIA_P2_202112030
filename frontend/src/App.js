import React, { useState } from 'react';
import './App.css';
import Sidebar from './Components/Sidebar';
import Console from './Components/Console';
import Explorer from './Components/Explorer';

const App = () => {
    const [selectedOption, setSelectedOption] = useState('Consola');

    const handleOptionSelect = (option) => {
        setSelectedOption(option);
    };

    return (
        <div className="appColumns">
            <Sidebar onOptionSelect={handleOptionSelect} />
            <div className="content">
                {selectedOption === 'Consola' ?
                    <Console/>
                : (selectedOption === 'Explorador' ?
                    <Explorer/>
                :
                    <div>ReportesReportesReportesReportesReportesReportesReportesReportesReportesReportesReportesReportesReportesReportes</div>
                )}
            </div>
        </div>
    );
}

export default App;
