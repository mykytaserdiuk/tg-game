import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Game from './components/Game';
import {WebAppProvider, useInitData} from '@vkruglikov/react-telegram-web-app'

function App() {

  const [initDataUnsafe] = useInitData();  

  return (
    <WebAppProvider value={initDataUnsafe}>
      <BrowserRouter>
        <Routes>
          <Route path='/' element={<Game />}>
          </Route>
        </Routes>
      </BrowserRouter>
      {window?.Telegram?.WebApp?.platform}
      {initDataUnsafe.user.first_name} {initDataUnsafe.user.last_name}
    </WebAppProvider>
  );
}

export default App;
