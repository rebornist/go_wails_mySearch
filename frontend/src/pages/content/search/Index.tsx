import React, {useState} from "react";

import {LoaderOff} from '../../../components/Loader';
import Controller from "../../../web/Controller";
import SearchResult from "./searchResult";

import {ChatGPT, GoogleTranslate, GoogleSearch} from '../../../../wailsjs/go/main/App';

import TranslateResult from "./translateResult";

export default function Search() {

    const target: HTMLElement | null = document.getElementById('search-inputs');
    const [resultText, setResultText] = useState('');
    const [search, setSearch] = useState(''); // 검색어 입력
    const [searchType, setSearchType] = useState(''); // [ 'gpt', '번역', '검색']
    const searchText = (e: any) => setSearch(e.target.value);

    const updateResultText = (result: string) => {
        setResultText(result);
        LoaderOff(target);
    }

    function gpt() {
        setResultText('');
        setSearchType('gpt');
        new Controller(target, resultText, searchType, ChatGPT(search).then(updateResultText));
    }

    function koToEnTranslate() {
        setResultText('');
        setSearchType('번역');
        new Controller(target, resultText, searchType, GoogleTranslate('en', search).then(updateResultText));
    }
    function enToKoTranslate() {
        setResultText('');
        setSearchType('번역');
        new Controller(target, resultText, searchType, GoogleTranslate('ko', search).then(updateResultText));
    }

    function gSearch() {
        setResultText('');
        setSearchType('검색');
        new Controller(target, resultText, searchType, GoogleSearch(search).then(updateResultText));
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
                (searchType === '번역') ?
                <TranslateResult resultText={resultText} /> :
                (searchType === 'gpt') ?
                <SearchResult searchText={search} searchType={searchType} resultText={resultText}/> :
                <SearchResult searchText={search} searchType={searchType} resultText={resultText}/>
            </div>
        </div>
    )
}