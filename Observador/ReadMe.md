El objetivo de este patron de diseno es registrar varios observadores a un sujeto para que cuando el sujeto cambie su estado notifique a los observadores, busca desacoplar el comportamiento de la notificacion, 

 ---------------------------                                                        ------------------
 |   Sujeto                |------------------------------------------------------->|  Osbervador    |
 |   +anadir(Observador)   |                                                        |  +actualizar() |
 |   +eliminar(Observador) |                                                        ------------------
 |   #notificar()          |  \    ---------------------------------                        ^
 ---------------------------   \   |  for all o in_observadores{   |                        |
            ^                   \  |       o.actualizar();         |                        |
            |                      |  }                            |                        |
            |                      ---------------------------------                        |
 -----------------------        _sujeto                                 -----------------------------
 |  SujetoConcreto     |   <--------------------------------------------|   ObservadorConcreto      |
 |  -_estadoSujeto     |                                                |   -_estadoObservador      |
 |  +getEstado()       |                                                |   +actualizar()           |
 |  +setEstado()       |\                                               -----------------------------
 -----------------------  \   
                            \ --------------------------                 --------------------------
                              |   return_estadoSujeto; |                 |   estadoObservador     |
                              --------------------------                 |   _sujeto.             |
                                                                         --------------------------
Elementos:
Sujeto:                 Consta de una interface la cual senala las funciones necesarias que debe tener un objeto clase o estructura es decir el sujeto concreto, para poder ser observada, estas funciones son de base addObserver para dar la capacidad de anadir observadores, un DelObserver para eliminar observadores, y para el ejemplo se nade la funcion de notify para notificar a los observadores de algun evento.
observador:              Consta de una interface que senala que funciones necesitan de tener todos los observadores para poder entrar dentro del patron, el caso mas comun  y para el ejemplo es la funcion notificar,
Sujeto concreto:        Son como tal las estructuras las cuales se desean ser observadas, e implementan la interface sujeto, y usan estos metodos paar notificar a los observadores, los sujetos deben de tener entre sus propiedades, elementos, o campos un slay o array tipo observador donde se guardara la lista de observadores que se vayan anadiendo
Observador concreto:    Son aquellos que estaran recibiendo la informacion o el proceso por el cual se enviara la notificacion, para el caso, se usa slack e email como ejemplos concretos, cada uno cuenta con sus propios metodos y singularidades pero deberan tener implementados los metodos de la interface para el ejemplo usan 

dentro del main es un ejemplo de la implementacion sin embargo se puede crear alguna logica o codigo a seguir que permite generar observadores bajo ciertas condiciones


                                                            