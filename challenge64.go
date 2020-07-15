package main

import (
	"errors"
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// write your error code here
		e := errors.New("I'm Going to get Go!!")
		e = fmt.Errorf("I will !! %v,", f)
		return 0, sqrtError{"34.62 N", "21.4 W", e}
	}
	return 42, nil
}
