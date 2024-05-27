import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Game from './components/Game';
import { WebAppProvider, useInitData } from '@vkruglikov/react-telegram-web-app'
import { createContext } from 'react';

export const WebAppContext = createContext()
function App() {
  const [initDataUnsafe] = useInitData();

  if (window?.Telegram?.WebApp?.platform === "android") {
    return (
      <div >
        <h1 style={{ color: 'blue' }}>FUUU ANDROID</h1>
      </div>
    )
  }

  if (!initDataUnsafe?.user?.id) {
    return (
      <div >
        <h1 style={{ color: 'blue', borderColor: 'blue' }}>You non person</h1>
      </div>
    )
  }

  return (
    <WebAppProvider >
      <WebAppContext.Provider value={initDataUnsafe}>
        <BrowserRouter>
          <Routes>
            <Route path='/' element={<Game />}>
            </Route>
          </Routes>
        </BrowserRouter>
        {initDataUnsafe?.user?.first_name} {initDataUnsafe?.user?.last_name}
      </WebAppContext.Provider>
    </WebAppProvider>
  );
}

export default App;
