package main

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	fs := flag.NewFlagSet("ExampleFunc", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)
	var ip net.IP
	fs.Func("ip", "`IP address` to parse", func(s string) error {
		ip = net.ParseIP(s)
		if ip == nil {
			return errors.New("could not parse IP")
		}
		return nil
	})
	fs.Parse([]string{"-ip", "127.0.0.1"})
	fmt.Printf("{ip: %v, loopback: %t}\n\n", ip, ip.IsLoopback()) // {ip: 127.0.0.1, loopback: true}

	// 256 is not a valid IPv4 component
	fs.Parse([]string{"-ip", "256.0.0.1"})
	fmt.Printf("{ip: %v, loopback: %t}\n\n", ip, ip.IsLoopback())
	// invalid value "256.0.0.1" for flag -ip: could not parse IP
	// Usage of ExampleFunc:
	// 	-ip IP address
	// 			IP address to parse
	// {ip: <nil>, loopback: false}
}
