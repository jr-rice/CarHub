package entity

type CarRequestData struct {
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
}

type ListedCar struct {
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	Stock        int    `json:"stock"`
}

type WantedCar struct {
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
}
