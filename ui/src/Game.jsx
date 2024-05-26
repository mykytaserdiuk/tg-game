
import axios from 'axios';
import { useSearchParams } from 'react-router-dom';
import React, { useState, useEffect } from 'react';

const [searchParams] = useSearchParams();
const [coin, setCoin] = useState(0);

function userId(){
  return searchParams.get("user_id")
} 

function getCoins() {
  axios.get(`tg-game-production-8e6f.up.railway.app/?user_id=${userId()}`)
  .then(response=> {
    setCoin(response.data)
  })
  .catch(res=>console.log(res))
}
function click() {
  axios.put(`tg-game-production-8e6f.up.railway.app/?user_id=${userId()}`)
  .catch(res=>console.log(res))
}

function Game() {
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

export default Game