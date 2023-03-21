import React from "react";
import ReactQuill from "react-quill";

type Props = {
    resultText: string
}

export default function TranslateResult(p: Props) {

    const modules = {
        toolbar: false
    }

    return (
        <div className="flex-row-fluid text-secondary">
            <ReactQuill className="mb-0" modules={modules} theme="snow" value={p.resultText}/>
        </div>
    )
}