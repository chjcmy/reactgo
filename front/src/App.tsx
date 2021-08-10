import React, {FC} from 'react';
import './App.css';
import SideBar from "./components/common/Sidebar";
import {BrowserRouter as Router, Route, Switch} from 'react-router-dom';
import Unit from "./components/common/unit/Unit";
import Home from "./components/common/home/Home";

const App: FC = () => {
    return (
        <Router>
            <SideBar/>
            <div className="overview">
                <Switch>
                    <Route exact path="/" component={Home}/>
                    <Route path="/unit/:unit" component={Unit}/>
                    <Route path="/all" component={Unit} />
                </Switch>
            </div>
        </Router>
    );
}

export default App;
