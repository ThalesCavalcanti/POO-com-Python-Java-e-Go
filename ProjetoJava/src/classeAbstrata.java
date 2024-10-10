abstract class Funcionario {
    abstract double calcularSalario();
}

class FuncionarioHorista extends Funcionario {
    private double horasTrabalhadas;
    private double valorHora;

    public FuncionarioHorista(double horasTrabalhadas, double valorHora) {
        this.horasTrabalhadas = horasTrabalhadas;
        this.valorHora = valorHora;
    }

    @Override
    public double calcularSalario() {
        return horasTrabalhadas * valorHora;
    }
}

class FuncionarioAssalariado extends Funcionario {
    private double salarioMensal;

    public FuncionarioAssalariado(double salarioMensal) {
        this.salarioMensal = salarioMensal;
    }

    @Override
    public double calcularSalario() {
        return salarioMensal;
    }

    public static void main(String[] args) {
        FuncionarioHorista horista = new FuncionarioHorista(160, 25);
        FuncionarioAssalariado assalariado = new FuncionarioAssalariado(5000);

        System.out.println("Salário horista: R$" + horista.calcularSalario());
    }
}