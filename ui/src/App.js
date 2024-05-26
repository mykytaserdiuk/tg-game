import logo from './logo.svg';
import './App.css';
import axios from 'axios';
import { useSearchParams } from 'react-router-dom';

const [searchParams] = useSearchParams();
const [coin, setCoin] = useState(0);

userId = searchParams.get("user_id")

function getCoins() {
  axios.get(`tg-game-production-8e6f.up.railway.app/?user_id=${userId}`)
  .then(response=> {
    setCoin(response.data)
  })
  .catch(res=>console.log(res))
}
function click() {
  axios.put(`tg-game-production-8e6f.up.railway.app/?user_id=${userId}`)
  .catch(res=>console.log(res))
}

function App() {
  return (
    <div className="App">
      <header className="App-header">
      </header>
      <p>
        Coins: {coin}
        ID: {userId}
      </p>
      <button onclick={click}>Add</button>
    </div>
  );
}

export default App;
