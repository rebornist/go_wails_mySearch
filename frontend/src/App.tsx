import React from 'react';
import './App.css';
import Header from './pages/header/Index';
import Searchbar from "./pages/content/searchbar/Index";

function App() {
    return (
        <div className="d-flex flex-row flex-column-fluid">
            <main className="d-flex flex-column flex-row-fluid">
                <Header />
                <Searchbar />
            </main>
        </div>
    )
}

export default App
