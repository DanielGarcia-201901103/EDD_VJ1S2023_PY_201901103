import Container from 'react-bootstrap/Container';
import Navbar from 'react-bootstrap/Navbar';
import Badge from 'react-bootstrap/Badge';
import Table from 'react-bootstrap/Table';
//import { useState } from 'react';

export const HistoryFact = () => {
    const usuarioIniciado = localStorage.getItem('current');
    /*
      const [userLogin, setUsuario] = useState()
      const [passwordLogin, setPassword] = useState()
      //const [imagen, setImagen] = useState('https://yakurefu.com/wp-content/uploads/2020/02/Chi_by_wallabby.jpg')
      const handleSubmit = async(e) => {
          e.preventDefault();
          await fetch('http://localhost:5000/loginAdmin',{
              method: 'POST',
              mode: 'cors',
              body: JSON.stringify({
                  Usuario: userLogin,
                  Password: passwordLogin
              }),
              headers:{
                  'Access-Control-Allow-Origin': '*',
                  'Content-Type': 'application/json'
              }
          })
      }*/

    const clickB = async (e) => {
        /*
        e.preventDefault();
        await fetch('http://localhost:5000/Reportes', {
            method: 'GET',
            mode: 'cors',
            headers: {
                'Access-Control-Allow-Origin': '*',
                'Content-Type': 'application/json'
            }
        })*/
        alert("hola click")
    }
    /*Aplicación Filtros
    Generar Factura
    Historial de facturas
    Ventas completadas */
    return (
        <>
            <div >
                <Navbar className="bg-body-tertiary" style={{ background: '#bdc3c7', background: '-webkit-linear-gradient(to right, #bdc3c7, #2c3e50)', background: 'linear-gradient(to right, #bdc3c7, #2c3e50)', minheight: '100vh' }}>
                    <Container>
                        <Navbar.Brand href="/empleado">Regresar</Navbar.Brand>
                        <Navbar.Toggle />
                        <Navbar.Collapse className="justify-content-center">
                            <Navbar.Text>
                                Sesión: {usuarioIniciado}
                            </Navbar.Text>
                        </Navbar.Collapse>
                    </Container>
                </Navbar>
            </div>
            <div
                style={{
                    padding: "10px",
                    paddingTop: "75px",
                    backgroundColor: "#282c34",
                    paddingLeft: "400px",
                    minheight: "100vh",
                    alignitems: "center",
                }}
            >
                <h1><Badge bg="secondary">Facturas Generadas por Empleado {usuarioIniciado} </Badge></h1>

            </div>
            <div
                style={{
                    padding: "50px",
                    paddingTop: "100px",
                    backgroundColor: "#282c34",
                    display: "flex",
                    paddingLeft: "200px",
                    flexdirection: "column",
                    minheight: "100vh",
                    alignitems: "center",
                }}
            >
                <Table striped bordered hover variant="dark">
                    <thead>
                        <tr>
                            <th>#</th>
                            <th>ID Cliente</th>
                            <th>ID Factura</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <td>1</td>
                            <td>Mark</td>
                            <td>Otto</td>
                        </tr>
                        <tr>
                            <td>2</td>
                            <td>Jacob</td>
                            <td>Thornton</td>
                        </tr>
                        <tr>
                            <td>3</td>
                            <td colSpan={2}>Larry the Bird</td>
                        </tr>
                    </tbody>
                </Table>
            </div>
            <div
                style={{
                    paddingTop: "200px",
                    backgroundColor: "#282c34",
                    display: "flex",
                    paddingLeft: "200px",
                    flexdirection: "column",
                    minheight: "100vh",
                    alignitems: "center",
                }}
            ></div>
        </>
    );
};