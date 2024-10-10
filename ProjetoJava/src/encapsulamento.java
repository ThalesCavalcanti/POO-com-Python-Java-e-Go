class ContaBancaria {
    private double saldo;
    private String titular;

    public ContaBancaria(String titular, double saldoInicial) {
        this.titular = titular;
        this.saldo = saldoInicial;
    }

    public void depositar(double valor) {
        saldo += valor;
    }

    public void sacar(double valor) {
        if (valor <= saldo) {
            saldo -= valor;
        } else {
            System.out.println("Saldo insuficiente.");
        }
    }

    public void exibirSaldo() {
        System.out.println("Saldo atual: R$" + saldo);
    }

    public static void main(String[] args) {
        ContaBancaria2 conta = new ContaBancaria2("João", 1000);
        conta.depositar(500);
        try {
            conta.sacar(200);
        } catch (SaldoInsuficienteException e) {
            throw new RuntimeException(e);
        }
        conta.exibirSaldo();
    }
}
