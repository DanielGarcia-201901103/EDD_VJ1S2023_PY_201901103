import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import Card from 'react-bootstrap/Card';
//import { useState } from 'react';
import Container from 'react-bootstrap/Container';
import Navbar from 'react-bootstrap/Navbar';

export const LoginAdmin = () => {
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
    const cargarPedidos = (event) => {
        const file = event.target.files[0];
        const filename = file.name;
        console.log('Archivo cargado:', event.target.files[0].name);
    }

    const cargarEmpleados = () => {
        alert("Cargando Empleados")

    }
    const obtenerReportes = () => {
        alert("Reportes")

    }

    return (
        <>
            <div >
                <Navbar className="bg-body-tertiary" style={{background: '#bdc3c7', background: '-webkit-linear-gradient(to right, #bdc3c7, #2c3e50)', background: 'linear-gradient(to right, #bdc3c7, #2c3e50)',minheight: '100vh' }}>
                    <Container>
                        <Navbar.Brand href="/">Cerrar Sesión</Navbar.Brand>
                        <Navbar.Toggle />
                        <Navbar.Collapse className="justify-content-center">
                            <Navbar.Text>
                                Signed in as: Administrador 201901103
                            </Navbar.Text>
                        </Navbar.Collapse>
                    </Container>
                </Navbar>
            </div>
            <div style={{padding: "200px", paddingTop: '75px', paddingLeft: '300px',backgroundColor: '#282c34', display: 'flex', flexdirection: 'column', minheight: '100vh', alignitems: 'center' }}>
                <Card style={{margin: '0 1.5%', width: '25rem', background: '#D3CBB8', display: 'inline-block'}}>
                    <Card.Body>
                        <Card.Title>Pedidos</Card.Title>
                        <Form onSubmit={cargarPedidos}>
                            <br />
                            <input type="file" name="mensaje" id="mensaje" required></input>
                            <br />
                            <br />
                            <Button className="w-100 btn btn-lg btn-primary" variant='dark' type="submit">
                                Cargar Pedidos
                            </Button>
                        </Form>
                    </Card.Body>
                </Card>

                <Card style={{margin: '0 1.5%', width: '25rem', background: '#D3CBB8', display: 'inline-block'}}>
                    <Card.Body>
                        <Card.Title>Empleados</Card.Title>
                        <Form onSubmit={cargarEmpleados}>
                            <br />
                            <input type="file" name="mensaje1" id="mensaje1" required></input>
                            <br />
                            <br />
                            <Button className="w-100 btn btn-lg btn-primary" variant='dark' type="submit">
                                Cargar Empleados
                            </Button>
                        </Form>
                    </Card.Body>
                </Card>
            </div>
            <div style={{padding: "200px", paddingTop: '1px',paddingLeft: '300px',backgroundColor: '#282c34', display: 'flex', flexdirection: 'column', minheight: '100vh', alignitems: 'center' }}>
            <Card style={{margin: '0 1.5%', width: '25rem', background: '#D3CBB8', display: 'inline-block'}}>
                    <Card.Body>
                        <Card.Title>Reportes</Card.Title>
                        <Form>
                            <br />
                            <Button className="w-100 btn btn-lg btn-primary" variant='dark' onClick={obtenerReportes}>
                                Reportes
                            </Button>
                        </Form>
                    </Card.Body>
                </Card>
            </div>
        </>
    );
}