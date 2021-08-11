import React, {FC} from 'react';
import './App.css';
import SideBar from "./components/common/Sidebar";
import {BrowserRouter as Router, Route, Switch} from 'react-router-dom';
import Unit from "./components/common/unit/Unit";
import Home from "./components/common/home/Home";
import Book from "./components/common/Book/Book";
import BookWrite from "./components/common/Book/BookWrite";

const App: FC = () => {
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
