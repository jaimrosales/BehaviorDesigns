Este patron de diseno es muy similar al observador

mientras el patron de diseno observador permite notificar a todos los observadores cuando un cambio se a realizado
La cadena de responsabilidad hace algo similar sin embargo lo hace de manera serial, dando a cada elemento dentro de la seria la responsabilidad de evaluar si debe de o no se pasar la notificacion
                Ejemplo 
            Solicitud de un prestamo
Asesor---->Coordinador----->Gerente---->Directivo       responsable
2000        5000            10000       200000          Prestamo deseado

utilizando este patron te puedes llegar a evitar las cadenas de elseif