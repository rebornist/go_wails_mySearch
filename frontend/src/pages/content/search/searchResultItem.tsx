import React from "react";
import remarkGfm from "remark-gfm";
import ReactMarkdown from "react-markdown";


type Props = {
    content: string
}
export default function SearchResultItem(p: Props) {
    return (
        <div className="list-group-item p-4 list-group-item-action">
            <div className="d-flex align-items-center">
                <img src="assets/media/icons/search-icon.jpg" className="flex-shrink-0 rounded-3 width-60" alt="" />
                <div className="ps-3 flex-grow-1 overflow-hidden text-secondary">
                    <ReactMarkdown children={p.content} remarkPlugins={[remarkGfm]} />
                </div>
            </div>
        </div>
    )
}