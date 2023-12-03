import fs from "fs";

interface Entry {
    game: number;
    red: number;
    green: number;
    blue: number;
}

const result = fs.readFileSync(process.argv[2], "utf-8")
    .split(/\r?\n/)
    .filter(line => line.length > 0)
    .map(line => {
        return parseLine(line)
    })
    .reduce((acc: number, current: Entry) => {
        return acc + current.red * current.green * current.blue;;
    }, 0)

console.log(result);

function parseLine(line: string) {
    const [game, data] = line.split(/\s*:\s*/);
    const result: Entry = {
        game: Number(game.substring(5)),
        red: 0,
        green: 0,
        blue: 0
    }
    data.split(/\s*;\s*/)
        .forEach(set => {
            set.split(/\s*,\s*/)
                .forEach(entry => {
                    const [count, color] = entry.split(" ");
                    if (isValidColor(color)) {
                        result[color] = Math.max(result[color], Number(count));
                    }
                });
            return result;
        });

    return result;
}

function isValidColor(color: string): color is keyof Entry {
    return color === "red" || color === "green" || color === "blue";
}
