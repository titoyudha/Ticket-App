package service

// func CreateOrder(orderDto dto.CreateOrder) (entity.Order, error) {
// 	var order entity.Order
// 	orderDto.DepartureDate = time.Now().Format(time.RFC3339) // format tanggal
// 	orderDto.CheckIn = time                                  // 1 jam sebelum departure date
// 	orderDto.DepartureDate = time                            //format waktu jam

// 	order.ID = uuid.New().String()
// 	order.OrderCode = orderDto.OrderCode
// 	order.DeparturePlace = orderDto.DeparturePlace
// 	order.SeatCode = orderDto.SeatCode
// 	order.Destination = orderDto.Destination
// 	order.DepartureDate = orderDto.DepartureDate

// 	// Relasi tabel order + passenger
// 	err := config.ConnectDB().Association(&order).Find(&entity.Passenger).Error()

// }

// func UpdateOrder(updateOrderDto dto.UpdateOrder) (entity.Order, error) {}

// func Invoice() {}
