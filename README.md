# Bitcoin Wallet Generator for Testnet

Este projeto em Go demonstra a criação de uma carteira para Bitcoin na rede Testnet, implementando as normas BIP39 e BIP32/BIP49 para geração de chaves hierárquicas determinísticas (HD Wallets).

## Fluxo de Atividades

1. **Geração de Entropia**: Criação de uma sequência aleatória de dados (128 bits) para ser a base do mnemônico.
2. **Geração de Mnemônico**: Transformação da entropia em uma frase mnemônica de fácil leitura.
3. **Geração de Seed**: Derivação de uma seed a partir do mnemônico usando uma senha.
4. **Geração de Chave Mestra**: Criação de uma chave mestre a partir da seed.
5. **Derivação de Chaves**: Uso da chave mestre para gerar chaves subordinadas seguindo a estrutura BIP49.
6. **Geração do Endereço Público**: Conversão da chave pública derivada em um endereço público Bitcoin.
7. **Conversão para WIF**: Transformação da chave privada em formato WIF para ser usado em outras carteiras.

## Descrição Detalhada

### 1. Geração de Entropia
A entropia é gerada para criar uma base segura e aleatória para o mnemônico. Este processo utiliza funções criptográficas para garantir que a entropia seja suficientemente aleatória.

### 2. Geração de Mnemônico
O mnemônico é criado a partir da entropia gerada. Ele converte a informação binária em palavras selecionadas de uma lista predeterminada, facilitando a anotação e memorização.

### 3. Geração de Seed
A seed é derivada usando o mnemônico e uma senha. Este processo usa o PBKDF2 para misturar o mnemônico com a senha e produzir uma seed que serve como a raiz de todas as futuras chaves da carteira.

### 4. Geração de Chave Mestra
A chave mestre é criada a partir da seed. Esta chave será a base para todas as chaves subsequentes geradas pela carteira, seguindo o esquema hierárquico determinístico.

### 5. Derivação de Chaves
A chave mestre é usada para derivar novas chaves de forma segura e organizada. A especificação BIP49 é seguida para garantir que as chaves sejam compatíveis com carteiras SegWit.

### 6. Geração do Endereço Público
A partir da chave pública derivada, um endereço Bitcoin público é gerado. Este endereço pode ser usado para receber transações na rede Bitcoin Testnet.

### 7. Conversão para WIF
A chave privada selecionada é convertida para o formato Wallet Import Format (WIF), permitindo sua utilização em diversas carteiras de software e hardware.

## Instalação e Uso

### Pré-requisitos
- Instale Go no seu sistema se ainda não estiver instalado. Verifique com `go version`.

### Configuração do Projeto

1. Clone o repositório:
```
git clone https://github.com/felipesilvaibi/go-btc-wallet-generator.git
```

2. Navegue até o diretório do projeto:
```
cd go-btc-wallet-generator
```

3. Inicialize um módulo Go:
```
go mod init go-btc-wallet-generator
```

4. Instale as dependências:
```
go get github.com/btcsuite/btcutil
go get github.com/tyler-smith/go-bip39
go get github.com/btcsuite/btcutil/hdkeychain
```

5. Execute o projeto:
```
go run src/main.go
```