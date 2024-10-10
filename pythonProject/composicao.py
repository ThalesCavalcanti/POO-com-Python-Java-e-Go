class Motor:
    def __init__(self, potencia):
        self.potencia = potencia

class Carro:
    def __init__(self, marca, modelo, motor):
        self.marca = marca
        self.modelo = modelo
        self.motor = motor

    def exibir_detalhes(self):
        print(f'Marca: {self.marca}, Modelo: {self.modelo}, Motor: {self.motor.potencia} cavalos')

# Testando a composição
motor_v8 = Motor(450)
carro = Carro("Ford", "Mustang", motor_v8)
carro.exibir_detalhes()
