import React from "react";
import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";
import ReactQuill from "react-quill";

import icon from "../assets/images/search-icon.jpg";

type Props = {
    data: any[]
}

type ItemProps = {
    title: string,
    link: string,
    content: string
}
function SearchResultItem(p: ItemProps) {
    return (
        <div className="list-group-item p-4 list-group-item-action">
            <div className="d-flex align-items-center">

                <img src={icon} className="flex-shrink-0 rounded-3 width-60" alt=""/>
                <div className="ps-3 flex-grow-1 overflow-hidden text-secondary">
                    {(p.link) &&
                        <a href={p.link} target={"_blank"} className="text-decoration-none text-dark">
                            <span className="fw-semibold d-block mb-1 text-truncate">
                                <svg xmlns="http://www.w3.org/2000/svg"
                                  width="24" height="24"
                                  viewBox="0 0 24 24" fill="none"
                                  stroke="currentColor" stroke-width="2"
                                  stroke-linecap="round"
                                  stroke-linejoin="round"
                                  className="feather feather-external-link fe-1x me-2">
                                    <path d="M18 13v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h6"></path>
                                    <polyline points="15 3 21 3 21 9"></polyline>
                                    <line x1="10" y1="14" x2="21" y2="3"></line>
                                </svg>
                                {p.link}
                            </span>
                        </a>
                    }
                    <h5 className="mb-1">{p.title}</h5>
                    <ReactMarkdown children={p.content} remarkPlugins={[remarkGfm]}/>
                </div>

            </div>
        </div>
    )
}

export function SearchResult(p: Props) {

    const rendering = () => {
        const result = [];
        for (const d of p.data) {
            const data = JSON.parse(d);
            result.push(<SearchResultItem title={data.title} link={data.link} content={data.content}/>);
        }
        return result;
    };

    return (
        <div className="list-group text-start mb-4">
            {rendering()}
        </div>
    );

}

export function TranslateResult(p: Props) {

    const modules = {
        toolbar: false
    }

    const data = JSON.parse(p.data[0]);

    return (
        <div className="flex-row-fluid text-secondary">
            <ReactQuill readOnly className="mb-0" modules={modules} theme="snow" value={data.content}/>
        </div>
    )
}
