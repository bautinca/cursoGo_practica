# DIFERENCIA ENTRE CONCURRENCIA Y PARALELISMO

Lo que hace fuerte al lenguaje GO es justamente que es ideal para esto ultimo que vimos, osea para la CONCURRENCIA
Nosotros ya lo vimos pero la diferencia es:

- **CONCURRENCIA**: Sucede cuando en UN MISMO PROCESADOR da la sensacion de que ejecuta 2 hilos simultaneamnete cuando en realidad lo que hace el procesador es ir saltando entre hilos para ejecutar cachitos de cada hilo, es decir, lo que siempre sucede en la concurrencia es que el 'camino' de ejecucion siempre es distinto pero el resultado final SIEMPRE ES EL MISMO

- **PARALELISMO**: Muchas veces se dice que los hilos se ejecutan en paralelo en la concurrencia, pero en realidad no es asi como aclaramos anteriormente, cuando tenemos paralelismo es porque tenemos mas de un procesador en el sistema entonces cada procesador agarra instrucciones por lo tanto las instrucciones se van ejecutando paralelamente ya que tenemos mas de 1 procesador para llevarlo a cabo. Es decir, maquinas de 1 solo procesador es imposible que tengan paralelismo, si pueden tener concurrencia

La **CONCURRENCIA** nos vende la ilusion de que esta haciendo varias cosas simultaneamente cuando en realidad esta haciendo una cosa, ya que al tener 1 solo procesador el procesadro va leyendo instruccion por instruccion, pero lo que hace es que va leyendo las instrucciones saltando de un hilo a otro hilo, es por ello que en el archivo *gorutinas.go* colocamos un *sleep* para que se pueda visualizar bien al ejecutar el programa como se va imprimiendo linea a linea de un hilo y del otro hilo, pero siempre se imprime linea por linea



## DIFERENCIA ENTRE HILOS Y GORUTINAS

¿Como? Los hilos y las gorutinas son distintos? Son conceptos casi identicos, pero difieren en minimas cosas. En cualquier otro lenguaje que maneje concurrencia de manera natural o artificial vamos a notar que un mismo proceso (programa) puede lanzar varios hilos de ejecucion, entonces cuando se lanzaba un hilo era para ejecutar una sola cosa, y asi..., es decir, es como que podiamos hacer una cosa en cada hilo

Ahora la novedad que introduce GO, y que por eso es tan potente, es que EN UN HILO SE PUEDEN EJECUTAR VARIAS GORUTINAS, es decir, las gorutinas se administran a nivel GO (el scheduler que tiene GO se encarga de las gorutinas en cada hilo), en cambio para los hilos tradicionales como comunmente lo conocemos ahi si se encarga el SO (Sistema Operativo).

Por tanto, en termino de recursos, las gorutinas son mas ligeros que los hilos comunes y corrientes que genera cada proceso, es por ello que se dice que las gorutinas son **HILOS LIGEROS**

Por otro lado, las gorutinas en GO estan diseñados para trabajar con **CANALES** (luego veremos mas adelante). Las gorutinas se comunican y sincronizan mediante el uso de canales o **CHANNELS** lo que simplifica la comunicacion entre las gorutinas

**NOTA**: Para que se entienda bien, el SO tiene el Scheduler o planificador (Esto lo vimos en la materia Sistemas Operativos), donde el Scheduler lo que hace es manejar o administrar la ejecucion de distintos procesos o tasks que representa cada programa, entonces el Scheduler es el que vende la ilusion de que se estan ejecutando programas simultaneamente cuando en realidad sabemos que no es asi. Luego cada programa o proceso tiene HILOS, donde estos hilos son los que conocemos, los hilos se desprenden de los procesos para que haya una concurrencia en la ejecucion del programa, luego la novedad que introduce GO es que ahora CADA HILO TIENE GORUTINAS. Y tambien lo importante es que el SO manejaba la creacion de procesos nuevos e hilos comunes y corrientes, pero ahora la creacion de GORUTINAS esta bajo el mando del propio Scheduler de GO y no del SO. Ademas, si nos ponemos a pensar el Scheduler de un SO tambien aplica concurrencia, porque lo que hace es vender la ilusion de que esta ejecutando 2 programas o procesos simultaneamente



## DIFERENCIA ENTRE HILO Y PROCESO

Esto ya lo explicamos, pero ahora lo diremos aca para entenderlo bien:

- **PROCESO**: Es una instancia de ejecucion de un programa, uno usualmente dice que un programa tiene un proceso, pero en realidad un mismo programa puede emitir varios procesos, donde cada proceso tiene su respectivo PID, espacio de memoria, descriptores, etc. (Esto lo vimos en la materia SO). Recordemos que el SO crea procesos nuevos con la instruccion fork() en Linux. Ademas recordemos que los procesos se comunican entre si mediante el uso de PIPES, lo que suele ser muy costoso que se comuniquen procesos en terminos de rendimiento porque enviamos syscalls para que el Kernel habilite esta comunicacion, por lo tanto el SO esta constantemente activo ya que es el que habilita esta comunicacion, osea el Kernel del SO

- **HILO**: Es una entidad DENTRO DE UN PROCESO, en el que ejecuta instrucciones de manera independiente. Los hilos comparten el espacio de memoria que tiene el proceso al que pertenecen y tambien los recursos que el proceso tiene.
Lo bueno de los hilos en comparacion a los procesos es que son mas eficientes en terminos de rendimiento para que se comuniquen entre ellos, ya que justamente comparten el mismo espacio de memoria que el proceso en el que pertenencen, por ejemplo si en el proceso declaramos una variable 'x' y luego creamos 2 hilos entonces estos 2 hilos van a tener directo acceso a esa variable 'x'




Ahora vamos al archivo gorutinas3.go