# Serial Monitor (WIP)

Tired of manually checking your serial port data? This tool makes your life easier. It displays the data flowing through your port in a clear and concise way, allowing you to focus on the logic of your project. Customize the display to your needs and quickly discover any communication anomalies.

The application is designed to complement the development of applications on microcontrollers (Arduino, Raspberry Pi, etc.), allowing you to display serial data in a customized way. Select the display strategy that best suits your data flow: classic serial mode, single line with line break update, or page break with automatic console update.

## ToDo

- [ ] Display page of data, using a string as a break indicator.
- [ ] Hot keys to pause and resume the display.
- [ ] Auto reconnect ports.
- [ ] Automatic port scanning.
- [ ] Display multiple ports at once.

## Compilation

To compile the application, use the command `go build -o serial-monitor`

## Examples

### Help

```bash
serial-monitor help
```

### List ports

```bash
serial-monitor list
```

### Monitor a port

```bash
serial-monitor monitor -p /dev/ttyUSB0 -b 9600
```