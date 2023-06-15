import mustangGT from './assets/fordmustanggt.jpg'
import chargerhellcat from './assets/dodgechargerhellcat.jpg'
import elcamino from './assets/chevroletelcamino.jpg'
import notfound from './assets/notfound.jpg'

export async function fetchCarData(manufacturer: string, model: string, searchType: string): Promise<JSON | null> {
    const endpointURL: string = `https://carhub-api.onrender.com/data/${searchType}`
    try {
        const requestData: Response | undefined = await fetch(endpointURL, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: `{"manufacturer": "${manufacturer}", "model": "${model}"}`,
        })

        if (!requestData.ok) {
            throw new Error("Failed to send JSON Object.")
        }

        const jsonResult: JSON = await requestData.json()
        return jsonResult

    } catch (error) {
        return null
    }
}

export async function orderCarData(manufacturer: string, model: string): Promise<JSON | null> {
    const endpointURL: string = `https://carhub-api.onrender.com/data/request_wanted`
    try {
        if (manufacturer == "") {
            throw new Error("Order form was not filled correctly. Missing entry: manufacturer")
        } else if (model == "") {
            throw new Error("Order form was not filled correctly. Missing entry: model")
        }

        const requestData: Response | undefined = await fetch(endpointURL, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: `{"manufacturer": "${manufacturer}", "model": "${model}"}`,
        })

        if (!requestData.ok) {
            throw new Error("Failed to send JSON Object.")
        }
        const jsonResult: JSON = await requestData.json()
        return jsonResult

    } catch (error) {
        return null
    }
}

export function fetchListedCarImg(manufacturer: string, model: string): string {
    const listedCar = `${manufacturer} ${model}`
    switch (listedCar) {
        case "Ford Mustang GT":
            return mustangGT
        case "Dodge Charger Hellcat":
            return chargerhellcat
        case "Chevrolet El Camino":
            return elcamino
        default:
            return notfound
    }
}
