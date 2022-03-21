package printer

import (
	"encoding/hex"
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"go.uber.org/zap"
	"log"
	"os"
)

// SetUpPrinter 打印机串口通信测试
func SetUpPrinter() {
	// Set up options.
	path := "/dev/ttyUSB0"
	if _, err := os.Stat(path); err != nil {
		options := serial.OpenOptions{
			PortName: "/dev/ttyUSB0",
			BaudRate: 115200,
			DataBits: 8,
			StopBits: 1,
			MinimumReadSize: 4,
		}

		// Open the port.
		s, err := serial.Open(options)
		if err != nil {
			log.Fatalf("serial.Open: %v", err)
		}

		// Make sure to close it later.
		defer s.Close()

		// Write 4 bytes to the port.
		//b := []byte{0x00, 0x01, 0x02, 0x03}
		//b := []byte{bytes("G28\n")}
		//b := []byte("G28\n")
		// 转换的用的 byte数据
		byte_data := []byte(`G28\n`)
		// 将 byte 装换为 16进制的字符串
		hex_string_data := hex.EncodeToString(byte_data)
		// 将 16进制的字符串 转换 byte
		hex_data, _ := hex.DecodeString(hex_string_data)
		b := hex_data
		n, err := s.Write(b)
		if err != nil {
			//log.Fatalf("port.Write: %v", err)
			msg := fmt.Sprintf("%s",err)
			zap.L().Error(msg)
		}
		fmt.Println("Wrote", n, "bytes.")

		buf := make([]byte, 128)
		n, err = s.Read(buf)
		if err != nil {
			//log.Fatal(err)
			msg := fmt.Sprintf("%s",err)
			zap.L().Error(msg)
		}
		//log.Printf("%X\n", buf[:n])
		zap.L().Error(string(buf[:n]))
	} else {
		zap.L().Error("/dev/ttyUSB0 does not exits")
	}

}