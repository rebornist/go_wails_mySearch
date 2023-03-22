import React, {useState} from "react";

import {LoaderOff} from '../../../components/Loader';
import Controller from "../../../web/Controller";
import {SearchResult, TranslateResult} from "../../../components/SearchResults";

import {ChatGPT, GoogleTranslate, GoogleSearch} from '../../../../wailsjs/go/main/App';


export default function Search() {

    const target: HTMLElement | null = document.getElementById('search-inputs');
    const [resultText, setResultText] = useState<JSX.Element | null>(null);
    const [search, setSearch] = useState(''); // 검색어 입력

    const searchText = (e: any) => setSearch(e.target.value);


    function gpt() {
        new Controller(target, ChatGPT(search).then(updateResultText));
    }

    function koToEnTranslate() {
        new Controller(target, GoogleTranslate('en', search).then(updateResultText));
    }
    function enToKoTranslate() {
        new Controller(target, GoogleTranslate('ko', search).then(updateResultText));
    }

    function gSearch() {
        new Controller(target, GoogleSearch(search).then(updateResultText));
    }

    const updateResultText = (result: string) => {
        const resp =  JSON.parse(result);
        if (resp.code === 200) {
            setResultText(<SearchResult data={resp.data} />);
        } else {
            alert(resp.message);
        }
        LoaderOff(target);
    }

    return (
        <div className="sticky-top top-0">
            <div className="px-5 pt-5 pb-2">
                <div className="position-relative">
                    <div className="input-group" id="search-inputs">
                        <input id="search-text" type="text" className="form-control" placeholder="검색어를 입력하세요"
                               aria-label="Recipient's username with two button addons" onChange={searchText}/>
                        <button className="btn btn-outline-primary" type="button" onClick={koToEnTranslate}>한→영</button>
                        <button className="btn btn-outline-primary" type="button" onClick={enToKoTranslate}>영→한</button>
                        <button className="btn btn-outline-success" type="button" onClick={gpt}>ChatGPT</button>
                        <button className="btn btn-outline-danger" type="button" onClick={gSearch}>Google</button>
                    </div>
                </div>
            </div>
            (searchType) &&
            <div className="mx-5">
                {resultText}
            </div>
        </div>
    )
}