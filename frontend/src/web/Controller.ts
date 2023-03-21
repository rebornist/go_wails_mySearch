import {LoaderOn} from "../components/Loader";

// 기본 컨트롤러 클래스
export default class Controller {
    constructor(target: HTMLElement | null, resultText: string, searchType: string, func: any) {
        LoaderOn(target);
        func(resultText, searchType);
    }

}