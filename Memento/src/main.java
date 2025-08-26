public class main {
    public static void main(String[] args) {
        Caretaker caretaker = new Caretaker();
        Originator originator = new Originator();

        originator.setState("A");
        originator.setState("B");
        originator.setState("C");
        caretaker.addMemento(originator.save());
        Memento memento = caretaker.getMemento(0);
        System.out.println("memento state: " + memento.getState());
        originator.setState("D");
        originator.setState("E");
        caretaker.addMemento(originator.save());
        Memento memento2 = caretaker.getMemento(1);
        originator.setState("F");
        caretaker.addMemento(originator.save());
        Memento memento3 = caretaker.getMemento(2);
        System.out.println("memento state: " + memento2.getState());
        System.out.println("memento state: " + memento3.getState());

    }
}
