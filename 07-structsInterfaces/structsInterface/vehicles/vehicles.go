package vehicles

import "fmt"

type Vehicle interface {
	Distance() float64
}

// "CAR", "BIKE", "MOTORCYCLE", "BUS", "TRUCK"
const (
	CarVehicle        = "CAR"
	BikeVehicle       = "BIKE"
	MotorcycleVehicle = "MOTORCYCLE"
	BusVehicle        = "BUS"
	TruckVehicle      = "TRUCK"
)

func New(v string, time int) (Vehicle, error) {
	switch v {
	case CarVehicle:
		return &Car{Time: time}, nil
	case BikeVehicle:
		return &Bike{Time: time}, nil
	case MotorcycleVehicle:
		return &Motorcycle{Time: time}, nil
	case BusVehicle:
		return &Bus{Time: time}, nil
	case TruckVehicle:
		return &Truck{Time: time}, nil
	default:
		return nil, fmt.Errorf("unknown vehicle type: %s", v)
	}
}

type Car struct {
	Time int
}

func (c *Car) Distance() float64 {
	return 120.0 * (float64(c.Time) / 60.0)
}

type Bike struct {
	Time int
}

func (b *Bike) Distance() float64 {
	return 30.0 * (float64(b.Time) / 60.0)
}

type Motorcycle struct {
	Time int
}

func (m *Motorcycle) Distance() float64 {
	return 90.0 * (float64(m.Time) / 60.0)
}

type Bus struct {
	Time int
}

func (b *Bus) Distance() float64 {
	return 80.0 * (float64(b.Time) / 60.0)
}

type Truck struct {
	Time int
}

func (t *Truck) Distance() float64 {
	return 70.0 * (float64(t.Time) / 60.0)
}
