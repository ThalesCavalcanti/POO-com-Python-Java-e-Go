class SaldoInsuficienteException(Exception):
    pass

class ContaBancaria:
    def __init__(self, titular, saldo=0):
        self.__saldo = saldo
        self.titular = titular

    def sacar(self, valor):
        if valor > self.__saldo:
            raise SaldoInsuficienteException("Saldo insuficiente.")
        self.__saldo -= valor

#
