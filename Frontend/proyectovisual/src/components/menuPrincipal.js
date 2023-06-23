import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import Card from 'react-bootstrap/Card';
import { useState } from 'react';

export const Login = () => {
    const [userLogin, setUsuario] = useState()
    const [passwordLogin, setPassword] = useState()
    //const [imagen, setImagen] = useState('https://yakurefu.com/wp-content/uploads/2020/02/Chi_by_wallabby.jpg')
    const handleSubmit = async(e) => {
        e.preventDefault();
        await fetch('http://localhost:5000/login',{
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
    }
    /*
    const validar = (data) =>{
        //console.log(data)
        //setImagen(data.Imagenbase64)
        if (data.mensaje === "OK"){
            swal({
                title: "Session Iniciada",
                text: "Datos Correctos",
                icon: "success",
                button: "aceptar"
            }).then(respuesta => {
                if(respuesta){
                    console.log(respuesta);
                    localStorage.setItem('current',respuesta.data);
                    window.open("/inicio","_self");
                }
            })
            
        }else{
            swal({
                title: "Error en Credenciales",
                text: "Su usuario o contraseña son incorrectos",
                icon: "error",
                timer: "4000",
                buttons: false
            })
        }
    }*/
    return (
        <div style={{padding: '200px', backgroundColor: '#282c34', display: 'flex', paddingLeft: '600px', flexdirection: 'column' ,minheight: '100vh',alignitems: 'center'}}>
        <Card style={{ width: '18rem'}}>
        <Card.Body>
            <Card.Title>Iniciar Sesión</Card.Title>
            <Form onSubmit={handleSubmit} >
            <Form.Group className="mb-3" controlId="userEmpleado">
                <Form.Label>Usuario</Form.Label>
                <Form.Control type="text" placeholder="Enter User" required  
                    onChange={e => setUsuario(e.target.value)} 
                    value={userLogin} 
                    autoFocus/>
            </Form.Group>
            <Form.Group className="mb-3" controlId="userPassword">
                <Form.Label>Password</Form.Label>
                <Form.Control type="password" placeholder="Password" required
                onChange={e => setPassword(e.target.value)} 
                value={passwordLogin} />
            </Form.Group>
            <Button variant="primary" type="submit" >
                Login
            </Button>
            </Form>
        </Card.Body>
    </Card>
        </div>
    );
}
/*
https://create-react-app.dev/docs/adding-bootstrap/#using-a-custom-theme
https://react-bootstrap.github.io/docs/getting-started/introduction
*/