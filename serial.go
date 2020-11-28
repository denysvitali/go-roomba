// Provides low-level functions for interacting with Roomba port/socket/buffer.

package roomba

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"log"

	"github.com/tarm/goserial"
)

// Packs the given data as big endian bytes.
func Pack(data []interface{}) []byte {
	buf := new(bytes.Buffer)
	for _, v := range data {
		err := binary.Write(buf, binary.BigEndian, v)
		if err != nil {
			log.Fatal("failed packing bytes:", err)
		}
	}
	return buf.Bytes()
}

// Configures and opens the given serial port.
func (r *Roomba) Open(baud uint) error {
	if baud != 115200 && baud != 19200 {
		return errors.New(fmt.Sprintf("invalid baud rate: %d. Must be one of 115200, 19200", baud))
	}

	c := &serial.Config{Name: r.PortName, Baud: int(baud)}
	port, err := serial.OpenPort(c)

	if err != nil {
		log.Printf("failed to open serial port: %s", r.PortName)
		return err
	}
	r.S = port
	log.Printf("opened serial port: %s", r.PortName)
	return nil
}

// Writes the given opcode byte and a sequence of data bytes to the serial port.
func (r *Roomba) Write(opcode byte, p []byte) error {
	log.Printf("Writing opcode: %v, data %v", opcode, p)
	n, err := r.S.Write([]byte{opcode})
	if n != 1 || err != nil {
		return fmt.Errorf("failed writing opcode %d to serial interface",
			opcode)
	}
	n, err = r.S.Write(p)
	if n != len(p) || err != nil {
		return fmt.Errorf("failed writing command to serial interface: % d", p)
	}
	return nil
}

// Writes a single byte to the serial port.
func (r *Roomba) WriteByte(opcode byte) error {
	return r.Write(opcode, []byte{})
}

// Reads bytes from the serial port.
func (r *Roomba) Read(p []byte) (n int, err error) {
	return r.S.Read(p)
}
