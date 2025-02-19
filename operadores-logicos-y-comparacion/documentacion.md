
# Operadores lógicos y de comparación

Son operadores que nos permiten hacer una comparación de condiciones y en caso de cumplirse como sino ejecutarán un código determinado. Si se cumple es VERDADERO/TRUE y si no se cumple son FALSO/FALSE.

### 1. Operadores de comparación

Son aquellos que retornan TRUE o FALSE en caso de cumplirse o no una expresión. Son los siguientes:

* ***valor1*** == ***valor2***: Retorna TRUE si ***valor1*** y ***valor2*** son exactamente iguales.
* ***valor1*** != ***valor2***: Retorna TRUE si ***valor1*** es diferente de ***valor2***.
* ***valor1*** < ***valor2***: Retorna TRUE si ***valor1*** es menor que ***valor2***
* ***valor1*** > ***valor2***: Retorna TRUE si ***valor1*** es mayor que ***valor2***
* ***valor1*** >= ***valor2***: Retorna TRUE si ***valor1*** es igual o mayor que ***valor2***
* ***valor1*** <= ***valor2***: Retorna TRUE si ***valor1*** es menor o igual que ***valor2***.

### 2. Operadores lógicos
Son aquellos que retorna TRUE o FALSE si cumplen o no una condición utilizando puertas lógicas.

#### Operador AND
Este operador indica que todas las condiciones declaradas deben cumplirse para poderse marcar como TRUE. En Go, se utiliza este símbolo &&.

Ejemplo1: ```1>0 && 2 >0``` Esto retornará TRUE porque tanto la primera como la segunda condición son verdaderas.

Ejemplo2: ```2<0 && 1 > 0``` Esto retornará FALSE porque una de las condiciones no es verdadera.

#### Operador OR
Este operador indica que al menos una de las condiciones debe cumplirse para marcarse como TRUE. En Go, se representa con el símbolo ||.

Ejemplo: ``` 2<0 || 1 > 0 ``` Esto retornará TRUE porque la segunda condición se cumple, a pesar que la primera no.

#### Operador NOT
Este operador retornará el opuesto al boleano que está dentro de la variable. 

Ejemplo:
```go
myBool :=  true
fmt.Println(!myBool) // Esto retornará false
```
