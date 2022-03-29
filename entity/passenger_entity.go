package entity

type Passenger struct {
	IDPassenger      string `json:"id_passenger,omitempty"`
	PassengerName    string `json:"passenger_name,omitempty"`
	SeatCode         string `json:"seat_code,omitempty"`
	Age              string `json:"age,omitempty"`
	Gender           string `json:"gender,omitempty"`
	Phone            string `json:"phone,omitempty"`
	DeparturePlace   string `json:"departure_place,omitempty"`
	DestinationPlace string `json:"destination_place,omitempty"`
}
