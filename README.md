# command_ssh
Este é um aplicativo Go simples que conecta-se a um dispositivo remoto via SSH e executa uma série de comandos. O aplicativo utiliza credenciais de senha para autenticação SSH e lê os comandos de um arquivo. Ele executa os comandos em sequência e registra o resultado de cada execução.

Funcionalidades
Conexão SSH a um dispositivo remoto utilizando endereço IP/porta e autenticação via senha.
Leitura de comandos a partir de um arquivo.
Execução de múltiplos comandos via SSH.
Registra o status de cada comando (sucesso ou falha) no log.
Dependências
O aplicativo utiliza a biblioteca simplessh para gerenciar a conexão SSH. Certifique-se de instalá-la:

bash
Copy code
go get github.com/helloyi/go-sshclient
Instalação
Clone este repositório:

bash
Copy code
git clone https://github.com/seu-usuario/seu-repositorio.git
cd seu-repositorio
Compile o aplicativo:

bash
Copy code
go build -o ssh-command-executor
Uso
Prepare um arquivo com os comandos a serem executados, um por linha.
Execute o aplicativo, que se conectará ao dispositivo remoto via SSH e executará cada comando listado no arquivo.
Exemplo de uso:

bash
Copy code
./ssh-command-executor
Configuração no Código
Aqui está um resumo do que o código faz:

Conecta-se via SSH ao endereço localhost:500 usando as credenciais de root e a senha passwd.
Lê comandos de um arquivo utilizando a função utils.ReadFile().
Para cada comando no arquivo:
Executa o comando via SSH.
Registra no log se o comando foi executado com sucesso ou se houve falha.
Exemplo de Comandos no Arquivo
Seu arquivo de comandos deve conter os comandos que deseja executar, um por linha, por exemplo:

bash
Copy code
uptime
df -h
free -m
Código Exemplo
go
Copy code
package main

import (
    "fmt"
    "log"
    "github.com/helloyi/go-sshclient"
    "your_project/utils"
)

func main() {
    log.Println("Aplicação iniciada")
    var client *simplessh.Client
    var err error

    // Conecta ao servidor SSH
    if client, err = simplessh.ConnectWithPassword("localhost:500", "root", "passwd"); err != nil {
        log.Fatal(err)
    }
    defer client.Close()

    // Lê comandos do arquivo
    commandText := utils.ReadFile()

    // Executa cada comando lido
    for _, command := range commandText {
        if _, err := client.Exec(command); err != nil {
            log.Println(err)
        }
        a := fmt.Sprintf("O Comando %s foi executado com sucesso", command)
        log.Println(a)
    }
}
Função utils.ReadFile()
A função ReadFile() deve ler um arquivo de texto e retornar uma lista de strings, onde cada string representa um comando a ser executado. Um exemplo simples dessa função poderia ser:

go
Copy code
package utils

import (
    "bufio"
    "os"
    "log"
)

func ReadFile() []string {
    file, err := os.Open("commands.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var commands []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        commands = append(commands, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return commands
}
Logs
O aplicativo registra cada comando que é executado e se foi bem-sucedido ou não. Isso pode ser útil para auditorias ou depuração.

Contribuições
Sinta-se à vontade para abrir pull requests ou relatar problemas no repositório!

Licença
Este projeto está licenciado sob a MIT License.
