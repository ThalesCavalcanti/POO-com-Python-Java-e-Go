class ContaBancaria:
    def __init__(self, titular, saldo=0):
        self.__saldo = saldo  # Encapsulamento
        self.titular = titular

    def depositar(self, valor):
        self.__saldo += valor

    def sacar(self, valor):
        if valor <= self.__saldo:
            self.__saldo -= valor
        else:
            print("Saldo insuficiente.")

    def exibir_saldo(self):
        print(f'Saldo atual: R${self.__saldo:.2f}')

# Testando a ContaBancaria
conta = ContaBancaria("João", 1000)
conta.depositar(500)
conta.sacar(200)
conta.exibir_saldo()
