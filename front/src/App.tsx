import React from 'react';
import './App.css';
import Header from "./component/Header";
import Body from "./component/Body";
import Footer from "./component/Footer";
import SideBar from "./component/SideBar";

function App() {
    return (
        <div>
            <Header/>
            <SideBar/>
            <Body/>
            <Footer/>
        </div>
    );
}

export default App;
