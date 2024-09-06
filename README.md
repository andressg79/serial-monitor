

Here is the rewritten README.md file in English, with a tag for test results:


**Serial Monitor (WIP)**

Tired of manually checking your serial port data? This tool makes your life easier. It displays the data flowing through your port in a clear and concise way, allowing you to focus on the logic of your project. Customize the display to your needs and quickly discover any communication anomalies.


**Designed for Microcontroller Developers**

This application complements the development of applications on microcontrollers such as Arduino, Raspberry Pi, etc. It allows you to visualize serial data in a customized way, so you can focus on the logic of your project.


**Features**

* Displays serial port data in a clear and concise way
* Customizable display to suit your needs
* Quickly discover any communication anomalies


**Test Results**

[![Test Status](https://img.shields.io/badge/test-status-passing-brightgreen)](https://github.com/andressg79/serial-monitor/actions)

**ToDo**

* [ ] Display page of data, using a string as a break indicator.
* [ ] Hot keys to pause and resume the display.
* [ ] Auto reconnect ports.
* [ ] Automatic port scanning.
* [ ] Display multiple ports at once.


**Build**

To build the application, use the command `go build -o serial-monitor`


**Examples**

### Help

```bash
serial-monitor help
```

### List Ports

```bash
serial-monitor list
```

### Monitor a Port

```bash
serial-monitor monitor -p /dev/ttyUSB0 -b 9600
```
