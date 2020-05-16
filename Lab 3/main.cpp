#include<iostream>
#include<fstream>
#include<stdlib.h>
#include<string.h>
#include<ctype.h>

using namespace std;

int languageiden(char buffer[]){
    char languageiden[32][10]={"function","with","values", "return", "main"};

    int i, flag = 0;

    for(i = 0; i < 32; ++i){
        if(strcmp(languageiden[i], buffer) == 0){
            flag = 1;
            break;
        }
    }

    return flag;
}
int datatype(char buffer[]){
    char datatype[32][10]={"int","float","double", "char", "string"};

    int i, flag = 0;

    for(i = 0; i < 32; ++i){
        if(strcmp(datatype[i], buffer) == 0){
            flag = 1;
            break;
        }
    }

    return flag;
}


int forLibraries(char buffer[]){
    char forLibraries[32][10]={"include","LIBRARY1"};

    int i, flag = 0;

    for(i = 0; i < 32; ++i){
        if(strcmp(forLibraries[i], buffer) == 0){
            flag = 1;
            break;
        }
    }

    return flag;
}

int statements(char buffer[]){
    char statements[32][10]={"if","else","for","break","while", "return","start", "stop"};

    int i, flag = 0;

    for(i = 0; i < 32; ++i){
        if(strcmp(statements[i], buffer) == 0){
            flag = 1;
            break;
        }
    }

    return flag;
}

int numbers(char buffer[]){
    char numbers[32][10]={"0","1","2","3","4","5", "6","7", "8","9","10"};

    int i, flag = 0;

    for(i = 0; i < 32; ++i){
        if(strcmp(numbers[i], buffer) == 0){
            flag = 1;
            break;
        }
    }

    return flag;
}
int main(){
    char ch, buffer[15], operators[] = "+-*/%=";
    int i,j=0;
    ifstream infile;
    infile.open ("nfadfa.txt");

    if(!infile.is_open()){
        cout<<"error while opening the file\n";
        exit(0);
    }

    while(!infile.eof()){
        ch = infile.get();

        for(i = 0; i < 6; ++i){
            if(ch == operators[i])
                cout<<ch<<" is operator\n";
        }

        if(isalnum(ch)){
            buffer[j++] = ch;
        }
        else if((ch == ' ' || ch == '\n') && (j != 0)){
            buffer[j] = '\0';
            j = 0;

          if(languageiden(buffer) == 1)
                cout<<buffer<<" - keyword\n";
            else if(forLibraries(buffer)==1)
                    cout<<buffer<<" - identifier for libraries\n";
            else if(numbers(buffer)==1)
                cout<<buffer<<" - number\n";
            else if(statements(buffer) == 1)
                cout<<buffer<<" - statement\n";
            else if(datatype(buffer) == 1)
                cout<<buffer<<" - datatype\n";
            else
                cout<<buffer<<" - indentifier\n";
        }

    }


    infile.close();

    return 0;
}
