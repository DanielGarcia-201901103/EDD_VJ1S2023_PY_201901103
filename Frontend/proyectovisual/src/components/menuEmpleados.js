import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";
import Card from "react-bootstrap/Card";
import Container from 'react-bootstrap/Container';
import Navbar from 'react-bootstrap/Navbar';
//import { useState } from 'react';

export const LoginEmpleados = () => {
    const usuarioIniciado = localStorage.getItem('current');
    const clickAFiltros = async (e) => {
        window.open("/aplicacionFiltros", "_self");
    }
    const clickGFactura = async (e) => {
        window.open("/generarFactura", "_self");
    }
    const clickHFacturas = async (e) => {
        window.open("/historialFacturas", "_self");
    }
    const clickVeCompletadas = async (e) => {
        window.open("/ventasCompletadas", "_self");
    }

/*
Aplicación Filtros
	aplicar negativo
	aplicar escala de grises
	aplicar espejo x
	aplicar espejo y
	aplicar ambos espejos
	generar imagen con
	regresar al menu anterior
Generar Factura
	-fecha
	-empleado cobrador
	-usuario
	-pago
	Realizar pago
	regresar al menu anterior
Historial de facturas
	Mostrar tabla con la lista de cl antendidos id cliente, y id factura y el id del empleado
	Regresar al menu anterior
Ventas completadas
	- tabla con id cliente, nombre imagen, filtros elegidos
	Reporte
	Regresar al menu anterior
cerrar sesion
 */
    return (
        <>
            <div >
                <Navbar className="bg-body-tertiary" style={{ background: '#bdc3c7', background: '-webkit-linear-gradient(to right, #bdc3c7, #2c3e50)', background: 'linear-gradient(to right, #bdc3c7, #2c3e50)', minheight: '100vh' }}>
                    <Container>
                        <Navbar.Brand href="/">Cerrar Sesión</Navbar.Brand>
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
                <Card style={{margin: '0 1.5%', width: '25rem', background: '#D3CBB8', display: 'inline-block'  }}>
                    <Card.Body>
                        <Card.Title>Empleado: {usuarioIniciado}</Card.Title>
                        <Form>
                        <br />
                            <Button
                                className="w-100 btn btn-lg btn-primary"
                                variant='dark'
                                onClick={clickAFiltros}
                            >
                                Aplicación Filtros
                            </Button>
                            <br />
                            <br />
                            <Button
                                className="w-100 btn btn-lg btn-primary"
                                variant='dark'
                                onClick={clickGFactura}
                            >
                                Generar Factura
                            </Button>
                            <br />
                            <br />
                            <Button
                                className="w-100 btn btn-lg btn-primary"
                                variant='dark'
                                onClick={clickHFacturas}
                            >
                                Historial de Facturas
                            </Button>
                            <br />
                            <br />
                            <Button
                                className="w-100 btn btn-lg btn-primary"
                                variant='dark'
                                onClick={clickVeCompletadas}
                            >
                                Ventas Completadas
                            </Button>
                        </Form>
                    </Card.Body>
                </Card>
            </div>
        </>
    );
};
