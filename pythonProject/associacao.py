class Professor:
    def __init__(self, nome):
        self.nome = nome
        self.escolas = []

    def associar_escola(self, escola):
        self.escolas.append(escola)

class Escola:
    def __init__(self, nome):
        self.nome = nome
        self.professores = []

    def adicionar_professor(self, professor):
        self.professores.append(professor)

# Testando associação
escola = Escola("Escola Primária")
professor = Professor("Maria")
escola.adicionar_professor(professor)
professor.associar_escola(escola)

print(f'{professor.nome} leciona na {escola.nome}')
