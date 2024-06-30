Vamos a hablar u poco de los tipos de datos en GO, basicamente los tipos de datos son casi todos tal cual identidos a los de C, osea en el sentido que obviamente tenemos los tipos de datos mas clasicos de todos, entre ellos los que comparte GO son:

- INT
- NIL
- BOOL
- STRUCT

Eso no quita que no tengamos tipo de dato flotante, de hecho tenemos pero varia ligeramente en el nombre, veamos todos los tipos de datos mas comunes que tenemos en GO:

- **Tipos Numericos**
    - **int**: Es el entero con signo comun y corriente que usualmente ocupa 64 bits = 8 bytes, despues tenemos variantes que son int8, int16, int32 e int64 donde logicamente nosotros le especificamos cuanto queremos que ocupen

    - **uint**: Este es como el int pero SIN EL SIGNO, tambien se le puede especificar el tamaño ejemplo uint8, uint16, etc.

    - **float32**, **float64**: Son numeros de punto flotante

    - **complex64**, **complex128**: Numeros complejos!


- **Tipos de Texto**: Aca la gran diferencia es que no tenemos CHAR!
    - **string**: Es una secuencia de CHARs, osea un arreglo


- **Booleanos**
    - **bool**: Logicamente este es el unico, obviamente indica true/false


- **Tipos de datos Compuestos**
    - **array**: Es un conjunto de elementos de un mismo tipo, por ejemplo de CHARs, INTs, etc.

    - **slice**: Es un array pero dinamico

    - **map**: Es una coleccion no ordenada de pares clave:valor, osea como un diccionario

    - **struct**: Es el mismo concepto que en C, es decir, una coleccion de campos con nombres que pueden ser de distinto tipo


- **Punteros**: Asi es, obviamente en GO vamos a tener al igual que en C, punteros!
    - **pointer**: Es un puntero, osea un tipo de dato que almacena la direccion de memoria de una variable


- **Tipos de datos Especiales**
    - **nil**: Es la misma funcion que en C, es decir es un valor nulo, se utiliza para indicar la ausencia de valor


- **Tipos de datos Especiales de Funcion**: Estos ya los vimos, por ejemplo func es claramente un tipo de dato!

    - **func**: Tipo de dato que representa funcion
    
    - **interface**: Tipo de dato que define un conjunto de metodos

    - **channel**: TIpo de datos que facilita la comunicacion entre goroutines (luego veremos que es) en concurrencia





Ahora vamos a ir al archivo tipos_de_datos2.go

