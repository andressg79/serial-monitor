# Serial Monitor

¿Cansado de revisar manualmente los datos de tu puerto serie? Esta herramienta te simplifica la vida. Visualiza de manera clara y concisa los datos que circulan por tu puerto, permitiéndote enfocarte en la lógica de tu proyecto. Personaliza la visualización según tus necesidades y descubre rápidamente cualquier anomalía en la comunicación.

La aplicación está pensada para complementar el desarrollo de aplicaciones en microcontroladores (Arduino, Raspberry Pi, etc.), permite visualizar datos seriales de forma personalizada. Selecciona la estrategia de visualización que mejor se adapte a tu flujo de datos: modo serial clásico, línea única con actualización por salto de línea, o salto de página con actualización automática de consola.

## Instalación

Instala la aplicación con `go install github.com/andressg79/serial-monitor`

## Ejemplos

### Ayuda

```bash
serial-monitor help
```

### Lista los puertos

```bash
serial-monitor list
```

### Monitor un puerto

```bash
serial-monitor monitor -p /dev/ttyUSB0 -b 9600
```


