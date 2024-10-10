import java.util.ArrayList;
import java.util.List;

class Empregado {
    String nome;
    String cargo;
    double salario;

    public Empregado(String nome, String cargo, double salario) {
        this.nome = nome;
        this.cargo = cargo;
        this.salario = salario;
    }
}

class Empresa {
    String nome;
    List<Empregado> empregados = new ArrayList<>();

    public Empresa(String nome) {
        this.nome = nome;
    }

    public void adicionarEmpregado(Empregado empregado) {
        empregados.add(empregado);
    }

    public void listarEmpregados() {
        for (Empregado empregado : empregados) {
            System.out.println(empregado.nome + ", " + empregado.cargo + ", R$" + empregado.salario);
        }
    }

    public static void main(String[] args) {
        Empresa empresa = new Empresa("Tech Solutions");
        Empregado empregado1 = new Empregado("Ana", "Gerente", 8000);
        Empregado empregado2 = new Empregado("João", "Desenvolvedor", 5000);
        empresa.adicionarEmpregado(empregado1);
        empresa.adicionarEmpregado(empregado2);
        empresa.listarEmpregados();
    }
}
