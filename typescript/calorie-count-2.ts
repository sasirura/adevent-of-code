import * as fs from 'fs'
import * as path from 'path'

try {
    const data = fs.readFileSync(path.join(__dirname, 'calorie-count-input.txt'), {
        encoding: 'utf-8',})

    let topThree: number[] = []
    let current: number  = 0
    const lines = data.split("\n").concat("")
    lines.forEach((line) => {
       if(line.length > 0) {
           current = current + parseInt(line)
       } else {
           topThree.push(current)
           current = 0
       }
    })

    const final = topThree
    .sort((a,b) => b-a)
    .slice(0, 3)
    .reduce((a, b) => a + b)
    console.log(final)
} catch (error) {
    console.error(error)
}
