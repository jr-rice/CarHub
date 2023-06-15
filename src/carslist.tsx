import React from 'react'
import { fetchCarData, fetchListedCarImg } from './requestcar'

export default function CarsList(): JSX.Element {
    const   [manufacturer, setManufacturer] = React.useState(""),
            [model, setModel]               = React.useState(""),
            [error, setError]               = React.useState(""),
            [listedSearch, setListedSearch] = React.useState(false),
            [wantedSearch, setWantedSearch] = React.useState(false),
            [carListings, setCarListings]   = React.useState<JSON | null>(null)

    const setListed = (): void => {
        setListedSearch(true)
        setWantedSearch(false)
    }

    const setWanted = (): void => {
        setListedSearch(false)
        setWantedSearch(true)
    }
    
    async function sendCarRequest(): Promise<any> {
        let searchType: string

        try {
            if (listedSearch == true) {
                searchType = "search_listed"
            } else if (wantedSearch == true) {
                searchType = "search_wanted"
            } else {
                throw new Error("No boolean for search specified.")
            }

            const carData: JSON | null = await fetchCarData(manufacturer, model, searchType)
            setCarListings(carData)
        }

        catch (error) {
            const errorMsg: string = `Error in Search Request: ${error}\nIf you receive another error like this when you think you shouldn't be, please contact us and let us know!`
            setError(errorMsg)
        }
    }

    return (
        <>
            <div id="main_block">
                <h1>CarHub</h1>
                <h2>Car Search</h2>
                <p>Take a look to see if we have the perfect car for you!</p>
                <div id="search_elem_block">
                    <label id="search_label" htmlFor="manufacturer_search">Manufacturer </label>
                    <input type="text" id="manufacturer_search" name="manufacturer" value={manufacturer} onChange={(event) => setManufacturer(event.target.value)} />
                </div>
                <div id="search_elem_block">
                    <label id="search_label" htmlFor="model_search">Model </label>
                    <input type="text" id="model_search" name="model" value={model} onChange={(event) => setModel(event.target.value)} />
                </div>
                <div id="search_elem_block">
                    <label htmlFor="toggle_listed">Listed Search</label>
                    <input type="radio" id="toggle_listed" name="toggle" value="listed" onChange={setListed} />
                </div>
                <div id="search_elem_block">
                    <label htmlFor="toggle_wanted">Wanted Search</label>
                    <input type="radio" id="toggle_wanted" name="toggle" value="wanted" onChange={setWanted} />
                </div>
                <div id="search_elem_block">
                    <button type="button" onClick={sendCarRequest}>Search</button>
                </div>
                {carListings && Array.isArray(carListings) && (
                    <div id="search_elem_block">
                        <div id="car_listing">
                            {carListings.map((car: any, index: number) => (
                                <div key={index}>
                                    <h3>{car.manufacturer} {car.model}</h3>
                                    <img src={fetchListedCarImg(car.manufacturer, car.model)} alt="car_image" />
                                    {car.stock && <p>Stock: {car.stock}</p>}
                                </div>
                            ))}
                        </div>
                    </div>
                )}
                {error && (
                    <div id="search_elem_block">
                        <h3>{error}</h3>
                    </div>
                )}
            </div>
        </>
    )
}