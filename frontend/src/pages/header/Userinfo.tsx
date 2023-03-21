import React from 'react';


export default function Userinfo() {
    return (
        <>
            <a href="src/pages/header/Index#"
               className="nav-link dropdown-toggle rounded-pill p-1 lh-1 pe-1 pe-md-2 hover-bg-primary d-flex align-items-center justify-content-center"
               aria-expanded="false" data-bs-toggle="dropdown" data-bs-auto-close="outside">
                <div className="d-flex align-items-center">

                    <div className="avatar-status status-online me-sm-2 avatar xs">
                        <img src="assets/media/icons/search-icon.jpg" className="rounded-circle img-fluid"
                             alt=""/>
                    </div>
                    <span className="d-none d-md-inline-block">Yusol's father</span>
                </div>
            </a>
            <div className="dropdown-menu mt-0 p-0 dropdown-menu-end overflow-hidden">
                <div
                    className="position-relative overflow-hidden px-3 pt-4 pb-9 bg-gradient-primary text-white">
                    <svg style={{transform: "rotate(-180deg)"}} preserveAspectRatio="none"
                         className="position-absolute start-0 bottom-0 w-100 dropdown-wave"
                         fill="currentColor" height="24" viewBox="0 0 1200 120"
                         xmlns="http://www.w3.org/2000/svg">
                        <path
                            d="M0 0v46.29c47.79 22.2 103.59 32.17 158 28 70.36-5.37 136.33-33.31 206.8-37.5 73.84-4.36 147.54 16.88 218.2 35.26 69.27 18 138.3 24.88 209.4 13.08 36.15-6 69.85-17.84 104.45-29.34C989.49 25 1113-14.29 1200 52.47V0z"
                            opacity=".25"/>
                        <path
                            d="M0 0v15.81c13 21.11 27.64 41.05 47.69 56.24C99.41 111.27 165 111 224.58 91.58c31.15-10.15 60.09-26.07 89.67-39.8 40.92-19 84.73-46 130.83-49.67 36.26-2.85 70.9 9.42 98.6 31.56 31.77 25.39 62.32 62 103.63 73 40.44 10.79 81.35-6.69 119.13-24.28s75.16-39 116.92-43.05c59.73-5.85 113.28 22.88 168.9 38.84 30.2 8.66 59 6.17 87.09-7.5 22.43-10.89 48-26.93 60.65-49.24V0z"
                            opacity=".5"/>
                        <path
                            d="M0 0v5.63C149.93 59 314.09 71.32 475.83 42.57c43-7.64 84.23-20.12 127.61-26.46 59-8.63 112.48 12.24 165.56 35.4C827.93 77.22 886 95.24 951.2 90c86.53-7 172.46-45.71 248.8-84.81V0z"/>
                    </svg>
                    <div className="position-relative">
                        <h5 className="mb-1">Noah Pierre</h5>
                        <p className="text-muted small mb-0 lh-1">Full stack developer</p>
                    </div>
                </div>
                <div className="p-2">
                    <a href="profile.html" className="dropdown-item"> <i data-feather="user"
                                                                         className="fe-1x opacity-50 align-middle me-2"></i>Profile</a>
                    <a href="account-general.html" className="dropdown-item"> <i
                        data-feather="settings"
                        className="fe-1x opacity-50 align-middle me-2"></i>Settings</a>
                    <a href="page-tasks.html" className="dropdown-item"> <i data-feather="list"
                                                                            className="fe-1x opacity-50 align-middle me-2"></i>Tasks</a>
                    <hr className="mt-3 mb-1"/>
                    <a href="page-auth-signin.html"
                       className="dropdown-item d-flex align-items-center">
                        <i data-feather="log-out"
                           className="fe-1x opacity-50 align-middle me-2"></i>
                        Sign out
                    </a>
                </div>
            </div>
        </>
    )
}