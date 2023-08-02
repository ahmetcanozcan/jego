import custom from 'custom'
import fl from 'functional'
import echo from 'echo'
import mod from 'mod'


console.log("mod 27 % 15 = ", mod(27, 15))
console.log(echo("test"))
console.log("fl", JSON.stringify(fl))

export const foo = 'bar'

export default (v) =>  {
    return custom.double(v)
}