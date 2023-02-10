import { useState } from 'react';
import './App.css';

function App() {


  const [valor1, setValor1] = useState('');
  const [valor2, setValor2] = useState('');

  const verificarNumero1 = event => {
    const { value } = event.target;
    if (!isNaN(value)) {  //verificamos si valor no es un numero seteamos el valor anterior
      setValor1(value);    //es decir ignoramos el valor que se ingreso (una letra)
    }
  };

  const verificarNumero2 = event => {
    const { value } = event.target;
    if (!isNaN(value)) {  //verificamos si valor no es un numero seteamos el valor anterior
      setValor2(value);    //es decir ignoramos el valor que se ingreso (una letra)
    }
  };

  return (
    
    <div className="conteiner fondo" style={{display:'flex',alignItems:'center',justifyContent:'center'}}>
      <div className="fondo2">
        <div className="App-titulo">
          <label  >Calculadora Basica SO1 202010223</label>
          <div className="App-resultado">
            <label >Resultado</label>
          </div>
        </div>

        
        <header className="App-header">
        
          <div class="form-group align-items-center">
            <label style={{marginLeft:'10px'}}>Numero 1</label>
            <input id="num1" value={valor1} onChange={verificarNumero1} type="text" className="form-control" style={{width:'150px', height:'100px',fontSize:'40px'}}></input>
          </div> 
          
          <div style={{marginTop:'40px', marginBottom:'40px'}}>
            <div style={{marginLeft:'0px'}} role="group">
              <button type="button" class="btn btn-outline-primary letraBoton" style={{marginLeft:'10px'}}>+</button>
              <button type="button" class="btn btn-outline-success letraBoton" style={{marginLeft:'10px'}}>-</button>
            </div>
            <div style={{marginLeft:'90px', marginTop:'10px'}} role="group">
              <button type="button" class="btn btn-outline-danger letraBoton" style={{marginLeft:'10px'}}>*</button>
              <button type="button" class="btn btn-outline-warning letraBoton" style={{marginLeft:'10px'}}>/</button>
            </div>
          </div>


          <div class="form-group align-items-center">
            <label style={{marginLeft:'10px'}}>Numero 2</label>
            <input id="num2" value={valor2} onChange={verificarNumero2} type="text" className="form-control"  style={{width:'150px', height:'100px',fontSize:'40px'}}></input>
          </div>

        </header>
      </div>

      <div className="App-titulo">
          <label>Aqui va a ir la tabla XDXDXD</label>
      </div>
    </div>
  );
}

export default App;
