class Motor:
    def __init__(self, tipo, potencia):
        self.tipo = tipo
        self.potencia = potencia

    def ligar(self):
        print("carro ligado")

    def desligar(self):
        print("carro desligado")

class Pneu:
    def __init__(self,marca,tamanho):
        self.marca = marca
        self.tamanho = tamanho

    def inflar(self):
        print("O pneu inflou")

    def desinflar(self):
        print("O pneu desinflou")

class Carro:
    def __init__(self, marca, modelo):
        self.marca = marca
        self.modelo = modelo
        self.motor = Motor("Gasolina", 150)
        self.pneu = [Pneu("Pirelli", 18) for _ in range(4)]

    def ligar(self):
        self.motor.ligar()

    def desligar(self):
        self.motor.desligar()

    def trocar_pneu(self,indice,novo_pneu):
        self.pneu[indice] = novo_pneu
        print(f"Pneu {indice+1} foi trocado")

carro1 = Carro("Toyota", "Corolla")
carro1.ligar()
novo_pneu = Pneu("michellin", 16)
carro1.trocar_pneu(2,novo_pneu)
carro1.desligar()