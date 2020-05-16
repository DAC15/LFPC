#include <iostream>
#include <fstream>
#include <vector>
#include <string>
#include <ctype.h>

using namespace std;
struct word
{
  int token;
  string name;
};

enum tokens
{
  tok_keywords = -1,
  tok_operators = -2,
  tok_delimeters = -3,
  tok_identifire = -4,
  tok_numbers = -5,
};

vector<string> keywords = {"define","var","call","return","for","while","if","else","Adriana"};
vector<string> operators = {"+","-","/","*","=","<",">","==","++","--","<=",">="};
vector<string> delimeters = {"(",")","{","}",",","[","]"};

string readFIle()
{
  string myText;
  string line;
  ifstream myfile("program.txt");
    if(myfile.is_open())
    {
      while (getline(myfile, line))
      {
        if (line == "") continue;
        myText += line + " ";
      }
      myfile.close();
    }
  return myText;
}

vector<string> makeWords()
{
    int i = 0;
    vector<string> tmpVec;
    string tmp;
    while (i<readFIle().length())
    {
      if(isspace(readFIle()[i]))
      {
        tmpVec.push_back(tmp);
        tmp.clear();
        goto inc;
      }else if(ispunct(readFIle()[i]))
      {
        tmpVec.push_back(tmp);
        tmpVec.push_back(string(1, readFIle()[i]));
        tmp.clear();
        goto inc;
      }
      tmp+=readFIle()[i];
  inc:i++;
    }
    return tmpVec;
}

vector<word> Tokire()
{
  int i = 0;
  vector<word> Tokens;
  vector<string> tmpVec = makeWords();
  while(i < tmpVec.size())
  {
      //keywords
      for(int j = 0; j < keywords.size(); j++)
      {
        if(tmpVec[i]==keywords[j])
        {
          Tokens.push_back({tok_keywords,tmpVec[i]});
          goto out;
        }
      }
      //delimeters
      for(int k = 0; k < delimeters.size(); k++)
      {
        if(tmpVec[i]==delimeters[k])
        {
          Tokens.push_back({tok_delimeters,tmpVec[i]});
          goto out;
        }
      }
      //operators
      for(int y = 0; y < operators.size(); y++)
      {
        if(tmpVec[i]==operators[y])
        {
          Tokens.push_back({tok_operators,tmpVec[i]});
          goto out;
        }
      }
      //identifier
      if(tmpVec[i][0] >= 60 && tmpVec[i][0] <=90 || tmpVec[i][0] >= 97 && tmpVec[i][0] <=122){
          Tokens.push_back({tok_identifire,tmpVec[i]});
        }
      //number
      for(int z = 0; z < tmpVec[i].length(); z++){
            if(!(tmpVec[i][z] >=48 && tmpVec[i][z] <=57)) break;
            else if(z == tmpVec[i].length()-1)
            Tokens.push_back({tok_numbers,tmpVec[i]});
          }
       out: i++;
  }
  return Tokens;
}

int main () {
  vector<word> Tokens = Tokire();
    for(int test = 0; test<Tokens.size(); test++)
      cout<<"Value\t"<<Tokens[test].name<<"\t\tToken\t"<<Tokens[test].token<<"\tIndex\t"<<test<<endl;
  return 0;
}
