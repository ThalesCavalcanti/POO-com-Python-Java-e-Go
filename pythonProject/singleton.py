class Configuracao:
    _instancia = None

    def __new__(cls):
        if cls._instancia is None:
            cls._instancia = super().__new__(cls)
        return cls._instancia

# Testando o Singleton
config1 = Configuracao()
config2 = Configuracao()
print(config1 == config2)  # True, pois é a mesma instância
