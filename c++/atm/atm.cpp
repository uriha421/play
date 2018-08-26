#include <iostream>
#include "atm.h"

using namespace std;

int main() {
  Atm atm1; // a test of default constructor
  Atm atm2 = atm1; // a test of copy constructor
  cout << "atm1 = " << atm1 << '\n'; // tests of operator<< and print()
  cout << "atm1 = " << atm2 << '\n';

  atm1.deposit(10000); // a test of deposit()
  cout << "atm1 = " << atm1 << '\n';
  atm1.withdraw(4444); // a test of withdraw()
  cout << "atm1 = " << atm1 << '\n';

  atm2 = atm1; // a test of assignment operator
  cout << "atm2 = " << atm2 << '\n';
} // a test of deconstructor
