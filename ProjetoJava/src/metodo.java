class Carro2 {
    String marca;
    String modelo;
    int ano;
    int velocidade = 0;

    public Carro2(String marca, String modelo, int ano) {
        this.marca = marca;
        this.modelo = modelo;
        this.ano = ano;
    }

    public void acelerar(int quantidade) {
        velocidade += quantidade;
    }

    public void frear(int quantidade) {
        velocidade = Math.max(0, velocidade - quantidade);
    }

    public void exibirVelocidade() {
        System.out.println("Velocidade atual: " + velocidade + " km/h");
    }

    public static void main(String[] args) {
        Carro2 carro = new Carro2("Ford", "Mustang", 2022);
        carro.acelerar(50);
        carro.exibirVelocidade();
        carro.frear(20);
        carro.exibirVelocidade();
    }

    public void exibirDetalhes() {
        System.out.println("Marca: " + marca + ", Modelo: " + modelo + ", Ano: " + ano);
    }
}
