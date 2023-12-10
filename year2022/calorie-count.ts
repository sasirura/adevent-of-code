import * as fs from 'fs'
import * as path from 'path'

try {
    const data = fs.readFileSync(path.join(__dirname, 'calorie-count-input.txt'), {
        encoding: 'utf-8',})

    let max: number = 0
    let current: number  = 0
    const lines = data.split(/\r?\n/)
    lines.forEach((line) => {
       if(line.length > 0) {
           current = current + parseInt(line)
       } else {
           if(current > max) {
               max = current
           }
           current = 0 
       }
    })
    console.log(max)
} catch (error) {
    console.error(error)
}
