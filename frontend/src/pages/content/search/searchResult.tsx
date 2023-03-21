import React from "react";
import SearchResultItem from "./searchResultItem";

type Props = {
    searchText: string,
    searchType: string,

    resultText: string
}

export default function SearchResult(p: Props) {

    return (
        <div className="list-group text-start mb-4">
            <SearchResultItem content={p.resultText} />
        </div>
    );
}