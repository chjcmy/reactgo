import React from 'react';
import './App.css';
import {BrowserRouter as Router, Route, Switch} from 'react-router-dom';
import Home from "./components/home/home";
import Unit from "./components/Unit/Unit";
import Book from "./components/Book/Book";
import BookWrite from "./components/Book/BookWrite";
import SideBar from "./components/Sidebar/SideBar";

function App() {
  return (
      <Router>
        <SideBar/>
        <div className="overview">
          <Switch>
            <Route exact path="/" component={Home}/>
            <Route path="/unit/:unit" component={Unit}/>
            <Route path="/all" component={Unit} />
            <Route path="/book/:id" component={Book}/>
            <Route path="/bookwrite" component={BookWrite} />
          </Switch>
        </div>
      </Router>
  );
}

export default App;
