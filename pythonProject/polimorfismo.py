from heranca import Cachorro, Gato


def emitir_som(animais):
    for animal in animais:
        animal.som()

# Testando o polimorfismo
animais = [Cachorro(), Gato()]
emitir_som(animais)
