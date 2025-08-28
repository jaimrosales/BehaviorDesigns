package Template.src;

public class CasoConcretoB extends ClaseAbstracta {

    public CasoConcretoB() {
        System.out.println("Constructor Caso Concreto B");
    }

    @Override
    public int sumar(int n) {
        return n + 15;
    }

    @Override
    public int multiplicar(int n) {
        return n * 1000;
    }

    
}