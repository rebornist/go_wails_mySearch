import React from 'react';
import Userinfo from "./Userinfo";


export default function Header() {
    return (
        <header className="navbar py-0 page-header navbar-expand navbar-light px-4 border-1">
            <ul className="navbar-nav ms-auto d-flex align-items-center h-100">
                <li className="nav-item dropdown d-flex align-items-center justify-content-center flex-column h-100">
                    <Userinfo />
                </li>
                <li
                    className="nav-item dropdown ms-1 ms-lg-3 d-flex d-lg-none align-items-center justify-content-center flex-column h-100">
                    <a href="javascript:void(0)"
                       className="nav-link sidebar-trigger-lg-down hover-bg-primary size-35 p-0 d-flex align-items-center justify-content-center rounded-2">
                        <i data-feather="menu" className="fe-2x"></i>
                    </a>
                </li>
            </ul>
        </header>
    )
}