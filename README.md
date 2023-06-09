Readme

Este é um programa em Go que publica mensagens em uma exchange do RabbitMQ e consome mensagens de uma fila do RabbitMQ. Ele também fornece um servidor HTTP para lidar com solicitações recebidas.

Pré-requisitos
Go instalado em sua máquina
Servidor RabbitMQ em execução localmente na porta padrão (5672)
Primeiros passos
Clone o repositório ou baixe os arquivos do código-fonte.

Instale as dependências necessárias executando o seguinte comando no terminal:

shell
Copy code
go get -u github.com/gorilla/handlers
go get -u github.com/gorilla/mux
go get -u github.com/rabbitmq/amqp091-go
Compile e execute o programa usando o seguinte comando:

shell
Copy code
go run main.go
Isso iniciará o servidor HTTP na porta 8084 e estabelecerá uma conexão com o servidor RabbitMQ.

Uso
Publicando Mensagens
A função Publich é responsável por publicar mensagens na exchange do RabbitMQ. Você pode personalizar a mensagem modificando a struct skimas.WhoisData passada para a função. Por padrão, ela publica uma struct WhoisData de exemplo com os seguintes valores:

go
Copy code
skimas.WhoisData{
    Domain:       "teste.com",
    Name:         "teste",
    Email:        "mail@gmail.com",
    Phone:        "123456789",
    Country:      "BR",
    Organization: "teste",
    CNPJ:         "123456789",
}
Para publicar mensagens diferentes, modifique a struct WhoisData ou substitua a chamada da função Publich na função main.

Consumindo Mensagens
O programa consome mensagens de uma fila do RabbitMQ chamada "concurrentInformation". Ele inicia várias goroutines para lidar com o processamento das mensagens de forma concorrente. Cada goroutine consome mensagens da fila, converte a carga útil JSON em uma struct WhoisData e envia um e-mail usando a função service.SendEmail.

Você pode personalizar a lógica do consumidor modificando as goroutines na função main. Por exemplo, você pode alterar o número de goroutines ou adicionar etapas adicionais de processamento para as mensagens recebidas.

Servidor HTTP
O programa inicia um servidor HTTP usando o roteador Gorilla Mux na porta 8084. Ele configura cabeçalhos CORS e um middleware para definir o cabeçalho "Referrer-Policy".

O servidor está configurado para lidar com solicitações HTTP recebidas. Você pode modificar o roteador e adicionar suas próprias rotas e manipuladores de solicitação para atender às necessidades do seu aplicativo.

Licença
Este projeto está licenciado sob a Licença MIT.
