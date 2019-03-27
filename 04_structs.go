package main

import ("fmt")

const mph_over_kmh float64 = 1.60934

type car struct {
    gas_pedal uint16
    brake_pedal uint32
    steering_wheel int
    top_speed_kmph float64
}

func (c car) mph() float64 {
    result := c.top_speed_kmph * mph_over_kmh;
    c.top_speed_kmph = 0 // there must be no change
    fmt.Println("int mph value receiver method, the car instance pointer is: ", &c, " yeah," + 
                "it copies, not imposing a constraint instead, it sucks!" +
                "Haven't they tried Java? Even C++ has const shit& .")
    return result
}

func (c *car) changeSpeed(new_speed float64) {
    fmt.Println("int mph pointer receiver method, the car instance pointer is: ", &c)
    c.top_speed_kmph = new_speed
}

func mph_func(c car) float64{
    return c.mph()
}

func main() {
    var a car = car{gas_pedal:14,
                    brake_pedal: 131072,
                    steering_wheel: -19,
                    top_speed_kmph: 240}
    b := car{brake_pedal:1,
            gas_pedal: 2,
            top_speed_kmph: 3.0,
            steering_wheel: 4}
    c := car{1,2,3,4}

    fmt.Println(a.gas_pedal)
    fmt.Println(b)
    fmt.Println(c)

    fmt.Println(c.mph())
    c.changeSpeed(c.top_speed_kmph * 2)
    fmt.Println(c.mph())
    fmt.Println(mph_func(c))
}
