import React from 'react'
import ReactDOM from 'react-dom/client'
import { BrowserRouter, Routes, Route } from "react-router-dom"

import './index.css'
import Layout from './layout'
import Index from './index'
import CarsList from './carslist'
import OrderForm from './orderform'
import PageNotFound from './pagenotfound'

function Main(): JSX.Element {
  const Homepage: JSX.Element = (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<Index />} />
          <Route path="cars" element={<CarsList />} />
          <Route path="order" element={<OrderForm />} />
          <Route path="*" element={<PageNotFound />} />
        </Route>
      </Routes>
    </BrowserRouter>
  )
  return Homepage
}

const main: ReactDOM.Root = ReactDOM.createRoot(document.getElementById("root")!)
main.render(<Main />)