//import logo from './logo.svg';
import './App.css';
import {BrowserRouter as Router, Route, Routes} from 'react-router-dom';
import { Login } from './components/menuPrincipal';
import { LoginAdmin } from './components/menuAdministrador';
import { LoginEmpleados } from './components/menuEmpleados';
import { ApliFiltros } from './components/aplicacionFiltros';
import { GeneraFact } from './components/generarFactura';
import { HistoryFact } from './components/historialFacturas';
import { VentCompl } from './components/ventasCompletadas';

function App() {
  return (
    <Router>
      <Routes>
        <Route exact path='/' element={<Login/>} />
        <Route exact path='/admin' element={<LoginAdmin/>} />
        <Route exact path='/empleado' element={<LoginEmpleados/>} />
        <Route exact path='/aplicacionFiltros' element={<ApliFiltros/>} />
        <Route exact path='/generarFactura' element={<GeneraFact/>} />
        <Route exact path='/historialFacturas' element={<HistoryFact/>} />
        <Route exact path='/ventasCompletadas' element={<VentCompl/>} />
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
