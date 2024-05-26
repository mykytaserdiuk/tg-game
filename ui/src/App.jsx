import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Game from './Game';

function App() {
  return (
    <BrowserRouter>
    <Routes>
      <Route path='/' element={<Game/>}>
      </Route>
    </Routes>
    </BrowserRouter>
  );
}

export default App;
