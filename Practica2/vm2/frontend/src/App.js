import { useState, useEffect } from 'react';
import './App.css';
import { CircularProgressbar } from 'react-circular-progressbar';
import 'react-circular-progressbar/dist/styles.css';

import TreeViewComponent from './Tree';




function App() {
  const [ramPorcentaje, setRam] = useState(0);
  const [cpuPorcentaje, setCpu] = useState(0);
  const [procesos, setProcesos] = useState([]);
  let arregloPID = [];
  let arregloProcesos = [];

  


  useEffect(() => {
    const interval = setInterval(() => {
      logs();
    }, 1000);

    return () => clearInterval(interval);
  }, []);

  async function logs() {
    const response = await fetch('http://localhost:5000/all');
    const data = await response.json();
  
    // console.log(data); // imprime todo el objeto
  
    const ramData = JSON.parse(data[0].ram);
    // console.log(ramData.porcentaje);
    setRam(ramData.porcentaje);
    const cpuData = JSON.parse(data[0].cpu);
    // console.log(cpuData.procesos);
    
    setCpu(cpuData.cpu);
    arregloPID=[]
    arregloProcesos=[]
    repetidos(cpuData.procesos, arregloPID);
    setProcesos(arregloProcesos);
    //console.log(arregloProcesos);

  }

  function repetidos(data, pidArray) {
    data.forEach(item => {
      const pid = item.pid;
      if (!pidArray.includes(pid)) {
        if(item.usuario){
          pidArray.push(pid);
          arregloProcesos.push(item)
          //console.log("es padre")
        }else{
          //console.log("es hijo")
          pidArray.push(pid);
        }
        
      }
      if (item.hijos && item.hijos.length > 0) {
        item.nombre=item.nombre+" (PADRE)"
        repetidos(item.hijos, pidArray);
      }
    });
  }
  


  return (
    
    <div className="conteiner fondo" style={{display:'flex',alignItems:'center',justifyContent:'center'}}>
      <div className="fondo2">
        <div className="App-titulo">
          <label  >Monitoreo de recursos</label>
        </div>

        
        <header className="App-header">
        
          <div className="form-group align-items-center" style={{ display: 'flex' }}>
            <label style={{marginLeft:'10px', marginRight:'10px'}}>RAM</label>
            <div style={{ width: 300, height: 200 }}>
              <CircularProgressbar value={ramPorcentaje} text={`${ramPorcentaje}%`} />
            </div>
            <label style={{marginLeft:'20px', marginRight:'10px'}}>CPU</label>
            <div style={{ width: 300, height: 200 }}>
              <CircularProgressbar value={cpuPorcentaje} text={`${cpuPorcentaje}%`} />
            </div>
          </div> 



        </header>
      </div>
      <div className="fondo3">
        <TreeViewComponent data={procesos}/>
      </div>
    </div>
  );
}

export default App;
