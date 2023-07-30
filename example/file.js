import env from 'env'

console.log(Object.keys(env), "env")

export default (v) =>  {
    return v * 2
}