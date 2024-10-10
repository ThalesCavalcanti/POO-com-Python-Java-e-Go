class Funcionario(ABC):
    @abstractmethod
    def calcular_salario(self):
        pass

class FuncionarioHorista(Funcionario):
    def __init__(self, horas_trabalhadas, valor_hora):
        self.horas_trabalhadas = horas_trabalhadas
        self.valor_hora = valor_hora

    def calcular_salario(self):
        return self.horas_trabalhadas * self.valor_hora

class FuncionarioAssalariado(Funcionario):
    def __init__(self, salario_mensal):
        self.salario_mensal = salario_mensal

    def calcular_salario(self):
        return self.salario_mensal

# Testando classes abstratas
horista = FuncionarioHorista(160, 25)
assalariado = FuncionarioAssalariado(5000)
print(horista.calcular_salario())
print(assalariado.calcular_salario())
