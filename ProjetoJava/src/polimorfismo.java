class Animal2 {
    public void som() {}
}

class Cachorro2 extends Animal2 {
    public void som() {
        System.out.println("O cachorro faz: Au au!");
    }
}

class Gato2 extends Animal2 {
    public void som() {
        System.out.println("O gato faz: Miau!");
    }
}

class Polimorfismo {
    public static void emitirSom(Animal2[] animais) {
        for (Animal2 animal : animais) {
            animal.som();
        }
    }

    public static void main(String[] args) {
        Animal2[] animais = { new Cachorro2(), new Gato2() };
        emitirSom(animais);
    }
}
