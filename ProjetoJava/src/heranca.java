class Animal {
    public void som() {}
}

class Cachorro extends Animal2 {
    public void som() {
        System.out.println("O cachorro faz: Au au!");
    }
}

class Gato extends Animal2 {
    public void som() {
        System.out.println("O gato faz: Miau!");
    }

    public static void main(String[] args) {
        Cachorro2 cachorro = new Cachorro2();
        Gato2 gato = new Gato2();
        cachorro.som();
        gato.som();
    }
}
