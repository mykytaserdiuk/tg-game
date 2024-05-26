
import axios from 'axios';
import { useLocation } from 'react-router-dom';
import React, { useState} from 'react';


function Game() {
    function useQuery() {
        const { search } = useLocation();
      
        return React.useMemo(() => new URLSearchParams(search), [search]);
      }
    const query = useQuery();
    const [coin] = useState(0);
    
    function userId(){
      return query.get("user_id")
    } 
    
    // function getCoins() {
    //   axios.get(`tg-game-production-8e6f.up.railway.app/?user_id=${userId()}`)
    //   .then(response=> {
    //     setCoin(response.data)
    //   })
    //   .catch(res=>console.log(res))
    // }
    function click() {
      axios.put(`https://tg-game-production-8e6f.up.railway.app/coin?user_id=${userId()}`)
      .catch(res=>console.log(res))
    }
  return (
    <div className="App">
      <header className="App-header">
      </header>
      <p>
        Coins: {coin}
        ID: {userId()11}
      </p>
      <button onClick={click}>Add</button>
    </div>
  );
}

export default Game