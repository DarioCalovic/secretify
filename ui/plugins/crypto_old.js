import { AES, enc, SHA256 } from 'crypto-js'
import { customAlphabet } from 'nanoid'
import { nolookalikes } from 'nanoid-dictionary'

const encryptionKeyLength = 26
const createHash = (key) => enc.Base64.stringify(SHA256(key))

const encrypt = (payload, key) => {
  const hash = createHash(key)
  return AES.encrypt(payload, hash).toString()
}

const decrypt = (payload, key) => {
  const hash = createHash(key)
  const bytes = AES.decrypt(payload, hash)
  return bytes.toString(enc.Utf8)
}

const generateNanoId = (length) => {
  const nanoid = customAlphabet(nolookalikes, length)
  return nanoid()
}

const generateEncryptionKey = () => generateNanoId(encryptionKeyLength)

export default (ctx, inject) => {
  inject('crypto', {
    encrypt,
    decrypt,
    generateEncryptionKey,
  })
}
