class Calculadora {
    public int somar(int a, int b) {
        return a + b;
    }

    public int somar(int a, int b, int c) {
        return a + b + c;
    }

    public static void main(String[] args) {
        Calculadora calc = new Calculadora();
        System.out.println(calc.somar(2, 3));  // Dois parâmetros
        System.out.println(calc.somar(2, 3, 4));  // Três parâmetros
    }
}
