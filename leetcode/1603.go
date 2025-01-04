type ParkingSystem struct {
	data map[int]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	data := map[int]int{
		1: big,
		2: medium,
		3: small,
	}

	return ParkingSystem{
		data: data,
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.data[carType] == 0 {
		return false
	}
	this.data[carType]--
	return true
}

/**
* Your ParkingSystem object will be instantiated and called as such:
* obj := Constructor(big, medium, small);
* param_1 := obj.AddCar(carType);
 */