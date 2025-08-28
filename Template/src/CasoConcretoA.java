package Template.src;

public class CasoConcretoA extends ClaseAbstracta {

    public CasoConcretoA() {
        System.out.println("Constructor Caso Concreto A");
    }

    @Override
    public int sumar(int n) {
        return n + 1;
    }

    @Override
    public int multiplicar(int n) {
        return n * 0;
    }

    
}