package Template.src;

public class main {
    public static void main(String[] args) {
        //El ejemplo del patron template permite ejecutar un mismo algoritmo con diferentes implementaciones, en este caso
        //se tiene un algoritmo que realiza una suma y una multiplicacion, pero cada caso concreto define como se realiza cada operacion
        //de manera diferente
        CasoConcretoA casoA = new CasoConcretoA();
        CasoConcretoB casoB = new CasoConcretoB();
        int n = casoA.getNum(5); //Se manda a llamar el algoritmo encapsulado
        System.out.println("Resultado Caso A: " + n);
        int m = casoB.getNum(5); //Se manda a llamar el algoritmo encapsulado
        System.out.println("Resultado Caso B: " + m);

    }
}
