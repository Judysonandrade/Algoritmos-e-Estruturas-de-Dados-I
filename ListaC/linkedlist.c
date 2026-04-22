#include <stdio.h>
#include <stdlib.h>

typedef struct Node {
    int valor;
    struct Node *proximo;

} Node;

// Inserir no Inicio 

void inserir_no_inicio(Node **lista, int numero){
    Node *novo = malloc(sizeof(Node));

    if(novo){
        novo -> valor = numero;
        novo -> proximo = *lista;
        *lista = novo;
    } else {
        printf("Erro ao alocar memoria !\n");
    }
}

// inserir no final da lista

void inserir_no_fim(Node **lista, int numero){
    Node  *aux ,*novo = malloc(sizeof(Node));

    if(novo){
        novo->valor = numero;
        novo->proximo = NULL;
        
        // É O PRIMEIRO ?

        if(*lista==NULL){
            *lista = novo;
        } else {
             aux = *lista;
             while(aux->proximo){
                aux = aux -> proximo;
            
            }
            aux->proximo = novo;
        }

    } else {
        printf("Erro ao alocar memoria! \n");
    }
}

// INSERIR  NO MEIO

void inserir_no_meio(Node **lista, int numero, int anterior){
    Node *aux, *novo = malloc(sizeof(Node));

    if(novo){
        novo->valor = numero;
    
        // é o primeiro

        if(*lista == NULL){
            novo->proximo=NULL;
            *lista = novo;
        } else {
            aux = *lista;
            while(aux->valor != anterior && aux->proximo){
                aux = aux->proximo;
            }
            novo->proximo= aux->proximo;
            aux->proximo = novo;
        }

    } else {
        printf("Erro ao alocar memoria! \n");
    }
}

void imprimir(Node *no){
    printf("\n Lista: ");
    while(no){
        printf("%d", no->valor);
        no = no->proximo;
    }
    printf("\n\n");
}

int main(){
    int opcao, valor, anterior;
    Node *lista = NULL; // É fundamental inicializar a lista como NULL

    do {
        printf("----------------------------------\n");
        printf("1 - Inserir no Inicio\n");
        printf("2 - Inserir no Fim\n");
        printf("3 - Inserir no Meio\n");
        printf("4 - Imprimir Lista\n");
        printf("0 - Sair\n");
        printf("Escolha uma opcao: ");
        scanf("%d", &opcao);

        switch(opcao){
            case 1:
                printf("Digite o valor a ser inserido: ");
                scanf("%d", &valor);
                inserir_no_inicio(&lista, valor);
                break;
            case 2:
                printf("Digite o valor a ser inserido no fim: ");
                scanf("%d", &valor);
                inserir_no_fim(&lista, valor);
                break;
            case 3:
                printf("Digite o valor a ser inserido: ");
                scanf("%d", &valor);
                printf("Inserir apos qual valor? ");
                scanf("%d", &anterior);
                inserir_no_meio(&lista, valor, anterior);
                break;
            case 4:
                imprimir(lista);
                break;
            case 0:
                printf("Saindo do programa...\n");
                break;
            default:
                printf("Opcao invalida! Tente novamente.\n");
        }
    } while(opcao != 0);

    return 0;
}