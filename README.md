# SALARY CALC
This is a very basic app which let you calc your salary in Spain.

## Requirements
* Go
* Know your salary

## Usage
Now for usage just simply, where salary is your gross salary and irpf is your retention percentage:
```shell
go build salary.go
./salary --salary 20000 --irpf 20.20
```
or
```shell
./salary --salary 20000 --irpf 20.20 --pays 14
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

**All result are fixed to 2 decimals, but all calc aren't fixed in order to get most exactly result.**
