## Números enteros

| Tipo  | Detalles                                          |
|-------|---------------------------------------------------|
| int   | Depende del OS (32 o 64 bits)                     |
| int8  | 8 bits de -128 a 127                              |
| int16 | 16 bits de -2<sup>15</sup> a 2<sup>15</sup>-1     |
| int32 | 32 bits de -2<sup>31</sup> a 2<sup>31</sup>-1     |
| int64 | 64 bits de -2<sup>63</sup> a 2<sup>63</sup>-1     |

## Números enteros positivos
* Optimizar memoria cuando sabemos que el dato simpre va a ser positivo

| Tipo   | Detalles                            |
|--------|-------------------------------------|
| uint   | Depende del OS (32 o 64 bits)       |
| uint8  | 8 bits de 0 a 2<sup>8</sup>-1       |
| uint16 | 16 bits de 0 a 2<sup>16</sup>-1     |
| uint32 | 32 bits de 0 a 2<sup>33</sup>-1     |
| uint64 | 64 bits de 0 a 2<sup>64</sup>-1     |

## Números decimales
| Tipo    | Detalles                                                     |
|---------|--------------------------------------------------------------|
| float32 | 32 bits de +/- 1.18e<sup>-38</sup> a +/-3.4e<sup>38</sup>    |
| float64 | 64 bits de +/- 2.23e<sup>-308</sup> a +/-1.8e<sup>308</sup>  |

## Textos y boleanos
| Tipo    | Detalles           |
|---------|--------------------|
| string  | String vacío ("")  |
| bool    | true o false       |

## Números complejos
| Tipo       | Detalles                   |
|------------|----------------------------|
| Complex64  | Real e imaginario float32  |
| Complex128 | Real e imaginario float64  |

Ejemplo:
```go
c := 10 + 8i
```