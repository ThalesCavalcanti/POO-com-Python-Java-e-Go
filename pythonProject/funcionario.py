class Pessoa:
    def __init__(self,nome,idade):
        self.nome = nome
        self.idade = idade
        self.endereco = None

    def adicionar_end(self,endereco):
        self.endereco = endereco

    def mostrar_info(self):
        print(f"nome: {self.nome}, idade: {self.idade}")
        if self.endereco:
            print("endereco: ")
            self.endereco.mostrar_end()
        else:
            print("Endereco Indisponivel")

class Endereco:
    def __init__(self,rua,cidade,estado,cep):
        self.rua = rua
        self.cidade = cidade
        self.estado = estado
        self.cep = cep

    def mostrar_end(self):
        print(f"rua: {self.rua}, cidade: {self.cidade}, estado: {self.estado}, cep: {self.cep}")

class Empresa:
    def __init__(self,nome,cnpj):
        self.nome = nome
        self.cnpj = cnpj
        self.funcionarios = []

    def contratar_func(self, funcionario):
        self.funcionarios.append(funcionario)

    def listar_funcionarios(self):
        print(f"funcionarios da empresa: {self.nome}")
        for funcionario in self.funcionarios:
            print(funcionario.nome)

endereco1 = Endereco("av.pinheiro","osasco", "sp", "12355-688")
pessoa1 = Pessoa("Carlos", "20")
pessoa1.adicionar_end(endereco1)

endereco2 = Endereco("Rua secunda", "cidade 2", "bc", "32156-121")
pessoa2 = Pessoa("Babun", "31")
pessoa2.adicionar_end(endereco2)

empresa = Empresa("Empresa ABC", "12312312312")
empresa.contratar_func(pessoa1)
empresa.contratar_func(pessoa2)

empresa.listar_funcionarios()