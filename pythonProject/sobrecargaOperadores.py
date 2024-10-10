class Produto:
    def __init__(self, nome, preco):
        self.nome = nome
        self.preco = preco

    def __add__(self, outro_produto):
        return self.preco + outro_produto.preco

# Testando a sobrecarga
produto1 = Produto("Produto A", 50)
produto2 = Produto("Produto B", 70)
print(produto1 + produto2)  # Soma dos preços
