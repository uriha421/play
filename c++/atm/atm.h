#include <iostream>

using namespace std;

class Atm {
public:
  Atm();

public:
  void print(ostream *os);
  void deposit(int m);
  void withdraw(int m);

private:
  int money;
};

Atm::Atm() {
  money = 0;
  cout << "Create a new Atm object";
}

void Atm::print(ostream *os) {
  *os << "Atm{" << money << '}';
}

ostream &operator<<(ostream &os, Atm &theAtm) {
  theAtm.print(&os);
  return os;
}

void Atm::deposit(int m) {
  money += m;
}

void Atm::withdraw(int m) {
  money -= m;
}
