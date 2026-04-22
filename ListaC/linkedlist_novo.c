#include <stdio.h>
#include <stdlib.h>

typedef struct Node {
    int valor;
    struct Node *proximo;
} Node;

typedef struct{
    Node *inicio;
    int tam;
} Lista;

void criar_lista(Lista *lista){
    lista->inicio = NULL;
    lista->tam = 0;
} 

// INSERIR NO INICIO
void inserir_no_inicio(Lista *lista, int numero){
    Node *novo = malloc(sizeof(Node));

    if(novo){
        novo->valor = numero;
        novo->proximo = lista->inicio; // CORREÇÃO 1: estava lista->novo
        lista->inicio = novo;
        lista->tam++;
    } else {
        printf("Erro ao alocar memoria! \n");
    }
}

// INSERIR NO FINAL 
void inserir_no_fim(Lista *lista, int numero){
    Node *aux ,*novo = malloc(sizeof(Node));

    if(novo){
        novo->valor = numero;
        novo->proximo = NULL;
        
        if(lista->inicio == NULL){ // CORREÇÃO 2: estava = em vez de ==
            lista->inicio = novo;
        } else {
            aux = lista->inicio;
            while(aux->proximo){
                aux = aux->proximo;
            }
            aux->proximo = novo;
        }
        lista->tam++;

    } else {
        printf("Erro ao alocar memoria! \n");
    }
}

// INSERIR NO MEIO
void inserir_no_meio(Lista *lista, int numero, int anterior){
    Node *aux, *novo = malloc(sizeof(Node));

    if(novo){
        novo->valor = numero;
        
        if(lista->inicio == NULL){ // CORREÇÃO 3: estava = em vez de ==
            novo->proximo = NULL;
            lista->inicio = novo;
        } else {
            aux = lista->inicio;
            while(aux->valor != anterior && aux->proximo){
                aux = aux->proximo;
            }
            novo->proximo = aux->proximo;
            aux->proximo = novo;
        }
        lista->tam++;

    } else {
        printf("Erro ao alocar memoria! \n");
    }
}

void inserir_ordenado(Lista *lista, int numero){
    Node *aux, *novo = malloc(sizeof(Node));

    if(novo){
        novo->valor = numero;

        // TA VAZIA?

        if(lista->inicio == NULL){
            novo->proximo = NULL;
            lista->inicio = novo;

        // É O MENOR?
        } else if(novo->valor < lista->inicio->valor){
            novo->proximo = lista->inicio;
            lista->inicio = novo;
        } else {
            aux = lista->inicio;
            while(aux->proximo && novo->valor > aux->proximo->valor){
                aux = aux->proximo;
            }
            novo->proximo = aux->proximo;
            aux->proximo = novo;
        }

        lista->tam++;

    } else {
        printf("Erro ao alocar a memoria");
    }
}

// REMOVER Elemento

Node *remover(Lista *lista, int numero){
    Node *aux ,*remover = NULL; 
    if(lista->inicio){
        if(lista->inicio->valor == numero){
            remover = lista->inicio;
            lista->inicio = remover->proximo;
            lista->tam--;
        } else {
            aux = lista->inicio;
            while(aux->proximo && aux->proximo->valor !=numero){
                aux = aux->proximo;

            }
            if(aux->proximo){
                remover = aux->proximo;
                aux->proximo = remover->proximo;
                lista->tam--;
            }
        }
    }
    return remover;

}

Node *buscar(Lista *lista, int numero){
    Node *aux ,*no = NULL;

    aux = lista->inicio;
    while(aux && aux->valor !=numero){
        aux = aux->proximo;
    }
    if(aux){
        no = aux;
    }

    return no;
}

// IMPRIMIR A LISTA
void imprimir(Lista *lista){
    Node *no_removido = lista->inicio;
    
    // Aproveitando que a struct guarda o tamanho, vamos exibi-lo!
    printf("\nLista (Tamanho: %d): ", lista->tam); 
    
    while(no_removido){
        printf("%d ", no_removido->valor);
        no_removido = no_removido->proximo;
    }
    printf("\n\n");
}

// FUNÇÃO PRINCIPAL (MAIN)
int main() {
    int opcao, valor, anterior;
    Lista lista;
    Node *removido, *encontrado; // Cria a variável do tipo Lista
    
    criar_lista(&lista); // Inicializa a lista com NULL e tam = 0

    do {
        printf("----------------------------------\n");
        printf("1 - Inserir no Inicio\n");
        printf("2 - Inserir no Fim\n");
        printf("3 - Inserir no Meio\n");
        printf("4 - Inserir Ordenado\n");
        printf("5 - Remover\n");
        printf("6 - Buscar Elemento\n");
        printf("7 - Imprimir Lista\n");
        printf("0 - Sair\n");
        printf("Escolha uma opcao: ");
        scanf("%d", &opcao);

        switch(opcao){
            case 1:
                printf("Digite o valor a ser inserido no inicio: ");
                scanf("%d", &valor);
                inserir_no_inicio(&lista, valor);
                break;
            case 2:
                printf("Digite o valor a ser inserido no fim: ");
                scanf("%d", &valor);
                inserir_no_fim(&lista, valor);
                break;
            case 3:
                printf("Digite o novo valor a ser inserido: ");
                scanf("%d", &valor);
                printf("Inserir apos qual valor existente? ");
                scanf("%d", &anterior);
                inserir_no_meio(&lista, valor, anterior);
                break;
            case 4: 
                printf("Digite o novo valor a ser inserido: ");
                scanf("%d", &valor);
                inserir_ordenado(&lista, valor);
                break;
            case 5:
                printf("Digite o novo valor a ser removido: ");
                scanf("%d", &valor);  
                
                removido = remover(&lista, valor);

                if(removido){
                    printf("Elemento removido com sucesso");
                    free(removido);
                } else {
                    printf("Elemento não encontrado");
                }
                break;
            
            case 6:
                printf("Digite o valor a ser buscado: ");
                scanf("%d", &valor);

                encontrado = buscar(&lista, valor);
                if(encontrado){
                    printf("Elemento encontrado");

                } else {
                    printf("Não achou");
                }
                break;

            case 7:
                // Note que enviamos o endereço de memória (&lista)
                imprimir(&lista); 
                break;
            case 0:
                printf("Saindo...\n");
                break;
            default:
                printf("Opcao invalida!\n");
        }
    } while(opcao != 0);

    return 0;
}