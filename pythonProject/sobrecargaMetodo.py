class Calculadora:
    def somar(self, a, b, c=None):
        if c is None:
            return a + b
        else:
            return a + b + c

# Testando a Calculadora
calc = Calculadora()
print(calc.somar(2, 3))  # Dois parâmetros
print(calc.somar(2, 3, 4))  # Três parâmetros
