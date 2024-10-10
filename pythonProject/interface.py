from abc import ABC, abstractmethod

class Imprimivel(ABC):
    @abstractmethod
    def imprimir(self):
        pass

class Relatorio(Imprimivel):
    def imprimir(self):
        print("Imprimindo relatório...")

class Contrato(Imprimivel):
    def imprimir(self):
        print("Imprimindo contrato...")

# Testando interfaces
documentos = [Relatorio(), Contrato()]
for doc in documentos:
    doc.imprimir()
