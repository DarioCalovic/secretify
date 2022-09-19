<template>
  <div>{{ status }}</div>
</template>
<script>
export default {
  name: 'Test',
  data() {
    return {
      status: 'fail',
    }
  },
  mounted() {
    let cipher = 's3cretpSttt'
    const passphrases = ['Test123', '!Q@W#E$R%T6y7u8i9o0p', '😁😂🤣']

    passphrases.forEach(async (passphrase) => {
      try {
        if (passphrase) {
          passphrase = this.$crypto.createHashWithoutPadding(passphrase)
        }
        const key = await this.$crypto.generateEncryptionKeyString()

        if (passphrase) {
          cipher = await this.$crypto.encryptString(cipher, passphrase)
        }
        cipher = await this.$crypto.encryptString(cipher, key)

        console.log(cipher)
      } catch (e) {
        console.error(e)
      }
    })
  },
}
</script>
