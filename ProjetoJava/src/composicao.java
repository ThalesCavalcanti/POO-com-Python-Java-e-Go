class Motor {
    int potencia;

    public Motor(int potencia) {
        this.potencia = potencia;
    }
}

class Carro3 {
    String marca;
    String modelo;
    Motor motor;

    public Carro3(String marca, String modelo, Motor motor) {
        this.marca = marca;
        this.modelo = modelo;
        this.motor = motor;
    }

    public void exibirDetalhes() {
        System.out.println("Marca: " + marca + ", Modelo: " + modelo + ", Motor: " + motor.potencia + " cavalos");
    }

    public static void main(String[] args) {
        Motor motorV8 = new Motor(450);
        Carro3 carro = new Carro3("Ford", "Mustang", motorV8);
        carro.exibirDetalhes();
    }
}
