class SaldoInsuficienteException extends Exception {
    public SaldoInsuficienteException(String mensagem) {
        super(mensagem);
    }
}

class ContaBancaria2 {
    private double saldo;
    private String titular;

    public ContaBancaria2(String titular, double saldo) {
        this.titular = titular;
        this.saldo = saldo;
    }

    public void sacar(double valor) throws SaldoInsuficienteException {
        if (valor > saldo) {
            throw new SaldoInsuficienteException("Saldo insuficiente.");
        }
        saldo -= valor;
    }

    public static void main(String[] args) {
        ContaBancaria2 conta = new ContaBancaria2("João", 1000);
        try {
            conta.sacar(1500);
        } catch (SaldoInsuficienteException e) {
            System.out.println(e.getMessage());
        }
    }

    public void depositar(double valor) {
        saldo += valor;
    }

    public void exibirSaldo() {System.out.println("Saldo atual: R$" + saldo);}

}
