import {
  generateNanoId,
  decryptData,
  generateEncryptionKeyString,
  encryptFile,
  decryptFile,
  encryptString,
  decryptString,
  createHash,
  createHashWithoutPadding,
} from 'secretify-crypto'

export default (ctx, inject) => {
  inject('crypto', {
    generateNanoId,
    decryptData,
    generateEncryptionKeyString,
    encryptFile,
    decryptFile,
    encryptString,
    decryptString,
    createHash,
    createHashWithoutPadding,
  })
}
