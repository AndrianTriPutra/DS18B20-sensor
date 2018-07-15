# go-pi_DS18B20
in this case i try to read tempeture use DS18B20 with pi. I learn from adafruit tutorial and i try to translate or make to code golang. thanks

raspi and DS18B20 with golang

running
1. sudo nano /boot/config.txt 
tambahkan "dtoverlay=w1-gpio,gpiopin=2" di line terakhir
sesuaikan gpio yang ingin km gunakan
2.reboot your raspi
3.sudo modprobe w1-gpio
4.sudo modprobe w1-therm
5.cd /sys/bus/w1/devices
6.ls
7.cd 28-xxxx (sesuaikan dengan yang tebaca di raspi kamu)
8.cat w1_slave

9.copi paste file go DS18B20.go
10.sesuaikan address 28-xx nya, lalu coba run

referensi:
https://cdn-learn.adafruit.com/downloads/pdf/adafruits-raspberry-pi-lesson-11-ds18b20-temperature-sensing.pdf

