package internal

type RouteCreatedEvent struct {
	EventName  string       `json:"event_name"`
	RouteId    string       `json:"id"`
	Distance   int          `json:"distance"`
	Directions []Directions `json:"directions"`
}

type FreightCalculatedEvent struct {
	EventName string  `json:"event"`
	RouteId   string  `json:"route_id"`
	Amount    float64 `json:"amount"`
}

func NewFreightCalculatedEvent(routeId string, amount float64) *FreightCalculatedEvent {
	return &FreightCalculatedEvent{
		EventName: "freightCalculated",
		RouteId:   routeId,
		Amount:    amount,
	}
}

func RouteCreatedHandler(event RouteCreatedEvent, routeService *RouteService) (*FreightCalculatedEvent, error) {
	route := NewRoute(event.RouteId, event.Distance, event.Directions)
	routeCreated, err := routeService.CreateRoute(route)
	if err != nil {
		return nil, err
	}
	freightCalculatedEvent := NewFreightCalculatedEvent(routeCreated.ID, routeCreated.FreightPrice)

	return freightCalculatedEvent, nil
}
