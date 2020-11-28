package main

import (
	"flag"
	"github.com/xa4a/go-roomba"
	"log"
	"time"
)

const (
	defaultPort = "/dev/cu.usbserial-FTTL3AW0"
)

var (
	portName = flag.String("port", defaultPort, "roomba's serial port name")
)

func main() {
	flag.Parse()
	r, err := roomba.MakeRoomba(*portName)
	if err != nil {
		log.Fatal("Making roomba failed")
	}

	_ = r.Start()
	_ = r.Safe()
	_ = r.MotorControl(true, true, true, true, roomba.DefaultDirection)
	err = r.MotorPWM(0, 10, 0)
	_ = r.LEDs(false, true, false, true, 0, 0)
	timer := time.Tick(2 * time.Second)
	<- timer
	_ = r.MotorControl(false, false, false, false, roomba.DefaultDirection)
	_ = r.Stop()
	_ = r.Reset()
}