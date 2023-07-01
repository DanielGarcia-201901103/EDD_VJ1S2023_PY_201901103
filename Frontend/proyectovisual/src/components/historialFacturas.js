import Container from 'react-bootstrap/Container';
import Navbar from 'react-bootstrap/Navbar';
import Badge from 'react-bootstrap/Badge';
import Table from 'react-bootstrap/Table';
import { useState,useEffect } from 'react';

export const HistoryFact = () => {
    const usuarioIniciado = localStorage.getItem('current');
    const [obtenidoFactura, setFact] = useState([])

    useEffect(() => {
        mostrarTabla()
    },[])

    const mostrarTabla = () => {
        fetch('http://localhost:5000/obTabla', {
        })
        .then(response => response.json())
        .then(data => validar(data))
    }

    const validar = (data) =>{
        console.log(data.data)
        setFact(data.data)
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
                    {
                                obtenidoFactura.map((element, j) => {
                                    if (element.Id_Cliente != '') {
                                        return <>
                                        <tr key={"fact"+j}>
                                            <th scope="row">{j+1}</th>
                                            <td>{element.Id_Cliente}</td>
                                            <td>{element.Id_Factura}</td>
                                        </tr>
                                    </>
                                    }
                                })
                            }
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