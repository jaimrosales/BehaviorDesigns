

public class main {
    public static void main(String[] args) {
        Contexto contexto = new Contexto();
        contexto.mostrar();
        contexto.setEstado(new EstadoConcretoB());
        contexto.mostrar();
        contexto.setEstado(new EstadoConcretoC());
        contexto.mostrar();
        contexto.setEstado(new EstadoConcretoA());
        contexto.mostrar();


    }

}
