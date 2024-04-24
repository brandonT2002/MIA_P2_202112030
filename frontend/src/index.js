import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import { MessageProvider } from './Components/MessageContext';

ReactDOM.render(
  <React.StrictMode>
    <MessageProvider>
      <App />
    </MessageProvider>
  </React.StrictMode>,
  document.getElementById('root')
);
