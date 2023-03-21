import React from 'react';
import './App.css';
import Header from './pages/header/Index';
import Search from "./pages/content/search/Index";

function App() {
    let test = document.body.getElementsByTagName('button');
    for (let i = 0; i < test.length; i++) {
        console.log(test[i]?.className);
    }

    return (
        <div id="content" className="d-flex flex-row flex-column-fluid">
            <main className="d-flex flex-column flex-row-fluid">
                <Header/>
                <Search/>
            </main>
        </div>
    )
}

export default App
