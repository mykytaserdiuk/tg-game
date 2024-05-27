
import axios from 'axios';
import { useLocation } from 'react-router-dom';
import React, { useState, useEffect } from 'react';

const Game = () => {
  function useQuery() {
    const { search } = useLocation();
    return React.useMemo(() => new URLSearchParams(search), [search]);
  }

  useEffect(c => {
    getCoins()
  }, [])

  const query = useQuery();
  const [coin, setCoin] = useState(0);
  function userId() {
    return query.get("user_id")
  }

  function getCoins() {
    axios.get(`tg-game-production-8e6f.up.railway.app/coin/?user_id=${userId()}`)
      .then(response => {
        setCoin(response.data)
      })
      .catch(res => console.log(res))
  }
  
  function click() {
    axios.put(`https://tg-game-production-8e6f.up.railway.app/coin?user_id=${userId()}`)
      .catch(res => console.log(res))
  }

  return (
    <div className="App">
      <header className="App-header">
      </header>
      <p>
        Coins: {coin}
        ID: {userId()}
        {this.context}
      </p>
      <button onClick={click}>Add</button>
    </div>
  );
}

export default Game
