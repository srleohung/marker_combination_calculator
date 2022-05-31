import React from 'react';
import logo from './logo.svg';
import { Display } from './features/display/Display';
import './App.css';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <Display />
      </header>
    </div>
  );
}

export default App;
