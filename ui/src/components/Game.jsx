import "../App.css"
import axios from 'axios';
import React, { useState, useEffect, useContext } from 'react';
import {WebAppContext} from '../App.jsx'

const Game = () => {
  useEffect(() => {
    getCoins()
  })

  const [coin, setCoin] = useState(0);
  const [webApp] = useContext(WebAppContext)
  function userId() {
    return webApp?.user?.id
  }
  let backUrl = process.env.back_end_url
  function getCoins() {
    axios.get(`${backUrl}/coin/?user_id=${userId()}`)
      .then(response => {
        setCoin(response.data)
      })
      .catch(res => console.log(res))
  }
  
  function click() {
    axios.put(`${backUrl}/coin?user_id=${userId()}`)
      .catch(res => console.log(res))
  }

  return (
    <div className="App">
      <header className="App-header">
      </header>
      <p>
        Coins: {coin}
        ID: {userId()}
      </p>
      <button onClick={click}>Add</button>
    </div>
  );
}

export default Game
