class Matematica:
    @staticmethod
    def fatorial(n):
        if n == 0:
            return 1
        else:
            return n * Matematica.fatorial(n-1)

# Testando o método estático
print(Matematica.fatorial(5))
