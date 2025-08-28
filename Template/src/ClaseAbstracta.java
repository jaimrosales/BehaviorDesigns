package Template.src;

public abstract class ClaseAbstracta {
    
    //Ejemplo del algoritmo a encapsular
    public int getNum(int num) {  
        this.mostrar();
        int n = this.sumar(num);
        n = this.multiplicar(n);
        return n;
    }

    public void mostrar() {
        System.out.println("Iniciando el algoritmo...");
    }
    public abstract int sumar(int n);
    public abstract int multiplicar(int n);
}




