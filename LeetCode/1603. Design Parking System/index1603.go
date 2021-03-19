package main

import "fmt"

func main() {
	obj := Constructor(1, 1, 0)
	param_1 := obj.AddCar(1)
	param_2 := obj.AddCar(2)
	param_3 := obj.AddCar(3)
	param_4 := obj.AddCar(1)

	fmt.Println(param_1, param_2, param_3, param_4)
}

type ParkingSystem struct {
	big    int
	medium int
	small  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		big:    big,
		medium: medium,
		small:  small,
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if this.big > 0 {
			this.big = this.big - 1
			return true
		}
		return false
	case 2:
		if this.medium > 0 {
			this.medium = this.medium - 1
			return true
		}
		return false
	case 3:
		if this.small > 0 {
			this.small = this.small - 1
			return true
		}
		return false
	default:
		return false
	}
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
