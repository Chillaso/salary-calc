# SALARY CALC
This is a very basic app which let you calc your salary in Spain.

## Requirements
### Node
First, we are going to install nvm, it'll let us install and manage node as simply as possible. Take care of nvm version. You can check the latest version [here](https://github.com/nvm-sh/nvm)
```shell
wget -qO- https://raw.githubusercontent.com/nvm-sh/nvm/v0.35.2/install.sh | bash
```
then close your terminal or source your shell config and run
```shell
nvm install --lts
```
Update npm
```shell
npm install npm@latest -g
```
## Download and install
Clone repo and just run it like this:
```shell
git clone https://github.com/Chillaso/salary-calc.git && \
cd salary-calc && \
node index.js
```
## Usage
Now for usage just simply, where salary is your gross salary and irpf is your retention percentage:
```shell
node index.js --salary 20000 --irpf 20.20
```
or
```shell
node index.js --salary 20000 --irpf 20.20 --pays 14
```
## Business Logic
Basically we get your gross salary and multiply by two constants:
* "Contingencias comunes"
* "Base de accidente" 
> Sorry, I haven't fucking idea of that constants traslations.
Take your gross salary again and multiply by IRPF variable.
**Example**
* Step 1 (20000 * CONTINGENCIAS_COMUNES)  
* Step 2 (20000 * BASE_DE_ACCIDENTES)
* Step 3 (20000 * irpfVariable)

Just sum that 3 values and rest from your gross salary divided by pays, by default 12. So the final result is:

(salary / pays) - ((salary * CONTINGENCIAS_COMUNES) + (salary * BASE_DE_ACCIDENTES) + (salary * irpfVariable))

All result are fixed to 2 decimals, but all calc aren't fixed in order to get most exactly result.
