class Matematica {
    public static int fatorial(int n) {
        if (n == 0) {
            return 1;
        } else {
            return n * fatorial(n - 1);
        }
    }

    public static void main(String[] args) {
        System.out.println("Fatorial de 5: " + Matematica.fatorial(5));
    }
}
