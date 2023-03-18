import { useState, useEffect } from 'react';
import './App.css';
import { CircularProgressbar } from 'react-circular-progressbar';
import 'react-circular-progressbar/dist/styles.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faRunning, faStopCircle, faSkullCrossbones, faBed, faCalculator} from '@fortawesome/free-solid-svg-icons';


import TreeViewComponent from './Tree';




function App() {
  const [ramPorcentaje, setRam] = useState(0);
  const [cpuPorcentaje, setCpu] = useState(0);
  const [ejecutandose, setEjecutandose] = useState(0);
  const [suspendidos, setSuspendidos] = useState(0);
  const [detenidos, setDetenidos] = useState(0);
  const [zombie, setZombie] = useState(0);
  const [totales, setTotales] = useState(0);
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
    arregloPID = []
    arregloProcesos = []
    repetidos(cpuData.procesos, arregloPID);
    setProcesos(arregloProcesos);
    //console.log(cpuData.procesos);
    contadorProcesos(cpuData.procesos)

  }

  function repetidos(data, pidArray) {
    data.forEach(item => {
      const pid = item.pid;
      if (!pidArray.includes(pid)) {
        if (item.usuario) {
          pidArray.push(pid);
          arregloProcesos.push(item)
          //console.log("es padre")
        } else {
          //console.log("es hijo")
          pidArray.push(pid);
        }

      }
      if (item.hijos && item.hijos.length > 0) {
        item.nombre = item.nombre + " (PADRE)"
        repetidos(item.hijos, pidArray);
      }
    });
  }

  function contadorProcesos(data) {
    let ejecutandose = 0;
    let suspendidos = 0;
    let detenidos = 0;
    let zombie = 0;
    data.forEach(item => {
      const estado = item.estado;
      if (estado === "suspendido") {
        suspendidos += 1
      } else if (estado === "ejecutandose") {
        ejecutandose += 1
      } else if (estado === "zombie") {
        zombie += 1
      } else if (estado === "detenido") {
        detenidos += 1
      }

    });

    setDetenidos(detenidos);
    setEjecutandose(ejecutandose);
    setZombie(zombie)
    setSuspendidos(suspendidos)
    setTotales(detenidos+ejecutandose+zombie+suspendidos)
  }



  return (

    <div className="conteiner fondo" style={{ display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
      <div className="fondo2">
        <div className="App-titulo">
          <label  >Monitoreo de recursos</label>
        </div>


        <header className="App-header">

          <div className="form-group align-items-center" style={{ display: 'flex' }}>
            <label style={{ marginLeft: '10px', marginRight: '10px' }}>RAM</label>
            <div style={{ width: 300, height: 200 }}>
              <CircularProgressbar value={ramPorcentaje} text={`${ramPorcentaje}%`} />
            </div>
            <label style={{ marginLeft: '20px', marginRight: '10px' }}>CPU</label>
            <div style={{ width: 300, height: 200 }}>
              <CircularProgressbar value={cpuPorcentaje} text={`${cpuPorcentaje}%`} />
            </div>
          </div>



        </header>
      </div>
      

      <div className="fondo3">
      <div className="process-counter">
        <h2>Contadores de procesos</h2>
        <div className="process-counter-row">
          <div className="process-counter-icon"><FontAwesomeIcon icon={faRunning} /></div>
          <div className="process-counter-label">En ejecución:</div>
          <div className="process-counter-value">{ejecutandose}</div>
        </div>
        <div className="process-counter-row">
          <div className="process-counter-icon"><FontAwesomeIcon icon={faStopCircle} /></div>
          <div className="process-counter-label">Detenidos:</div>
          <div className="process-counter-value">{detenidos}</div>
        </div>
        <div className="process-counter-row">
          <div className="process-counter-icon"><FontAwesomeIcon icon={faSkullCrossbones} /></div>
          <div className="process-counter-label">Zombie:</div>
          <div className="process-counter-value">{zombie}</div>
        </div>
        <div className="process-counter-row">
          <div className="process-counter-icon"><FontAwesomeIcon icon={faBed} /></div>
          <div className="process-counter-label">Suspendidos:</div>
          <div className="process-counter-value">{suspendidos}</div>
        </div>
        <div className="process-counter-row">
        <div className="process-counter-icon"><FontAwesomeIcon icon={faCalculator} /></div>
        <div className="process-counter-label">Total:</div>
        <div className="process-counter-value">{totales}</div>
      </div>
      </div>
        <TreeViewComponent data={procesos} />
      </div>
    </div>
  );
}

export default App;
