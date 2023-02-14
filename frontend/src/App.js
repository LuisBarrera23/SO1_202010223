import { useRef, useState } from 'react';
import './App.css';

function App() {


  const [valor1, setValor1] = useState('');
  const [valor2, setValor2] = useState('');
  const LabelResultado=useRef(null);

  const verificarNumero1 = event => {
    const { value } = event.target;
    if (!isNaN(value) || value ==='-') {  //verificamos si valor no es un numero seteamos el valor anterior
      setValor1(value);    //es decir ignoramos el valor que se ingreso (una letra)
    }
  };

  const verificarNumero2 = event => {
    const { value } = event.target;
    if (!isNaN(value) || value ==='-') {  //verificamos si valor no es un numero seteamos el valor anterior
      setValor2(value);    //es decir ignoramos el valor que se ingreso (una letra)
    }
  };

  async function suma() {

    if(valor1 === ""&&valor2 === ""){
      alert("Por favor ingrese ambos numeros")
    }else if(valor1 === ""){
      alert("Por favor ingrese el primer numero")
    }else if(valor2 === ""){
      alert("Por favor ingrese el segundo numero")
    }else{
      const jsonDatos={"Simbolo":"+","Numero1":valor1,"Numero2":valor2}
      console.log(jsonDatos)

      try {
        const response = await fetch('http://localhost:5000/suma', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' ,'Access-Control-Allow-Origin': '*',},
          body: JSON.stringify(jsonDatos)
        });
        const result = await response.json();
        LabelResultado.current.innerHTML=result.Resultado;
        //console.log('Success:', result);
      } catch (error) {
        console.error('Error:', error);
      }
    }

    
  }

  async function resta() {

    if(valor1 === ""&&valor2 === ""){
      alert("Por favor ingrese ambos numeros")
    }else if(valor1 === ""){
      alert("Por favor ingrese el primer numero")
    }else if(valor2 === ""){
      alert("Por favor ingrese el segundo numero")
    }else{
      const jsonDatos={"Simbolo":"-","Numero1":valor1,"Numero2":valor2}
      console.log(jsonDatos)

      try {
        const response = await fetch('http://localhost:5000/resta', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' ,'Access-Control-Allow-Origin': '*',},
          body: JSON.stringify(jsonDatos)
        });
        const result = await response.json();
        LabelResultado.current.innerHTML=result.Resultado;
        //console.log('Success:', result);
      } catch (error) {
        console.error('Error:', error);
      }
    }

    
  }

  async function multiplicacion() {

    if(valor1 === "" && valor2 === ""){
      alert("Por favor ingrese ambos numeros")
    }else if(valor1 === ""){
      alert("Por favor ingrese el primer numero")
    }else if(valor2 === ""){
      alert("Por favor ingrese el segundo numero")
    }else{
      const jsonDatos={"Simbolo":"*","Numero1":valor1,"Numero2":valor2}
      console.log(jsonDatos)

      try {
        const response = await fetch('http://localhost:5000/multiplicacion', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' ,'Access-Control-Allow-Origin': '*',},
          body: JSON.stringify(jsonDatos)
        });
        const result = await response.json();
        LabelResultado.current.innerHTML=result.Resultado;
        //console.log('Success:', result);
      } catch (error) {
        console.error('Error:', error);
      }
    }

    
  }

  async function division() {

    if(valor1 === ""&&valor2 === ""){
      alert("Por favor ingrese ambos numeros")
    }else if(valor1 === ""){
      alert("Por favor ingrese el primer numero")
    }else if(valor2 === ""){
      alert("Por favor ingrese el segundo numero")
    }else{
      const jsonDatos={"Simbolo":"/","Numero1":valor1,"Numero2":valor2}
      console.log(jsonDatos)

      try {
        const response = await fetch('http://localhost:5000/division', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' ,'Access-Control-Allow-Origin': '*',},
          body: JSON.stringify(jsonDatos)
        });
        const result = await response.json();
        LabelResultado.current.innerHTML=result.Resultado;
        //console.log('Success:', result);
      } catch (error) {
        console.error('Error:', error);
      }
    }

    
  }

  const [dataLogs, setLogs] = useState(null);
  const tablaLogs=useRef(null);

  async function logs() {
    const response = await fetch('http://localhost:5000/registros');
    const data = await response.json();
    setLogs(data);
    console.log(dataLogs)
    var tableBody=""
    data.forEach(element => {
      tableBody+=`<tr>
                    <td>${element.Numero1}</td>
                    <td>${element.Numero2}</td>
                    <td>${element.Operacion}</td>
                    <td>${element.Resultado}</td>
                    <td>${element.Fecha}</td>
                  </tr>`
    });

    tablaLogs.current.innerHTML=tableBody;
  }


  return (
    
    <div className="conteiner fondo" style={{display:'flex',alignItems:'center',justifyContent:'center'}}>
      <div className="fondo2">
        <div className="App-titulo">
          <label  >Calculadora Basica SO1 202010223</label>
          <div className="App-resultado">
            <label id='LabelResultado' ref={LabelResultado}>Resultado</label>
          </div>
        </div>

        
        <header className="App-header">
        
          <div class="form-group align-items-center">
            <label style={{marginLeft:'10px'}}>Numero 1</label>
            <input id="num1" value={valor1} onChange={verificarNumero1} type="text" className="form-control" style={{width:'150px', height:'100px',fontSize:'40px'}}></input>
          </div> 
          
          <div style={{marginTop:'40px', marginBottom:'40px'}}>
            <div style={{marginLeft:'0px'}} role="group">
              <button onClick={suma} type="button" class="btn btn-outline-primary letraBoton" style={{marginLeft:'10px'}}>+</button>
              <button onClick={resta} type="button" class="btn btn-outline-success letraBoton" style={{marginLeft:'10px'}}>-</button>
            </div>
            <div style={{marginLeft:'90px', marginTop:'10px'}} role="group">
              <button onClick={multiplicacion} type="button" class="btn btn-outline-danger letraBoton" style={{marginLeft:'10px'}}>*</button>
              <button onClick={division} type="button" class="btn btn-outline-warning letraBoton" style={{marginLeft:'10px'}}>/</button>
            </div>
          </div>


          <div class="form-group align-items-center">
            <label style={{marginLeft:'10px'}}>Numero 2</label>
            <input id="num2" value={valor2} onChange={verificarNumero2} type="text" className="form-control"  style={{width:'150px', height:'100px',fontSize:'40px'}}></input>
          </div>

        </header>
      </div>
      <button onClick={logs} type="button" class="btn btn-warning letraBoton2" style={{marginLeft:'10px'}}>Logs almacenados</button>
      <div style={{width:"80%", marginTop:"30px"}}>
        <table class="table table-dark table-hover" >
          <thead>
            <tr>
              <th scope="col">Numero 1</th>
              <th scope="col">Numero 2</th>
              <th scope="col">Operación</th>
              <th scope="col">Resultado</th>
              <th scope="col">Fecha y hora</th>
            </tr>
          </thead>
          <tbody id="tregistros" ref={tablaLogs}></tbody>
        </table>
      </div>
    </div>
  );
}

export default App;
