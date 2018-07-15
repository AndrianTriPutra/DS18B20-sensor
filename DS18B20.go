/*
	DS18B20 and Raspi with go
	* from ATP to Maker
*/
package main

import (
	"fmt"

	"strconv"	
	"strings"
	"os/exec"

)

func main (){
		gpio, _ := exec.Command("sh", "-c", "sudo modprobe w1-gpio").Output()
		fmt.Printf("%s",gpio)
		
		therm, _ := exec.Command("sh", "-c", "sudo modprobe w1-therm").Output()
		fmt.Printf("%s",therm)
		
		for {
			out, _ := exec.Command("sh", "-c", "date +\"%Y-%m-%d %H:%M:%S %Z\"").Output()
			buffer := string(out)
			buffer=strings.TrimSuffix(buffer,"\n")
			Clock:=buffer
			fmt.Printf("Clock:%s\t",Clock)
			
			ReadDS18B20, _ := exec.Command("sh", "-c", "cat /sys/bus/w1/devices/28-001413f3f2ff/w1_slave ").Output()
			//fmt.Printf("ReadDS18B20:%s\t",ReadDS18B20)
			bufferReadDS18B20 := string(ReadDS18B20)
			buffertrimDS18B20 := strings.TrimSuffix(bufferReadDS18B20,"\n")
				
			index := strings.LastIndex(buffertrimDS18B20,"t")
			trimDS18B20 := buffertrimDS18B20[index+2:len(buffertrimDS18B20)]
			//fmt.Printf("trimDS18B20:%s",trimDS18B20)
			//fmt.Printf("<=\t")
			
			stringtofloat,_ := strconv.ParseFloat(trimDS18B20,32)
			//fmt.Printf("stringtofloat:%v\t",stringtofloat)
			
			tempcelcius := stringtofloat/1000.00
			fmt.Printf("tempcelcius:%v\n",tempcelcius)
		}	
}
