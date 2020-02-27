'use strict';
var argv = require('minimist')(process.argv.slice(2));
const COMMON_REDUCTIONS = 0.047;
const SS = 0.0165;

const init = (salary, irpf, pays) => {

  console.log("\x1b[4m", '===== Visit https://www2.agenciatributaria.gob.es/wlpl/PRET-R200/index.zul to calc your right IRPF retention ====', "\x1b[0m\n");
  controlParams(salary, irpf);

  pays = pays == undefined  ? 12 : pays;

  const monthSalary = salary / pays;
  var salaryIrfp = monthSalary * (irpf / 100);
  var salaryContingencias = monthSalary * COMMON_REDUCTIONS;
  var salarySS = monthSalary * SS;

  const totalReduction = (salaryIrfp + salaryContingencias + salarySS).toFixed(2);
  const totalReductionYear = ((salaryIrfp + salaryContingencias + salarySS) * pays).toFixed(2);

  var netSalary = (monthSalary - salaryIrfp - salarySS - salaryContingencias).toFixed(2);

  printWithColor("Total reductions per year", totalReductionYear);
  printWithColor("Gross salary per month", monthSalary.toFixed(2));
  printWithColor("Total reductions per month", totalReduction);
  printWithColor("Your net salary is", netSalary);
};

const controlParams = (salary, irpf) => {
  if(salary == undefined || irpf == undefined)
  {
    console.log("use --salary XXXX --irpf XXXX --pays 12 params to calc your salary");
    console.log("if you dont define --pays will be automatically set to 12.")
    process.exit(1);
  }   
};

/**
 * Print with yellow color
 * @param {String} text 
 * @param {String} toColor 
 */
const printWithColor = (text, toColor) => {
  console.log(text, "\x1b[33m-", toColor, "\x1b[0m");
};

init(argv.salary, argv.irpf, argv.pays);