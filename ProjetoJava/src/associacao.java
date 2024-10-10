import java.util.ArrayList;
import java.util.List;

class Professor {
    String nome;
    List<Escola> escolas = new ArrayList<>();

    public Professor(String nome) {
        this.nome = nome;
    }

    public void associarEscola(Escola escola) {
        escolas.add(escola);
    }
}

class Escola {
    String nome;
    List<Professor> professores = new ArrayList<>();

    public Escola(String nome) {
        this.nome = nome;
    }

    public void adicionarProfessor(Professor professor) {
        professores.add(professor);
    }

    public static void main(String[] args) {
        Escola escola = new Escola("Escola Primária");
        Professor professor = new Professor("Maria");
        escola.adicionarProfessor(professor);
        professor.associarEscola(escola);

        System.out.println(professor.nome + " leciona na " + escola.nome);
    }
}
