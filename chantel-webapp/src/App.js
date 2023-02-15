import logo from './logo.svg';
import './App.css';
import { useState, useEffect } from 'react'
import axios from 'axios'

function App() {
  const [date, setDate] = useState("");
  const [error, setError] = useState(null);

  useEffect(() => {
    axios("http://api.chantel.sandbox/date")
        .then((response) => {
            console.log(response.data)
          setDate(response.data.message);
          setError(null);
        })
        .catch(setError);
  }, []);

  if (error) {
      return (
          <div>
              <p>An error occurred</p>
              <div><pre>{JSON.stringify(error, null, 2)}</pre></div>
          </div>
      )

  }

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
            Current date: {date}
        </p>
      </header>
    </div>
  );
}

export default App;
