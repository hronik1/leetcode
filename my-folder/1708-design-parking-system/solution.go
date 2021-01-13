type ParkingSystem struct {
    capacities map[int]int
    usage map[int]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
    return ParkingSystem {
        capacities: map[int]int{1:big, 2:medium, 3:small},
        usage: map[int]int{1:0, 2:0, 3:0},
    }
}

func (this *ParkingSystem) AddCar(carType int) bool {
    if this.usage[carType] < this.capacities[carType] {
        this.usage[carType]++
        return true
    }
    
    return false
}


/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
