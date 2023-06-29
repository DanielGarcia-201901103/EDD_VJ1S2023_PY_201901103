import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";
import Card from "react-bootstrap/Card";
import Container from 'react-bootstrap/Container';
import Navbar from 'react-bootstrap/Navbar';
import { useState } from 'react';

export const ApliFiltros = () => {
    const usuarioIniciado = localStorage.getItem('current');
    var [imagenActual, setImagenAc] = useState("");
    var obtenIm = localStorage.getItem('nombreImagen');
    const apliNegativo = async (e) => {
        console.log(imagenActual)
        e.preventDefault();
        await fetch('http://localhost:5000/filtroNegativo', {
            method: 'POST',
            mode: 'cors',
            body: JSON.stringify({
                Tipo: "escalaNegativo",
                Imagen: obtenIm,
            }),
            headers: {
                'Access-Control-Allow-Origin': '*',
                'Content-Type': 'application/json'
            }
        })
    }

    const apliGrises = async (e) => {
        e.preventDefault();
        await fetch('http://localhost:5000/filtroNegativo', {
            method: 'POST',
            mode: 'cors',
            body: JSON.stringify({
                Tipo: "escalaGris",
                Imagen: obtenIm,
            }),
            headers: {
                'Access-Control-Allow-Origin': '*',
                'Content-Type': 'application/json'
            }
        })
    }

    const apliX = async (e) => {
        e.preventDefault();
        await fetch('http://localhost:5000/filtroNegativo', {
            method: 'POST',
            mode: 'cors',
            body: JSON.stringify({
                Tipo: "espejoX",
                Imagen: obtenIm,
            }),
            headers: {
                'Access-Control-Allow-Origin': '*',
                'Content-Type': 'application/json'
            }
        })
    }

    const apliY = async (e) => {
        e.preventDefault();
        await fetch('http://localhost:5000/filtroNegativo', {
            method: 'POST',
            mode: 'cors',
            body: JSON.stringify({
                Tipo: "espejoY",
                Imagen: obtenIm,
            }),
            headers: {
                'Access-Control-Allow-Origin': '*',
                'Content-Type': 'application/json'
            }
        })
    }

    const apliAmbos = async (e) => {
        e.preventDefault();
        await fetch('http://localhost:5000/filtroNegativo', {
            method: 'POST',
            mode: 'cors',
            body: JSON.stringify({
                Tipo: "dobleEspejo",
                Imagen: obtenIm,
            }),
            headers: {
                'Access-Control-Allow-Origin': '*',
                'Content-Type': 'application/json'
            }
        })
    }

    const obtenerCliente = async (e) => {
        e.preventDefault();
        await fetch('http://localhost:5000/clienteObtener', {
            method: 'GET',
            mode: 'cors',
            headers: {
                'Access-Control-Allow-Origin': '*',
                'Content-Type': 'application/json'
            }
        })
            .then(response => response.json())
            .then(data => {
                //console.log(data.data)
                //console.log(data.imagen)
                imagenActual = data.imagen
                localStorage.setItem('nombreImagen', imagenActual);
                setImagenAc(data.imagen)
            });
            //setClienteAc("")
    }

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
                    padding: "200px",
                    backgroundColor: "#282c34",
                    display: "flex",
                    paddingLeft: "600px",
                    flexdirection: "column",
                    minheight: "100vh",
                    alignitems: "center",
                }}
            >
                <Card style={{ margin: '0 1.5%', width: '25rem', background: '#D3CBB8', display: 'inline-block' }}>
                    <Card.Body>
                        <Card.Title>Filtros</Card.Title>
                        <Form>
                            <br />
                            <Button
                                className="w-100 btn btn-lg btn-primary"
                                variant='dark'
                                onClick={apliNegativo}
                            >
                                Aplicar Negativo
                            </Button>
                            <br />
                            <br />
                            <Button
                                className="w-100 btn btn-lg btn-primary"
                                variant='dark'
                                onClick={apliGrises}
                            >
                                Aplicar Escala de Grises
                            </Button>
                            <br />
                            <br />
                            <Button
                                className="w-100 btn btn-lg btn-primary"
                                variant='dark'
                                onClick={apliX}
                            >
                                Aplicar Espejo X
                            </Button>
                            <br />
                            <br />
                            <Button
                                className="w-100 btn btn-lg btn-primary"
                                variant='dark'
                                onClick={apliY}
                            >
                                Aplicar Espejo Y
                            </Button>
                            <br />
                            <br />
                            <Button
                                className="w-100 btn btn-lg btn-primary"
                                variant='dark'
                                onClick={apliAmbos}
                            >
                                Aplicar Ambos Espejos
                            </Button>
                            <br />
                            <br />
                            <Button
                                className="w-100 btn btn-lg btn-primary"
                                variant='success'
                                onClick={obtenerCliente}
                            >
                                Obtener Imagen del cliente
                            </Button>
                            <Form.Group className="mb-3">
                                <Form.Label>Empleado Cobrador</Form.Label>
                                <Form.Control type="text" placeholder={imagenActual} aria-label="Disabled input example" disabled readOnly />
                            </Form.Group>
                        </Form>

                    </Card.Body>
                </Card>
            </div>
        </>
    );
};