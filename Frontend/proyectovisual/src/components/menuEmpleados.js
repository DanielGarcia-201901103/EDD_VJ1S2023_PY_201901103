import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";
import Card from "react-bootstrap/Card";
//import { useState } from 'react';

export const LoginEmpleados = () => {
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
    return (
        <>
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
                <Card style={{ width: "18rem" }}>
                    <Card.Body>
                        <Card.Title>Empleado faltaObtenerUser</Card.Title>
                        <Form>
                            <Button
                                className="w-100 btn btn-lg btn-primary"
                                onClick={alert("Pendiente1")}
                            >
                                Cargar Pedidos
                            </Button>
                            <br />
                            <Button
                                className="w-100 btn btn-lg btn-primary"
                                onClick={alert("Pendiente2")}
                            >
                                Cargar Empleados
                            </Button>
                            <br />
                            <Button
                                className="w-100 btn btn-lg btn-primary"
                                onClick={alert("Pendiente3")}
                            >
                                Reportes
                            </Button>
                        </Form>
                    </Card.Body>
                </Card>
            </div>
        </>
    );
};
