import React, {useState} from "react";

import {Greet} from '../../../../wailsjs/go/main/App';

export default function Searchbar() {

    const [resultText, setResultText] = useState("Please enter your search term.");
    const [search, setSearch] = useState('');
    const searchText = (e: any) => setSearch(e.target.value);
    const updateResultText = (result: string) => setResultText(result);

    function greet() {
        Greet(search).then(updateResultText);
    }

    return (
        <div className="sticky-top top-0">
            <div className="p-5">
                <div className="position-relative">
                    <span
                        className="d-flex size-20 rounded-circle ms-3 align-items-center justify-content-center position-absolute start-0 top-50 translate-middle-y">
                        <i className="bi bi-search text-black"></i>
                    </span>
                    <input type="text" className="form-control shadow py-4 form-control-lg ps-9 text-secondary" placeholder="검색어를 입력하세요." onChange={searchText} />
                    <button className="btn  btn-primary" onClick={greet}>검색</button>
                </div>
                <div className="mt-5">
                    <h3 className="text-dark-75 mb-0">{resultText}</h3>
                </div>
            </div>
        </div>
    )
}