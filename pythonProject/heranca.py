class Animal:
    def som(self):
        pass

class Cachorro(Animal):
    def som(self):
        print("O cachorro faz: Au au!")

class Gato(Animal):
    def som(self):
        print("O gato faz: Miau!")

# Testando herança
cachorro = Cachorro()
gato = Gato()
cachorro.som()
gato.som()
