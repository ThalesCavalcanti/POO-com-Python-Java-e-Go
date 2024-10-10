class Carro:
    def __init__(self, marca, modelo, ano):
        self.marca = marca
        self.modelo = modelo
        self.ano = ano
        self.velocidade = 0

    def acelerar(self, quantidade):
        self.velocidade += quantidade

    def frear(self, quantidade):
        self.velocidade = max(0, self.velocidade - quantidade)

    def exibir_velocidade(self):
        print(f'Velocidade atual: {self.velocidade} km/h')

# Testando os métodos
carro = Carro("Ford", "Mustang", 2022)
carro.acelerar(50)
carro.exibir_velocidade()
carro.frear(20)
carro.exibir_velocidade()
