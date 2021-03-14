import React from "react";
import Login from "./pages/Login";
import Resister from "./pages/Register";
import Home from "./pages/Home";
import Nav from "./components/Nav";
import "./App.css";
import { BrowserRouter, Route } from "react-router-dom";

function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <Nav />

        <main className="form-signin">
          <Route path="/" exact component={Home} />
          <Route path="/login" component={Login} />
          <Route path="/register" component={Resister} />
        </main>
      </BrowserRouter>
    </div>
  );
}

export default App;
