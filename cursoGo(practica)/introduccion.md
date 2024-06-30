# CURSO LENGUAJE GO

## INTRODUCCION

Go se lo conoce como un lenguaje de programacion donde su especialidad es la **CONCURRENCIA**, que es compilado en la sintaxis de C. 

Es un lenguaje desarrollado por Google y sus diseñadores son:

- Robert Griesemer
- Rob Pike
- Ken Thompson

GO tiene las caracteristicas de ser:

- **LENGUAJE COMPILADO**: Es aquel lenguaje cuyo codigo fuente se traduce completamente a codigo maquina (archivo objeto) antes de ser ejecutado. Esta traduccion se realiza gracias al compilador, por lo que en ese sentido es igual al lenguaje C. Obviamente tambien C es un lenguaje compilado, C++, GO, Rust, etc. Y obviamente Python NO ES UN LENGUAJE COMPILADO claramente

- **CONCURRENTE**: GO es un lenguaje de programacion tambien CONCURRENTE, es decir, se puede escribir codigo en el que se ejecuten multiples hilos de manera simultanea en lugar de esperar a que una tarea finalice para arrancar otra. Como dijimos, esto se logra mediante el uso de hilos **(THREADS)**. Otros lenguajes que tambien tienen concurrencia, ademas de GO, son; Java, Python! y C#

- **IMPERATIVO**: GO, es imperativo, esto es, se escribe como una serie de instrucciones que indican como realizar una tarea especifica. Estas instrucciones se centran en como se deben llevar a cabo las operaciones y manipulacion de datos en lugar de enfocarse en que operaciones se deben realizar, por ejemplo C es un lenguaje imperativo!. La antitesis del lenguaje imperativo es el lenguaje **DECLARATIVO** en donde este se centra en el que se debe hacer en lugar de como hacerlo, osea el lenguaje declarativo se centra en la logica dejando que el sistema determine la mejor manera de llevar a cabo dichas instrucciones

- **ESTRUCTURADO**: GO se basa en la programacion estructurada, osea un lenguaje que se centra en el uso de estructuras de control como secuencias, selecciones (if), iteraciones(for, while) para controlar el flujo del programa. Tambien tiene funciones que se utilizan para organizar el codigo en bloques mas grande y reutilizables, lo que promueve justamente una estructuracion del programa. Por ejemplo C y Python claramente son estructurados

- **POO**: GO tambien tiene la antites de la programacion Estructurada que es la POO (Programacion Orientada a Objetos), osea creacion de objetos que representarian entidades del mundo real, estos objetos tienen propiedades (atributos) y comportamientos (metodos) asociados. Un claro ejemplo del lenguaje fundador del POO es Smalltalk!, POO se basa en la encapsulacion, herencia, polimorfismo para organizar los programas. Tambien Python tiene POO, y C no!, en realidad si se le puede implementar pero de manera artificial, es decir, la POO no es natural en el lenguaje C, o por lo menos, no se la penso para ello, pero a la fuerza se puede hacer que C piense como POO usando cosas que comunmente no fueron pensadas para POO y forzarlo a que funcionen para POO 

---

Lo que tiene mucho GO es que mezcla lo bueno del lenguaje Python y C, es decir, toma la gran potencia que tiene el lenguaje C, que compilan a codigo maquina y la gran versatilidad de Python

Ademas otra caracteristica que tiene GO y que no mencionamos es que tiene la capacidad de hacer **DUCK TYPING**, es decir, tiene la capacidad de determinar el tipo de dato o comportamiento de un objeto basandose en sus metodos y atributos en lugar de su tipo especifico, es decir si por ejemplo nosotros hacemos *x=10* es bastante logico que la variable 'x' es un INT sin necesidad de especificar el tipo de dato

**NOTA**: En las primeras versiones de GO se usaba el mismo compilador que el lenguaje C, pero ya en versiones mas recientes GO tiene su propio compilador

---

Bueno, ya dejamos la introduccion, ahora vamos directamente al lenguaje GO, para ello vamos al archivo **primer_programa.go** donde a partir de ahora comentaremos ahi...

