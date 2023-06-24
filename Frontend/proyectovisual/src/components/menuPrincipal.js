import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import Card from 'react-bootstrap/Card';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

export const Login = () => {
    const navigate = useNavigate();
    const [userLogin, setUsuario] = useState("")
    const [passwordLogin, setPassword] = useState("")
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

        .then(response => response.json())
        .then(data => validar(data));
        /*
        .then(function(response) {
        if (response.ok) {
            return response.json();
        }
        throw new Error('Error en la solicitud POST');
        })
        .then(function(responseData) {
            // Aquí puedes acceder a la respuesta del backend (responseData)
        })
        .catch(function(error) {
            console.log('Error:', error.message);
        });*/
    }
    
    const validar = (data) =>{
        console.log(data)
        //setImagen(data.Imagenbase64)
        if (data.data === "Administrador"){
            window.localStorage.setItem("Administrador","201901103")
            //window.open("/admin","_self")
            navigate('/admin')
            console.log("estoy en admin")
        }else if (data.data === "SI"){
            /*
            swal({
                title: "Session Iniciada",
                text: "Datos Correctos",
                icon: "success",
                button: "aceptar"
            }).then(respuesta => {
                if(respuesta){
                    console.log(respuesta);*/
                    localStorage.setItem('current', userLogin);
                    window.open("/empleado","_self");
            //    }
            //})
            console.log("estoy en cualquier usuario")
        }else{
            alert("Contraseña o usuario incorrecto")
        }
    }

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