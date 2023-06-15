import React from 'react'
import { orderCarData } from './requestcar'

export default function OrderForm(): JSX.Element {
    const   [manufacturer, setManufacturer] = React.useState(""),
            [model, setModel]               = React.useState(""),
            [orderResult, setOrderResult]   = React.useState("")

    async function sendOrderRequest(): Promise<any> {
        try {       
            const carData: JSON | null = await orderCarData(manufacturer, model)
            if (carData != null) {
                setOrderResult(`Order for ${manufacturer} ${model} was successful!`)
            } else {
                throw new Error("JSON was not properly processed.")
            }
        }
        
        catch (error) {
            setOrderResult(`Error making Order Request: ${error}\nIf you receive another error like this when you think you shouldn't be, please contact us and let us know!`)
        }
    }

    return (
        <>
            <div id="main_block">
                <h1>CarHub</h1>
                <h2>Order Form</h2>
                <p>We don't have the car you want? Then order the car of your dreams here!</p>
                <div id="request_elem_block">
                    <label id="search_label" htmlFor="manufacturer_search">Manufacturer </label>
                    <input type="text" id="manufacturer_search" name="manufacturer" value={manufacturer} onChange={(event) => setManufacturer(event.target.value)} />
                </div>
                <div id="request_elem_block">
                    <label id="search_label" htmlFor="model_search">Model </label>
                    <input type="text" id="model_search" name="model" value={model} onChange={(event) => setModel(event.target.value)} />
                </div>
                <div id="request_elem_block">
                    <button type="button" onClick={sendOrderRequest}>Order</button>
                </div>
                {orderResult && (
                    <div id="request_elem_block">
                        <h3>{orderResult}</h3>
                    </div>
                )}
            </div>
        </>
    )
}