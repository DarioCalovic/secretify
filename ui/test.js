import { enc, SHA256 } from 'crypto-js'

createHashWithoutPadding = (val) => {
    return createHash(val).replace(/=/g, '')
}

createHash = (val) => {
    return enc.Base64.stringify(SHA256(val))
}

console.log('test')