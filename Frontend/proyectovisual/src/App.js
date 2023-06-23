//import logo from './logo.svg';
import './App.css';
import {BrowserRouter as Router, Route, Routes} from 'react-router-dom';
import { Login } from './components/menuPrincipal';

function App() {
  return (
    <Router>
      <Routes>
        <Route exact path='/' element={<Login/>} />
      </Routes>
    </Router>
  );
  /*
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
      <footer className="App-footer">
        Hi estoy probando
      </footer>
    </div>
    
  );
  */
}

export default App;
