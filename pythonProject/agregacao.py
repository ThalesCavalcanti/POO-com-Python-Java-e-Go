class Empregado:
    def __init__(self, nome, cargo, salario):
        self.nome = nome
        self.cargo = cargo
        self.salario = salario

class Empresa:
    def __init__(self, nome):
        self.nome = nome
        self.empregados = []

    def adicionar_empregado(self, empregado):
        self.empregados.append(empregado)

    def listar_empregados(self):
        for empregado in self.empregados:
            print(f'{empregado.nome}, {empregado.cargo}, R${empregado.salario:.2f}')

# Testando agregação
empresa = Empresa("Tech Solutions")
empregado1 = Empregado("Ana", "Gerente", 8000)
empregado2 = Empregado("João", "Desenvolvedor", 5000)
empresa.adicionar_empregado(empregado1)
empresa.adicionar_empregado(empregado2)
empresa.listar_empregados()
