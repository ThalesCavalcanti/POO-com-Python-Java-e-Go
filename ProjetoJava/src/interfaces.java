interface Imprimivel {
    void imprimir();
}

class Relatorio implements Imprimivel {
    public void imprimir() {
        System.out.println("Imprimindo relatório...");
    }
}

class Contrato implements Imprimivel {
    public void imprimir() {
        System.out.println("Imprimindo contrato...");
    }

    public static void main(String[] args) {
        Imprimivel[] documentos = { new Relatorio(), new Contrato() };
        for (Imprimivel doc : documentos) {
            doc.imprimir();
        }
    }
}
