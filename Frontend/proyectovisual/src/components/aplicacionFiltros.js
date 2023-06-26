import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";
import Card from "react-bootstrap/Card";
import Container from 'react-bootstrap/Container';
import Navbar from 'react-bootstrap/Navbar';
//import { useState } from 'react';

export const ApliFiltros = () => {
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
                        <Card.Title>Filtros</Card.Title>
                        <Form>
                        <br />
                            <Button
                                className="w-100 btn btn-lg btn-primary"
                                variant='dark'
                                onClick={clickB}
                            >
                                Aplicar Negativo
                            </Button>
                            <br />
                            <br />
                            <Button
                                className="w-100 btn btn-lg btn-primary"
                                variant='dark'
                                onClick={clickB}
                            >
                                Aplicar Escala de Grises
                            </Button>
                            <br />
                            <br />
                            <Button
                                className="w-100 btn btn-lg btn-primary"
                                variant='dark'
                                onClick={clickB}
                            >
                                Aplicar Espejo X
                            </Button>
                            <br />
                            <br />
                            <Button
                                className="w-100 btn btn-lg btn-primary"
                                variant='dark'
                                onClick={clickB}
                            >
                                Aplicar Espejo Y
                            </Button>
                            <br />
                            <br />
                            <Button
                                className="w-100 btn btn-lg btn-primary"
                                variant='dark'
                                onClick={clickB}
                            >
                                Aplicar Ambos Espejos
                            </Button>
                            <br />
                            <br />
                            <Button
                                className="w-100 btn btn-lg btn-primary"
                                variant='success'
                                onClick={clickB}
                            >
                                Generar Imagen con Filtros
                            </Button>
                        </Form>
                        
                    </Card.Body>
                </Card>
            </div>
        </>
    );
};