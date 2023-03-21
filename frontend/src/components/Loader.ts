export function LoaderOn(target: HTMLElement | null): void {
    if (target !== null) {
        // 컨텐츠 내 모든 버튼 찾기
        const buttons: HTMLCollectionOf<HTMLButtonElement> = target.getElementsByTagName('button');

        // 모든 버튼 비활성화
        for (let i = 0; i < buttons.length; i++) {
            if (!buttons[i].disabled) {
                buttons[i].disabled = true;
            }
        }

        // 로더 엘리먼트 선택
        let loaderElem: HTMLElement | null = document.getElementById('loader');

        // 로더가 보이지 않는 상태라면 로더 보이게 하고
        if (loaderElem?.classList.contains('d-none')) {
            loaderElem?.classList.remove('d-none');
        }
    }
}

export function LoaderOff(target: HTMLElement | null): void {
    if (target !== null) {
        // 컨텐츠 내 모든 버튼 찾기
        const buttons: HTMLCollectionOf<HTMLButtonElement> = target.getElementsByTagName('button');

        // 모든 버튼 활성화
        for (let i = 0; i < buttons.length; i++) {
            if (buttons[i].disabled) {
                buttons[i].removeAttribute('disabled');
            }
        }

        // 로더 엘리먼트 선택
        const loaderElem: HTMLElement | null = document.getElementById('loader');

        // 로더가 보이는 상태라면 로더 숨기고
        if (!loaderElem?.classList.contains('d-none')) {
            loaderElem?.classList.add('d-none');
        }
    }

}