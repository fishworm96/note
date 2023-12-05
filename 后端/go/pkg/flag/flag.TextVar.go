package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	fs := flag.NewFlagSet("ExampleTextVar", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)
	var ip net.IP
	fs.TextVar(&ip, "ip", net.IPv4(192, 168, 0, 100), "`IP address` to parse")
	fs.Parse([]string{"-ip", "127.0.0.1"})
	fmt.Printf("{ip: %v}\n\n", ip) // {ip: 127.0.0.1}

	// 256 is not a valid IPv4 component
	ip = nil
	fs.Parse([]string{"-ip", "256.0.0.1"})
	fmt.Printf("{ip: %v}\n\n", ip)
	// invalid value "256.0.0.1" for flag -ip: invalid IP address: 256.0.0.1
	// Usage of ExampleTextVar:
	// 	-ip IP address
	// 			IP address to parse (default 192.168.0.100)
	// {ip: <nil>}
}
