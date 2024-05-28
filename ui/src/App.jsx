import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Game from './components/Game';
import { WebAppProvider, useInitData, useShowPopup, useCloudStorage } from '@vkruglikov/react-telegram-web-app'
import { useEffect, useState } from 'react';
import { CreateWallet } from './network/request/walletService';

function App() {
  const [initDataUnsafe] = useInitData();
  const cloud = useCloudStorage()
  const [isLogin, setLogin] = useState(false)

  const showPopup = useShowPopup()
  useEffect(() => {
    if (!initDataUnsafe?.user?.id) {
      return
    }
    async function init() {
      cloud.getItem("__have_account").then(async s => {
        if (s) {
          setLogin(true)
          return
        } else {
          await createWallet()
        }
      })
    }
    init()
  })
  async function createWallet() {
    await CreateWallet(initDataUnsafe.user.id)
      .then(x => x.status === 200
        ?
        setLogin(true)
        :
        showPopup({ message: 'Error on creating account: ' + JSON.stringify(x.data) }))
      .catch(res => {
        showPopup({ message: 'Error on creating account: ' + JSON.stringify(res) })
      })
  }

  if (!isLogin) {
    return (
      <div >
        <h1 style={{ color: 'blue', borderColor: 'blue' }}>You non person</h1>
      </div>
    )
  }

  return (
    <WebAppProvider >
      {window?.Telegram?.platform === "android" ? <>Android</> : <></>}
      <BrowserRouter>
        <Routes>
          <Route path='/' element={<Game />}>
          </Route>
        </Routes>
      </BrowserRouter>
    </WebAppProvider>
  );
}

export default App;
