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
      {initDataUnsafe.user.id}
    </WebAppProvider>
  );
}

export default App;
