import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import Card from 'react-bootstrap/Card';
import { useState } from 'react';
import Container from 'react-bootstrap/Container';
import Navbar from 'react-bootstrap/Navbar';

export const LoginAdmin = () => {
    const [rutaPedidos, setRutaP] = useState("");
    const [rutaEmpleados, setrutaEmpleados] = useState("");
    const [imagen, setImagen] = useState('./guitarra.jpg')
    const [imagenOtro, setImagenOtro] = useState('./guitarra1.jpg')
    const cargaPedidos = async (e) => {
        e.preventDefault();
        await fetch('http://localhost:5000/cargarPedidos', {
            method: 'POST',
            mode: 'cors',
            body: JSON.stringify({
                Ruta: rutaPedidos
            }),
            headers: {
                'Access-Control-Allow-Origin': '*',
                'Content-Type': 'application/json'
            }
        })
            .then(response => response.json())
            .then(data => validar(data));
        setRutaP("")
        setrutaEmpleados("")
    }
    const validar = (data) => {
        console.log(data)
    }
    const cargarEmpleados = async (e) => {
        e.preventDefault();
        await fetch('http://localhost:5000/cargaEmpleados', {
            method: 'POST',
            mode: 'cors',
            body: JSON.stringify({
                Ruta: rutaEmpleados
            }),
            headers: {
                'Access-Control-Allow-Origin': '*',
                'Content-Type': 'application/json'
            }
        })
            .then(response => response.json())
            .then(data => validar1(data));
        setrutaEmpleados("")
    }
    const validar1 = (data) => {
        console.log(data)
    }

    const obtenerReportes = async (e) => {
        e.preventDefault();

        await fetch('http://localhost:5000/Reportes', {
            method: 'GET',
            mode: 'cors',
            headers: {
                'Access-Control-Allow-Origin': '*',
                'Content-Type': 'application/json'
            }
        })
        .then(response => response.json())
        .then(data => validar2(data));
    }
    const validar2 = (data) => {
        setImagen(data.data)
    }

    const obtenerReportesOtro = async (e) => {
        e.preventDefault();
        await fetch('http://localhost:5000/reporteBloquePago', {
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

    return (
        <>
            <div >
                <Navbar className="bg-body-tertiary" style={{ background: '#bdc3c7', background: '-webkit-linear-gradient(to right, #bdc3c7, #2c3e50)', background: 'linear-gradient(to right, #bdc3c7, #2c3e50)', minheight: '100vh' }}>
                    <Container>
                        <Navbar.Brand href="/">Cerrar Sesión</Navbar.Brand>
                        <Navbar.Toggle />
                        <Navbar.Collapse className="justify-content-center">
                            <Navbar.Text>
                                Usuario: Administrador 201901103
                            </Navbar.Text>
                        </Navbar.Collapse>
                    </Container>
                </Navbar>
            </div>
            <div style={{ padding: "200px", paddingTop: '75px', paddingLeft: '300px', backgroundColor: '#282c34', display: 'flex', flexdirection: 'column', minheight: '100vh', alignitems: 'center' }}>
                <Card style={{ margin: '0 1.5%', width: '25rem', background: '#D3CBB8', display: 'inline-block' }}>
                    <Card.Body>
                        <Card.Title>Pedidos</Card.Title>
                        <Form onSubmit={cargaPedidos}>
                            <Form.Group className="mb-3">
                                <Form.Label>Ruta del Archivo</Form.Label>
                                <Form.Control id='fileruta' type="text" placeholder="archivo.json" required
                                    onChange={e => setRutaP(e.target.value)}
                                    value={rutaPedidos}
                                    autoFocus />
                            </Form.Group>
                            <Button className="w-100 btn btn-lg btn-primary" variant='dark' type="submit">
                                Cargar Pedidos
                            </Button>
                        </Form>
                    </Card.Body>
                </Card>

                <Card style={{ margin: '0 1.5%', width: '25rem', background: '#D3CBB8', display: 'inline-block' }}>
                    <Card.Body>
                        <Card.Title>Empleados</Card.Title>
                        <Form onSubmit={cargarEmpleados}>
                            <Form.Group className="mb-3">
                                <Form.Label>Ruta del Archivo</Form.Label>
                                <Form.Control id='fileR' type="text" placeholder="archivo.csv" required
                                    onChange={e => setrutaEmpleados(e.target.value)}
                                    value={rutaEmpleados}
                                />
                            </Form.Group>
                            <Button className="w-100 btn btn-lg btn-primary" variant='dark' type="submit">
                                Cargar Empleados
                            </Button>
                        </Form>
                    </Card.Body>
                </Card>
            </div>
            <div style={{ padding: "200px", paddingTop: '1px', paddingLeft: '150px', backgroundColor: '#282c34', display: 'flex', flexdirection: 'column', minheight: '100vh', alignitems: 'center' }}>
                <Card style={{ margin: '0 1.5%', width: '55rem', background: '#D3CBB8', display: 'inline-block' }}>
                    <Card.Img variant="top" src= {imagen} width="300" height="400" alt='reporte arbol AVL'/>
                    <Card.Body>
                        <Card.Title>Reporte</Card.Title>
                        <Form>
                            <br />
                            <Button className="w-100 btn btn-lg btn-primary" variant='dark' onClick={obtenerReportes}>
                                Reporte de Arbol
                            </Button>
                        </Form>
                    </Card.Body>
                </Card>
                <Card style={{ margin: '0 1.5%', width: '55rem', background: '#D3CBB8', display: 'inline-block' }}>
                    <Card.Img variant="top" src= {imagenOtro} width="300" height="400" alt='reporte pagos'/>
                    <Card.Body>
                        <Card.Title>Reporte</Card.Title>
                        <Form>
                            <br />
                            <Button className="w-100 btn btn-lg btn-primary" variant='dark' onClick={obtenerReportesOtro}>
                                Reporte de Pagos
                            </Button>
                        </Form>
                    </Card.Body>
                </Card>
            </div>
        </>
    );
}