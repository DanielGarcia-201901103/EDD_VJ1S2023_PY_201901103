import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";
import Card from "react-bootstrap/Card";
import Container from 'react-bootstrap/Container';
import Navbar from 'react-bootstrap/Navbar';
import { useState } from 'react';

export const GeneraFact = () => {
    const usuarioIniciado = localStorage.getItem('current');
    const clienteObtenido = localStorage.getItem('cliente');
    const moment = require('moment');
    const fechaHora = moment().format('DD-MM-YYYY-::HH:mm:ss');
    const [pagoR, setPago] = useState("");
    //DD-MM-YY-::HH:MM:SS
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
        alert(clienteObtenido+"  "+usuarioIniciado)
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
                        <Card.Title alignitems= "center">Generar Factura</Card.Title>
                        <Form onSubmit={clickB}>
                            <Form.Group className="mb-3">
                                <Form.Label>Fecha</Form.Label>
                                <Form.Control type="text" placeholder={fechaHora} aria-label="Disabled input example" disabled readOnly required/>
                            </Form.Group>
                            <Form.Group className="mb-3">
                                <Form.Label>Empleado Cobrador</Form.Label>
                                <Form.Control type="text" placeholder={usuarioIniciado} aria-label="Disabled input example" disabled readOnly required/>
                            </Form.Group>
                            <Form.Group className="mb-3">
                                <Form.Label>Usuario</Form.Label>
                                <Form.Control type="text" placeholder={clienteObtenido} aria-label="Disabled input example" disabled readOnly required/>
                            </Form.Group>
                            <Form.Group className="mb-3">
                                <Form.Label>Pago</Form.Label>
                                <Form.Control type="text" placeholder="Q 00.00" required autoFocus
                                onChange={e => setPago(e.target.value)}
                                value={pagoR} />
                            </Form.Group>
                            <Button
                                className="w-100 btn btn-lg btn-primary"
                                variant='dark'
                                type="submit"
                            >
                                Realizar Pago
                            </Button>
                        </Form>
                    </Card.Body>
                </Card>
            </div>
        </>
    );
};