class Carro {
    String marca;
    String modelo;
    int ano;

    public Carro(String marca, String modelo, int ano) {
        this.marca = marca;
        this.modelo = modelo;
        this.ano = ano;
    }

    public void exibirDetalhes() {
        System.out.println("Marca: " + marca + ", Modelo: " + modelo + ", Ano: " + ano);
    }

    public static void main(String[] args) {
        Carro2 carro1 = new Carro2("Toyota", "Corolla", 2020);
        Carro2 carro2 = new Carro2("Honda", "Civic", 2019);
        Carro2 carro3 = new Carro2("Ford", "Mustang", 2022);

        carro1.exibirDetalhes();
        carro2.exibirDetalhes();
        carro3.exibirDetalhes();
    }
}
