import React from 'react'
import { Link, Outlet } from 'react-router-dom'

export default function Layout(): JSX.Element {
    return (
        <>
            <div id="layout_block">
                <ul>
                    <li>
                        <Link to="/">Index</Link>
                    </li>
                    <li>
                        <Link to="/cars">Car Search</Link>
                    </li>
                    <li>
                        <Link to ="/order">Order Form</Link>
                    </li>
                </ul>
            </div>
            <Outlet />
        </>
    )
}