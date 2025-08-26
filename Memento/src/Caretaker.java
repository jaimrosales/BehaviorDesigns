import java.util.ArrayList;

public class Caretaker {
    private ArrayList<Memento> mementoList = new ArrayList<>();
    
    public void addMemento(Memento memento){
        mementoList.add(memento);
    }

    public Memento getMemento(int index){
        return mementoList.get(index);
    }

}
