// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main

import (
    "fmt"
    "encoding/binary"
)
func IsLittleEndian() bool {
	// https://stackoverflow.com/a/51332762
	// Not the best answer, But may help you in your research.
	return true
}


func Prefix_Packet_Header(bufferLength int) []byte {
	var outputPacket []byte = make([]byte, 4)
	if(IsLittleEndian() == true) {
		binary.LittleEndian.PutUint32(outputPacket, uint32(bufferLength))
	} else {
		binary.BigEndian.PutUint32(outputPacket, uint32(bufferLength))
	}
	return outputPacket
}

func Read_PacketLenght(packetBuffer []byte) int {
	if(IsLittleEndian() == true) {
		return int(binary.LittleEndian.Uint32(packetBuffer))
	} else {
		return int(binary.BigEndian.Uint32(packetBuffer))
	}

}

func main() {

	var packetLengthBuffer = Prefix_Packet_Header(128)
	var packetLengthInt = Read_PacketLenght(packetLengthBuffer)
	
    	fmt.Println(packetLengthInt)
}

