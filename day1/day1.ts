import fs from "fs";

const numberMapping : Record<string, number> = {
    one: 1,
    two: 2,
    three: 3,
    four: 4,
    five: 5,
    six: 6,
    seven: 7,
    eight: 8,
    nine: 9,
    zero: 0
}

const validValues = [...Object.keys(numberMapping), ...Object.values(numberMapping)];

const result = fs.readFileSync(process.argv[2], "utf-8")
    .split(/\r?\n/)
    .map(line => {
        let result: number[] = [];
        validValues.forEach((v:any) => {
            let index = -1;
            do {
                index = line.indexOf(v, index+1);
                if (index !== -1) {
                    result[index] = (typeof v === "number") ? v : numberMapping[v];
                }
            } while(index !== -1)
        })
        result = result.filter(v => v !== undefined);
        return Number(result.length === 0 ? 0 : result[0] * 10 + result[result.length-1]);
    })
    .reduce((acc, value) => acc + value);

console.log(result);
