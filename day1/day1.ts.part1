import fs from "fs";

const result = fs.readFileSync(process.argv[2], "utf-8")
    .split(/\r?\n/)
    .map(line => {
        line = line.replace(/\D/g, "");
        return Number(line.length === 0 ? 0 : `${line[0]}${line[line.length-1]}`);
    })
    .reduce((acc, value) => acc + value);

console.log(result);
