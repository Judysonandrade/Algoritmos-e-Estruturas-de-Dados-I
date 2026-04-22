#include <stdio.h>
#include <stdlib.h> 

typedef struct Node{
    int valor;
    struct Node *proximo;
    struct Node *anterior;
} Node;

typedef struct{
    Node *inicio;
    Node *fim;
    int tam;
} Lista;

void criar_lista(Lista *lista){
    lista->inicio = NULL;
    lista->fim = NULL;
    lista->tam = 0;
}

// INSERIR NO INICIO
void inserir_no_inicio(Lista *lista, int numero){
    Node *novo = malloc(sizeof(Node)); // Corrigido: era *no

    if(novo){
        novo->valor = numero;
        novo->proximo = lista->inicio;
        novo->anterior = NULL;
        
        if(lista->inicio){ // Corrigido: era lista-inicio
            lista->inicio->anterior = novo;
        } else {
            // Se a lista estava vazia, o novo nó é o início E o fim!
            lista->fim = novo; 
        }
        
        lista->inicio = novo;
        lista->tam++;
    } else {
        printf("Erro ao alocar memoria!\n");
    }
}

// INSERIR NO FIM (Otimizado usando lista->fim)
void inserir_no_fim(Lista *lista, int numero){
    Node *novo = malloc(sizeof(Node)); // Corrigido: era *no

    if(novo){
        novo->valor = numero;
        novo->proximo = NULL;
    
        if(lista->inicio == NULL){ // Corrigido: era = NULL
            novo->anterior = NULL;
            lista->inicio = novo;
            lista->fim = novo;
        } else {
            // Não precisamos mais do while! Usamos o lista->fim direto
            novo->anterior = lista->fim;
            lista->fim->proximo = novo;
            lista->fim = novo; // Atualiza o novo fim da lista
        }
        lista->tam++;
    } else {
        printf("Erro ao alocar memoria!\n");
    }
}

// INSERIR NO MEIO
void inserir_no_meio(Lista *lista, int numero, int anterior){
    Node *aux, *novo = malloc(sizeof(Node)); // Corrigido: era *no

    if(novo){
        novo->valor = numero;

        if(lista->inicio == NULL){ // Corrigido: era = NULL
            novo->proximo = NULL;
            novo->anterior = NULL;
            lista->inicio = novo;
            lista->fim = novo; // Atualiza o fim também
        } else {
            aux = lista->inicio;
            while(aux->valor != anterior && aux->proximo){
                aux = aux->proximo;
            }
            
            novo->proximo = aux->proximo;
            novo->anterior = aux;
            
            // Proteção crucial: só altera o proximo->anterior se o proximo existir!
            if(aux->proximo){
                aux->proximo->anterior = novo;
            } else {
                lista->fim = novo; // Se inseriu no final, atualiza o fim
            }
            
            aux->proximo = novo;
        }
        lista->tam++;
    } else {
        printf("Erro ao alocar memoria!\n");
    }
}

// INSERIR ORDENADO
void inserir_ordenado(Lista *lista, int numero){
    Node *aux, *novo = malloc(sizeof(Node)); // Corrigido: era *no

    if(novo){
        novo->valor = numero;
        
        // TA VAZIA?
        if(lista->inicio == NULL){
            novo->proximo = NULL;
            novo->anterior = NULL;
            lista->inicio = novo;
            lista->fim = novo; // Atualiza o fim
            
        // É O MENOR?
        } else if(novo->valor < lista->inicio->valor){
            novo->proximo = lista->inicio;
            novo->anterior = NULL;
            lista->inicio->anterior = novo;
            lista->inicio = novo; // Corrigido: era lista-> novo;
            
        // INSERE NO MEIO OU FIM
        } else {
            aux = lista->inicio;
            while(aux->proximo && novo->valor > aux->proximo->valor){
                aux = aux->proximo;
            }
            
            novo->proximo = aux->proximo;
            novo->anterior = aux;
            
            // Proteção crucial igual ao inserir_no_meio
            if(aux->proximo){
                aux->proximo->anterior = novo;
            } else {
                lista->fim = novo; // Atualiza o fim
            }
            aux->proximo = novo;
        }
        lista->tam++;

    } else {
        printf("Erro ao alocar memoria!\n");
    }
}

// REMOVER Elemento
Node *remover(Lista *lista, int numero){
    Node *aux, *no_remover = NULL;

    if(lista->inicio){
        if(lista->inicio->valor == numero){
            no_remover = lista->inicio;
            lista->inicio = no_remover->proximo;
            
            if(lista->inicio){
                lista->inicio->anterior = NULL;
            } else {
                lista->fim = NULL; // Se removeu o único elemento, a lista esvaziou
            }
            lista->tam--;
        } else {
            aux = lista->inicio;
            while(aux->proximo && aux->proximo->valor != numero){
                aux = aux->proximo;
            }
            
            if(aux->proximo){
                no_remover = aux->proximo;
                aux->proximo = no_remover->proximo;
                
                if(aux->proximo){
                    aux->proximo->anterior = aux;
                } else {
                    lista->fim = aux; // Se removeu o último, o aux passa a ser o fim
                }
                lista->tam--;
            }
        }
    }
    return no_remover; 
}

// BUSCAR Elemento
Node *buscar(Lista *lista, int numero){
    Node *aux = lista->inicio;
    
    while(aux && aux->valor != numero){
        aux = aux->proximo;
    }
    
    return aux;
}

// IMPRIMIR A LISTA
void imprimir(Lista *lista){
    Node *aux = lista->inicio; // Mudei o nome para aux para ficar mais legível
    
    printf("\nLista (Tamanho: %d): ", lista->tam); 
    
    while(aux){
        printf("%d ", aux->valor);
        aux = aux->proximo;
    }
    printf("\n\n");
}

// FUNÇÃO PRINCIPAL (MAIN)
int main() {
    int opcao, valor, anterior;
    Lista lista;
    Node *removido, *encontrado; 
    
    criar_lista(&lista); 

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
                    printf("Elemento %d removido com sucesso!\n", removido->valor);
                    free(removido);
                } else {
                    printf("Elemento nao encontrado!\n");
                }
                break;
            case 6:
                printf("Digite o valor a ser buscado: ");
                scanf("%d", &valor);

                encontrado = buscar(&lista, valor);
                if(encontrado){
                    printf("Elemento %d encontrado na lista!\n", encontrado->valor);
                } else {
                    printf("Elemento nao achou...\n");
                }
                break;
            case 7:
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