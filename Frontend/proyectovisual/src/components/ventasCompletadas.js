import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";
import Card from "react-bootstrap/Card";
import Container from 'react-bootstrap/Container';
import Navbar from 'react-bootstrap/Navbar';
import Badge from 'react-bootstrap/Badge';
import Table from 'react-bootstrap/Table';
import { useState, useEffect } from 'react';

export const VentCompl = () => {
    const usuarioIniciado = localStorage.getItem('current');
    const [obtenidoFactura, setFact] = useState([]);
    const [imagenOtro, setImagenOtro] = useState('./guitarra1.jpg')

    useEffect(() => {
        mostrarTabla()
    },[])

    const mostrarTabla = () => {
        fetch('http://localhost:5000/solicitudClientes', {
        })
        .then(response => response.json())
        .then(data => validar(data))
    }

    const validar = (data) =>{
        console.log(data.data)
        setFact(data.data)
    }
    
    const obtenerReportesOtro = async (e) => {
        e.preventDefault();
        await fetch('http://localhost:5000/reporteGrafo', {
            method: 'GET',
            mode: 'cors',
            headers: {
                'Access-Control-Allow-Origin': '*',
                'Content-Type': 'application/json'
            }
        })
        .then(response => response.json())
        .then(data => validar3(data));
    }
    const validar3 = (data) => {
        setImagenOtro(data.data)
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
                <h1><Badge bg="secondary">Solicitudes de Filtros por Clientes </Badge></h1>

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
                            <th>ID Cliente</th>
                            <th>Imagen</th>
                            <th>Filtros</th>
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
                            <td>Jacob</td>
                            <td>Thornton</td>
                        </tr>
                    </tbody>
                </Table>
            </div>
            <div
                style={{
                    padding: "200px",
                    backgroundColor: "#282c34",
                    display: "flex",
                    paddingLeft: "300px",
                    flexdirection: "column",
                    minheight: "100vh",
                    alignitems: "center",
                }}
            >
                <Card style={{ margin: '0 1.5%', width: '55rem', background: '#D3CBB8', display: 'inline-block' }}>
                    <Card.Img variant="top" src= {imagenOtro} width="300" height="400" alt='reporte pagos'/>
                    <Card.Body>
                        <Card.Title>Reporte</Card.Title>
                        <Form>
                            <br />
                            <Button className="w-100 btn btn-lg btn-primary" variant='dark' onClick={obtenerReportesOtro}>
                                Reporte Grafo
                            </Button>
                        </Form>
                    </Card.Body>
                </Card>
            </div>
        </>
    );
};