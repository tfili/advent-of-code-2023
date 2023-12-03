import { fork } from "child_process"
import path from "path"

const [,,day,...args] = process.argv
const modulePath = path.join(__dirname, `day${day}`, `day${day}`);

const child = fork(modulePath, args);
child.on("error", console.log);
